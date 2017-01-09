package main

import (
	"errors"
	"fmt"
	//"io"
	"os"
	"path"
	"runtime"
	"sync"
	"time"
)

//多线程控制数
var runNum int
var runTimeS int

//文件信息
type fileInfoSt struct {
	filePath string
	fileInfo os.FileInfo
}

//多线程数据通道
var fileInfoStChan chan fileInfoSt

//多线数量程控制
type empty struct{}

var limit chan empty

//记录执行完成的线程数
var runTimePast chan empty

//控制可以启用的多余线程
var nextTimeNow chan empty

func main() {
	t1 := time.Now()

	//设置CPU数
	runNum = runtime.NumCPU()
	if runNum != 1 {
		runtime.GOMAXPROCS(runNum - 1)
	}

	//最多线程数
	runTimeS = 1
	if runNum > 2 {
		runTimeS = runNum - 1
	}

	//runTimeS = 200
	//遍历所有文件
	//startGetDirFile("D:/学习")
	startGetDirFile("D:/WEB/www")
	//startGetDirFile("D:/tmp/MYPG/go_backup/file")
	//startGetDirFile("D:/WEB/")
	i := 0
LOOP:
	for {
		select {
		case fileInfoSt := <-fileInfoStChan:
			if fileInfoSt.filePath == "!END!" {
				break LOOP
			} else {
				//fmt.Println(fileInfoSt.filePath)
				if !fileInfoSt.fileInfo.IsDir() {
					i++

					if i%1000 == 0 {
						fmt.Println(i)
						fmt.Println(len(fileInfoStChan))
						fmt.Println(fileInfoSt.filePath)
					}

				}
			}
		}
	}
	fmt.Println(i)
	fmt.Println(len(fileInfoStChan))
	fmt.Println(time.Now().Sub(t1))
}

//为遍历所有文件做准备并且控制子进程
func startGetDirFile(dirPath string) {

	fileInfoStChan = make(chan fileInfoSt, 2048)
	dirPath = path.Clean(dirPath)

	if !Exists(dirPath) {
		panic(errors.New(dirPath + ",文件夹路径不存在"))
	}

	// 获取文件或目录信息
	fi, err := os.Stat(dirPath)
	if err != nil {
		panic(err)
	}

	if fi.IsDir() {

		limit = make(chan empty, runTimeS)

		nextTimeNow = make(chan empty, runTimeS)
		for j := 0; j < runTimeS; j++ {
			//标记无空隙进程
			nextTimeNow <- empty{}
		}

		go startRunTime(dirPath)

	} else {
		panic(errors.New(dirPath + ",文件夹路径不存在"))
	}
}

func startRunTime(dirPath string) {
	//打开文件夹
	dir, err := os.Open(dirPath)
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	//读取文件列表
	fis, err := dir.Readdir(0) //n<=0 读取所有 n>0 读取n个
	if err != nil {
		panic(err)
	}

	var dirsPath []string

	for _, fi := range fis {
		if fi.IsDir() {
			dirsPath = append(dirsPath, dirPath+string(os.PathSeparator)+fi.Name())
		}
		//记录文件信息
		fileInfoStChan <- fileInfoSt{fi.Name(), fi}
	}

	dirsPathLen := len(dirsPath)
	runTimePast = make(chan empty, dirsPathLen)
	if dirsPathLen != 0 {
		var wg sync.WaitGroup

		wg.Add(dirsPathLen)

		for _, nextDirPath := range dirsPath {
			go func(nextDirPath string, dirsPathLen int) {
				limit <- empty{}
				fileInfoStGetAll := getAllFile(nextDirPath)
				for _, fileInfo := range fileInfoStGetAll {
					fileInfoStChan <- fileInfo
				}
				runTimePast <- empty{}
				//获取多余线程
				overLen := dirsPathLen - len(runTimePast)
				if overLen < runTimeS {
					select {
					case <-nextTimeNow:
					default:
					}

				}
				<-limit
				wg.Done()
			}(nextDirPath, dirsPathLen)
		}

		wg.Wait()

	}
	fileInfoStChan <- fileInfoSt{"!END!", nil}
}

//输出指定目录下的所有文件
func getAllFile(dirPath string) []fileInfoSt {

	var fileInfoStGetAll []fileInfoSt

	//打开文件夹
	dir, err := os.Open(dirPath)
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	//读取文件列表
	fis, err := dir.Readdir(0) //n<=0 读取所有 n>0 读取n个
	if err != nil {
		panic(err)
	}

	filePathString := dirPath + string(os.PathSeparator)

	//当前循环中执行多线程个数
	runNow := 0
	//记录循环中多线程的结果
	runNowFileInfo := make(chan []fileInfoSt, runTimeS)

	for _, fi := range fis {
		//记录文件信息
		if fi.IsDir() {
			select {
			case nextTimeNow <- empty{}:
				runNow++
				//fmt.Println(runNow)
				//fmt.Println(filePathString + fi.Name())
				go func(path string, runNowFileInfo chan []fileInfoSt) {
					runNowFileInfo <- getAllFile(path)
				}(filePathString+fi.Name(), runNowFileInfo)
			default:
				fileInfoStGetAll = append(fileInfoStGetAll,
					getAllFile(filePathString+fi.Name())...)
			}
		}
		fileInfoStGetAll = append(fileInfoStGetAll,
			fileInfoSt{filePathString + fi.Name(), fi})
	}

	if runNow > 0 {
	LOOP:
		for {
			select {
			case fileInfo := <-runNowFileInfo:

				runNow--
				fileInfoStGetAll = append(fileInfoStGetAll, fileInfo...)
				if runNow == 0 {
					break LOOP
				}
			}
		}
	}

	return fileInfoStGetAll
}

// 判断档案是否存在
func Exists(name string) bool {
	_, err := os.Stat(name)
	return err == nil || os.IsExist(err)
}
