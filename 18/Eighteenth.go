package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	mx    sync.Mutex
}

func counter(wg *sync.WaitGroup, c *Counter) {
	defer wg.Done()
	c.mx.Lock()
	c.count = c.count + 1
	fmt.Println(c.count)
	c.mx.Unlock()
}

func nDigit() (int, error) {
	fmt.Println("И сколько мы будем считать?")
	var n int
	_, err := fmt.Scanf("%d", &n)
	return n, err
}
func main() {
	c := &Counter{}
	wg := sync.WaitGroup{}
	n, err := nDigit()
	if err != nil {
		return
	}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go counter(&wg, c)
	}
	wg.Wait()
	fmt.Println("\n", c.count)
}
