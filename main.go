package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello Http Router")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	server.ListenAndServe()
}
