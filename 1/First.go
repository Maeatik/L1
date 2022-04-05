package main

import "fmt"

type Human struct {
	name string
	age int
}

func (h Human) say(){
	fmt.Println("I'm", h.name, "i'm", h.age, "years old")
}

type Action struct {
	Human
}

func main() {
	action := Action{Human{"marat", 21}}
	action.say()
}
