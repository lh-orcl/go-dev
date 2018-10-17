package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Remove("nukeme")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Deleting directory AND CONTENTS")
		err2 := os.RemoveAll("nukeme")
		if err2 != nil {
			panic(err)
		}
	}
}
