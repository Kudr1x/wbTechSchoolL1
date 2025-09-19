package main

import "fmt"

func setBit(n int64, i uint, v uint) int64 {
	if v == 0 {
		return n &^ (1 << i)
	}
	return n | (1 << i)
}

func main() {
	var n int64 = 5
	fmt.Println(setBit(n, 0, 0))
	fmt.Println(setBit(n, 2, 0))
}
