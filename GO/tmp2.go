package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"path"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	skillfolder := `C:\tmp\mt`

	keyMap := make(map[int]int)

	files, _ := ioutil.ReadDir(skillfolder)
	for _,file := range files {
		if file.IsDir() {
			continue
		} else {
			//fmt.Println(file.Name())
			fileobj, _ := os.Open(path.Join(skillfolder, file.Name()))

			re := bufio.NewScanner(fileobj)
			for re.Scan() {
				str := re.Text()
				strs := strings.Split(str, " ")
				if len(strs) == 4 {
					key, _ := strconv.Atoi(strs[3])
					keyMap[key]++
				}

			}
		}
	}
	fmt.Println(len(keyMap))
	fmt.Println("------------------------")
	for key, value := range keyMap {

		if value > 1 {
			fmt.Println(key, ":", value)
		}
	}
}
