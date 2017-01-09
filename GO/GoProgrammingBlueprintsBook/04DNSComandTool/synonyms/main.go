package main

import (
	"os"
	"./thesaurus"
	"bufio"
	"log"
	"fmt"
)

func main() {
	apiKey := os.Getenv("BHT_APIKEY")
	thesaurus := &thesaurus.BigHuge{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalf("%q检索失败:%v\n", word, err)
		}
		if len(syns) == 0 {
			log.Fatal("%q类似词不存在")
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
