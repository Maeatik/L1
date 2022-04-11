package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "abcd"
	runes := []rune(str)
	answer := make(map[rune]int)

	for _, r := range runes {
		answer[unicode.To(unicode.LowerCase, r)]++
	}
	flag := true
	for _, i := range answer {
		if i >= 2 {
			fmt.Println(false)
			flag = false
			break
		}
	}
	if flag {
		fmt.Println(true)
	}
}
