package main

import (
	"os"
	"log"
	"syscall"
	"fmt"
)

const APPEND_DATA int64 = 1024 * 1024

func main() {
	var err error
	dateFile, err := os.Create("test.data")
	if err != nil {
		log.Fatalln(err)
	}
	defer dateFile.Close()

	dateFile.Write([]byte("aaaaaal"))

	f, err := dateFile.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	_, err = syscall.Mmap(int(dateFile.Fd()), 0, int(f.Size() + APPEND_DATA), syscall.PROT_READ | syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		log.Fatalln(err)
	}

}
