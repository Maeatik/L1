package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func add(i int, num int64) {
	rand.Seed(time.Now().UnixNano())
	//Добавление либо 1, либо 0
	adding := rand.Intn(2)
	//Если выпала 1, используется побитовая дизъюнкция в место добавления 1
	if adding == 1 {
		num |= 1 << i
		//Если выпала не 1, тогда используется побитовая конъюнкция на операцию XOR
	} else {
		num &= ^(1 << i)
	}
	fmt.Println(adding, num)
}

func numOfBit() (int, error) {
	fmt.Println("Введите в какой бит вставить 1 или 0")

	var n int
	_, err := fmt.Scanf("%d", &n)

	return n, err
}

func main() {
	var num int64 = 15

	i, err := numOfBit()
	if err != nil {
		return
	}

	add(i, num)
}
