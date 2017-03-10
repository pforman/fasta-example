package fastaexample

import (
	"testing"
)

func TestAssembleFile(t *testing.T) {

	cases := []struct {
		file string
		want bool
	}{
		{"testdata/1clean.txt", true},
		{"testdata/2dirty.txt", false},
		{"testdata/match-success", true},
		{"testdata/match-fail", false},
	}
	for _, v := range cases {
		_, err := AssembleFile(v.file)
		if v.want && err != nil {
			t.Errorf("readFile(%s) produced an error on a correct case: %v\n", v.file, err)
		}
		if !v.want && err == nil {
			t.Errorf("readFile(%s) failed to produce an error on an incorrect case", v.file)
		}
	}
}

func TestReadFile(t *testing.T) {

	cases := []struct {
		file   string
		length int
		want   bool
	}{
		{"testdata/1clean.txt", 1, true},
		{"testdata/2dirty.txt", 0, false},
		{"testdata/3trailingspace.txt", 0, false},
		{"testdata/4blanklines.txt", 4, true},
		{"testdata/5multiline.txt", 5, true},
		{"testdata/6emptysequence.txt", 0, false},
		// this file should not exist
		{"testdata/nonexistant", 0, false},
	}
	for _, v := range cases {
		result, err := readFile(v.file)
		if v.want && err != nil {
			t.Errorf("readFile(%s) produced an error on a correct case: %v\n", v.file, err)
		}
		if v.want && v.length != len(result) {
			t.Errorf("readFile(%s) wanted %d successes, got %d", v.file, v.length, len(result))
		}
		if !v.want && err == nil {
			t.Errorf("readFile(%s) failed to produce an error on an incorrect case", v.file)
		}
	}
}

func TestFragFromChunk(t *testing.T) {

	cases := []struct {
		chunk string
		want  bool
	}{
		{"ok\nGATT", true},
		{"singleline", false},
		{"multiline\nACGTACGT\nACGTG\n", true},
		{"dirty\nACGTAXGC\nACGT\n", false},
	}
	for _, v := range cases {
		_, err := fragFromChunk(v.chunk)
		if v.want && err != nil {
			t.Errorf("fragFromChunk(%s) produced an error on a correct case: %v\n", v.chunk, err)
		}
		if !v.want && err == nil {
			t.Errorf("fragFromChunk(%s) failed to produce an error on an incorrect case", v.chunk)
		}
	}
}

func TestSanityCheckSequence(t *testing.T) {

	cases := []struct {
		data string
		want bool
	}{
		{"AGTC", true},
		{"AAAAAAAAAAAAAA", true},
		{"AGTCAGTTTCAGTC", true},
		{"AXGTC", false},
		{"XAGTC$..GTC", false},
	}
	for _, v := range cases {
		success := sanityCheckSequence(v.data)
		// t.Logf("%: wanted %b got: %b\n", v.data, v.want, success)
		if !success && v.want {
			t.Errorf("sanityCheckSequence(%s) failed on a correct case", v.data)
		}
		if success && !v.want {
			t.Errorf("sanityCheckSequence(%s) incorrectly succeeded", v.data)
		}
	}
}
