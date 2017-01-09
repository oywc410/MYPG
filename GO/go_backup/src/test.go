package main

import (
	"compress/gzip"
	"fmt"
	"io"
	getAllFiles "oGetAllFiles"
	tarFiles "oTarFiles"
	unTarFiles "oUnTarFiles"
	"os"
	"path"
	"runtime"
	"time"
)

func main() {

	t1 := time.Now()
	//main1()
	mainTar("F:/BaiduYunDownload")
	//mainUnTar("test.tar.gz")
	fmt.Println(time.Now().Sub(t1))

}

//压缩文件
func mainTar(dirPath string) {

	i := 0
	dirPath = path.Clean(dirPath)
	getAllFileObj := getAllFiles.SetDirPath(dirPath)
	getAllChan := getAllFileObj.GetFileInfoChan()

	fw, err := os.Create("test.tar.gz")
	defer fw.Close()
	if err != nil {
		fw.Close()
		panic(err)
	}

	//gizp write
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	tarObj := tarFiles.CreateTar(func() io.Writer {
		return gw
	})

	tarEndChan := make(chan string)
	tarFileChan := make(chan getAllFiles.FileInfoSt, 100)
	var tarFileArr []getAllFiles.FileInfoSt

	//异步方式添加打包文件
	go func() {
	LOOP2:
		for {
			select {
			case fileInfo := <-tarFileChan:
				if fileInfo.GetFilePath() == "!END!" { //判断目标目录信息是否收集完成

					if len(tarFileArr) > 0 {
						//读取缓存信息,将缓存信息中的文件进行打包
						for _, fileInfo := range tarFileArr {
							if fileInfo.GetFilePath() == "!END!" {

							} else {

								fmt.Println(fileInfo.GetFilePath(), "[", fileInfo.GetFileInfo().IsDir(), "]>>>>>>>>>>>>to_tar2")

								srcRelative := fileInfo.GetFilePath()
								catSrcRelative := catPath(srcRelative, len(dirPath))
								//fmt.Println(catSrcRelative, ">>>>>>>>>uri")
								tarObj.AddFileInfoSt(fileInfo, catSrcRelative)
							}
						}
					}
					//发送打包结束信息
					tarEndChan <- "end"
					fmt.Println("break LOOP2")
					break LOOP2
				}

				//将数据写入包中
				fmt.Println(fileInfo.GetFilePath(), "[", fileInfo.GetFileInfo().IsDir(), "]>>>>>>>>>>>>to_tar1")

				srcRelative := fileInfo.GetFilePath()
				catSrcRelative := catPath(srcRelative, len(dirPath))

				//fmt.Println(catSrcRelative, ">>>>>>>>>uri")
				tarObj.AddFileInfoSt(fileInfo, catSrcRelative)

				//runtime.Gosched()
			}
		}
	}()

	//读取文件信息
LOOP:
	for {
		select {
		case fileInfo := <-getAllChan:
			//目标文件夹数据收集完毕
			if fileInfo.GetFilePath() == "!END!" {
			LOOP3:
				//向打包进程发送结束命令
				for {
					select {
					case tarFileChan <- getAllFiles.FileInfoSt{"!END!", nil}:
						break LOOP3
					default:
					}
				}
				fmt.Println("break LOOP1")
				break LOOP
			} else {
				if !fileInfo.GetFileInfo().IsDir() {
					i++
				}
			}

			runtime.Gosched()

			select {
			case tarFileChan <- fileInfo: //发送数据 打包多线程
				fmt.Println(fileInfo.GetFilePath(), "[", fileInfo.GetFileInfo().IsDir(), "]>>>>>>>>>>>>to_tar_chan")
			default: //通道满时把数据存入缓存
				fmt.Println(fileInfo.GetFilePath(), "[", fileInfo.GetFileInfo().IsDir(), "]>>>>>>>>>>>>to_arr_cache")
				tarFileArr = append(tarFileArr, fileInfo)
			}
		default:
		}
	}

	fmt.Println("-----------------------------------------------------")

	//判断打包结束
LOOP4:
	for {
		select {
		case <-tarEndChan:
			break LOOP4
		}
	}

	defer tarObj.CloseFile()

	fmt.Println(i)
}

