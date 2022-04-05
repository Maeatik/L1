package main

import (
	"fmt"
	"os"
	"sync"
)

var m sync.Mutex

func pow(digit int, c chan int)  {
	m.Lock()
	defer m.Unlock()
	digitPow := digit * digit
	c <- digitPow

}

func main() {
	digit_array := []int{2,4,6,8,10}
	c := make(chan int)
	defer close(c)
	for _,digit := range digit_array{
		go pow(digit, c)
		fmt.Fprintln(os.Stdout, <-c)
	}
}
