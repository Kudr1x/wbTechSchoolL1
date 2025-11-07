package main

import "fmt"

func deleteElement(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		fmt.Println("Ошибка: индекс вне диапазона")
		return slice
	}

	copy(slice[i:], slice[i+1:])

	slice[len(slice)-1] = 0

	return slice[:len(slice)-1]
}

func main() {

	numbers := []int{10, 20, 30, 40, 50, 60, 70, 80}
	fmt.Printf("Исходный слайс: %v (длина: %d, ёмкость: %d)\n\n", numbers, len(numbers), cap(numbers))

	indexToDelete := 3
	fmt.Printf("Удаляем элемент с индексом %d (значение: %d)\n", indexToDelete, numbers[indexToDelete])
	numbers = deleteElement(numbers, indexToDelete)
	fmt.Printf("Результат: %v (длина: %d, ёмкость: %d)\n\n", numbers, len(numbers), cap(numbers))
}
