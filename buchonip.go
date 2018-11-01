package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func parse_ip(remoteAddr string) string {
	semicolonIndex := strings.LastIndex(remoteAddr, ":")
	return remoteAddr[:semicolonIndex]
}

func homeHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Server", "BuchonIP/0.1")
	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	remote_ip := parse_ip(req.RemoteAddr)

	log.Printf("Incoming request from %s", remote_ip)
	fmt.Fprintf(res, `<p>Su IP es <b>%s</b></p>
Disponible en <a href="/txt">txt</a> y <a href="/json">json</a>`, remote_ip)
}

func jsonHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Server", "BuchonIP/0.1")
	res.Header().Set("Content-Type", "application/json")

	remote_ip := parse_ip(req.RemoteAddr)

	log.Printf("Incoming request from %s", remote_ip)
	fmt.Fprintf(res, "{\"ip\": \"%s\"}", remote_ip)
}

func txtHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Server", "BuchonIP/0.1")
	res.Header().Set("Content-Type", "text/plain")

	remote_ip := parse_ip(req.RemoteAddr)

	log.Printf("Incoming request from %s", remote_ip)
	fmt.Fprintf(res, remote_ip)
}

func main() {
	host := flag.String("host", "127.0.0.1", "Host")
	port := flag.Int("port", 8080, "Port")
	flag.Parse()

	address := fmt.Sprintf("%s:%d", *host, *port)
	log.Printf("Listening at http://%s\n", address)

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/txt", txtHandler)
	log.Fatal(http.ListenAndServe(address, nil))
}
