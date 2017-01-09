package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://127.0.0.1/",
		"http://127.0.0.1/",
		"http://127.0.0.1/",
	}

	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}

			defer res.Body.Close()
			fmt.Println(url, res.Status)
		}(url)
	}

	//执行等待
	time.Sleep(time.Second)
}
