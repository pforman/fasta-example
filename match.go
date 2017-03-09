package fastaexample

import (
	"fmt"
	//"strings"
)

// Match attempts to assemble the array of fragments
// Match utilizes recurseMatch to do the assembly, but hides the implementation
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

// recurseMatch iterates through the list of fragments, attempting to match one
// onto the previously assembled base.  On success, recurse with the matched
// fragment removed from the list.  When the list is empty, assembly is complete.
// If the end of the list is reached with no match, assembly has failed.
func recurseMatch(base *FastaFrag, frags []*FastaFrag) (bool, *FastaFrag) {
	var err error
	var res *FastaFrag

	// If we've reached an empty fragment list, we're done.
	if len(frags) == 0 {
		return true, base
	}

	// TODO: debug info only
	fmt.Printf("rM: len %d, base: %d\n", len(frags), len(base.Data))

	for i, f := range frags {
		if len(f.Data) < len(base.Data) {
			res, err = assemble(base, f, (len(f.Data)/2)+1)
		} else {
			// unlikely, other than on the first match,
			// but matchPairs() is strict about s1 >= s2
			res, err = assemble(f, base, (len(base.Data)/2)+1)
		}
		if err != nil {
			// TODO: quiet this or put it in debug
			fmt.Printf("trouble matching index %d of %d, skipping\n", i, len(frags))
			fmt.Printf("error was: %s\n", err)
			continue
		}
		// continue on with the matched fragment placed on the base
		// and removed from the list
		complete, final := recurseMatch(res, append(frags[:i], frags[i+1:]...))
		if complete {
			return complete, final
		}
	}

	// If we reach this, we have unmatchable fragments.
	return false, nil
}
