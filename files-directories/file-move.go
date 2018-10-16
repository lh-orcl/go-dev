package main

import (
	"io"
	"os"
)

func main() {
	original, err := os.Open("original.txt")
	if err != nil {
		panic(err)
	}
	defer original.Close()
	original_copy, err2 := os.Create("target/copy.txt")
	if err2 != nil {
		panic(err)
	}
	defer original_copy.Close()
	_, err3 := io.Copy(original_copy, original)
	if err3 != nil {
		panic(err)
	}
	original.Close()
	os.Remove("original.txt")
}
