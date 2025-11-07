package main

import (
	"fmt"
	"strings"
)

func AreAllCharactersUnique(str string) bool {
	str = strings.ToLower(str)

	seen := make(map[rune]bool)

	for _, char := range str {
		if seen[char] {
			return false
		}
		seen[char] = true
	}

	return true
}

func main() {
	testCases := []string{
		"abcd",
		"abCdefAaf",
		"aabcd",
		"hello",
		"world",
		"",
		"a",
		"ABC",
		"12345",
		"112233",
	}

	for _, testStr := range testCases {
		result := AreAllCharactersUnique(testStr)
		fmt.Printf("\"%s\" -> %t\n", testStr, result)
	}
}
