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

// SecondHand is the unit vector of the second hand of a analogue clock at the time `t`
// represented as a Point.
func SecondHand(t time.Time) Point {
	return Point{150, 60}
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
	return Point{}
}
