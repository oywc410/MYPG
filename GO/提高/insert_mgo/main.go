package main

import (
	"time"
	"strconv"
	"fmt"
	"runtime"
	"gopkg.in/mgo.v2"
	"sync"
	"log"
	"flag"
	"os"
	"runtime/pprof"
	"runtime/debug"
	"net/http"
	_ "net/http/pprof"
)


var cpuprofile = flag.String("cpuprofile", "cpuprofile.p", "")
var memprofile = flag.String("memprofile", "memprofile.p", "")
var blockprofile = flag.String("blockprofile", "blockprofile.p", "")
var goroutineprofile = flag.String("goroutineprofile", "goroutineprofile.p", "")
var heapdumpfile = flag.String("heapdumpfile", "heapdumpfile.p", "")

type T struct {
	B_id string
	B_coins int
}

var db *mgo.Database
var c *mgo.Collection
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//http://127.0.0.1:6060/debug/pprof/

	go func() {
		runtime.SetBlockProfileRate(1)
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()

	flag.Parse()
	//CPU追踪
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("start cpu write heap profile....")
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	//内存追踪
	if *memprofile != "" {
		var err error
		memFile, err := os.Create(*memprofile)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("start mem write heap profile....")
			pprof.WriteHeapProfile(memFile)
			defer memFile.Close()
		}
	}

	//协程堵塞追踪
	if *blockprofile != "" {
		blockFile, err := os.Create(*blockprofile)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("start block write heap profile....")
			runtime.SetBlockProfileRate(1)
			defer pprof.Lookup("block").WriteTo(blockFile, 0)
		}
	}

	//协程运行数
	if *goroutineprofile != "" {
		goFile, err := os.Create(*goroutineprofile)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("start goroutine write heap profile....")

			pprof.Lookup("goroutine").WriteTo(goFile, 0)
			defer goFile.Close()
		}
	}

	//堆倾卸器
	if *heapdumpfile != "" {
		heapFile, err := os.Create(*heapdumpfile)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("start heapdump write heap profile....")

			debug.WriteHeapDump(heapFile.Fd())
			defer heapFile.Close()
		}
	}

	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	db = session.DB("test")
	c = db.C("t")
	insertData()
}
func insertData() {
	count := 50000
	//insertSync(count)
	insertAsync(count)
}
func insertAsync(count int) {
	time1 := time.Now()
	var wg sync.WaitGroup
	lim := make(chan struct{}, 100)

	for index := 0; index < count; index++ {
		wg.Add(1)
		lim <- struct {}{}
		go func(i int) {
			t := &T{B_id: "b" + strconv.Itoa(i), B_coins: i}
			err := c.Insert(t)
			if err != nil {
				log.Println(err)
			}
			<- lim
			wg.Done()
		}(index)
	}
	wg.Wait()
	time2 := time.Now()
	fmt.Printf("async insert cost time: %s \n", time2.Sub(time1))
}
func insertSync(count int) {
	time1 := time.Now()
	for index := 0; index < count; index++ {
		t := &T{B_id: "b" + strconv.Itoa(index), B_coins: index}
		c.Insert(t)
	}
	time2 := time.Now()
	fmt.Printf("sync insert cost time: %s \n", time2.Sub(time1))
}