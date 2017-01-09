package main

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"log"
	"net"
)

func main() {
	cert, err := tls.LoadX509KeyPair("../certs/CARoot.crt", "../certs/CARoot.key")
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	config.Rand = rand.Reader
	service := "0.0.0.0:9999"
	listener, err := tls.Listen("tcp", service, &config)
	if err != nil {
		log.Fatalf("server: listen: %s", err)
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
		tlscon, ok := conn.(*tls.Conn)
		if ok {
			log.Print("ok=true")
			state := tlscon.ConnectionState()
			for _, v := range state.PeerCertificates {
				log.Print(x509.MarshalPKIXPublicKey(v.PublicKey))
			}
		}
		handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 25966)

	log.Print("server: conn: waiting")
	n, err := conn.Read(buf)

	if err != nil {
		if err != nil {
			log.Printf("server: conn: read: %s", err)
		}
	}
	log.Printf("server: conn: echo %q\n", string(buf[:n]))

	n, err = conn.Read(buf)
	if err != nil {
		if err != nil {
			log.Printf("server: conn: read: %s", err)
		}
	}
	log.Printf("server: conn: echo %q\n", string(buf[:n]))
	n, err = conn.Read(buf)
	if err != nil {
		if err != nil {
			log.Printf("server: conn: read: %s", err)
		}
	}
	log.Printf("server: conn: echo %q\n", string(buf[:n]))
	n, err = conn.Read(buf)
	if err != nil {
		if err != nil {
			log.Printf("server: conn: read: %s", err)
		}
	}
	log.Printf("server: conn: echo %q\n", string(buf[:n]))
	n, err = conn.Read(buf)
	if err != nil {
		if err != nil {
			log.Printf("server: conn: read: %s", err)
		}
	}
	log.Printf("server: conn: echo %q\n", string(buf[:n]))
	n, err = conn.Read(buf)
	if err != nil {
		if err != nil {
			log.Printf("server: conn: read: %s", err)
		}
	}
	log.Printf("server: conn: echo %q\n", string(buf[:n]))
	conn.Write([]byte("aaaaaa\n"))

	log.Println("server: conn: closed")
}
