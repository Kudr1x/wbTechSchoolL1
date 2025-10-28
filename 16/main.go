package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	return quickSortHelper(arr, 0, len(arr)-1)
}

func quickSortHelper(arr []int, low, high int) []int {
	if low < high {

		pivotIndex := partition(arr, low, high)

		quickSortHelper(arr, low, pivotIndex-1)
		quickSortHelper(arr, pivotIndex+1, high)
	}
	return arr
}

func partition(arr []int, low, high int) int {
	mid := low + (high-low)/2
	pivot := arr[mid]

	arr[mid], arr[high] = arr[high], arr[mid]

	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90, 5}
	fmt.Println("Исходный массив:", arr)

	sorted := quickSort(arr)
	fmt.Println("Отсортированный массив:", sorted)
}
