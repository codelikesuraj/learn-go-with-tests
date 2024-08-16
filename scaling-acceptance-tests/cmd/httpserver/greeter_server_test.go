package main_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/alecthomas/assert/v2"
	"github.com/codelikesuraj/learn-go-with-tests/scaling-acceptance-tests/specifications"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestGreeterServer(t *testing.T) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    "../../.",
			Dockerfile: "./cmd/httpserver/Dockerfile",
			KeepImage: true,
		},
		ExposedPorts: []string{"8080/tcp"},
		WaitingFor:   wait.ForHTTP("/"),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	assert.NoError(t, err)

	ip, err := container.Host(ctx)
	assert.NoError(t, err)

	mappedPort, err := container.MappedPort(ctx, "8080")
	assert.NoError(t, err)

	t.Cleanup(func() {
		assert.NoError(t, container.Terminate(ctx))
	})

	driver := specifications.Driver{
		BaseURL: fmt.Sprintf("http://%s:%s", ip, mappedPort.Port()),
		Client:  &http.Client{Timeout: 1 * time.Second},
	}
	specifications.GreetSpecification(t, driver)
}
