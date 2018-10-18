package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		channel1 <- "Hello from channel1"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		channel2 <- "Hello from channel2"
	}()
	var result string
	//	for {
	for i := 0; i < 2; i++ {
		select {
		case result = <-channel1:
			fmt.Println("result ", result)
		case result = <-channel2:
			fmt.Println("result ", result)
			//		case <-quit:
			//			return
		}
	}
}
