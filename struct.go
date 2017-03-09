package fastaexample

// FastaFrag contains the information for a fragment, Title and Data
type FastaFrag struct {
	Title string
	Data  string
}

// Noop exists to hook into main without warnings, remove later
func Noop() error {
	return nil
}
