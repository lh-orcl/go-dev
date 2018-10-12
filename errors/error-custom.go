package main

import "fmt"

type MyError struct {
	ShortMessage    string
	DetailedMessage string
	//Name string
	//Age int
}

func (e *MyError) Error() string {
	return e.ShortMessage + "\n" + e.DetailedMessage
}

func main() {
	err := doSomething()
	fmt.Print(err)
}

func doSomething() error {
	//Something, anything?!
	return &MyError{
		ShortMessage:    "Woo, something happened!",
		DetailedMessage: "File not found",
	}
}
