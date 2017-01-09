package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"sync"
)

func main() {

	//2种多线程实例

	//获取CPU数
	cpus := runtime.NumCPU()
	//指定多线程时CPU数
	runtime.GOMAXPROCS(cpus)
	fmt.Println(cpus)
	go1()
	go2()
}

func go1() {
	wait := new(sync.WaitGroup)
	urls := []string{
		"https://golang.org/",
		"https://golang.org/",
		"https://golang.org/",
	}

	for _, url := range urls {
		// waitGroupに追加
		wait.Add(1)
		// 取得処理をゴルーチンで実行する
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			fmt.Println(url, res.Status)
			// waitGroupから削除
			wait.Done()
		}(url)
	}
	// 待ち合わせ(3个进程进行等待)
	wait.Wait()

}

func go2() {
	urls := []string{
		"https://golang.org/",
		"https://golang.org/",
		"https://golang.org/",
	}
	statusChan := make(chan string)
	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			statusChan <- res.Status
		}(url)
	}

	//输出最快执行完成结果
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-statusChan)
	}

}
