package main

import (
	"log"
	"github.com/sshindesiddesh/distributed-go/examples/prolog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}