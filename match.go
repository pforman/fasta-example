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

	final, err := recurseMatch(base, frags)
	if err != nil {
		return "", err
	}
	if debug {
		logger.Printf("final path: %v\n", final.Title)
	}
	return final.Data, nil
}

// recurseMatch iterates through the list of fragments, attempting to match one
// onto the previously assembled base.  On success, recurse with the matched
// fragment removed from the list.  When the list is empty, assembly is complete.
// If the end of the list is reached with no match, assembly has failed.
func recurseMatch(base *fastaFrag, frags []*fastaFrag) (*fastaFrag, error) {
	var err error
	var res *fastaFrag

	// If we've reached an empty fragment list, we're done.
	if len(frags) == 0 {
		return base, nil
	}

	if debug {
		logger.Printf("recurseMatch: frags %d, current base length: %d\n", len(frags), len(base.Data))
	}

	for i, f := range frags {
		if len(f.Data) < len(base.Data) {
			res, err = assemble(base, f, (len(f.Data)/2)+1)
		} else {
			// Unlikely, other than on the first match,
			// but assemble() is strict about s1 >= s2
			res, err = assemble(f, base, (len(base.Data)/2)+1)
		}
		if err != nil {
			// Check if this is a simple failure to place the fragment
			e, ok := err.(placeError)
			if ok {
				if debug {
					logger.Printf("%s: index %d of %d cannot be placed yet, skipping...\n", e, i, len(frags))
				}
				continue
			}
			// If it's not a placeError, it's a pretty serious error that won't get
			// better later on.  Get out.
			// This is very hard to reach in test coverage, the rest of the code
			// guards against it ever happening.
			return nil, err
		}
		// Continue on with the matched fragment placed on the base
		// and removed from the list.
		final, err := recurseMatch(res, append(frags[:i], frags[i+1:]...))
		// Whatever we get back, pass it upward
		return final, err
	}

	// If we reach this, we have unmatchable fragments.
	return nil, fmt.Errorf("unmatchable fragments, full assembly is impossible")
}
