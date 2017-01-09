package getAllFiles

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"sync"
)

//多线程控制数
var runNum int
var runTimeS int

//外部引用类
type Object struct {
	dirPath        string
	fileInfoStChan chan FileInfoSt
	isEnd          bool
	runTimePast    chan empty
	nextTimeNow    chan empty
	limit          chan empty
}

func SetDirPath(dirPath string) *Object {
	obj := &Object{dirPath: dirPath,
		fileInfoStChan: make(chan FileInfoSt, 2048),
		isEnd:          false}

	return obj
}

func (object *Object) GetFileInfoChan() <-chan FileInfoSt {
	object.startGetDirFile()
	return object.fileInfoStChan
}

//多线数量程控制
type empty struct{}

func init() {

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
}

func SetRunTimeS(run int) {
	runTimeS = run
}

//为遍历所有文件做准备并且控制子进程
func (object *Object) startGetDirFile() {
	object.isEnd = false
	object.dirPath = path.Clean(object.dirPath)

	if !Exists(object.dirPath) {
		panic(errors.New(object.dirPath + ",文件夹路径不存在"))
	}

	// 获取文件或目录信息
	fi, err := os.Stat(object.dirPath)
	if err != nil {
		panic(err)
	}

	if fi.IsDir() {

		object.limit = make(chan empty, runTimeS)

		object.nextTimeNow = make(chan empty, runTimeS)
		for j := 0; j < runTimeS; j++ {
			//标记无空隙进程
			object.nextTimeNow <- empty{}
		}

		go object.startRunTime()

	} else {
		panic(errors.New(object.dirPath + ",文件夹路径不存在"))
	}
}

func (object *Object) startRunTime() {
	//打开文件夹
	dir, err := os.Open(object.dirPath)
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
	dirPaths := object.dirPath + string(os.PathSeparator)

	for _, fi := range fis {
		if fi.IsDir() {
			dirsPath = append(dirsPath,
				dirPaths+fi.Name())
		}
		//记录文件信息
		//fmt.Println(fi.Name())
		object.fileInfoStChan <- FileInfoSt{dirPaths + fi.Name(), fi}
	}

	//异步读取子文件夹信息
	dirsPathLen := len(dirsPath)
	object.runTimePast = make(chan empty, dirsPathLen)
	if dirsPathLen != 0 {
		var wg sync.WaitGroup

		wg.Add(dirsPathLen)

		for _, nextDirPath := range dirsPath {
			go func(nextDirPath string, dirsPathLen int) {
				object.limit <- empty{}
				fileInfoStGetAll := object.getAllFile(nextDirPath)
				for _, fileInfo := range fileInfoStGetAll {
					object.fileInfoStChan <- fileInfo
				}
				object.runTimePast <- empty{}
				//获取多余线程
				overLen := dirsPathLen - len(object.runTimePast)
				if overLen < runTimeS {
					select {
					case <-object.nextTimeNow:
					default:
					}

				}
				<-object.limit
				wg.Done()
			}(nextDirPath, dirsPathLen)
		}

		wg.Wait()

	}
	object.isEnd = true
	object.fileInfoStChan <- FileInfoSt{"!END!", nil}
}

//输出指定目录下的所有文件
func (object *Object) getAllFile(dirPath string) []FileInfoSt {

	var fileInfoStGetAll []FileInfoSt

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
	runNowFileInfo := make(chan []FileInfoSt, runTimeS)

	//异步读取子文件夹信息
	for _, fi := range fis {
		//记录文件信息
		if fi.IsDir() {
			select {
			case object.nextTimeNow <- empty{}:
				runNow++
				fmt.Println(runNow)
				//fmt.Println(filePathString + fi.Name())
				go func(path string, runNowFileInfo chan []FileInfoSt, object *Object) {
					runNowFileInfo <- object.getAllFile(path)
				}(filePathString+fi.Name(), runNowFileInfo, object)
			default:
				fileInfoStGetAll = append(fileInfoStGetAll,
					object.getAllFile(filePathString+fi.Name())...)
			}
		}
		fileInfoStGetAll = append(fileInfoStGetAll,
			FileInfoSt{filePathString + fi.Name(), fi})
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
