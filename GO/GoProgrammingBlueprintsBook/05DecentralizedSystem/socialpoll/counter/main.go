package main
import (
	"fmt"
	"flag"
	"os"
	"log"
	"labix.org/v2/mgo"
	"sync"
	"github.com/bitly/go-nsq"
	"time"
	"labix.org/v2/mgo/bson"
	"os/signal"
	"syscall"
)

//1秒间隔执行
const updateDuration  = 1 * time.Second

var fataErr error

func fatal (e error) {
	fmt.Println(e)
	flag.PrintDefaults()
	fataErr = e
}

func main() {
	defer func() {
		if fataErr != nil {
			os.Exit(1)
		}
	}()

	log.Println("数据库链接中")
	db, err := mgo.Dial("localhost")
	if err != nil {
		fatal(err)
		return
	}

	defer func() {
		log.Println("数据库链接中断")
		db.Close()
	}()

	pollData := db.DB("ballots").C("polls")

	var countsLock sync.Mutex
	var counts map[string]int

	//读取nsq上votes信息
	log.Println("NSQ链接中")
	q, err := nsq.NewConsumer("votes", "counter", nsq.NewConfig())
	if err != nil {
		fatal(err)
		return
	}

	q.AddHandler(nsq.HandlerFunc(func(m *nsq.Message) error {
		countsLock.Lock()
		defer countsLock.Unlock()
		if counts == nil {
			counts = make(map[string]int)
		}
		vote := string(m.Body)
		counts[vote]++
		return nil
	}))

	if err := q.ConnectToNSQLookupd("localhost:4161"); err != nil {
		fatal(err)
		return
	}

	log.Println("等待NSQ投票结果")

	//定期反复执行
	var updater *time.Timer
	updater = time.AfterFunc(updateDuration, func() {
		countsLock.Lock()
		defer countsLock.Unlock()
		if len(counts) == 0 {
			log.Println("没有投票更新,跳过数据库更新步骤")
		} else {
			log.Println("更新数据库中")
			log.Println(counts)
			ok := true
			for option, count := range counts {
				sel := bson.M{"options": bson.M{"$in": []string{option}}}
				up := bson.M{"$inc": bson.M{"results." + option: count}}
				if _, err := pollData.UpdateAll(sel, up); err != nil {
					log.Println("更新失败:", err)
					ok = false
					continue
				}
				counts[option] = 0
			}

			if ok {
				log.Println("数据库更新完毕")
				counts = nil
			}
		}

		updater.Reset(updateDuration)
	})

	//捕获Ctrl+C结束
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	for {
		select {
		case <-termChan://结束信号
			//结束数据库更新
			updater.Stop()
			q.Stop()
		case <-q.StopChan://等待数据库结束 跳出程序
			return
		}
	}
}
