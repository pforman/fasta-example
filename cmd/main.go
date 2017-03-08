package main

import (
	"fmt"
	"os"

	"github.com/pforman/fasta-example"
)

func main() {
	fastaexample.Noop()
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s filename\n", os.Args[0])
		os.Exit(1)
	}
	x, err := fastaexample.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
	for i, f := range x {
		fmt.Printf("Sequence %d (%s)\n%s\nEnd Sequence %d (%d pairs)\n", i, f.Title, f.Data, i, f.Length)
	}
}
