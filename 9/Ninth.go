package main

import (
	"fmt"
	"sync"
)

//Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
//во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

func toChannel(wg *sync.WaitGroup, arr []int, ch chan int) {
	defer wg.Done()
	for _, i := range arr {
		ch <- i
	}
	close(ch)
}
func fromChannel(wg *sync.WaitGroup, ch1 chan int, ch2 chan int) {
	defer wg.Done()

	for {
		x, flag := <-ch1
		if !flag {
			break
		}
		ch2 <- x * 2
	}
	close(ch2)
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	wg := sync.WaitGroup{}

	ch1 := make(chan int)
	ch2 := make(chan int)
	wg.Add(2)

	go toChannel(&wg, arr, ch1)
	go fromChannel(&wg, ch1, ch2)

	go func() {
		for {
			fmt.Printf("%d\n", <-ch2)
		}
	}()
	wg.Wait()
}
