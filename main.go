//
// github.com/egdx/server1
//

// A basic HTTP server.
// By default, it serves the current working directory on port 8080.
//
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

// usage:
// server1 -dir ./
// server1 -dir=./
// server1 -port=3000 -dir=../dir1/mywebapp

var _appVersion = "v1.0.3"

var (
	port = flag.String("port", "8080", "port number")

	// ~/dir1/dir2/dir3/project/ -- does not work!
	dir = flag.String("dir", ".", "directory folder to serve")
)

func main() {
	fmt.Printf("server1 %s\n\n", _appVersion)
	flag.Parse()
	fmt.Printf("  Simple HTTP server for quick testing\n\n  optional flags: -dir -port\n\n  example: server1 -dir=../dir1/webapp -port=8042\n\n")
	fmt.Printf("  http://localhost:%s\n\n", *port)

	log.Printf("listening on \":%s\"\n", *port)
	log.Printf("dir on %q\n", *dir)

	err := http.ListenAndServe(":"+(*port), http.FileServer(http.Dir(*dir)))
	log.Fatalln(err)
}
