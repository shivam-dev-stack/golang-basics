package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string { return "Woof!" }

type Cat struct{}

func (c Cat) Speak() string { return "Meow!" }

func MakeItSpeak(s Speaker) {
	fmt.Println(s.Speak())
}

type PaymentProcessor interface {
	Pay(amount float64) string
}

type PayPal struct{}

func (p PayPal) Pay(a float64) string { return "Paid via PayPal" }

type Stripe struct{}

func (s Stripe) Pay(a float64) string { return "Paid via Stripe" }

func MakePayment(p PaymentProcessor, amount float64) {
	fmt.Println(p.Pay(amount))
}

func main() {
	MakeItSpeak(Dog{}) // Woof!
	MakeItSpeak(Cat{}) // Meow!
	MakePayment(PayPal{}, 99.99)
	MakePayment(Stripe{}, 50.00)
}
