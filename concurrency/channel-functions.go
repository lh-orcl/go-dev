package main

import (
	"fmt"
	"time"
)

func main() {
	nameChannel := make(chan string)

	go func() {
		names := []string{"luke", "alice", "matt", "oliver"}
		for _, name := range names {
			time.Sleep(1 * time.Second)
			nameChannel <- name
		}
		close(nameChannel)
	}()

	for data := range nameChannel {
		fmt.Println(data)
	}
}
