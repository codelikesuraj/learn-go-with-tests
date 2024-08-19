package main_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/codelikesuraj/learn-go-with-tests/scaling-acceptance-tests/adapters"
	"github.com/codelikesuraj/learn-go-with-tests/scaling-acceptance-tests/adapters/httpserver"
	"github.com/codelikesuraj/learn-go-with-tests/scaling-acceptance-tests/specifications"
)

func TestGreeterServer(t *testing.T) {
	host, port := adapters.StartDockerServer(t, "8080", "./cmd/httpserver/Dockerfile")
	driver := httpserver.Driver{
		BaseURL: fmt.Sprintf("http://%s:%s", host, port),
		Client:  &http.Client{Timeout: 1 * time.Second},
	}

	specifications.GreetSpecification(t, driver)
}
