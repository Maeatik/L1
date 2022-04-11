package main

//Написать программу, которая конкурентно рассчитает значение квадратов
//чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

import (
	"fmt"
	"os"
)

func pow(digit int, c chan int) {
	digitPow := digit * digit
	//квадрат числа передаем в канал
	c <- digitPow
}

func main() {
	digit_array := []int{2, 4, 6, 8, 10}
	//Создадим канал с типом int
	c := make(chan int)
	//Закроем его после выполнения программы
	defer close(c)
	//Пробегаемся по каждой цифре
	for _, digit := range digit_array {
		//запускаем метод возведения в степень в горутину
		go pow(digit, c)
		//Получаем данные из канала и выводим их через stdout
		fmt.Fprintln(os.Stdout, <-c)
	}
}
