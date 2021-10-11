package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yuekcc/fns/web"
	"github.com/yuekcc/fns/webapi"
)

var (
	hostFlag string
)

func init() {
	flag.StringVar(&hostFlag, "addr", "0.0.0.0:10086", "set web server host")
}

func main() {
	flag.Parse()

	router := mux.NewRouter()
	router.PathPrefix("/api").Handler(webapi.Router())
	router.PathPrefix("/").Handler(http.FileServer(http.FS(web.Assets)))

	server := &http.Server{
		Handler: router,
		Addr:    hostFlag,
	}

	log.Printf("Server host on: %s\n", hostFlag)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
