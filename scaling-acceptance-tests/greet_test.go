package scaling_acceptance_tests

import (
	"testing"

	"github.com/codelikesuraj/learn-go-with-tests/scaling-acceptance-tests/specifications"
)

func TestGree(t *testing.T) {
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(Greet),
	)
}
