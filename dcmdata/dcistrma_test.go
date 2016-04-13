package dcmdata

import (
	"testing"
)

func TestNewDcmInputStream(t *testing.T) {
	cases := []struct {
		want *DcmInputStream
	}{
		{&DcmInputStream{}},
	}
	for _, c := range cases {
		got := NewDcmInputStream()
		if *got != *c.want {
			t.Errorf("NewDcmInputStream() == want %v got %v", c.want, got)
		}
	}

}

/*
func TestDcmInputStreamGood(t *testing.T) {
	cases := []struct {
		in   *DcmInputStream
		want bool
	}{
		{NewDcmInputStream(), false},
	}
	for _, c := range cases {
		got := c.in.Good()

		if got != c.want {
			t.Errorf(" %v Good() == want %v got %v", c.in, c.want, got)
		}
	}

}
*/
