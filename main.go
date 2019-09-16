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
// server1 -dir ../
// server1 -dir=../
// server1 -listen=3000

var _appVersion = "v1.0.1"

var (
	listen = flag.String("listen", "8080", "listen address")

	// ~/dir1/dir2/dir3/project/ -- does not work!
	dir = flag.String("dir", ".", "directory to serve")
)

func main() {
	fmt.Printf("server1 %s\n\n", _appVersion)
	flag.Parse()
	fmt.Printf("  Simple HTTP server for quick testing\n\n")
	fmt.Printf("  http://localhost:%s\n\n", *listen)

	log.Printf("listening on \":%s\"\n", *listen)
	log.Printf("dir on %q\n", *dir)

	err := http.ListenAndServe(":"+(*listen), http.FileServer(http.Dir(*dir)))
	log.Fatalln(err)
}
