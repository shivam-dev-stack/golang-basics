package main

import "fmt"

type Car struct {
	Brand string
	Year  int
}

func (c Car) Info() {
	fmt.Printf("Car: %s (%d)\n", c.Brand, c.Year) // Method on struct
}

func main() {
	myCar := Car{"Toyota", 2020}
	myCar.Info() // Calls method on struct
}
