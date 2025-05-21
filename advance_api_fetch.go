package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Define struct to match JSON response
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	// API URL
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// Make GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read Error:", err)
		return
	}

	// Unmarshal JSON into struct
	var post Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		fmt.Println("JSON Unmarshal Error:", err)
		return
	}

	// Print result
	fmt.Println("Title:", post.Title)
	fmt.Println("Body:", post.Body)
}
