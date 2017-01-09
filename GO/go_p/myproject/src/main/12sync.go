package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	//多线程超时

	timeout := time.After(time.Second) //时间计时器(指针)
	//time.Sleep(time.Second)
	urls := []string{
		"https://golang.org/",
		"https://golang.org/",
		"https://golang.org/",
	}
	statusChan := getStatus(urls)

LOOP: //名称任意
	for {
		select {
		case status := <-statusChan:
			fmt.Println(status)
		case <-timeout:
			break LOOP //跳出for/select
		}
	}
}

func getStatus(urls []string) <-chan string {
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

	return statusChan
}
