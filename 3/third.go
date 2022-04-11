package main

import "fmt"

// Дана последовательность чисел: 2,4,6,8,10.
//Найти сумму их квадратов(22+32+42….)
//с использованием конкурентных вычислений.

//Принцип точно такой же как и в 2 задании
func sumOfPow(digit_array []int) int {
	c := make(chan int)
	defer close(c)
	sum := 0
	for _, digit := range digit_array {
		go pow(digit, c)
		//Передаем данные из канала и добавляем их к сумме
		sum += <-c
	}
	return sum
}

func pow(digit int, c chan int) {
	digitPow := digit * digit
	c <- digitPow

}
func main() {
	digit_array := []int{2, 4, 6, 8, 10}
	//Выводим в консоль сумму корней
	fmt.Println(sumOfPow(digit_array))
}
