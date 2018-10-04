package main

import (
	"fmt"
	"strconv"
)

func main() {
	isNew := "j"
	isNewBool, err := strconv.ParseBool(isNew)
	if err != nil {
		fmt.Println("failed")
	} else {
		if isNewBool {
			fmt.Print("IsNew")
		} else {
			fmt.Println("Not new")
		}
	}
}
