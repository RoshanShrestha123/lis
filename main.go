package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

const KB = 1024

func main() {

	dir, err := os.ReadDir("./")
	if err != nil {
		log.Fatal("Something went wrong", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "Names\tSize(KB)\t")
	fmt.Fprintln(w, "-\t-\t")

	for _, value := range dir {
		file, err := value.Info()
		if err != nil {
			log.Fatal("something went wrong", err)
		}

		fmt.Fprintf(w, "%s\t%f\t\n", file.Name(), float32(file.Size())/KB)

	}

	w.Flush()

}
