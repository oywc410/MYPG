package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stdout, resp.Status)
		fmt.Fprintf(os.Stdout, resp.Header.Get("Content-Length"))

		b, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stdout, "fetch: reading %s: %v\n", url, err)

			os.Exit(1)
		}

		outFile, err := os.Create("out.html")
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		n, err := outFile.WriteString(string(b))
		defer outFile.Close()

		fmt.Printf("%s", b)

		fmt.Println(n, err)
	}
}
