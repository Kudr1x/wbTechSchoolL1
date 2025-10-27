package main

import (
	"fmt"
)

func detectTypeGeneral(value interface{}) string {
	switch v := value.(type) {
	case int:
		return fmt.Sprintf("Тип: int, значение: %d", v)
	case string:
		return fmt.Sprintf("Тип: string, значение: \"%s\"", v)
	case bool:
		return fmt.Sprintf("Тип: bool, значение: %t", v)
	case chan int, chan string, chan bool, chan interface{}:
		return fmt.Sprintf("Тип: channel (%T), значение: %v", v, v)
	default:
		return fmt.Sprintf("Неизвестный тип: %T, значение: %v", v, v)
	}
}

func main() {

	var intVar interface{} = 42
	var strVar interface{} = "Hello, Go!"
	var boolVar interface{} = true
	var chanIntVar interface{} = make(chan int)
	var chanStrVar interface{} = make(chan string, 5)

	fmt.Println(detectTypeGeneral(intVar))
	fmt.Println(detectTypeGeneral(strVar))
	fmt.Println(detectTypeGeneral(boolVar))
	fmt.Println(detectTypeGeneral(chanIntVar))
	fmt.Println(detectTypeGeneral(chanStrVar))

	close(chanIntVar.(chan int))
	close(chanStrVar.(chan string))
}
