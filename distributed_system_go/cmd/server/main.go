package main

import (
	"distributed_system_go/internal/server"
	"log"
)

func main() {
	srv := server.NewHttpServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
