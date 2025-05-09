package main

import (
	"fmt"
	"net/http"
	"sync"
)

func fetch(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	res, _ := http.Get(url)
	fmt.Println(url, res.Status)
}

func main() {
	var wg sync.WaitGroup
	urls := []string{"https://python.org", "https://golang.org"}

	for _, url := range urls {
		wg.Add(1)
		go fetch(url, &wg)
	}

	wg.Wait()
}
