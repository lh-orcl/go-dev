package main

import (
	"fmt"
	"strconv"
)

func main() {
	isNew := true
	isNewStr := strconv.FormatBool(isNew)
	message := "Purchased item is " + isNewStr
	fmt.Println(message)
}
