package oTsl

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
)

type Message struct {
	by     []byte
	refunc func(net.Conn, error) error
}

type TslClient struct {
	config     *tls.Config
	oTslConfig *oTslConfig
	message    chan (Message)
}

func NewClient() *TslClient {
	return &TslClient{message: make(chan (Message), 10)}
}

func (obj *TslClient) SetConfig(config *oTslConfig) error {
	tlsConfig, err := config.ToClientTslConfig()
	if err != nil {
		log.Printf("server: loadKeys: %s", err)
		return err
	}
	obj.config = tlsConfig
	obj.oTslConfig = config

	return nil
}

func (obj *TslClient) Dial() error {

	conn, err := tls.Dial(obj.oTslConfig.GetListenT(), obj.oTslConfig.GetListenIp(), obj.config)

	if err != nil {
		return err
	}

	//连接错误发生时
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			conn.Close()
			os.Exit(1)
		}
	}()

	log.Println("client: connected to: ", conn.RemoteAddr())

	state := conn.ConnectionState()

	for _, v := range state.PeerCertificates {
		fmt.Println(x509.MarshalPKIXPublicKey(v.PublicKey))
		fmt.Println(v.Subject)
	}
	log.Println("client: handshake: ", state.HandshakeComplete)
	log.Println("client: mutual: ", state.NegotiatedProtocolIsMutual)

	if !state.HandshakeComplete || !state.NegotiatedProtocolIsMutual {
		return errors.New("no tsl")
	}

	for {
		select {
		case message := <-obj.message: //发送命令
			_, err = conn.Write(message.by)
			if err != nil {
				log.Println("client: send message: ", string(message.by))
			}

			err = message.refunc(conn, err)
			if err != nil {
				log.Println("client: send message re func error: ", err)
			}

		default: //接收命令
			/*
				reply := make([]byte, 256)
				n, err := conn.Read(reply)

				if err != nil {
					return err
				}

				log.Printf("client: read %q (%d bytes)", string(reply[:n]), n)
			*/
		}
		/*
			message := "Hello"
			n, err := io.WriteString(conn, message)
			if err != nil {
				log.Printf("client: write: %s", err)
				break
			}
			log.Printf("client: wrote %q (%d bytes)", message, n)

			reply := make([]byte, 256)
			n, err = conn.Read(reply)

			if err != nil {
				break
			}

			log.Printf("client: read %q (%d bytes)", string(reply[:n]), n)
		*/
	}

	return nil
}

func (obj *TslClient) SendMess(by []byte, fun func(net.Conn, error) error) {
	obj.message <- Message{by: by, refunc: fun}
}
