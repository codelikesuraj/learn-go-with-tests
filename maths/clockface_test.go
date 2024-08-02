package maths

import (
	"bytes"
	"encoding/xml"
	"log"
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

func TestMinuteHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(0, 15, 0), Point{1, 0}},
		{simpleTime(0, 30, 0), Point{0, -1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := MinuteHandPoint(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Fatalf("Wanted %v Point, but got %v", c.point, got)
			}
		})
	}
}

func TestHourHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(3, 0, 0), Point{1, 0}},
		{simpleTime(6, 0, 0), Point{0, -1}},
		{simpleTime(9, 0, 0), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := HourHandPoint(c.time)
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
			point:       Point{X: 150, Y: 150 - SecondHandLength},
		},
		{
			description: "at 15 seconds",
			time:        simpleTime(0, 0, 15),
			point:       Point{X: 150 + SecondHandLength, Y: 150},
		},
		{
			description: "at 30 seconds",
			time:        simpleTime(0, 0, 30),
			point:       Point{X: 150, Y: 150 + SecondHandLength},
		},
		{
			description: "at 45 seconds",
			time:        simpleTime(0, 0, 45),
			point:       Point{X: 150 - SecondHandLength, Y: 150},
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

func TestMinuteHand(t *testing.T) {
	cases := []struct {
		description string
		time        time.Time
		point       Point
	}{
		{
			description: "at 0 minutes",
			time:        simpleTime(0, 0, 0),
			point:       Point{X: 150, Y: 150 - MinuteHandLength},
		},
		{
			description: "at 15 minutes",
			time:        simpleTime(0, 15, 0),
			point:       Point{X: 150 + MinuteHandLength, Y: 150},
		},
		{
			description: "at 30 minutes",
			time:        simpleTime(0, 30, 0),
			point:       Point{X: 150, Y: 150 + MinuteHandLength},
		},
		{
			description: "at 45 minutes",
			time:        simpleTime(0, 45, 0),
			point:       Point{X: 150 - MinuteHandLength, Y: 150},
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			got := MinuteHand(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Errorf("Got %v, wanted %v", got, c.point)
			}
		})
	}
}

func TestHourHand(t *testing.T) {
	cases := []struct {
		description string
		time        time.Time
		point       Point
	}{
		{
			description: "at 0 hours",
			time:        simpleTime(0, 0, 0),
			point:       Point{X: 150, Y: 150 - HourHandLength},
		},
		{
			description: "at 3 hours",
			time:        simpleTime(3, 0, 0),
			point:       Point{X: 150 + HourHandLength, Y: 150},
		},
		{
			description: "at 6 hours",
			time:        simpleTime(6, 0, 0),
			point:       Point{X: 150, Y: 150 + HourHandLength},
		},
		{
			description: "at 9 hours",
			time:        simpleTime(9, 0, 0),
			point:       Point{X: 150 - HourHandLength, Y: 150},
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			got := HourHand(c.time)
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
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 15), (math.Pi / 30) * 15},
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 45), (math.Pi / 30) * 45},
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

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 15), (math.Pi / (30 * 60)) * 15},
		{simpleTime(0, 0, 30), math.Pi / 60},
		{simpleTime(0, 0, 45), (math.Pi / (30 * 60)) * 45},

		{simpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},

		{simpleTime(0, 15, 0), (math.Pi / 30) * 15},
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 45, 0), (math.Pi / 30) * 45},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := MinutesInRadians(c.time)
			if !roughlyEqualFloat64(got, c.angle) {
				t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
			}
		})
	}
}

func TestHoursInRadians(t *testing.T) {
	log.Println(time.Now().Hour())
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 0), 0},
		{simpleTime(3, 0, 0), math.Pi * 0.5},
		{simpleTime(6, 0, 0), math.Pi},
		{simpleTime(9, 0, 0), math.Pi * 1.5},
		{simpleTime(12, 0, 0), 0},
		{simpleTime(15, 0, 0), math.Pi * 0.5},
		{simpleTime(18, 0, 0), math.Pi},
		{simpleTime(21, 0, 0), math.Pi * 1.5},
		{simpleTime(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := HoursInRadians(c.time)
			if !roughlyEqualFloat64(got, c.angle) {
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
		{
			simpleTime(0, 0, 30),
			Line{150, 150, 150, 240},
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

func TestSVGWriterMinuteHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			simpleTime(0, 0, 0),
			Line{150, 150, 150, 70},
		},
	}

	for _, c := range cases {
		b := bytes.Buffer{}
		SVGWriter(&b, c.time)

		svg := SVG{}
		xml.Unmarshal(b.Bytes(), &svg)

		if !containsLine(c.line, svg.Line) {
			t.Errorf("expected to find the minute hand line %v in the SVG lines %v", c.line, svg.Line)
		}
	}
}

func TestSVGWriterHourHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			simpleTime(6, 0, 0),
			Line{150, 150, 150, 200},
		},
	}

	for _, c := range cases {
		b := bytes.Buffer{}
		SVGWriter(&b, c.time)

		svg := SVG{}
		xml.Unmarshal(b.Bytes(), &svg)

		if !containsLine(c.line, svg.Line) {
			t.Errorf("expected to find the hour hand line %v in the SVG lines %v", c.line, svg.Line)
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
