package main

//Разработать программу, которая перемножает, делит, складывает,
//вычитает две числовых переменных a,b, значение которых > 2^20.

import (
	"fmt"
	"math/big"
)

const low float64 = 2e20

func whatNum() (int, error) {
	fmt.Println("Введите число больше 1")
	var n int
	_, err := fmt.Scanf("%d \n", &n)
	return n, err
}

func main() {
	//Ввод числа А
	a, err := whatNum()
	if err != nil {
		fmt.Println(err)
		return
	}
	// переведем его во float64, чтобы можно было умножить его на low
	floatA := float64(a)
	//Аналогично с Б
	b, err1 := whatNum()
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	floatB := float64(b)
	//Если оба числа больше единицы, тогда они переводятся в big, путем их умножения на low
	if a > 1 && b > 1 {
		bigA := big.NewFloat(floatA * low)
		fmt.Println("Число 1:", bigA)
		bigB := big.NewFloat(floatB * low)
		fmt.Println("Число 2:", bigB)
		//для соверщения операций над большими числами создаем переменную с указателем на тип, с которым будем работать
		operations := new(big.Float)

		sum := operations.Add(bigA, bigB)
		fmt.Println("Сумма:", sum)

		mult := operations.Mul(bigA, bigB)
		fmt.Println("Произведение:", mult)

		raz := operations.Add(bigA, operations.Neg(bigB))
		fmt.Println("Разность:", raz)

		div := operations.Quo(bigA, bigB)
		fmt.Println("Частное:", div)
	} else {
		fmt.Println("Больше единицы надо было вводить число")
		return
	}
}
