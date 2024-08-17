package main_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/codelikesuraj/learn-go-with-tests/scaling-acceptance-tests/adapters"
	"github.com/codelikesuraj/learn-go-with-tests/scaling-acceptance-tests/adapters/httpserver"
	"github.com/codelikesuraj/learn-go-with-tests/scaling-acceptance-tests/specifications"
)

func TestGreeterServer(t *testing.T) {
	baseURL := adapters.StartDockerServer(t, "8080", "./cmd/httpserver/Dockerfile")
	driver := httpserver.Driver{
		BaseURL: baseURL,
		Client:  &http.Client{Timeout: 1 * time.Second},
	}

	specifications.GreetSpecification(t, driver)
}
