package main

import "fmt"

//Имеется последовательность строк - (cat, cat, dog, cat, tree)
//создать для нее собственное множество.
func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	mn := make(map[string]bool)

	for _, i := range arr {
		mn[i] = true
	}
	fmt.Println(mn)
}
