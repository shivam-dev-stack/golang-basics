package main

import "fmt"

// Switch statement

func main() {

	x := -5
	switch {
	case x > 0:
		fmt.Println("Positive")
	case x < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}
}
