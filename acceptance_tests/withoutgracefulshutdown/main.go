package main

import (
	"log"
	"net/http"

	"github.com/codelikesuraj/learn-go-with-tests/acceptance_tests"
)

func main() {
	server := &http.Server{Addr: ":8081", Handler: http.HandlerFunc(acceptance_tests.SlowHandler)}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
