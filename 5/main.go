package main

import (
	"fmt"
	"time"
)

func main() {
	N := 3
	ch := make(chan int)

	go func() {
		defer close(ch)
		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()

		timeout := time.After(time.Duration(N) * time.Second)
		i := 1
		for {
			select {
			case <-timeout:
				return
			case <-ticker.C:
				ch <- i
				i++
			}
		}
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
