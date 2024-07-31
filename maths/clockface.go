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
	CenterX           = 150
	CenterY           = 150
	SecondsHandLength = 90

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
	p := SecondHandPoint(tm)
	p = Point{p.X * SecondsHandLength, p.Y * SecondsHandLength}
	p = Point{p.X, -p.Y}
	p = Point{p.X + CenterX, p.Y + CenterY}
	return p
}

func SecondHandSVG(w io.Writer, p Point) {
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
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

func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, svgBezel)
	SecondHandSVG(w, SecondHand(t))
	io.WriteString(w, svgEnd)
}
