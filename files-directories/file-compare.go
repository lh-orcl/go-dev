package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	one, err := ioutil.ReadFile("one.txt")
	if err != nil {
		panic(err)
	}
	two, err2 := ioutil.ReadFile("two.txt")
	if err2 != nil {
		panic(err)
	}
	same := bytes.Equal(one, two)
	fmt.Println(same)
}
