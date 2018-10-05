package main

import (
	"fmt"
	"time"
)

func main() {
	current := time.Now()
	nextMonth := current.AddDate(0, 1, 0)
	nextYear := current.AddDate(1, 0, 0)

	fmt.Println(current.String())
	fmt.Println(nextMonth.String())
	fmt.Println(nextYear.String())

	lastYear := current.AddDate(-1, 0, 0)
	fmt.Println(lastYear.String())

	nextMinute := current.Add(1 * time.Minute)
	fmt.Println(nextMinute.String())
}
