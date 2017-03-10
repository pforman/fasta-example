package fastaexample

import (
	"os"
	"testing"
)

func TestMatch(t *testing.T) {

	cases := []struct {
		file     string
		sequence string
		want     bool
	}{
		{"testdata/match-success", "ATTAGACCTGCCGGAATAC", true},
		{"testdata/match-fail", "", false},
	}
	for _, v := range cases {
		frags, err := readFile(v.file)
		if err != nil {
			t.Errorf("TestMatch generated unrelated error in ReadFile\n")
		}
		result, err := match(frags)
		if v.want && err != nil {
			t.Errorf("match(%s) produced an error on a correct case: %v\n", v.file, err)
		}
		if v.want && v.sequence != result {
			t.Errorf("match(%s) expected %s got %s", v.file, v.sequence, result)
		}
		if !v.want && err == nil {
			t.Errorf("match(%s) failed to produce an error on an incorrect case", v.file)
		}
	}
}

func TestRecurseMatchWithBadData(t *testing.T) {

	base := fastaFrag{
		Title: "base",
		Data:  "ACGTTACGTTACG",
	}
	// in order to induce a failure in assemble, we use an empty string
	// recurseMatch will pass that with a threshold of len/2 +1 = 1.
	frags := []*fastaFrag{
		{
			Title: "good",
			Data:  "AACGTTGA",
		},
		{
			Title: "bad",
			Data:  "",
		},
	}

	_, err := recurseMatch(&base, frags)
	if err == nil {
		t.Errorf("TestRecurseMatch failed to generate an error on bad data\n")
	}
}

func TestMatchWithDebug(t *testing.T) {

	cases := []struct {
		file     string
		sequence string
		want     bool
	}{
		{"testdata/match-success", "ATTAGACCTGCCGGAATAC", true},
		{"testdata/match-fail", "", false},
	}
	// Turn debug on to cover those code paths
	SetDebug(true)
	// Quiet the logger during testing
	devnull, _ := os.Open("/dev/null")
	logger.SetOutput(devnull)
	for _, v := range cases {
		frags, err := readFile(v.file)
		if err != nil {
			t.Errorf("TestMatch generated unrelated error in ReadFile\n")
		}
		result, err := match(frags)
		if v.want && err != nil {
			t.Errorf("match(%s) produced an error on a correct case: %v\n", v.file, err)
		}
		if v.want && v.sequence != result {
			t.Errorf("match(%s) expected %s got %s", v.file, v.sequence, result)
		}
		if !v.want && err == nil {
			t.Errorf("match(%s) failed to produce an error on an incorrect case", v.file)
		}
	}
	// Turn the logger back on
	logger.SetOutput(os.Stderr)
}
