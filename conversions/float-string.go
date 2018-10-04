package main

import (
	"fmt"
	"strconv"
)

func main() {
	number := 100
	numberStr := strconv.Itoa(number)
	fmt.Println(numberStr)

	numberFloat := 123456.789
	numberFloatStr := strconv.FormatFloat(numberFloat, 'f', -1, 64)
	fmt.Println(numberFloatStr)
}
