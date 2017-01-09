package main

import (
	"fmt"
	"log"
	"net"
	"oTsl"
)

func main() {
	config := oTsl.NewConfig()
	config.SetListenT("tcp")
	config.SetListenIp("127.0.0.1:9999")
	config.SetCrtPath("certs/CARoot2.crt")
	config.SetKeyPath("certs/CARoot2.key")

	tslClient := oTsl.NewClient()
	tslClient.SetConfig(config)

	tslClient.SendMess([]byte("Hello"), func(conn net.Conn, err error) error {
		if err == nil {
			reply := make([]byte, 256)
			n, err := conn.Read(reply)

			if err != nil {
				return err
			}

			log.Printf("client: read %q (%d bytes)", string(reply[:n]), n)
		}

		return nil
	})

	tslClient.SendMess([]byte("Hello2"), func(conn net.Conn, err error) error {
		if err == nil {
			reply := make([]byte, 256)
			n, err := conn.Read(reply)

			if err != nil {
				return err
			}

			log.Printf("client: read %q (%d bytes)", string(reply[:n]), n)
		}

		return nil
	})

	err := tslClient.Dial()
	if err != nil {
		fmt.Println(err)
	}

}
