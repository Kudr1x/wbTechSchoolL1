package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) Speak() {
	fmt.Printf("Меня зовут %s, мне %d\n", h.Name, h.Age)
}

func (h Human) Move(dest string) {
	fmt.Printf("%s идет в %s\n", h.Name, dest)
}

func (h *Human) Birthday() {
	h.Age++
}

type Action struct {
	Human
	Verb string
}

func (a Action) Speak() {
	fmt.Printf("%s выполняет действие: %s\n", a.Name, a.Verb)
}

func (a *Action) Do(dest string) {
	a.Move(dest)
	a.Birthday()
	a.Human.Speak()
	a.Speak()
}

func main() {
	a := Action{
		Human: Human{Name: "Алиса", Age: 20},
		Verb:  "бег",
	}

	a.Speak()
	a.Human.Speak()
	a.Move("университет")
	a.Birthday()
	fmt.Println("Возраст:", a.Age)

	a.Do("домой")
}
