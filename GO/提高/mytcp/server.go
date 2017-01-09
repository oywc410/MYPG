package main

import (
	"./mytcp"
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"runtime"
	"runtime/debug"
	_ "net/http/pprof"
	"net/http"
)

var cpuprofile = flag.String("cpuprofile", "", "")
var memprofile = flag.String("memprofile", "", "")
var blockprofile = flag.String("blockprofile", "", "")
var goroutineprofile = flag.String("goroutineprofile", "", "")
var heapdumpfile = flag.String("heapdumpfile", "", "")

func main() {

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

	mytcp.ServerStart()
}
