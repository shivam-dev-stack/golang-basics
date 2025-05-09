package main

import (
	"fmt"
	"sync"
)

// func main() {

// 	go hello()
// 	fmt.Println("Hello from main")
// 	time.Sleep(1 * time.Second)

// }

func main() {
	// ch := make(chan string)

	ch := make(chan int, 2)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
	}()

	// msg := <-ch
	// fmt.Println(msg)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	close(ch) // close the channel

	var wg sync.WaitGroup

	wg.Add(2) // Wait for 2 goroutines

	go func() {
		fmt.Println("Task 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("Task 2")
		wg.Done()
	}()

	wg.Wait()

}

func hello() {
	fmt.Println("hello") // Start goroutine

}
