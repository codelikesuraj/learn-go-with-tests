package maths

import (
	"math"
	"time"
)

const (
	CenterX           = 150
	CenterY           = 150
	SecondsHandLength = 90
)

type Point struct {
	X, Y float64
}

func SecondHand(tm time.Time) Point {
	p := SecondHandPoint(tm)

	// scale
	p = Point{p.X * SecondsHandLength, p.Y * SecondsHandLength}

	// // flip (SVG has origin at top left)
	p = Point{p.X, -p.Y}

	// // translate
	p = Point{p.X + CenterX, p.Y + CenterY}

	return p
}

func SecondHandPoint(tm time.Time) Point {
	angle := SecondsInRadians(tm)
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{math.Trunc(x), math.Trunc(y)}
}

func SecondsInRadians(t time.Time) float64 {
	return math.Pi * (float64(t.Second()) / 30)
}
