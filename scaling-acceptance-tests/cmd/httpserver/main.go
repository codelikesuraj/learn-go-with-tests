package main

import (
	"log"
	"net/http"

	"github.com/codelikesuraj/learn-go-with-tests/scaling-acceptance-tests/adapters/httpserver"
)

func main() {
	handler := http.HandlerFunc(httpserver.Handler)

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
