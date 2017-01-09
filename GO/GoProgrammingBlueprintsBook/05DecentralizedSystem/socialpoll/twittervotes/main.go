package main

import (
	"gopkg.in/mgo.v2"
	"log"
	"sync"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	if err := dialdb(); err != nil {
		log.Fatalln("MongoDb写入失败")
	}
	defer closedb()

	var stoplock sync.Mutex
	stop := false
	stopChan := make(chan struct{}, 1)
	//捕获程序停止(Ctrl+C)
	sigalChan := make(chan os.Signal, 1)
	go func() {
		<-sigalChan//等待程序结束
		stoplock.Lock()
		stop = true
		stoplock.Unlock()
		log.Println("终止中")
		stopChan <- struct{}{}
		closeConn()

	}()
	signal.Notify(sigalChan, syscall.SIGINT, syscall.SIGTERM)

	//处理开始
	votes := make(chan string)
	publisherStoppedChan := publishVotes(votes)
	twitterStoppendChan := startTwitterStream(stopChan, votes)

	go func() {
		for {
			time.Sleep(1 * time.Minute)
			//每一分钟断开链接
			closeConn()
			stoplock.Lock()
			if stop {
				//捕获停止运行时跳出寻循环
				stoplock.Unlock()
				break
			}
			stoplock.Unlock()
		}
	}()

	<-twitterStoppendChan
	close(votes)
	<-publisherStoppedChan
}

var db *mgo.Session

func dialdb() error {
	var err error
	log.Println("MongoDB链接中: localhost")
	db, err = mgo.Dial("localhost")
	return err
}

func closedb() {
	db.Close()
	log.Println("数据库链接中断")
}

type poll struct {
	Options []string
}

func loadOptions() ([]string, error) {
	var options []string
	iter := db.DB("ballots").C("polls").Find(nil).Iter()
	var p poll
	for iter.Next(&p) {
		options = append(options, p.Options...)
	}
	iter.Close()
	return options, iter.Err()
}
