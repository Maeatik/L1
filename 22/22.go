package main

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
	a, err := whatNum()
	if err != nil {
		fmt.Println(err)
		return
	}

	floatA := float64(a)

	b, err1 := whatNum()
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	floatB := float64(b)

	if a >= 1 && b >= 1 {
		bigA := big.NewFloat(floatA * low)
		fmt.Println("Число 1:", bigA)
		bigB := big.NewFloat(floatB * low)
		fmt.Println("Число 2:", bigB)
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
