package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	cert2_b, _ := ioutil.ReadFile("../ca/CARoot.crt")

	cert := tls.Certificate{
		Certificate: [][]byte{cert2_b},
	}

	config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
	conn, err := tls.Dial("tcp", "127.0.0.1:7777", &config)
	if err != nil {
		log.Fatalf("client: dial: %s", err)
	}
	defer conn.Close()
	log.Println("client: connected to: ", conn.RemoteAddr())

	state := conn.ConnectionState()
	for _, v := range state.PeerCertificates {
		fmt.Println(x509.MarshalPKIXPublicKey(v.PublicKey))
		fmt.Println(v.Subject)
	}

	reply := make([]byte, 256)
	n, err := conn.Read(reply)
	log.Printf("====== client: read %q size :%d bytes)", string(reply[:n]), n)

}
