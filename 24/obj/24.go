package obj

import (
	"fmt"
	"math"
)

// Point так как в задании требовались инкапсулированные параметры и конструктор,
//то они указаны с большой буквы - объявление их глобальности,
//тк они находятся в отдельном пакете.
type Point struct {
	x int
	y int
}

//Создние нового объекта структуры Point
func NewPoint(x, y int) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

//Определение расстояния через теорему Пифагора
func Distance(point1, point2 *Point) {
	f := math.Sqrt(math.Pow(math.Abs(float64(point1.x-point2.x)), 2) + math.Pow(math.Abs(float64(point1.y-point2.y)), 2))
	fmt.Println(f)
}
