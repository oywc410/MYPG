package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		//并行处理信息
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // 2分钟无接收信息超时
	//func (c *TCPConn) SetReadDeadline(t time.Time) error
	//func (c *TCPConn) SetWriteDeadline(t time.Time) error

	defer conn.Close()

	for {
		request := make([]byte, 128)
		//128 位读取信息
		read_len, err := conn.Read(request)

		if err != nil {
			fmt.Println(err)
			break
		}

		if read_len == 0 {
			break // connection already closed by client
		} else if string(request) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		} else {
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
		}

	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, "Fatal err: %s", err.Error())
		os.Exit(1)
	}
}
