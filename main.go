package main

import (
	"os"
	"time"

	"github.com/codelikesuraj/learn-go-with-tests/mocking"
)

func main() {
	mocking.Countdown(os.Stdout, &mocking.ConfigurableSleeper{Duration: 5 * time.Second, Sleep_: time.Sleep})
}
