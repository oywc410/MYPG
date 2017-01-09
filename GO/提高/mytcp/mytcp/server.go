package mytcp

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var isEnd = false

var exitChan chan bool

var logWrith io.Writer
var logObj *log.Logger

func setFileLog(mess interface{}) {

	if logObj == nil {
		logFileWrith, err := os.OpenFile("test.log", os.O_WRONLY, os.ModePerm)
		//写入文件的最后一行
		logFileWrith.Seek(0, os.SEEK_END)

		if err != nil && os.IsNotExist(err) {
			fmt.Println(err)
			logFileWrith, err = os.Create("test.log")
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
				return
			}
		}

		logWrith = logFileWrith
		logObj = log.New(logFileWrith, "[log]", log.Ldate|log.Ltime|log.Lshortfile)
	}

	logObj.Println(mess)
}

func ServerStart() {

	exitChan = make(chan bool)

	var tcpAddr *net.TCPAddr

	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")

	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)

	defer func() {
		tcpListener.Close()

		if closeIo, ok := logWrith.(*os.File); ok {
			closeIo.Close()
		}
	}()

	fmt.Println("start server")
	setFileLog("start server")

	go func() {
		for {
			tcpConn, err := tcpListener.AcceptTCP()

			if err != nil {
				fmt.Println(err)
				setFileLog(err)
				continue
			}

			fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())
			setFileLog("A client connected : " + tcpConn.RemoteAddr().String())
			go tcpPipe(tcpConn)
		}
	}()

	<-exitChan

	fmt.Println("Server End")
	setFileLog("Server End")
	//tcpListener.SetDeadline(time.Now().Add(time.Second * 2))
}

func tcpPipe(conn *net.TCPConn) {

	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnected :" + ipStr)
		setFileLog("disconnected :" + ipStr)
		err := conn.Close()
		fmt.Println(err)
	}()

	reader := bufio.NewReader(conn)

	for {

		if isEnd {
			return
		}

		messages, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		//fmt.Println(string(messages))

		if string(messages) == "end\n" {
			isEnd = true
			exitChan <- false
			return
		}

		msg := "test\n"
		b := []byte(msg)
		conn.Write(b)
	}
}
