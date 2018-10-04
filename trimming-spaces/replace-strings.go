package main

import (
	"fmt"
	"strings"
)

func main() {
	helloWorld := "Hello, World, World, World"
	helloMars := strings.Replace(helloWorld, "World", "Mars", 2)
	helloLots := strings.Replace(helloWorld, "World", "Mars", -1)
	fmt.Println(helloMars)
	fmt.Println(helloLots)
}
