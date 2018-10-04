package main

import (
	"fmt"
	"strings"
)

func main() {
	helloWorld := "hello world, how are you?"
	helloWorldTitle := strings.Title(helloWorld)
	fmt.Println(helloWorldTitle)
	helloWorldUpper := strings.ToUpper(helloWorld)
	fmt.Println(helloWorldUpper)
}
