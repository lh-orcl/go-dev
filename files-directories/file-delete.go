package main

import (
	"os"
)

func main() {
	err := os.Remove("toremove.txt")
	if err != nil {
		panic(err)
	}
}
