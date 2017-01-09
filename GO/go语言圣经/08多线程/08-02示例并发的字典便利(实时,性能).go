package main
import (
	"io/ioutil"
	"os"
	"path/filepath"
	"fmt"
	"flag"
	"time"
	"sync"
)


func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}

	return entries
}

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {

	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan int64)
	var n sync.WaitGroup
	go func() {
		for _, root := range roots {
			n.Add(1)
			go walkDir(root, &n, fileSizes)
		}

	}()

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	//每500ms生成事件
	var tick <- chan time.Time

	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nfiles, nbytes int64

loop:
	for {
		select {
		case size, ok := <- fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <- tick:
			printDiskUsage(nfiles, nbytes)
		}
	}

	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %%.1f GB\n", nfiles, float64(nbytes)/1e9)
}