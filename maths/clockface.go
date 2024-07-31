package maths

import (
	"math"
	"time"
)

type Point struct {
	X, Y float64
}

const (
	CenterX           = 150
	CenterY           = 150
	SecondsHandLength = 90
)

func SecondHand(tm time.Time) Point {
	p := SecondHandPoint(tm)
	p = Point{p.X * SecondsHandLength, p.Y * SecondsHandLength}
	p = Point{p.X, -p.Y}
	p = Point{p.X + CenterX, p.Y + CenterY}
	return p
}

func SecondHandPoint(t time.Time) Point {
	angle := SecondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

func SecondsInRadians(t time.Time) float64 {
	return math.Pi / (30 / float64(t.Second()))
}
