package main

import (
	"fmt"
	"strconv"
)

func main() {
	number := int64(100)
	numberStr := strconv.FormatInt(number, 10)
	fmt.Println(numberStr)
}
