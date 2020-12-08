package library

import (
	"math"
)

type Point struct {
	X float64
	Y float64
}

type LineProperty struct {
	length float64
	angle  float64
}

func PointsLength(points []Point) float64 {
	length := 0.
	if len(points) < 1 {
		return length
	}
	current_point := points[0]
	for idx, next_point := range points {
		if idx == 0 {
			continue
		}
		length += math.Hypot(next_point.X-current_point.X, next_point.Y-current_point.Y)
		current_point = next_point
	}

	return length
}

func PropertiesOfLine(start, end Point) LineProperty {
	lengthX := end.X - start.X
	lengthY := end.Y - start.Y
	return LineProperty{math.Sqrt(math.Pow(lengthX, 2) + math.Pow(lengthY, 2)), math.Atan2(lengthY, lengthX)}
}

// https://medium.com/@francoisromain/smooth-a-svg-path-with-cubic-bezier-curves-e37b49d46c74
func ControlPoint(current Point, previous Point, next Point, reverse bool) Point {

	//fmt.Printf("CURRNT.x:%f y:%f Prev.x:%f y:%f Next.x:%f y:%f\n", current.X, current.Y, previous.X, previous.Y, next.X, next.Y)
	p := previous
	if p.X == 0 && p.Y == 0 {
		//	fmt.Println("UF, PREV is empty - making P=CURRENT")
		p = current
	}
	n := next
	if n.X == 0 && n.Y == 0 {
		//	fmt.Println("UF, NEXT is empty - making n=CURRENT")
		n = current
	}
	//fmt.Printf("Previous X:%f Y:%f // Next.X:%f Y:%f\n", p.X, p.Y, n.X, n.Y)
	const smoothing = 0.2
	o := PropertiesOfLine(p, n)
	angle := o.angle
	if reverse {
		angle += math.Pi
	}
	length := o.length * smoothing
	x := current.X + math.Cos(angle)*length
	y := current.Y + math.Sin(angle)*length
	return Point{x, y}
}

func SplitPointsSlice(points []Point) ([]int, []int) {
	x_pos := make([]int, 0)
	y_pos := make([]int, 0)
	for i := 0; i < len(points); i++ {
		x_pos = append(x_pos, int(points[i].X))
		y_pos = append(y_pos, int(points[i].Y))
	}

	return x_pos, y_pos
}
