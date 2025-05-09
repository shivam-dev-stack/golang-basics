package main

import "fmt"

func main() {

	type Person struct {
		Name string
		Age  int
	}

	p := Person{Name: "Alice", Age: 30}
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}
