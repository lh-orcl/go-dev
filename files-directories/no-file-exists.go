package main

import (
	"fmt"
	"os"
)

func main() {
	if _, err := os.Stat("Log.txt"); os.IsNotExist(err) {
		fmt.Println("Log.txt file does not exist")
	} else {
		fmt.Println("Log.txt file exists")
	}
}
