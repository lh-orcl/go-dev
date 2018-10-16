package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("names.txt")
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	//Look at this later: https://www.joeshaw.org/dont-defer-close-on-writable-files/
	defer file.Close()
	fmt.Println(lineCount)
}
