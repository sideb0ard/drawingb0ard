package main

import (
	"log"
	"os"

	svg "github.com/ajstarks/svgo"
	"github.com/sideb0ard/drawingb0ard/library"
)

var (
	width  = 700
	height = 700
	canvas = svg.New(os.Stdout)
)

func NewFile(filename string) *os.File {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func Map(pointArray []int, f func(int) int) []int {
	mapped_array := make([]int, len(pointArray))
	for i, v := range pointArray {
		mapped_array[i] = f(v)
	}
	return mapped_array
}

func DrawPolyLine(x int, y int, squig []library.Point, color string) {
	xpos, ypos := library.SplitPointsSlice(squig)
	tx := Map(xpos, func(val int) int { return val + x })
	ty := Map(ypos, func(val int) int { return val + y })
	canvas.Polyline(tx, ty, "fill:none; stroke:"+color)
}

func AdjustPoint(point library.Point, x int, y int) library.Point {
	point.X += float64(x)
	point.Y += float64(y)
	return point
}

func DrawCurvedLine(x int, y int, squig []library.Point, color string) {
	for i := 0; i < len(squig); i++ {

		start_point := AdjustPoint(squig[i], x, y)

		var previous library.Point
		var next library.Point
		// Start Control Point
		if i > 0 {
			previous = AdjustPoint(squig[i-1], x, y)
		}
		if i < len(squig)-1 {
			next = AdjustPoint(squig[i+1], x, y)
		}
		start_control := library.ControlPoint(start_point, previous, next, false)

		end_point := start_point // default

		var end_next library.Point
		if i < len(squig)-2 {
			end_point = AdjustPoint(squig[i+1], x, y)
			end_next = AdjustPoint(squig[i+2], x, y)
		} else if i < len(squig)-1 {
			end_point = AdjustPoint(squig[i+1], x, y)
		}
		end_control := library.ControlPoint(end_point, start_point, end_next, true)

		canvas.Bezier(int(start_point.X), int(start_point.Y), int(start_control.X), int(start_control.Y), int(end_control.X), int(end_control.Y), int(end_point.X), int(end_point.Y), "fill:none; stroke:"+color)
	}
}

func main() {

	colors := [...]string{"chartreuse", "greenyellow", "lawngreen", "fuchsia", "yellow", "black"}
	canvas.Start(width, height)
	canvas.RGB(0, 0, 0)

	squig := library.Squiggle(50, 500)
	third := width / 3
	DrawPolyLine(third, height/2, squig, colors[2])
	DrawCurvedLine(third*2, height/2, squig, colors[3])

	canvas.End()
}
