package main

import "fmt"

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	examples := []string{
		"бебебе с бабаба",
		//"Улыбок тебе дед Макар",
		"Hello, 世界",
		"абоба 😎",
		"Ångström",
	}

	for _, s := range examples {
		reversed := ReverseString(s)
		fmt.Printf("Оригинал: %s -> Развернуто: %s\n", s, reversed)
	}
}
