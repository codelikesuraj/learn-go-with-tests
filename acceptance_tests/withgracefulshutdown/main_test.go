package main

import (
	"testing"
	"time"

	"github.com/codelikesuraj/learn-go-with-tests/acceptance_tests"
	"github.com/quii/go-graceful-shutdown/assert"
)

const (
	port = "8080"
	url  = "http://localhost:" + port
)

func TestGracefulShutdown(t *testing.T) {
	cleanup, sendInterrupt, err := acceptance_tests.LaunchTestProgram(port)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(cleanup)

	// just check the serve works before we shut things down
	assert.CanGet(t, url)

	// fire off a reques, and before it has a chance to respond, send SIGTERM.
	time.AfterFunc(50*time.Millisecond, func() {
		assert.NoError(t, sendInterrupt())
	})

	// without graceful shutdown, this would fail
	assert.CanGet(t, url)

	// after interrupt, the server should be shutdown, and no more requests will work
	assert.CantGet(t, url)
}
