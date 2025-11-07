package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}

func main() {
	input := "snow dog sun"
	output := reverseWords(input)

	fmt.Printf("Входная строка: %s\n", input)
	fmt.Printf("Выходная строка: %s\n", output)
}
