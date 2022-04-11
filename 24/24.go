package main

import (
	"L1/24/obj"
	"fmt"
)

func whatPoint(i int, coord string) (int, error) {
	fmt.Println("Введите координаты", coord, i, "-ой точки")
	var n int
	_, err := fmt.Scanf("%d \n", &n)
	return n, err
}
func errCheck(err error) error {
	if err != nil {
		return err
	}
	return nil
}
func main() {
	x1, err := whatPoint(1, "x")
	errCheck(err)
	y1, err := whatPoint(1, "y")
	errCheck(err)
	x2, err := whatPoint(2, "x")
	errCheck(err)
	y2, err := whatPoint(2, "y")
	errCheck(err)
	point1 := obj.NewPoint(x1, y1)
	point2 := obj.NewPoint(x2, y2)

	obj.Distance(point1, point2)

}
