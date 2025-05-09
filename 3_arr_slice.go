package main

import "fmt"

func main() {
	var a = [3]int{1, 2, 3}

	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])
	fmt.Println(a[0:2])
	fmt.Println(a[1:3])
	fmt.Println(a[0:3])
	fmt.Println(a[1:3])
	fmt.Println(a[0:2:2])
	fmt.Println(a[1:3:3])

	// a = append(a, 4)  // not possible

	b := []int{1, 2, 3, 4, 5}
	b = append(b, 6)
	fmt.Println(b)

}
