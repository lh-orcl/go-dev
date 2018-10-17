package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		names := []string{"luke", "alice", "matt", "oliver"}

		for _, name := range names {
			time.Sleep(1 * time.Second)
			fmt.Println(name)
		}
	}()

	go func() {
		ages := []int{1, 2, 3, 4, 5}

		for _, age := range ages {
			time.Sleep(1 * time.Second)
			fmt.Println(age)
		}
	}()
	time.Sleep(10 * time.Second)
}
