package main

import (
	"github.com/yuekcc/fns/webapi"
	"log"
	"net/http"
)

const (
	ADDR = "0.0.0.0:10089"
)

func main() {
	router := webapi.Router()

	server := &http.Server{
		Handler: router,
		Addr:    ADDR,
	}

	log.Printf("Server host on: %s\n", ADDR)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
