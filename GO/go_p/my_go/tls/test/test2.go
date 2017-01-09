kage main
import (
        "crypto/tls"
        "fmt"
        "log"
)
func main() {
        cert, err := tls.LoadX509KeyPair("newcert.pem.crt", "key.pem.unencrypted")
        if err != nil {
                log.Fatalf("server: loadkeys: %s", err)
        }
        config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
        conn, err := tls.Dial("tcp", "ngote.cnnic.cn:3121", &config)
        if err != nil {
                log.Fatalf("client: dial: %s", err)
        }
        defer conn.Close()
        log.Println("client: connected to: ", conn.RemoteAddr())
        state := conn.ConnectionState()
        for _, v := range state.PeerCertificates {
                fmt.Println(v.Subject)
        }
        log.Println("client: handshake: ", state.HandshakeComplete)
        log.Println("client: mutual: ", state.NegotiatedProtocolIsMutual)
}