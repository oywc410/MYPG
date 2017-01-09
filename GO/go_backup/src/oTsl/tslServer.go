package oTsl

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net"
)

type TslServer struct {
	config     *tls.Config
	oTslConfig *oTslConfig
}

func NewTslServer() *TslServer {
	return &TslServer{}
}

func (obj *TslServer) SetConfig(config *oTslConfig) error {
	tlsConfig, err := config.ToTslConfig()
	if err != nil {
		log.Printf("server: loadKeys: %s", err)
		return err
	}

	obj.config = tlsConfig
	obj.oTslConfig = config

	return nil
}

func (obj *TslServer) Start() error {
	listener, err := tls.Listen(obj.oTslConfig.GetListenT(), obj.oTslConfig.GetListenIp(), obj.config)
	if err != nil {
		log.Printf("server: listen: %s", err)
		return err
	}

	log.Print("server: listening")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("server: accept: %s", err)
			break
		}
		defer conn.Close()
		log.Printf("server: accepted from %s", conn.RemoteAddr())
		//验证接口
		tlscon, ok := conn.(*tls.Conn)
		if ok {
			log.Print("ok=true")
			state := tlscon.ConnectionState()
			for _, v := range state.PeerCertificates {
				log.Print(x509.MarshalPKIXPublicKey(v.PublicKey))
			}

			go obj.handleClient(conn)
		}
	}

	return nil
}

func (obj *TslServer) handleClient(conn net.Conn) {
	defer conn.Close()
	remoteIp := conn.RemoteAddr().String()

	log.Printf("server: conn: waiting | remote %s.", remoteIp)

	for {
		//接收信息
		err := obj.receiveMess(conn, remoteIp)
		if err != nil {
			break
		}
		//发送信息
	}
}

func (obj *TslServer) receiveMess(conn net.Conn, remoteIp string) error {
	buf := make([]byte, 1024)
	//接收命令
	n, err := conn.Read(buf)
	if err != nil {
		log.Printf("server: conn: read: %s | remote %s.", err, remoteIp)
		return err
	}

	if n > 0 {
		mess := string(buf[:n])
		log.Printf("server: conn: echo %s | remote %s.", mess, remoteIp)

		fun, ok := obj.oTslConfig.GetCommand().GetCommand(mess)

		if ok {
			err := fun(conn)
			if err != nil {
				log.Printf("server: conn: func: %s | remote %s.", err, remoteIp)
				return err
			}
		}

	}

	return nil
}
