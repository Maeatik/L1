package main
//Реализовать все возможные способы остановки выполнения горутины.
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
	fmt.Println(<-ch) //Получаем данные из канала, после получения - канал закрывается
	fmt.Println(<-ch) //горутина закрыта
}

//Реализовать все возможные способы остановки выполнения горутины
func one() {
	//создание канала, который при каждом вызове увеличивается на 1
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
			//выполняется тогда, когда из канала просят данные
			case ch <- n:
				n++
			//выполняется когда закрывается канал
			case <-ch:
				return
			}
		}
	}()
	return ch
}

func two() {
	//создается канал булевых выражений
	ch := make(chan bool)

	go func() {
		for {
			select {
			//вызывается, когда горутина закрывается(когда в нее передаются данные)
			case <-ch:
				fmt.Println("stop the goroutine...")
				return

			default:
				//выполняется каждую секунду
				time.Sleep(time.Second)
				fmt.Println("Doria`s sheep is eating grass.")
			}
		}
	}()
	//через пять секунд в канал пойдет передача
	time.Sleep(5 * time.Second)
	ch <- true
}

var lastName string = ""

func three() {
	//создаение "канала" через контекст, с методом для закрытия горутины
	ctx, cancel := context.WithCancel(context.Background())
	//передача в "канал" данные - под lastName будет упоминаться Старк
	ctxHasValue := context.WithValue(ctx, lastName, "Stark")
	go sayHello(ctxHasValue, "Snow")
	//горутина выполняется 5 секунд
	time.Sleep(5 * time.Second)
	//Через 5 секунд вызывается метод контекста, для остановки
	cancel()
	//Через 2 секунды придет сообщение об остановке
	time.Sleep(2 * time.Second)

	fmt.Println("main ends...")
}
func sayHello(ctx context.Context, name string) {
	for {
		select {
		//выполняется в случае вызова метода cancel()
		case <-ctx.Done():
			//получение значения из канала контекста по ключу lastName
			fmt.Printf("%s says: Goodbye, Arya %s...\n", name, ctx.Value(lastName))
			return
		//выполняется каждую секунду
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
