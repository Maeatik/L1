package main

import "fmt"

//Имеется последовательность строк - (cat, cat, dog, cat, tree)
//создать для нее собственное множество.
func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	//создается мапа с ключом типа string и булевым значением
	mn := make(map[string]bool)

	for _, i := range arr {
		//по ключу добавляется слово, и булевой переменной указывается, что такое слово в мапе есть
		mn[i] = true
	}
	fmt.Println(mn)
}
