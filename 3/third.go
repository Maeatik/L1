package main

import (
	"fmt"
	"sync"
)

var m sync.Mutex

func sumOfPow(digit_array []int) int  {
	c := make(chan int)
	defer close(c)
	sum := 0
	for _,digit := range digit_array{
		go pow(digit, c)
		sum += <- c
	}
	return sum
}

func pow(digit int, c chan int)  {
	m.Lock()
	defer m.Unlock()
	digitPow := digit * digit
	c <- digitPow

}
func main() {
	digit_array := []int{2,4,6,8,10}
	fmt.Println(sumOfPow(digit_array))
}
