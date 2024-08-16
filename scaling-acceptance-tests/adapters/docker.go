package adapters

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/alecthomas/assert/v2"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func StartDockerServer(t testing.TB, exposedPort, dockerFilePath string) (url string) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:       "../../.",
			Dockerfile:    dockerFilePath,
			PrintBuildLog: true,
		},
		ExposedPorts: []string{exposedPort},
		WaitingFor:   wait.ForListeningPort(nat.Port(exposedPort)).WithStartupTimeout(5 * time.Second),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	assert.NoError(t, err)

	ip, err := container.Host(ctx)
	assert.NoError(t, err)

	port, err := container.MappedPort(ctx, nat.Port(exposedPort))
	assert.NoError(t, err)

	url = fmt.Sprintf("http://%s:%s", ip, port.Port())

	t.Cleanup(func() {
		assert.NoError(t, container.Terminate(ctx))
	})

	return
}
