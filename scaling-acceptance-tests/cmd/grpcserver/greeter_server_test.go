package main_test

import (
	"testing"

	"github.com/codelikesuraj/learn-go-with-tests/scaling-acceptance-tests/adapters"
	"github.com/codelikesuraj/learn-go-with-tests/scaling-acceptance-tests/adapters/grpcserver"
	"github.com/codelikesuraj/learn-go-with-tests/scaling-acceptance-tests/specifications"
)

func TestGreeterServer(t *testing.T) {
	baseURL := adapters.StartDockerServer(t, "8080", "./cmd/grpcserver/Dockerfile")
	driver := grpcserver.Driver{
		Addr: baseURL,
	}

	specifications.GreetSpecification(t, driver)
}
