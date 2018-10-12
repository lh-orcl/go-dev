package main

import (
	"fmt"
	"time"
)

func main() {
	parsedDate, err := time.Parse("2006", "2018 abc")
	if err != nil {
		fmt.Println("An error occured", err.Error())
	} else {
		fmt.Println(parsedDate)
	}
}
