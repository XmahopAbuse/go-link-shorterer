package main

import (
	"go-link-shorterer/app/apiserver"
	"log"
)

func main() {
	server := apiserver.New(":80")
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}


