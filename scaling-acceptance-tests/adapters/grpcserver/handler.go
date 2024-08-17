package grpcserver

import (
	"fmt"
	"net/http"

	"github.com/codelikesuraj/learn-go-with-tests/scaling-acceptance-tests/domain/interactions"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, interactions.Greet(name))
}
