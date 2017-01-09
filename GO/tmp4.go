package main

import (
	"os/exec"
	"log"
	"os"
	"path/filepath"
	"strings"
	"github.com/howeyc/fsnotify"
)


func main() {

	plgPath := filepath.Dir(os.Args[0])

	arr := strings.Split(plgPath, string(filepath.Separator))
	fileNameTop := arr[len(arr)-1]
	outFile := "C:/tmp/" + fileNameTop

	cmd := exec.Command("C:/Program Files/7-Zip/7z.exe", "a" , "-ttar" , outFile)
	//cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		log.Println(err)
		return
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println(err)
	}

	done := make(chan bool)

	// Process events
	go func() {
		log.Println(outFile)
		for {
			select {
			case ev := <-watcher.Event:
				arr2 := strings.Split(ev.Name, string(filepath.Separator))
				fileName := arr2[len(arr2)-1]
				if(fileName != fileNameTop + ".tar.tmp") {

					cmd := exec.Command("C:/Program Files/7-Zip/7z.exe", "a" , "-ttar" , outFile)
					err := cmd.Run()

					if err != nil {
						log.Println(err)
						return
					}

					log.Println("文件更新")
				}
				log.Println(ev)
			case err := <-watcher.Error:
				log.Println(err)
			}
		}
	}()

	log.Println(plgPath)

	err = watcher.Watch(plgPath)
	if err != nil {
		log.Println(err)
	}

	// Hang so program doesn't exit
	<-done

	/* ... do stuff ... */
	watcher.Close()
}