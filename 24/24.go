package main

//Разработать программу нахождения расстояния между двумя точками,
//которые представлены в виде структуры Point с
//инкапсулированными параметрами x,y и конструктором.

import (
	"L1/24/obj"
	"fmt"
)

//определение точки для объекта струтуры Point
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
	//вводятся точки х1,у1 и х2,у2
	x1, err := whatPoint(1, "x")
	errCheck(err)
	y1, err := whatPoint(1, "y")
	errCheck(err)
	x2, err := whatPoint(2, "x")
	errCheck(err)
	y2, err := whatPoint(2, "y")
	errCheck(err)
	//создаются новые объекты структуры Point
	point1 := obj.NewPoint(x1, y1)
	point2 := obj.NewPoint(x2, y2)

	obj.Distance(point1, point2)

}
