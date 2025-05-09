package main

import (
	"fmt"
	"reflect"
)

func main() {

	var x int = 10              // Explicit type
	y := 20.0                   // Type inference
	const PI float64 = 22.0 / 7 // Constants

	fmt.Println("x:", reflect.TypeOf(x).String())
	fmt.Println("y:", reflect.TypeOf(y).String())
	fmt.Println("PI:", PI)
	fmt.Println("x + y:", add(float64(x), y))
	age(28)
	fmt.Println()
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	counter := 5
	for counter < 10 {
		fmt.Println("Counter:", counter)
		counter++
	}

}

func add(a float64, b float64) float64 {
	return a + b
}

func age(age int) {
	if age > 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}

}
