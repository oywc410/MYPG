package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"regexp"
	"path"
	"os"
	"io"
	"log"
)

var digitsRegexp = regexp.MustCompile(`(/user_data.*jpg)`)
var digitsRegexp2 = regexp.MustCompile(`(/user_data.*gif)`)
var digitsRegexp3 = regexp.MustCompile(`(/user_data.*png)`)

func main() {
	resp, err := http.Get("http://local.toraichi.com/local.toraichi.com-1461568160849.log")
	//resp, err := http.Get("http://www.asahimatsu-shop.com/")
	if err != nil {
		// handle error
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	bodys := string(body)


	all := digitsRegexp.FindAllString(bodys, -1)

	for _, value := range all {
		//value = value[2:len(value)-1]
		value = value[1:]
		fmt.Println(value)

		r, err := http.Get("http://www.toraichi-shop.com/" + value)
		if err != nil {
			// handle error
			log.Fatalln(err)
		}

		os.MkdirAll(path.Dir(value), os.ModePerm)
		f, _ := os.Create(value)

		defer f.Close()
		_, err = io.Copy(f, r.Body)

	}

	all = digitsRegexp2.FindAllString(bodys, -1)

	for _, value := range all {
		//value = value[2:len(value)-1]
		value = value[1:]
		fmt.Println(value)

		r, err := http.Get("http://www.toraichi-shop.com/" + value)
		if err != nil {
			// handle error
			log.Fatalln(err)
		}

		os.MkdirAll(path.Dir(value), os.ModePerm)
		f, _ := os.Create(value)

		defer f.Close()
		_, err = io.Copy(f, r.Body)

	}

	all = digitsRegexp3.FindAllString(bodys, -1)

	for _, value := range all {
		//value = value[2:len(value)-1]
		value = value[1:]
		fmt.Println(value)

		r, err := http.Get("http://www.toraichi-shop.com/" + value)
		if err != nil {
			// handle error
			log.Fatalln(err)
		}

		os.MkdirAll(path.Dir(value), os.ModePerm)
		f, _ := os.Create(value)

		defer f.Close()
		_, err = io.Copy(f, r.Body)

	}
}
