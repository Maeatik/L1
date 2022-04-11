package main

//Реализовать собственную функцию sleep.

import (
	"fmt"
	"time"
)

//Функция Sleep которая ждет X секунд
func Sleep(x int) {
	<-time.After(time.Second * time.Duration(x))
}
func main() {
	fmt.Println("Через 5 секунд уснем")
	Sleep(5)
	fmt.Println("Пора спать")
}
