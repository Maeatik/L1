package main

import "fmt"

//Реализовать пересечение двух неупорядоченных множеств.

func main() {
	mn1 := map[int]bool{4: true, 3: true, 99: true, 0: true}
	mn2 := map[int]bool{6: true, 4: true, 100: true, 99: true}

	mn3 := make(map[int]bool)

	for i := range mn1 {
		_, flag := mn2[i]
		if flag {
			mn3[i] = true
		}
	}
	fmt.Println(mn3)
}
