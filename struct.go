package fastaexample

type FastaFrag struct {
	Title string
	Data  string
	// not necessary with strings, might be needed if we use []byte
	Length int
}

// Exists to hook into main without warnings, remove later
func Noop() error {
	return nil
}
