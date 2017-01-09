package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	service := "127.0.0.1:9999"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	//func DialTCP(net string, laddr, raddr *TCPAddr) (c *TCPConn, err os.Error) 创建客户端连接
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	//接收返回数据
	conn.Write([]byte("client")) //非堵塞操作
	conn.Write([]byte("client"))
	conn.Write([]byte("client"))
	conn.Write([]byte("client"))

	fmt.Println("----------------------------------------")
	bytes := make([]byte, 1024)
	n, err := conn.Read(bytes) //非堵塞操作
	fmt.Println(string(bytes[:n]))

	bytes2, err := ioutil.ReadAll(conn) //堵塞操作 直到连接断开才结束
	fmt.Println(string(bytes2))

	fmt.Println("----------------------------------------")

	os.Exit(0)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
