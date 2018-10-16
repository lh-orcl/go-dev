package main

import (
	"fmt"
	"os"
)

func main() {
	if _, err := os.Stat("rename1.txt"); err == nil {
		os.Rename("rename1.txt", "rename2.txt")
	} else if _, err2 := os.Stat("rename2.txt"); err2 == nil {
		os.Rename("rename2.txt", "rename1.txt")
	} else {
		fmt.Println("file does not exist, please create either \"rename1.txt\" or \"rename2.txt\"")
	}
}
