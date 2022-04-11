package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "snow dog sun 6"
	sentenceArr := strings.Split(sentence, " ")
	for i, j := 0, len(sentenceArr)-1; i < j; i, j = i+1, j-1 {
		sentenceArr[i], sentenceArr[j] = sentenceArr[j], sentenceArr[i]
	}
	fmt.Println(sentenceArr)
}
