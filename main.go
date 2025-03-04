package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	dir, err := os.ReadDir("./")
	if err != nil {
		log.Fatal("Something went wrong", err)
	}

	for _, value := range dir {
		fmt.Println(value)
	}

}
