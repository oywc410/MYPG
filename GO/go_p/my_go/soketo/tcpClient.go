package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	/*
		type TCPAddr struct {
		    IP IP
		    Port int
		}
	*/
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	//func DialTCP(net string, laddr, raddr *TCPAddr) (c *TCPConn, err os.Error) 创建客户端连接
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	//写入数据
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	//接收返回数据
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
