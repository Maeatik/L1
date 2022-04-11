package main

import (
	"fmt"
	"time"
)

func Sleep(x int) {
	<-time.After(time.Second * time.Duration(x))
}
func main() {
	fmt.Println("Через 5 секунд уснем")
	Sleep(5)
	fmt.Println("Пора спать")
}
