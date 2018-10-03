package main

import "fmt"

func main() {
	var a int = 5
	fmt.Println(a)

	var b = 10
	fmt.Println(b)

	c := 15
	fmt.Println(c)

	var d byte = 20
	fmt.Println(d)

	var e int
	fmt.Println(e)
	e = 25
	fmt.Println(e)

	f := 30
	f = 35
	// Go variables must be *read*, else:
	// ./var_numbers.go:23:2: f declared and not used
	fmt.Println(f)
}
