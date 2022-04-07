package main

import (
	"fmt"
	"sync"
)

func main() {
	cMap := map[int]int{}
	mx := sync.Mutex{}
	wg := sync.WaitGroup{}
	n := 10

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			mx.Lock()
			cMap[i] = i
			mx.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println(len(cMap))

	for i := range cMap {
		fmt.Println(cMap[i])
	}

}
