package main

import (
	"log"
	"net/http"

	scaling_acceptance_tests "github.com/codelikesuraj/learn-go-with-tests/scaling-acceptance-tests"
)

func main() {
	handler := http.HandlerFunc(scaling_acceptance_tests.Handler)

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
