package fastaexample

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// AssembleFile reads a file and assembles the contents
func AssembleFile(file string) (string, error) {
	frags, err := readFile(file)
	if err != nil {
		return "", err
	}
	seq, err := match(frags)
	if err != nil {
		return "", err
	}
	return seq, nil
}

// readFile reads a given file and converts FASTA sequence to an array for assembly
func readFile(file string) ([]*fastaFrag, error) {
	var frags []*fastaFrag

	// Slurp in the whole file.  Line-by-line parsing might
	// be an improvement here
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// Find fragments based on the header, first is empty
	// Convert these to our FastaFrag struct
	chunks := strings.Split(string(data), ">")[1:]
	for _, c := range chunks {
		f, err := fragFromChunk(c)
		if err != nil {
			return nil, err
		}
		frags = append(frags, f)
	}
	return frags, nil
}

func fragFromChunk(c string) (*fastaFrag, error) {
	var f fastaFrag

	// clean up the data, strip off the title
	data := strings.SplitN(c, "\n", 2)
	if len(data) != 2 {
		return nil, fmt.Errorf("failure to split chunk '%s'", c)
	}
	t := strings.TrimSpace(data[0])
	seq := strings.Replace(data[1], "\n", "", -1)

	// sanity check the data against ACGT
	// This could be much smarter, but extracted for later update
	if !sanityCheckSequence(seq) {
		return nil, fmt.Errorf("unparsed data in sequence %s", t)
	}

	f.Title = t
	f.Data = seq

	return &f, nil
}

func sanityCheckSequence(d string) bool {
	sane := strings.Replace(
		strings.Replace(
			strings.Replace(
				strings.Replace(d, "A", "", -1), "C", "", -1), "G", "", -1), "T", "", -1)
	if len(sane) > 0 {
		return false
	}
	return true
}
