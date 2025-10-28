package main

import "fmt"

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{5, 11, 12, 22, 25, 34, 64, 90}

	target := 22
	index := binarySearch(arr, target)

	if index != -1 {
		fmt.Printf("Элемент %d найден на индексе %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден\n", target)
	}

	target = 100
	index = binarySearch(arr, target)

	if index != -1 {
		fmt.Printf("Элемент %d найден на индексе %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден\n", target)
	}
}
