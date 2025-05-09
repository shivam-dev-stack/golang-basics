package main

import (
	"fmt"
)

func main() {

	var x int = 10
	var p *int = &x // p now holds the memory address of x

	fmt.Println(x)  // 10
	fmt.Println(p)  //  (address)
	fmt.Println(*p) // 10 (value at address)

	changeValue(&x)
	fmt.Println(x)

	person := Person{"Alice", 29}
	birthday(&person)
	fmt.Println(person.Age) // Output: 30

	z := new(int) // allocates zeroed int, returns pointer
	// *z = 42
	*z = 42         // dereference and assign
	fmt.Println(*z) // 42

}

func changeValue(n *int) {
	*n = 100 // dereference and update
}

type Person struct {
	Name string
	Age  int
}

func birthday(p *Person) {
	p.Age++
}
