package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	//创建监听端口
	ln, err := net.Listen("tcp", ":5000")

	fmt.Println("Listen 5000")

	if err != nil {

	}

	for {
		//获取客户链接
		conn, err := ln.Accept()
		if err != nil {
			//避开错误处理
			continue
		}
		fmt.Println(&conn)
		//进入数据处理  go 非堵塞
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	myIp := conn.LocalAddr().String()
	fmt.Println("Server: ", myIp)
	clientIp := conn.RemoteAddr().String()
	//fmt.Printf("%T", clientIp)

	_, w1_err := conn.Write([]byte("client: " + clientIp + "\n"))
	if w1_err != nil {

	}

	_, w2_err := conn.Write([]byte("server: " + myIp + "\n"))
	if w2_err != nil {

	}

	conn.SetDeadline(time.Now().AddDate(0, 0, 1)) //设置超时时间 0为不超时

	nameInfo := make([]byte, 512) //生成一个缓存数组
	_, err := conn.Read(nameInfo)
	checkError(err)

	stringBuf := ""

	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf) //读取客户机发的消息
		flag := checkError(err)
		if flag != 1 {
			break
		}

		if buf[0] == 13 || buf[0] == 10 { // 用户输入换行判断 \r or \n
			fmt.Println(stringBuf)

			if stringBuf == "end" {
				conn.Write([]byte("\n---------------\n"))
				conn.Write([]byte("\n88\n"))
				conn.Close()
			}

			stringBuf = ""
		} else {
			for _, t := range buf {
				if t == 0 {
					break
				}
				stringBuf = stringBuf + string(t)
			}
		}
	}

	//conn.Close()
}

//错误判断
func checkError(err error) int {
	if err != nil {
		if err.Error() == "EOF" {
			fmt.Println("用户退出了")
			return 0
		}
		fmt.Println(err.Error())
		//log.Fatal("an error!", err.Error())
		return -1
	}
	return 1
}
