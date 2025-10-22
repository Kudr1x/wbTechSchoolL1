package main

import (
	"fmt"
)

func intersectionMap(a, b []int) []int {
	set := make(map[int]bool)
	for _, val := range a {
		set[val] = true
	}

	result := []int{}
	for _, val := range b {
		if set[val] {
			result = append(result, val)
			delete(set, val)
		}
	}

	return result
}

func intersectionLoop(a, b []int) []int {
	result := []int{}

	for _, val1 := range a {
		for _, val2 := range b {
			if val1 == val2 {
				result = append(result, val1)
				break
			}
		}
	}

	return result
}

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}

	fmt.Println("Множество A:", A)
	fmt.Println("Множество B:", B)
	fmt.Println("Пересечение (map):", intersectionMap(A, B))
	fmt.Println("Пересечение (loop):", intersectionLoop(A, B))
}
