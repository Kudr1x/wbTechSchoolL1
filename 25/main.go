package main

import (
	"fmt"
	"time"
)

func Sleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	fmt.Println("Начало выполнения...")
	start := time.Now()

	Sleep(2 * time.Second)

	elapsed := time.Since(start)
	fmt.Printf("Завершено. Прошло времени: %v\n", elapsed)
}
