package mocking

import (
	"fmt"
	"io"
	"time"
)

const sleep = "sleep"
const write = "write"

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type DefaultSleeper struct{}

type SpyCountdownOperations struct {
	Calls []string
}

type ConfigurableSleeper struct {
	Duration time.Duration
	Sleep_   func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.Sleep_(c.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

const (
	countdownStart int    = 3
	finalWord      string = "Go!"
)

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprintf(out, "%s", finalWord)
}
