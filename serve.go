/*
Serve - Simple HTTP File Server

Usage: serve [port]
    port is optional, default 9000


__version__:   1.0.0
__license__:   MIT
__author__:    Ye Liu
__contact__:   yeliu@instast.com
__copyright__: Copyright (c) 2014 Ye Liu
*/

package main

import (
    "log"
    "net"
    "net/http"
    "os"
)

const (
    HOST = "localhost"
    DEFAULT_PORT = "9000"
)

func main() {
    port := DEFAULT_PORT
    if (len(os.Args) > 1) {
        port = os.Args[1]
    }
    log.Printf("Starting server on %s port %s ...\n", HOST, port)
    log.Fatal(http.ListenAndServe(net.JoinHostPort(HOST, port), http.FileServer(http.Dir("."))))
}
