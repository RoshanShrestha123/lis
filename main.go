package main

import (
	"fmt"
	"log"
	"os"
)

const KB = 1024

func main() {

	dir, err := os.ReadDir("./")
	if err != nil {
		log.Fatal("Something went wrong", err)
	}

	for _, value := range dir {
		file, err := value.Info()
		if err != nil {
			log.Fatal("something went wrong", err)
		}
		fmt.Printf(">> %s - %.2f KB\n", file.Name(), float64(file.Size())/KB)
	}

}
