package main

import "fmt"

type PaymentProcessor interface {
	ProcessPayment(amount float64) bool
}

type Service1 struct{}

func (s *Service1) ProcessPayment(amount float64) bool {
	fmt.Printf("Обработка платежа через Сервис1: %.2f руб.\n", amount)
	return true
}

type Service2 struct{}

func (s *Service2) ChargeCard(cardNumber string, amount float64) bool {
	fmt.Printf("Списание с карты %s через Сервис2: %.2f руб.\n", cardNumber, amount)
	return true
}

type Service2Adapter struct {
	service    *Service2
	cardNumber string
}

func NewService2Adapter(service *Service2, cardNumber string) *Service2Adapter {
	return &Service2Adapter{
		service:    service,
		cardNumber: cardNumber,
	}
}

func (a *Service2Adapter) ProcessPayment(amount float64) bool {
	return a.service.ChargeCard(a.cardNumber, amount)
}

func processClientPayment(processor PaymentProcessor, amount float64) {
	success := processor.ProcessPayment(amount)
	if success {
		fmt.Println("Платёж успешно обработан\n")
	} else {
		fmt.Println("Ошибка обработки платежа\n")
	}
}

func main() {
	fmt.Println("=== Демонстрация паттерна Адаптер ===\n")

	service1 := &Service1{}
	fmt.Println("1. Платёж через Сервис1:")
	processClientPayment(service1, 1500.0)

	service2 := &Service2{}
	service2Adapter := NewService2Adapter(service2, "4532-****-****-1234")
	fmt.Println("2. Платёж через Сервис2 (через адаптер):")
	processClientPayment(service2Adapter, 2500.0)
}
