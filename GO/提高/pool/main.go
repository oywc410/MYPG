package main

import (
	"sync"
	"bytes"
	"fmt"
	"runtime"
	"time"
)

func main() {
	test1()
}

func test1() {
	//Pool内存缓存 GC时将被清除

	var bp sync.Pool
	var looks sync.RWMutex
	bp.New = func() interface{} {
		return &bytes.Buffer{}
	}

	buf := bp.Get().(*bytes.Buffer)
	buf.WriteString("test")
	fmt.Println(buf.String())
	bp.Put(buf)//记录至缓存

	limit := make(chan struct{}, 10)

	go func() {
		buf := bp.Get().(*bytes.Buffer)
		buf.WriteString("tttt")
		looks.Lock()
		bp.Put(buf)
		looks.Unlock()
	}()
	time.Sleep(1 * time.Second)
	for {
		limit <- struct {}{}
		go func() {
			//提取以后必须再次Put 否则将被清除
			buf := bp.Get().(*bytes.Buffer)
			if buf.Len() != 0 {
				fmt.Println(buf.String())
				looks.Lock()
				bp.Put(buf)
				looks.Unlock()
				//runtime.GC()
			}
			<- limit
		}()
	}

	return
	runtime.GC()//手动启动GC

	buf = bp.Get().(*bytes.Buffer)//由于缓存被清空 反问内容为空
	fmt.Println(buf.String())
}

