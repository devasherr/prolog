package main

import (
	"log"

	"github.com/devasherr/prolog/internal/server"
)

func main() {
	srv := server.NewHttpServer(":7777")
	log.Fatal(srv.ListenAndServe())
}
