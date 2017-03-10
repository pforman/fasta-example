package fastaexample

import (
	"fmt"
)

// Match attempts to assemble the array of fragments
// Match utilizes recurseMatch to do the assembly, but hides the implementation
func match(frags []*fastaFrag) (string, error) {

	// begin at the beginning
	base := frags[0]
	frags = frags[1:]

	complete, final := recurseMatch(base, frags)
	if complete {
		if debug {
			logger.Printf("final path: %v\n", final.Title)
		}
		return final.Data, nil
	}

	return "", fmt.Errorf("no full match found")
}

// recurseMatch iterates through the list of fragments, attempting to match one
// onto the previously assembled base.  On success, recurse with the matched
// fragment removed from the list.  When the list is empty, assembly is complete.
// If the end of the list is reached with no match, assembly has failed.
func recurseMatch(base *fastaFrag, frags []*fastaFrag) (bool, *fastaFrag) {
	var err error
	var res *fastaFrag

	// If we've reached an empty fragment list, we're done.
	if len(frags) == 0 {
		return true, base
	}

	if debug {
		//fmt.Fprintf(os.Stderr, "recurseMatch: frags %d, current base length: %d\n", len(frags), len(base.Data))
		logger.Printf("recurseMatch: frags %d, current base length: %d\n", len(frags), len(base.Data))
	}

	for i, f := range frags {
		if len(f.Data) < len(base.Data) {
			res, err = assemble(base, f, (len(f.Data)/2)+1)
		} else {
			// unlikely, other than on the first match,
			// but matchPairs() is strict about s1 >= s2
			res, err = assemble(f, base, (len(base.Data)/2)+1)
		}
		if err != nil {
			if debug {
				// TODO: error-type this or signal with nil
				//fmt.Fprintf(os.Stderr, "trouble matching index %d of %d, skipping\n", i, len(frags))
				//fmt.Fprintf(os.Stderr, "error was: %s\n", err)
				logger.Printf("trouble matching index %d of %d, skipping\n", i, len(frags))
				logger.Printf("error was: %s\n", err)
			}
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
