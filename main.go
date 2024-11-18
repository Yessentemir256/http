package main

import (
	"log"
	"net"
	"net/http"
	"os"
)

type handler struct{}

func main() {
	host := "0.0.0.0"
	port := "9999"

	if err := execute(host, port); err != nil {
		os.Exit(1)
	}
}

func execute(host string, port string) (err error) {
	srv := &http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: &handler{},
	}
	return srv.ListenAndServe()
}

// ServeHTTP
func (h *handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	_, err := writer.Write([]byte("Hello world"))
	if err != nil {
		log.Print(err)
	}
}
