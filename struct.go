package fastaexample

type FastaFrag struct {
	Title  string
	Data   string
	Length int
}

// Exists to hook into main without warnings, remove later
func Noop() error {
	return nil
}
