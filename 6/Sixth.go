package main

import (
	"context"
	"fmt"
	"time"
)

func zero() {
	ch := make(chan string)
	go func() {
		ch <- "Hello!"
		close(ch)
	}()
	fmt.Println(<-ch)
	fmt.Println(<-ch) //горутина закрыта
}

//Реализовать все возможные способы остановки выполнения горутины
func one() {
	num := counter()
	for i := 0; i < 5; i++ {
		fmt.Println(<-num)
	}
	close(num)
}
func counter() chan int {
	ch := make(chan int)
	go func() {
		n := 1
		for {
			select {
			case ch <- n:
				n++
			case <-ch:
				return
			}
		}
	}()
	return ch
}

func two() {
	ch := make(chan bool)

	go func() {
		for {
			select {
			case <-ch:
				fmt.Println("stop the goroutine...")
				return
			default:
				time.Sleep(time.Second)
				fmt.Println("Doria`s sheep is eating grass.")
			}
		}
	}()

	time.Sleep(5 * time.Second)
	ch <- true
}

var lastName string = ""

func three() {

	ctx, cancel := context.WithCancel(context.Background())
	ctxHasValue := context.WithValue(ctx, lastName, "Stark")
	go sayHello(ctxHasValue, "Snow")

	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)

	fmt.Println("main ends...")
}
func sayHello(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s says: Goodbye, Arya %s...\n", name, ctx.Value(lastName))
			return
		default:
			time.Sleep(time.Second)
			fmt.Printf("Hello, Arya %s, I`m %s\n", ctx.Value(lastName), name)
		}
	}
}

func main() {
	//0
	zero()
	//1
	one()
	//2
	two()
	//3
	three()

}
