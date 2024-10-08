package interactions

import (
	"testing"

	"github.com/codelikesuraj/learn-go-with-tests/scaling-acceptance-tests/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(Greet),
	)
}
