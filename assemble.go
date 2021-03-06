package fastaexample

import (
	"fmt"
	"strings"
)

// assemble attempts to match s2 onto s1, with a match of at least th in
// length.  s1 must be at least the length of s2, not shorter.
func assemble(s1, s2 *fastaFrag, th int) (*fastaFrag, error) {

	// Safety check
	if len(s1.Data) < len(s2.Data) {
		return nil, fmt.Errorf("assemble called with misordered arguments, len(s1) < len(s2)")
	}

	// Threshold check
	if len(s2.Data) < th {
		return nil, fmt.Errorf("sequence '%s' in fragment %s is shorter than threshold %d", s2.Data, s2.Title, th)
	}

	// Check for fully contained match first
	if strings.Contains(s1.Data, s2.Data) {
		return s1, nil
	}

	prefix := s2.Data[:th]
	suffix := s2.Data[len(s2.Data)-th:]

	// Check for a prefix match, ie S1..S2
	i := strings.LastIndex(s1.Data, prefix)
	if i != -1 {
		// check for a full match, based on finding the threshold
		match := s1.Data[i:]
		if strings.HasPrefix(s2.Data, match) {
			// Concat the titles and the data
			title := fmt.Sprintf("%s+%s", s1.Title, s2.Title)
			data := s1.Data[:i] + s2.Data

			return &fastaFrag{
				Title: title,
				Data:  data,
			}, nil
		}
	}

	// Check for a suffix match, ie S2..S1
	i = strings.Index(s1.Data, suffix)
	if i != -1 {
		// check for a full match, based on finding the threshold
		match := s1.Data[:i] + suffix
		if strings.HasSuffix(s2.Data, match) {
			// Concat the titles and the data
			title := fmt.Sprintf("%s+%s", s2.Title, s1.Title)
			data := s2.Data[:len(s2.Data)-len(match)] + s1.Data

			return &fastaFrag{
				Title: title,
				Data:  data,
			}, nil
		}
	}

	// This is a non-fatal error, we can try to place this fragment later
	return nil, placeError("placement failed")
}
