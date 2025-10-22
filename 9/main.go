package main

import (
	"fmt"
)

func generate(numbers []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, x := range numbers {
			out <- x
		}
	}()
	return out
}

func multiply(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for x := range in {
			out <- x * 2
		}
	}()
	return out
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	firstChannel := generate(numbers)
	secondChannel := multiply(firstChannel)

	for result := range secondChannel {
		fmt.Println(result)
	}
}
