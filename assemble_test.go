package fastaexample

import (
	"testing"
)

func TestWrongOrderMatch(t *testing.T) {

	s1 := FastaFrag{
		Title: "test1",
		Data:  "ACTGGTCA",
	}

	s2 := FastaFrag{
		Title: "test2",
		Data:  "CTGG",
	}

	f, err := assemble(&s2, &s1, 4)
	if err == nil && f != nil {
		t.Errorf("assemble failed to produce an error on an incorrect case")
	}
}

func TestShortMatch(t *testing.T) {

	s1 := FastaFrag{
		Title: "test1",
		Data:  "ACTGGTCA",
	}

	s2 := FastaFrag{
		Title: "test2",
		Data:  "CTGG",
	}

	f, err := assemble(&s2, &s1, 5)
	if err == nil && f != nil {
		t.Errorf("assemble failed to produce an error on an incorrect case")
	}
}

func TestContainsMatch(t *testing.T) {

	s1 := FastaFrag{
		Title: "test1",
		Data:  "ACTGGTCA",
	}

	s2 := FastaFrag{
		Title: "test2",
		Data:  "CTGG",
	}

	f, err := assemble(&s1, &s2, 4)
	if err != nil {
		t.Errorf("assemble produced an error on a correct case: %v", err)
	}
	if f.Data != s1.Data {
		t.Errorf("assemble expected %s, got %s", s1.Data, f.Data)
	}
}

func TestPrefixMatch(t *testing.T) {

	s1 := FastaFrag{
		Title: "test1",
		Data:  "ACTGGTCAAGGTCA",
	}

	s2 := FastaFrag{
		Title: "test2",
		Data:  "AAGGTCAGG",
	}

	expect := "ACTGGTCAAGGTCAGG"

	f, err := assemble(&s1, &s2, 4)
	if err != nil {
		t.Errorf("assemble produced an error on a correct case: %v", err)
	}
	if f.Data != expect {
		t.Errorf("assemble expected %s, got %s", expect, f.Data)
	}
	// t.Logf("found %s, expected %s, seems legit", f.Data, expect)
}

func TestFalsePrefixMatch(t *testing.T) {

	// Close but not quite.  Matches threshold, but fails a complete match
	// AAGGTCG.. != ..AAGGTCA
	s1 := FastaFrag{
		Title: "test1",
		Data:  "ACTGGTCAAGGTCA",
	}

	s2 := FastaFrag{
		Title: "test2",
		Data:  "AAGGTCGGG",
	}

	f, err := assemble(&s1, &s2, 5)
	if err == nil && f != nil {
		t.Errorf("assemble failed to produce an error on an incorrect case")
	}
}

func TestSuffixMatch(t *testing.T) {

	s1 := FastaFrag{
		Title: "test1",
		Data:  "ACTGGTCAAGGTCG",
	}

	s2 := FastaFrag{
		Title: "test2",
		Data:  "AAGGTCACTGG",
	}

	expect := "AAGGTCACTGGTCAAGGTCG"

	f, err := assemble(&s1, &s2, 4)
	if err != nil {
		t.Errorf("assemble produced an error on a correct case: %v", err)
	}
	if f.Data != expect {
		t.Errorf("assemble expected %s, got %s", expect, f.Data)
	}
	// t.Logf("found %s, expected %s, seems legit", f.Data, expect)
}

func TestFalseSuffixMatch(t *testing.T) {

	// Close but not quite.  Matches threshold, but fails a complete match
	// TACTGG.. != ..CACTGG
	s1 := FastaFrag{
		Title: "test1",
		Data:  "TACTGGTCAAGGTCG",
	}

	s2 := FastaFrag{
		Title: "test2",
		Data:  "AAGGTCACTGG",
	}

	f, err := assemble(&s1, &s2, 4)
	if err == nil && f != nil {
		t.Errorf("assemble failed to produce an error on an incorrect case")
	}
}
