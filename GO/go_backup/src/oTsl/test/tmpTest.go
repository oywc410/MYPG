package main

import (
	"fmt"
	"net"
	"oTsl"
	"time"
)

func main() {
	config := oTsl.NewConfig()
	config.SetListenT("tcp")
	config.SetListenIp("0.0.0.0:9999")
	config.SetCrtPath("certs/CARoot.crt")
	config.SetKeyPath("certs/CARoot.key")

	command := oTsl.NewCommand()
	config.SetCommand(command)

	command.AddCommand("Hello", func(conn net.Conn) error {
		time.Sleep(10 * time.Second)
		conn.Write([]byte("Hello"))
		return nil
	})

	command.AddCommand("Hello2", func(conn net.Conn) error {
		conn.Write([]byte("Hello2"))
		return nil
	})

	tslServer := oTsl.NewTslServer()
	tslServer.SetConfig(config)
	err := tslServer.Start()
	if err != nil {
		fmt.Print(err)
	}
}
