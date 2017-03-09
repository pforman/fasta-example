package fastaexample

import (
	"testing"
)

func TestSetDebug(t *testing.T) {

	cases := []struct {
		val  bool
		want bool
	}{
		{true, true},
		{false, false},
	}

	// Verify we default to false
	if debug {
		t.Errorf("SetDebug() expected default false, got %t\n", debug)
	}
	for _, v := range cases {
		SetDebug(v.val)
		if v.want != debug {
			t.Errorf("SetDebug(%v) expected %v, got %v\n", v.val, v.want, debug)
		}
	}
}
