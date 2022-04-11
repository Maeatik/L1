package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.
//
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
func workers(n int) {
	//создается канал для чисел
	c := make(chan int)
	//Создание N воркеров и передача им данных из канала
	for i := 1; i <= n; i++ {
		go worker(i, c)
	}
	//Передача в канал случайных чисел для воркеров каждую секунду
	go func() {
		rand.Seed(time.Now().UnixNano())
		for {
			c <- rand.Intn(100000)
			time.Sleep(1000 * time.Millisecond)
		}
	}()

}

func worker(index int, work <-chan int) {
	for w := range work {
		fmt.Fprintf(os.Stdout, "Воркер №%d: получил %d рублей\n", index, w)
	}
}

func nWorker() (int, error) {
	fmt.Println("Введите кол-во воркеров")
	var n int
	_, err := fmt.Scanf("%d", &n)
	return n, err
}
func main() {
	//Введем количество воркеров
	n, err := nWorker()
	if err != nil {
		return
	}
	//Запускаем постоянную запись в канал
	workers(n)
	//Создаем канал, который будет ждать остановки программы
	c := make(chan os.Signal)
	//Ctrl+C
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	fmt.Fprintf(os.Stdout, "Выходной")

}
