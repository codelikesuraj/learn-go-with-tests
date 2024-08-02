package maths

import (
	"encoding/xml"
	"fmt"
	"io"
	"math"
	"time"
)

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

type Point struct {
	X, Y float64
}

const (
	CenterX          = 150
	CenterY          = 150
	SecondHandLength = 90
	MinuteHandLength = 80
	HourHandLength   = 50

	SecondsInHalfClock = 30
	SecondsInClock     = 2 * SecondsInHalfClock
	MinutesInHalfClock = 30
	MinutesInClock     = 2 * MinutesInHalfClock
	HoursInHalfClock   = 6
	HoursInClock       = 2 * HoursInHalfClock

	svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`
	svgBezel = `<!-- bezel -->
  <circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`
	svgEnd = `</svg>`
)

func SecondHand(tm time.Time) Point {
	return makeHand(SecondHandPoint(tm), SecondHandLength)
}

func MinuteHand(tm time.Time) Point {
	return makeHand(MinuteHandPoint(tm), MinuteHandLength)
}

func HourHand(tm time.Time) Point {
	return makeHand(HourHandPoint(tm), HourHandLength)
}

func makeHand(p Point, length float64) Point {
	p = Point{p.X * length, p.Y * length}
	p = Point{p.X, -p.Y}
	p = Point{p.X + CenterX, p.Y + CenterY}
	return p
}

func SecondHandSVG(w io.Writer, p Point) {
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func MinuteHandSVG(w io.Writer, p Point) {
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

func HourHandSVG(w io.Writer, p Point) {
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

func SecondHandPoint(t time.Time) Point {
	return angleToPoint(SecondsInRadians(t))
}

func MinuteHandPoint(t time.Time) Point {
	return angleToPoint(MinutesInRadians(t))
}

func HourHandPoint(t time.Time) Point {
	return angleToPoint(HoursInRadians(t))
}

func angleToPoint(angle float64) Point {
	return Point{
		X: math.Sin(angle),
		Y: math.Cos(angle),
	}
}

func SecondsInRadians(t time.Time) float64 {
	return math.Pi * (float64(t.Second()) / 30)
}

func MinutesInRadians(t time.Time) float64 {
	return (SecondsInRadians(t) / 60) + (math.Pi * (float64(t.Minute()) / 30))
}

func HoursInRadians(t time.Time) float64 {
	// return (MinutesInRadians(t) / 12) + (math.Pi / (6 / (float64(t.Hour() % 12))))
	// return (MinutesInRadians(t) / 12) + (math.Pi * ((float64(t.Hour() % 12) / 6)))
	return (MinutesInRadians(t) / 12) + (math.Pi * ((float64(t.Hour() % 12) / 6)))
	// return math.Pi * (float64(t.Hour()) / 6)
}

func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, svgBezel)
	SecondHandSVG(w, SecondHand(t))
	MinuteHandSVG(w, MinuteHand(t))
	HourHandSVG(w, HourHand(t))
	io.WriteString(w, svgEnd)
}
