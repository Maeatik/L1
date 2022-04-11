package main

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

import (
	"fmt"
	"strings"
)

func main() {
	//Аналогичное задание, однако вместо слайса рун, обрабатывается слайс из слов, который
	//образовался путем деления строки пробелами. Они так же меняются в цикле местами
	sentence := "snow dog sun 6"
	sentenceArr := strings.Split(sentence, " ")
	for i, j := 0, len(sentenceArr)-1; i < j; i, j = i+1, j-1 {
		sentenceArr[i], sentenceArr[j] = sentenceArr[j], sentenceArr[i]
	}
	fmt.Println(sentenceArr)
}
