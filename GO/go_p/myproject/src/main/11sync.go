package main

import (
	"fmt"
	"log"
	"net/http"
)

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

func main() {

	//多线程的封装

	urls := []string{
		"https://golang.org/",
		"https://golang.org/",
		"https://golang.org/",
	}
	statusChan := getStatus(urls)

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-statusChan)
	}

	/*
		ch1 := make(chan string)
		ch2 := make(chan string)
		for {
			select {
			case c1 := <-ch1:
			// ch1からデータを読み出したときに実行される
			case c2 := <-ch2:
			// ch2からデータを読み出したときに実行される
			case ch2 <- "c":
			// ch2にデータを書き込んだときに実行される
			default:
			// caseが実行されなかった場合に実行される
			}
		}
		//default被执行后会自动跳出select
	*/
}
