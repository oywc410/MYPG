package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	//当url执行多时make会对内存照成压力
	//所以对getStatus进行以下改良

	urls := []string{
		"https://golang.org/",
		"https://golang.org/",
		"https://golang.org/",
		"https://golang.org/",
		"https://golang.org/",
		"https://golang.org/",
		"https://golang.org/",
		"https://golang.org/",
		"https://golang.org/",
		"https://golang.org/",
		"https://golang.org/",
		"https://golang.org/",
		"https://golang.org/",
		"https://golang.org/",
		"https://golang.org/",
	}
	statusChan := getStatus(urls)

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-statusChan)
	}

}

var empty struct{} // サイズがゼロの構造体
func getStatus(urls []string) <-chan string {
	statusChan := make(chan string, 3) //固定使用内存大小
	// バッファを5に指定して生成
	limit := make(chan struct{}, 5)
	go func() {
		for _, url := range urls {
			select {
			case limit <- empty: //利用通道的缓存特性来限制go执行数量
				// limitに書き込みが可能な場合は取得処理を実施
				go func(url string) {
					// このゴルーチンは同時に5つしか起動しない
					res, err := http.Get(url)
					if err != nil {
						log.Fatal(err)
					}
					statusChan <- res.Status
					// 終わったら1つ読み出して空きを作る
					<-limit
				}(url)
			}
		}
	}()
	return statusChan
}
