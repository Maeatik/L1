package main

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре
//Action от родительской структуры Human (аналог наследования).
import "fmt"

//Создадим структуру Human
type Human struct {
	name string
	age  int
}

//Создадим метода структуры
func (h Human) say() {
	fmt.Println("I'm", h.name, "i'm", h.age, "years old")
}

//Объявим структуру Action, которая наследуется от структуры Human
type Action struct {
	Human
}

func main() {
	//Через структуру Action создаем объект структуры Human и вызовем ее метод
	action := Action{Human{"marat", 21}}
	action.say()
}
