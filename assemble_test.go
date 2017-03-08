package fastaexample

import (
	"testing"
)

func TestWrongOrderMatch(t *testing.T) {

	s1 := FastaFrag{
		Title:  "test1",
		Data:   "ACTGGTCA",
		Length: 8,
	}

	s2 := FastaFrag{
		Title:  "test2",
		Data:   "CTGG",
		Length: 4,
	}

	f, err := matchPairs(&s2, &s1, 4)
	if err == nil && f != nil {
		t.Errorf("matchPairs failed to produce an error on an incorrect case")
	}
}

func TestShortMatch(t *testing.T) {

	s1 := FastaFrag{
		Title:  "test1",
		Data:   "ACTGGTCA",
		Length: 8,
	}

	s2 := FastaFrag{
		Title:  "test2",
		Data:   "CTGG",
		Length: 4,
	}

	f, err := matchPairs(&s2, &s1, 5)
	if err == nil && f != nil {
		t.Errorf("matchPairs failed to produce an error on an incorrect case")
	}
}

func TestContainsMatch(t *testing.T) {

	s1 := FastaFrag{
		Title:  "test1",
		Data:   "ACTGGTCA",
		Length: 8,
	}

	s2 := FastaFrag{
		Title:  "test2",
		Data:   "CTGG",
		Length: 4,
	}

	f, err := matchPairs(&s1, &s2, 4)
	if err != nil {
		t.Errorf("matchPairs produced an error on a correct case: %v", err)
	}
	if f.Data != s1.Data {
		t.Errorf("matchPairs expected %s, got %s", s1.Data, f.Data)
	}
}

func TestPrefixMatch(t *testing.T) {

	s1 := FastaFrag{
		Title:  "test1",
		Data:   "ACTGGTCAAGGTCA",
		Length: 14,
	}

	s2 := FastaFrag{
		Title:  "test2",
		Data:   "AAGGTCAGG",
		Length: 9,
	}

	expect := "ACTGGTCAAGGTCAGG"

	f, err := matchPairs(&s1, &s2, 4)
	if err != nil {
		t.Errorf("matchPairs produced an error on a correct case: %v", err)
	}
	if f.Data != expect {
		t.Errorf("matchPairs expected %s, got %s", expect, f.Data)
	}
	// t.Logf("found %s, expected %s, seems legit", f.Data, expect)
}

func TestFalsePrefixMatch(t *testing.T) {

	// Close but not quite.  Matches threshold, but fails a complete match
	// AAGGTCG.. != ..AAGGTCA
	s1 := FastaFrag{
		Title:  "test1",
		Data:   "ACTGGTCAAGGTCA",
		Length: 14,
	}

	s2 := FastaFrag{
		Title:  "test2",
		Data:   "AAGGTCGGG",
		Length: 9,
	}

	f, err := matchPairs(&s1, &s2, 5)
	if err == nil && f != nil {
		t.Errorf("matchPairs failed to produce an error on an incorrect case")
	}
}

func TestSuffixMatch(t *testing.T) {

	s1 := FastaFrag{
		Title:  "test1",
		Data:   "ACTGGTCAAGGTCG",
		Length: 14,
	}

	s2 := FastaFrag{
		Title:  "test2",
		Data:   "AAGGTCACTGG",
		Length: 11,
	}

	expect := "AAGGTCACTGGTCAAGGTCG"

	f, err := matchPairs(&s1, &s2, 4)
	if err != nil {
		t.Errorf("matchPairs produced an error on a correct case: %v", err)
	}
	if f.Data != expect {
		t.Errorf("matchPairs expected %s, got %s", expect, f.Data)
	}
	// t.Logf("found %s, expected %s, seems legit", f.Data, expect)
}

func TestFalseSuffixMatch(t *testing.T) {

	// Close but not quite.  Matches threshold, but fails a complete match
	// TACTGG.. != ..CACTGG
	s1 := FastaFrag{
		Title:  "test1",
		Data:   "TACTGGTCAAGGTCG",
		Length: 14,
	}

	s2 := FastaFrag{
		Title:  "test2",
		Data:   "AAGGTCACTGG",
		Length: 11,
	}

	f, err := matchPairs(&s1, &s2, 4)
	if err == nil && f != nil {
		t.Errorf("matchPairs failed to produce an error on an incorrect case")
	}
}
