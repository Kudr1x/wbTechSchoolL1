package main

import "fmt"

func swapXOR(a, b int) (int, int) {
	fmt.Printf("До обмена (XOR): a = %d, b = %d\n", a, b)

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf("После обмена: a = %d, b = %d\n", a, b)
	return a, b
}

func main() {
	fmt.Println(swapXOR(15, 25))
}
