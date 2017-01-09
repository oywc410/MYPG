package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
)

func main() {
	service := ":7777"

	cert, err := tls.LoadX509KeyPair("../ca/CARoot.crt", "../ca/CARoot.key")
	checkError(err)

	config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
	listener, err := tls.Listen("tcp", service, &config)
	checkError(err)

	for {
		//开始侦听
		conn, err := listener.Accept()
		tlscon, ok := conn.(*tls.Conn)
		fmt.Println(ok)

		state := tlscon.ConnectionState()
		for _, v := range state.PeerCertificates {
			fmt.Println(x509.MarshalPKIXPublicKey(v.PublicKey))
		}

		if err != nil {
			fmt.Println(err)
			continue
		}

		//发送信息
		conn.Write([]byte("ye!"))
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
