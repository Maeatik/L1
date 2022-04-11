package main

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика.
import (
	"fmt"
	"sync"
)

//Counter создается структура счетчика, в которую входит само поле для счета
//и поле блокировки горутин
type Counter struct {
	count int
	mx    sync.Mutex
}

func counter(wg *sync.WaitGroup, c *Counter) {
	//отложенное уменьшение счетчика
	defer wg.Done()
	//блокировка горутин используя блокировщик структуры Counter
	c.mx.Lock()
	//поле для счета у счетчика увеличивается на 1
	c.count = c.count + 1
	fmt.Println(c.count)
	//разблокировка горутин
	c.mx.Unlock()
}

func nDigit() (int, error) {
	fmt.Println("И сколько мы будем считать?")
	var n int
	_, err := fmt.Scanf("%d", &n)
	return n, err
}
func main() {
	//создается новый счетчик и ожидание
	c := &Counter{}
	wg := sync.WaitGroup{}
	//вводится число n
	n, err := nDigit()
	if err != nil {
		return
	}
	for i := 0; i < n; i++ {
		//каждую иттерацию счетчик ожидания увеличивается
		wg.Add(1)
		go counter(&wg, c)
	}
	wg.Wait()
	fmt.Println("\n", c.count)
}
