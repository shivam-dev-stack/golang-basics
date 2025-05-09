package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"a": 1, "b": 2}
	fmt.Println(m["a"])
	fmt.Println(m["b"])
	fmt.Println(m["c"]) // returns 0 if not found
}
