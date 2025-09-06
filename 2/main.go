package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	for _, num := range numbers {

		wg.Go(func() {
			square := num * num
			fmt.Println(square)
		})
	}

	wg.Wait()
}
