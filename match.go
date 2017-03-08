package fastaexample

import (
	"fmt"
	//"strings"
)

func Match(frags []*FastaFrag) (string, error) {

	// begin at the beginning
	base := frags[0]
	frags = frags[1:]

	complete, final := recurseMatch(base, frags)
	if complete {
		return final.Data, nil
	}

	return "", fmt.Errorf("no full match found")
}

func recurseMatch(base *FastaFrag, frags []*FastaFrag) (bool, *FastaFrag) {
	var err error
	var res *FastaFrag

	// If we've reached an empty fragment list, we're done.
	if len(frags) == 0 {
		return true, base
	}

	fmt.Printf("rM: len %d, base: %d\n", len(frags), base.Length)

	for i, f := range frags {
		if f.Length < base.Length {
			res, err = matchPairs(base, f, (f.Length/2)+1)
		} else {
			// unlikely, other than on the first match,
			// but matchPairs() is strict about s1 > s2
			res, err = matchPairs(f, base, (base.Length/2)+1)
		}
		if err != nil {
			fmt.Printf("trouble matching index %d of %d, skipping\n", i, len(frags))
			fmt.Printf("error was: %s\n", err)
			continue
		}
		// continue on with the matched fragment placed and removed from the list
		complete, final := recurseMatch(res, append(frags[:i], frags[i+1:]...))
		if complete {
			return complete, final
		}
	}
	fmt.Printf("bottomed out, no good\n")
	return false, nil
}
