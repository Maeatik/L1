package main

import (
	"fmt"
	"math/big"
)

const low float64 = 2e20

func whatNum() (int, error) {
	fmt.Println("Введите число больше или равно 1")
	var n int
	_, err := fmt.Scanf("%d", &n)
	return n, err
}

func main() {
	a, err := whatNum()
	if err != nil {
		fmt.Println(err)
		return
	}
	floatA := float64(a)

	b, err := whatNum()
	if err != nil {
		return
	}
	floatB := float64(b)

	if a >= 1 && b >= 1 {
		bigA := big.NewFloat(floatA * low)
		fmt.Println(bigA)
		bigB := big.NewFloat(floatB * low)
		fmt.Println(bigB)
		operations := new(big.Float)
		sum := operations.Add(bigA, bigB)
		fmt.Println(sum)

	} else {
		return
	}
}
