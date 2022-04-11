package main

//Реализовать конкурентную запись данных в map.

import (
	"fmt"
	"sync"
)

func main() {
	//создание мапы
	cMap := map[int]int{}
	//создание блокировщика
	mx := sync.Mutex{}
	//создание переменной для ожидания
	wg := sync.WaitGroup{}
	n := 10

	for i := 0; i < n; i++ {
		//увеличиваем счетчик ожидания каждую иттерацию
		wg.Add(1)
		go func(i int) {
			//при вызове горутины блокируем все остальные горутины, так как мапы не потокобезопасны
			mx.Lock()
			//делаем запись в мапу
			cMap[i] = i
			//после записи в мапу, можно разблокировать горутины
			mx.Unlock()
			//уменьшаем счетчик ожидания
			wg.Done()
		}(i)
	}
	//ожидание завершения работы с мапой
	wg.Wait()
	//все элементы добавились, так как длина мапы - 10
	fmt.Println(len(cMap))
	//но так как мапа не безопасна для конкурентных вычислений - порядок записи чисел - случайный
	//какая горутина раньше запустилась, то и будет записываться раньше
	for i := range cMap {
		fmt.Println(cMap[i])
	}

}
