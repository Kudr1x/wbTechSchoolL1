package main

import (
	"fmt"
)

func createSet(words []string) []string {
	set := make(map[string]bool)

	for _, word := range words {
		set[word] = true
	}

	result := []string{}
	for word := range set {
		result = append(result, word)
	}

	return result
}

func createSetOrdered(words []string) []string {
	set := make(map[string]bool)
	result := []string{}

	for _, word := range words {
		if !set[word] {
			set[word] = true
			result = append(result, word)
		}
	}

	return result
}

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println("Исходная последовательность:", sequence)
	fmt.Println("Множество (неупорядоченное):", createSet(sequence))
	fmt.Println("Множество (сохранён порядок):", createSetOrdered(sequence))
}
