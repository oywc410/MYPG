package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":7777"
	//TCP地址对象 ()
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	//侦听对象
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	//确定 不可以同时执行任务
	for {
		//开始侦听
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		daytime := time.Now().String()
		fmt.Println(daytime)
		conn.Write([]byte("ye!"))
		//发送信息
		conn.Write([]byte(daytime))
		//关闭连接
		conn.Close()
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
