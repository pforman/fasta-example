package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pforman/fasta-example"
)

func main() {
	var debug bool
	var title string
	var wrap int

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s filename\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.BoolVar(&debug, "d", false, "Enable debug info")
	flag.StringVar(&title, "t", "", "Set FASTA title/description")
	flag.IntVar(&wrap, "w", 0, "Wrap lines at specified length")

	flag.Parse()

	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	// Turn on the debug logger if necessary
	fastaexample.SetDebug(debug)

	// Read the file and assemble the fragments, or fail trying
	seq, err := fastaexample.AssembleFile(flag.Args()[0])
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}

	// Add a title (description) before the sequence
	if title != "" {
		fmt.Printf(">%s\n", title)
	}

	// If we asked to linewrap, chunk up the data until we're under that length
	if wrap > 0 {
		for len(seq) > wrap {
			fmt.Println(seq[:wrap])
			seq = seq[wrap:]
		}
	}
	// Print whatever is left over, or everything if we didn't wrap.
	fmt.Println(seq)
}
