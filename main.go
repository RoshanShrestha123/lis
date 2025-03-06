package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

const KB = 1024

func main() {

	var Reset = "\033[0m"
	var Red = "\033[31m"
	var Green = "\033[32m"

	dir, err := os.ReadDir("./")
	if err != nil {
		log.Fatal("Something went wrong", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)
	fmt.Fprintln(w, Green+"Names\tSize(KB)"+Reset+"\t")

	for _, value := range dir {
		file, err := value.Info()
		if err != nil {
			log.Fatal("something went wrong", err)
		}

		var selectedColor string
		if file.IsDir() {
			selectedColor = Red
		} else {
			selectedColor = Green
		}

		fmt.Fprintf(w, selectedColor+"%s\t%.2f"+Reset+"\t\n", file.Name(), float32(file.Size())/KB)

	}

	w.Flush()

}
