package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/ajstarks/svgo"
)

type Point struct {
	x float64
	y float64
}

func points_length(points []Point) float64 {
	length := 0.
	if len(points) < 1 {
		return length
	}
	current_point := points[0]
	for idx, next_point := range points {
		if idx == 0 {
			continue
		}
		length += math.Hypot(next_point.x-current_point.x, next_point.y-current_point.y)
		current_point = next_point
	}

	fmt.Println("LENGTH YO!", length)
	return length
}

func squiggle(radius float64, length float64) []Point {
	points := make([]Point, 0)
	current_length := 0.
	current_angle := 0.

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	points = append(points, Point{radius, 0})
	for current_length < length {
		new_x := radius * math.Cos(current_angle)
		new_y := radius * math.Sin(current_angle)
		points = append(points, Point{new_x, new_y})
		current_length = points_length(points)
		angle_mod := r1.Intn(7) + 1
		fmt.Println("RAND ANGLE MOD:", angle_mod)
		current_angle += float64(angle_mod)
		radius_mod := r1.Intn(7) - 3
		fmt.Println("RAND RADIUS MOD:", radius_mod)
		radius += float64(radius_mod)
	}

	return points
}

func main() {
	width := 500
	height := 500
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Circle(width/2, height/2, 100)
	canvas.Text(width/2, height/2, "Hello, SVG", "text-anchor:middle;font-size:30px;fill:white")
	canvas.End()
	squiggle(10, 100)
}
