package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в канал,
//а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

func nSec() (int, error) {
	fmt.Println("Введите кол-во секунд для программы")
	var n int
	_, err := fmt.Scanf("%d", &n)
	return n, err
}

func messagePerSec(wg *sync.WaitGroup, n int) {
	rand.Seed(time.Now().UnixNano())
	c := make(chan int)
	go func() {
		for {
			c <- rand.Intn(100000)
			time.Sleep(1000 * time.Millisecond)
			wg.Done()
		}
	}()
	go func() {
		for {
			fmt.Printf("%d\n", <-c)
		}
	}()
	wg.Wait()

}

func main() {
	n, err := nSec()
	if err != nil {
		return
	}
	var wg sync.WaitGroup
	wg.Add(n)
	messagePerSec(&wg, n)
}