func mainUnTar(unTarFile string) {
	fr, err := os.Open(unTarFile)
	defer fr.Close()
	if err != nil {
		panic(err)
	}

	gr, err := gzip.NewReader(fr)
	defer gr.Close()
	if err != nil {
		panic(err)
	}

	unTar := unTarFiles.CreateUnTar(gr)
	unTar.GetAllFileChan()
	fmt.Println("----END----")
}

func catPath(str string, start int) string {

	rs := []rune(str)
	return string(rs[start:])
}

func main1() {

	getAllFileObj := getAllFiles.SetDirPath("D:/WEB/www")
	getAllChan := getAllFileObj.GetFileInfoChan()

	i := 0

LOOP:
	for {
		select {
		case fileInfo := <-getAllChan:
			//fmt.Println(fileInfo.FilePath)
			if fileInfo.GetFilePath() == "!END!" {
				break LOOP
			} else {
				if fileInfo.GetFileInfo().IsDir() {

				} else {
					i++
					if i%100 == 0 {
						fmt.Println(fileInfo.GetFilePath())
					}
				}
			}

		}
	}

	fmt.Println(i)
}

func main2() {
	getAllFileObj := getAllFiles.SetDirPath("D:/WEB/www")
	getAllFileObj2 := getAllFiles.SetDirPath("D:/WEB/www")
	getAllFileObj3 := getAllFiles.SetDirPath("D:/WEB/www")

	getAllChan := getAllFileObj.GetFileInfoChan()
	getAllChan2 := getAllFileObj2.GetFileInfoChan()
	getAllChan3 := getAllFileObj3.GetFileInfoChan()

	isOutLoop := false
	isOutLoop2 := false
	isOutLoop3 := false

	i, j, z := 1, 1, 1

LOOP:
	for {
		select {
		case fileInfo := <-getAllChan:
			//fmt.Println(fileInfo.FilePath)
			if fileInfo.GetFilePath() == "!END!" {
				isOutLoop = true
			} else {

				if fileInfo.GetFileInfo().IsDir() {
				} else {
					i++
					if i%1000 == 0 {
						fmt.Println(i)
						fmt.Println(fileInfo.GetFilePath())
					}
				}
			}
		case fileInfo := <-getAllChan2:
			//fmt.Println(fileInfo.FilePath)
			if fileInfo.GetFilePath() == "!END!" {
				isOutLoop2 = true
			} else {

				if fileInfo.GetFileInfo().IsDir() {
				} else {
					j++
					if j%1000 == 0 {
						fmt.Println(j)
						fmt.Println(fileInfo.GetFilePath())
					}
				}
			}
		case fileInfo := <-getAllChan3:
			//fmt.Println(fileInfo.FilePath)
			if fileInfo.GetFilePath() == "!END!" {
				isOutLoop3 = true
			} else {

				if fileInfo.GetFileInfo().IsDir() {
				} else {
					z++
					if z%1000 == 0 {
						fmt.Println(z)
						fmt.Println(fileInfo.GetFilePath())
					}
				}
			}

		default:
			if isOutLoop && isOutLoop2 && isOutLoop3 {
				fmt.Println(isOutLoop)
				fmt.Println(i)
				fmt.Println(isOutLoop2)
				fmt.Println(j)
				fmt.Println(isOutLoop3)
				fmt.Println(z)
				break LOOP
			}

		}
	}

}

func catNewDir(str string) string {
	rs := []rune(str)
	end := 0
	for keys, a := range rs {
		if a == 92 {
			end = keys
		}
	}

	return string(rs[:end])
}
