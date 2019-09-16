//
// github.com/egdx/server1
//

// A basic HTTP server.
// By default, it serves the current working directory on port 8080.
//
package main

import (
	"flag"
	"log"
	"net/http"
)

// server1 -dir ../
// server1 -dir=../
// server1 -listen=3000

var (
	listen = flag.String("listen", ":8080", "listen address")
	// ~/dir1/dir2/dir3/project/ -- does not work!
	dir = flag.String("dir", ".", "directory to serve")
)

func main() {
	flag.Parse()
	log.Printf("listening on %q\n", *listen)
	log.Printf("dir on %q\n", *dir)

	err := http.ListenAndServe(*listen, http.FileServer(http.Dir(*dir)))
	log.Fatalln(err)
}
