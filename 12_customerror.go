package main

import "fmt"

type MyError struct {
	Code int
	Msg  string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error [%d]: %s", e.Code, e.Msg)
}

func doSomething() error {
	return &MyError{Code: 404, Msg: "Not Found"}
}

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println(err)
	}
}
