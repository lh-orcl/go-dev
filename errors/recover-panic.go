package main

import (
	"fmt"
)

func main() {
	//Deffered function returns after everything else within the function
	defer func() {
		if r := recover(); r != nil {
			fmt.Print("Recovered in f", r)
		}
	}()
	//This is executed first
	fmt.Println("1")
	writeSomething()
}

//This is outside the main function, so it is executed last
func writeSomething() {
	//Anything
	panic("Write operation error")
}
