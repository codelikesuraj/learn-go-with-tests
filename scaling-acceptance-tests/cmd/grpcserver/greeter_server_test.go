package main_test

import (
	"fmt"
	"testing"

	"github.com/codelikesuraj/learn-go-with-tests/scaling-acceptance-tests/adapters"
	"github.com/codelikesuraj/learn-go-with-tests/scaling-acceptance-tests/adapters/grpcserver"
	"github.com/codelikesuraj/learn-go-with-tests/scaling-acceptance-tests/specifications"
)

func TestGreeterServer(t *testing.T) {
	host, port := adapters.StartDockerServer(t, "50051", "./cmd/grpcserver/Dockerfile")
	driver := grpcserver.Driver{
		Addr: fmt.Sprintf("%s:%s", host, port),
	}

	specifications.GreetSpecification(t, &driver)
}
