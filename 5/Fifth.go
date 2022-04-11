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
			//случайное число передается в канал
			c <- rand.Intn(100000)
			time.Sleep(1000 * time.Millisecond)
			//После каждой секунды - счетчик уменьшается
			wg.Done()
		}
	}()
	go func() {
		for {
			//печать числа из канала
			fmt.Printf("%d\n", <-c)
		}
	}()
	//Ожидание того, когда счетчик опустится до 0, то есть через n секунд
	wg.Wait()

}

func main() {
	n, err := nSec()
	if err != nil {
		return
	}
	//Создаем переменную для ожидания горутин
	var wg sync.WaitGroup
	//Увеличиваем счетчик ожидания горутин на n
	wg.Add(n)

	messagePerSec(&wg, n)
}
