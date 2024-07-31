package maths

import (
	"bytes"
	"encoding/xml"
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

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := SecondHandPoint(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Fatalf("Wanted %v Point, but got %v", c.point, got)
			}
		})
	}
}

func TestSecondHand(t *testing.T) {
	cases := []struct {
		description string
		time        time.Time
		point       Point
	}{
		{
			description: "at 0 seconds",
			time:        simpleTime(0, 0, 0),
			point:       Point{X: 150, Y: 150 - SecondsHandLength},
		},
		{
			description: "at 15 seconds",
			time:        simpleTime(0, 0, 15),
			point:       Point{X: 150 + SecondsHandLength, Y: 150},
		},
		{
			description: "at 30 seconds",
			time:        simpleTime(0, 0, 30),
			point:       Point{X: 150, Y: 150 + SecondsHandLength},
		},
		{
			description: "at 45 seconds",
			time:        simpleTime(0, 0, 45),
			point:       Point{X: 150 - SecondsHandLength, Y: 150},
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			got := SecondHand(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Errorf("Got %v, wanted %v", got, c.point)
			}
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
		t.Run(testName(c.time), func(t *testing.T) {
			got := SecondsInRadians(c.time)
			if got != c.angle {
				t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
			}
		})
	}
}

func TestSVGWriterSecondHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			simpleTime(0, 0, 0),
			Line{150, 150, 150, 60},
		},
	}

	for _, c := range cases {
		b := bytes.Buffer{}
		SVGWriter(&b, c.time)

		svg := SVG{}
		xml.Unmarshal(b.Bytes(), &svg)

		if !containsLine(c.line, svg.Line) {
			t.Errorf("expected to find the second hand line %v in the SVG lines %v", c.line, svg.Line)
		}
	}
}

func containsLine(line Line, lines []Line) bool {
	for _, l := range lines {
		if l == line {
			return true
		}
	}
	return false
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(1337, time.January, 1, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
