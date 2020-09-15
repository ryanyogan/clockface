package clockface

import (
	"math"
	"time"
)

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

// Point represents a two-dimensional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

// SecondsInRadians returns the angle of the second hand from 12 in radians
func SecondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

// MinutesInRadians returns the angle of the minute hand from 12 in radians
func MinutesInRadians(t time.Time) float64 {
	return (SecondsInRadians(t) / 60) +
		(math.Pi / (30 / float64(t.Minute())))
}

// HoursInRadians returns the angle of the hour hand from 12 in radians
func HoursInRadians(t time.Time) float64 {
	return (MinutesInRadians(t) / 12) +
		(math.Pi / (6 / float64(t.Hour()%12)))
}

// SecondHandPoint is the unit vector of the second hand at time T as a pont
func SecondHandPoint(t time.Time) Point {
	return angleToPoint(SecondsInRadians(t))
}

// MinuteHandPoint is the unit vector of the minute hand at time T as a pont
func MinuteHandPoint(t time.Time) Point {
	return angleToPoint(MinutesInRadians(t))
}

// HourHandPoint is the unit vector of the hour hand at time T as a pont
func HourHandPoint(t time.Time) Point {
	return angleToPoint(HoursInRadians(t))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
