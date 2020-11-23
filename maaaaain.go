package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"

	svg "github.com/ajstarks/svgo"
)

var (
	width  = 700
	height = 700
	s1     = rand.NewSource(time.Now().UnixNano())
	r1     = rand.New(s1)
)

type Point struct {
	x float64
	y float64
}

func split_points_slice(points []Point, offset_x int, offset_y int) ([]int, []int) {
	x_pos := make([]int, 0)
	y_pos := make([]int, 0)
	for i := 0; i < len(points); i++ {
		x_pos = append(x_pos, int(points[i].x)+offset_x)
		y_pos = append(y_pos, int(points[i].y)+offset_y)
	}

	return x_pos, y_pos
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

	return length
}

func squiggle(radius float64, length float64) []Point {
	points := make([]Point, 0)
	current_length := 0.
	current_angle := 0.

	//s1 := rand.NewSource(time.Now().UnixNano())
	//r1 := rand.New(s1)

	// points = append(points, Point{radius, 0})
	for current_length < length {
		new_x := radius * math.Cos(current_angle)
		new_y := radius * math.Sin(current_angle)
		points = append(points, Point{new_x, new_y})
		current_length = points_length(points)
		angle_mod := r1.Intn(7) + 1
		current_angle += float64(angle_mod)
		radius_mod := r1.Intn(7) - 3
		radius += float64(radius_mod)
	}

	return points
}

type LineProperty struct {
	length float64
	angle  float64
}

func propertiesOfLine(start, end Point) LineProperty {
	lengthX := end.x - start.x
	lengthY := end.y - start.y
	return LineProperty{math.Sqrt(math.Pow(lengthX, 2) + math.Pow(lengthY, 2)), math.Atan2(lengthY, lengthX)}
}

// https://medium.com/@francoisromain/smooth-a-svg-path-with-cubic-bezier-curves-e37b49d46c74
func controlPoint(current Point, previous Point, next Point, reverse bool) Point {

	p := previous
	if p.x == 0 && p.y == 0 {
		p = current
	}
	n := next
	if n.x == 0 && n.y == 0 {
		n = current
	}
	const smoothing = 0.2
	o := propertiesOfLine(p, n)
	angle := o.angle
	if reverse {
		angle += math.Pi
	}
	length := o.length * smoothing
	x := current.x + math.Cos(angle)*length
	y := current.y + math.Sin(angle)*length
	return Point{x, y}

}

func drawSquig(x int, y int, rad float64, lennt float64, canvas *svg.SVG, color string) {
	//colors := [...]string{"black", "greenyellow", "black"}
	//rand_color := colors[r1.Intn(len(colors))]
	//rand_color := colors[0]
	var squiggly_line []Point = squiggle(rad, lennt)
	x_pos, y_pos := split_points_slice(squiggly_line, x, y)
	//fmt.Printf("LEN OF POINTS IS %d\n", len(squiggly_line))
	for i := 0; i < len(squiggly_line); i++ {

		fmt.Printf("X: %d, Y: %d\n", x_pos[i], y_pos[i])

		// start control point
		//const [cpsX, cpsY] = controlPoint(a[i - 1], a[i - 2], point)
		// end control point
		//const [cpeX, cpeY] = controlPoint(point, a[i - 1], a[i + 1], true)
	}
	canvas.Polyline(x_pos, y_pos, "fill:none; stroke:"+color)
}

func OpenCreateFile(filename string) *os.File {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func main() {

	colors := [...]string{"chartreuse", "greenyellow", "lawngreen", "fuchsia", "yello", "black"}
	// num_draws := r1.Intn(18) + 3
	// num_draws := 3
	num_squigs := 3
	incr := width / num_squigs

	//for i := 0; i < num_draws; i++ {

	//fname := "blah" + strconv.Itoa(i) + ".svg"
	fname := "blah.svg"
	f := OpenCreateFile(fname)

	canvas := svg.New(f)
	canvas.Start(width, height)
	canvas.RGB(0, 0, 0)

	for j := 0; j < num_squigs; j++ {
		drawSquig(j*incr+10, j*incr+10, float64(50), float64(500), canvas, colors[3])
		drawSquig(width-j*incr, height-j*incr+50, float64(10), float64(1500), canvas, colors[5])
	}

	rand_color := colors[r1.Intn(len(colors))]
	canvas.Grid(0, 0, width, height, 10, "stroke:"+rand_color+"; opacity:0.1i; stroke-width="+strconv.Itoa(r1.Intn(10)))

	canvas.End()
	//}
}
