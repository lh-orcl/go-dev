package main

import (
	"fmt"
	"time"
)

func main() {

	names := []string{"luke", "alice", "matt", "oliver"}

	for _, name := range names {
		time.Sleep(1 * time.Second)
		fmt.Println(name)
	}

	ages := []int{1, 2, 3, 4, 5}
	for _, age := range ages {
		time.Sleep(1 * time.Second)
		fmt.Println(age)
	}
}
