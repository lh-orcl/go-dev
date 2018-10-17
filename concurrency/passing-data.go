package main

import (
	"fmt"
)

func main() {
	nameChannel := make(chan string)
	done := make(chan string)

	go func() {
		names := []string{"luke", "alice", "matt", "oliver"}
		for _, name := range names {
			// Something
			fmt.Println("Processing the first stage of: " + name)
			nameChannel <- name
		}
		close(nameChannel)
	}()

	go func() {
		for name := range nameChannel {
			fmt.Println("Processing the second stage of: " + name)
		}
		done <- ""
	}()
	<-done
}
