package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	//数组通道缓存 堵塞与非堵塞
	//make(chan string, n) n : 设置缓存

	urls := []string{
		"https://golang.org/",
		"https://golang.org/",
		"https://golang.org/",
	}
	statusChan := getStatus(urls)

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-statusChan)
	}
}

func main1() {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Second)
		ch <- "a" // 1秒後にデータを書き込む
	}()
	<-ch // 1秒後にデータが書き込まれるまでブロック
}

func main2() {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Second)
		<-ch // 1秒後にデータを読み出す
	}()
	ch <- "a" // 1秒後にデータが読み出されるまでブロック
}

func main3() {
	ch := make(chan string, 3)
	go func() {
		time.Sleep(time.Second)
		<-ch // 1秒後にデータを読み出す
	}()
	ch <- "a" // ブロックしない
	ch <- "b" // ブロックしない
	ch <- "c" // ブロックしない
	ch <- "d" // 1秒後にデータが読み出されるまでブロック
}

//改进之前的函数s
func getStatus(urls []string) <-chan string {
	// バッファをURLの数（3）に
	statusChan := make(chan string, len(urls))
	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			// main()の読み出しが遅くても
			// 3つのゴルーチンはすぐに終わる
			statusChan <- res.Status
		}(url)
	}
	return statusChan
}
