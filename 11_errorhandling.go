package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

// In this example, the `divide` function returns an error if the second argument is zero. The main function checks for this error and handles it appropriately.
// This is a simple demonstration of error handling in Go. In real-world applications,  typically log the error or take some other action based on the error type.
