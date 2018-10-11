package main

import (
	"fmt"
)

func main() {
	nameAges := map[string]int{
		"Tarik":   32,
		"Michael": 30,
		"Jon":     25,
		"Jessica": 20,
	}

	/*	fmt.Println(nameAges["Tarik"])

		value, exists := nameAges["Tarik"]
		fmt.Println(value)
		fmt.Println(exists)

		value2, exists2 := nameAges["Jessica"]
		fmt.Println(value2)
		fmt.Println(exists2)
	*/

	if value, exists := nameAges["Jessica"]; exists {
		fmt.Println(value)
	} else {
		fmt.Println("Jessica cannot be found")
	}
}
