package main

import (
	//"fmt"
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

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

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

func ten_lines_of_squiggles() {
	width := 500
	height := 500
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)

	div := 10
	draw_count := div
	incr := width / div

	cur_x := 0
	cur_y := incr

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < div; i++ {
		offset := 0
		for j := 0; j < draw_count; j++ {
			offset += r1.Intn(incr)
			//  fmt.Println("Drawing", draw_count, "at ", cur_x+offset, cur_y)

			var squiggly_line []Point = squiggle(10, 100)
			x_pos, y_pos := split_points_slice(squiggly_line, cur_x+offset, cur_y)
			// canvas.Polyline(x_pos, y_pos)
			canvas.Polygon(x_pos, y_pos)

		}
		cur_x += incr
		cur_y += incr
		draw_count--
	}

	canvas.End()
}

func make_squiggle_alphabet_and_numberline(div float64) [][]Point {
	dictionary := make([][]Point, 26+10) // alphanumeric
	for i := range dictionary {
		dictionary[i] = squiggle(div, 100)
	}
	return dictionary
}
func highpoint_lowlife() {
	width := 500
	height := 500
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)

	text_to_use := "the highpoint lowlife radio"
	//fmt.Println(text_to_use)

	text_len := len(text_to_use)
	//fmt.Println(text_to_use, text_len)

	div := 10
	incr := width / div
	alphabet := make_squiggle_alphabet_and_numberline(float64(incr))

	text_idx := 0
	for i := 0; i < div; i++ {

		for j := 0; j < div; j++ {
			char_to_print := text_to_use[text_idx]
			text_idx++
			if text_idx >= text_len {
				text_idx = 0
			}

			if char_to_print == ' ' {
				// no-op
			} else {

				x := i*incr + (incr / 2)
				y := j*incr + (incr / 2)
				squiggly_line := alphabet[char_to_print-'a']
				x_pos, y_pos := split_points_slice(squiggly_line, x, y)
				canvas.Polyline(x_pos, y_pos)

				//fmt.Println("X:", i*incr, "Y:", j*incr+incr, "Char:", string(char_to_print), char_to_print-'a')
			}
		}
	}

	canvas.End()
}

func golden_ratio() {
	width := 500
	height := 500
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)

	golden_ratio_percent := 61.8

	cur_width := float64(width)
	cur_x := int(cur_width - (cur_width / 100. * golden_ratio_percent))
	cur_y := cur_x
	for i := 0; i < 27; i++ {
		big_bit := cur_width / 100.0 * golden_ratio_percent
		wee_bit := cur_width - big_bit
		//fmt.Println("X:", cur_x, " Y:", cur_y, " BIG:", big_bit, " WEE:", wee_bit)
		for j := 0; j < width/(int(wee_bit)+1); j++ {
			var squiggly_line []Point = squiggle(wee_bit+10, big_bit)
			x_pos, y_pos := split_points_slice(squiggly_line, cur_x, cur_y)
			canvas.Polyline(x_pos, y_pos)
			cur_y += 10 + int(wee_bit*big_bit)
			cur_y = cur_y % width
		}
		cur_width = big_bit
		cur_x += 10 + int(big_bit)
		cur_x = cur_x % height
	}

	canvas.End()
}

func main() {
	golden_ratio()
}
