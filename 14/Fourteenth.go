package main

import (
	"fmt"
	"reflect"
)

//Разработать программу, которая в рантайме способна определить
//тип переменной: int, string, bool, channel из переменной типа interface{}.
func getType(i interface{}) {
	fmt.Println(reflect.TypeOf(i))
}
func main() {
	getType(1)
	getType("1")
	getType(true)
	getType(make(chan int))
	getType(make(chan string))
	getType(make(chan bool))
}
