package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func requestHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Server", "BuchonIP/0.1")

	semicolon_index := strings.LastIndex(req.RemoteAddr, ":")
	remote_ip := req.RemoteAddr[:semicolon_index]

	log.Printf("Incoming request from %s", remote_ip)
	fmt.Fprintf(res, remote_ip)
}

func main() {
	host := flag.String("host", "127.0.0.1", "Host")
	port := flag.Int("port", 8080, "Port")
	flag.Parse()

	address := fmt.Sprintf("%s:%d", *host, *port)
	log.Printf("Listening at http://%s\n", address)

	http.HandleFunc("/", requestHandler)
	log.Fatal(http.ListenAndServe(address, nil))
}
