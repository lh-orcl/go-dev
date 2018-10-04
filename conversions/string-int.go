package main

import (
	"fmt"
	"strconv"
)

func main() {
	//	number := "cheese"
	number := "2"
	valueInt, err := strconv.ParseInt(number, 10, 32)
	if err != nil {
		fmt.Print("Error happened")
	} else {
		if valueInt == 2 {
			fmt.Println("success")
		}
	}
}
