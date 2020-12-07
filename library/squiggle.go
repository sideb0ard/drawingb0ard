package library

import (
	"math"
	"math/rand"
	"time"
)

var (
	s1 = rand.NewSource(time.Now().UnixNano())
	r1 = rand.New(s1)
)

func Squiggle(radius float64, length float64) []Point {
	points := make([]Point, 0)
	current_length := 0.
	current_angle := 0.

	for current_length < length {
		new_x := radius * math.Cos(current_angle)
		new_y := radius * math.Sin(current_angle)
		points = append(points, Point{new_x, new_y})
		current_length = PointsLength(points)
		angle_mod := r1.Intn(7) + 1
		current_angle += float64(angle_mod)
		radius_mod := r1.Intn(7) - 3
		radius += float64(radius_mod)
	}

	return points
}
