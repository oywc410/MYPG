package main

import (
	"flag"
	"net/http"
	"log"
)

func main() {
	var addr = flag.String("addr", ":9001", "WEB ADDRS")
	flag.Parse()
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("public"))))
	log.Println("WEB SITE:", *addr)
	http.ListenAndServe(*addr, mux)
}
