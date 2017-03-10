package fastaexample

// fastaFrag contains the information for a fragment, Title and Data
type fastaFrag struct {
	Title string
	Data  string
}

// placeError occurs when assemble() cannot place a fragment
// These are recoverable, and simply indicate we should move on.
type placeError string

// matchError interface implementation
func (s placeError) Error() string {
	return string(s)
}
