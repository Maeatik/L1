package main

import (
	"fmt"
	"strings"
)

//К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
//Приведите корректный пример реализации.
//
//
//var justString string
//
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
//}
//
//func main() {
//	someFunc()
//}
func main() {
	n := 1 << 10
	var v strings.Builder

	v.Grow(n)
	for x := 0; x < n; x++ {
		v.WriteRune('1')
	}

	vs := v.String()
	fmt.Println(vs[:10])

	runes := []rune(vs)
	fmt.Println(string(runes[:10]))
}
