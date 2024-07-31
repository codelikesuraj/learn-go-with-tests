package maths

import (
	"math"
	"testing"
	"time"
)

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(0, 0, 15), Point{1, 0}},
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			assertEqual(t, SecondHandPoint(test.time), test.point)
		})
	}
}

func TestSecondHand(t *testing.T) {
	cases := []struct {
		description string
		seconds     time.Time
		want        Point
	}{
		{
			description: "at 0 seconds",
			seconds:     simpleTime(0, 0, 0),
			want:        Point{X: 150, Y: 150 - SecondsHandLength},
		},
		{
			description: "at 15 seconds",
			seconds:     simpleTime(0, 0, 15),
			want:        Point{X: 150 + SecondsHandLength, Y: 150},
		},
		{
			description: "at 30 seconds",
			seconds:     simpleTime(0, 0, 30),
			want:        Point{X: 150, Y: 150 + SecondsHandLength},
		},
		{
			description: "at 45 seconds",
			seconds:     simpleTime(0, 0, 45),
			want:        Point{X: 150 - SecondsHandLength, Y: 150},
		},
	}

	for _, test := range cases {
		t.Run(test.description, func(t *testing.T) {
			assertEqual(t, SecondHand(test.seconds), test.want)
		})
	}
}

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}

	for _, c := range cases {
		assertEqual(t, SecondsInRadians(c.time), c.angle)
	}
}

func assertEqual(t testing.TB, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(1337, time.January, 1, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
