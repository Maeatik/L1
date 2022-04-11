package main

import (
	"fmt"
)

//К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
//Приведите корректный пример реализации.

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

var justString string

func createHugeString(i int) string {
	s := ""
	for n := 0; n < i; n++ {
		s = s + "a"
	}
	return s
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:]
	fmt.Println(justString)
}
func someFunc_good() {
	v := createHugeString(1 << 10)
	tmp := make([]rune, 100)
	copy(tmp, []rune(v[0:100]))
	justString = string(tmp)
}
func main() {
	someFunc()
	someFunc_good()
}
