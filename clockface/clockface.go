package clockface

import (
	"math"
	"time"
)

// A point represents a two dimensional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

const (
	secondHandLength = 90
	clockCenterX     = 150
	clockCenterY     = 150
)

// SecondHand is the unit vector of the second hand of a analogue clock at the time `t`
// represented as a Point.
func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength}
	p = Point{p.X, -p.Y}
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}

	return p
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

func zero() float64 {
	return 0.0
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
