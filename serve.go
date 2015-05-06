package main

// Serve - Simple HTTP File Server
//
// Usage: serve [port]
//     port is optional, default 9000
//
//
// Version:   1.0
// License:   MIT
// Author:    Ye Liu
// Contact:   yeliu@instast.com
// Copyright: Copyright (c) 2015 Ye Liu

import (
	"log"
	"net"
	"net/http"
	"os"
)

const (
	HOST         = "localhost"
	DEFAULT_PORT = "9000"
)

func main() {
	port := DEFAULT_PORT
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	log.Printf("Starting server on %s port %s ...\n", HOST, port)
	log.Fatal(http.ListenAndServe(net.JoinHostPort(HOST, port), http.FileServer(http.Dir("."))))
}
