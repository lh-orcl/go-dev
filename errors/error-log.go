package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log_file, err := os.Create("log_file")
	if err != nil {
		fmt.Println("An error occured")
	}
	defer log_file.Close()
	log.SetOutput(log_file)

	log.Fatalln("Fatal: Application crashed!")
	log.Println("Doing some logging")
}
