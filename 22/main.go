package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("123456789012345678901234567890", 10)
	b.SetString("987654321098765432109876543210", 10)

	fmt.Printf("a = %s\n", a.String())
	fmt.Printf("b = %s\n\n", b.String())

	product := new(big.Int)
	product.Mul(a, b)
	fmt.Printf("Умножение (a * b) = %s\n\n", product.String())

	quotient := new(big.Int)
	quotient.Div(a, b)
	fmt.Printf("Деление (a / b) = %s\n\n", quotient.String())

	sum := new(big.Int)
	sum.Add(a, b)
	fmt.Printf("Сложение (a + b) = %s\n\n", sum.String())

	difference := new(big.Int)
	difference.Sub(a, b)
	fmt.Printf("Вычитание (a - b) = %s\n\n", difference.String())

	remainder := new(big.Int)
	remainder.Mod(a, b)
	fmt.Printf("Остаток (a %% b) = %s\n\n", remainder.String())

	power := new(big.Int)
	power.Exp(a, big.NewInt(3), nil)
	fmt.Printf("Возведение в степень (a^3) = %s\n\n", power.String())

	cmp := a.Cmp(b)
	fmt.Print("Сравнение a и b: ")
	switch cmp {
	case -1:
		fmt.Println("a < b")
	case 0:
		fmt.Println("a == b")
	case 1:
		fmt.Println("a > b")
	}
}
