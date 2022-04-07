package main

import "fmt"

//Поменять местами два числа без создания временной переменной.

func main() {
	arr := []int{1, 2, 3, 4}

	arr[0], arr[2] = arr[2], arr[0]
	fmt.Println(arr)
}
