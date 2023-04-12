package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
)

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("Handle request %s", request.RequestURI)
		writer.WriteHeader(200)
		writer.Write([]byte("Hello, World"))
	})
	mux.HandleFunc("/toomanyheaders", func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("Handle request %s", request.RequestURI)
		for i := 0; i < 100; i++ {
			writer.Header().Add(fmt.Sprintf("header_%d", i), strconv.Itoa(i))
		}
		writer.WriteHeader(200)
		writer.Write([]byte("Hello, World"))
	})
	mux.HandleFunc("/toomanyheaderbytes", func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("Handle request %s", request.RequestURI)
		s := ""
		for i := 0; i < 100*1024; i++ {
			s = fmt.Sprintf("%s%d", s, i%10)
		}
		writer.Header().Add("header", s)
		writer.WriteHeader(200)
		writer.Write([]byte("Hello, World"))
	})
	if listener, err := net.Listen("tcp", "0.0.0.0:8888"); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Listening on %s", listener.Addr())
		http.Serve(listener, nil)
	}
}
