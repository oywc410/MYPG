package main
import (
	"fmt"
	"github.com/golang-china/gopl-zh/vendor/gopl.io/ch5/links"
	"log"
	"os"
)

//控制线程数
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)

	tokens <- struct{}{}

	list, err := links.Extract(url)

	<- tokens

	if err != nil {
		log.Print(err)
	}

	return list
}

func main() {
	workList := make(chan []string)

	go func() {
		workList <- os.Args[1:]
	}()

	seen := make(map[string]bool)

	for list := range workList {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					workList <- crawl(link)
				}(link)
			}
		}
	}
}
