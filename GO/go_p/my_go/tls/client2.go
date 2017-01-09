package main

import (
	"fmt"
	"io"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Sprintf("client: dial: %s", err)
	}
	defer conn.Close()

	fmt.Println("client: connected to: ", conn.RemoteAddr())

	message := "Hello\n"
	n, err := io.WriteString(conn, message)
	if err != nil {
		fmt.Println("client: write: %s", err)
	}
	fmt.Println("client: wrote %q (%d bytes)", message, n)

	reply := make([]byte, 256)
	n, err = conn.Read(reply)
	fmt.Println("client: read %q (%d bytes)", string(reply[:n]), n)
	fmt.Println("client: exiting")
}
