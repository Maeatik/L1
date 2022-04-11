package obj

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) *Point {
	return &Point{
		x: x,
		y: y,
	}
}
func Distance(point1, point2 *Point) {
	f := math.Sqrt(math.Pow(math.Abs(float64(point1.x-point2.x)), 2) + math.Pow(math.Abs(float64(point1.y-point2.y)), 2))
	fmt.Println(f)
}
