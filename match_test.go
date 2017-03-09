package fastaexample

import (
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
		frags, err := ReadFile(v.file)
		if err != nil {
			t.Errorf("TestMatch generated unrelated error in ReadFile\n")
		}
		result, err := Match(frags)
		if v.want && err != nil {
			t.Errorf("Match(%s) produced an error on a correct case: %v\n", v.file, err)
		}
		if v.want && v.sequence != result {
			t.Errorf("Match(%s) expected %s got %s", v.file, v.sequence, result)
		}
		if !v.want && err == nil {
			t.Errorf("Match(%s) failed to produce an error on an incorrect case", v.file)
		}
	}
}
