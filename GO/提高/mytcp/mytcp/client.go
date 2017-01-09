package mytcp
import (
	"net"
	"bufio"
	"time"
	"fmt"
)


func ClientStart(endFunc func(code int8)) {


	var tcpAddr *net.TCPAddr
	var err error
	tcpAddr, err = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")

	if err != nil {
		fmt.Println(err)
		endFunc(1)
		return
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println(err)
		endFunc(2)
		return
	}
	defer conn.Close()
	fmt.Println("connectd!")

	b := []byte("time\n")
	_, err = conn.Write(b)
	if err != nil {
		endFunc(3)
	}

	onMessageRecoved(conn, endFunc)

	//endFunc(4)
}

func onMessageRecoved(conn *net.TCPConn, endFunc func(code int8)) {
	reader := bufio.NewReader(conn)
	for {
		_, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			endFunc(3)
			break
		}

		time.Sleep(time.Second)
		b := []byte("end1\n")

		_, err = conn.Write(b)
		if err != nil {
			fmt.Println(err)
			endFunc(3)
			break
		}
	}

}
