package dcmdata

import (
	"testing"
)

func TestNewDcmInputStream(t *testing.T) {
	cases := []struct {
		in   *DcmProducer
		want *DcmInputStream
	}{
		{new(DcmProducer), &DcmInputStream{}},
	}
	for _, c := range cases {
		got := NewDcmInputStream(c.in)
		if *got != *c.want {
			t.Errorf("NewDcmInputStream(%v) == want %v got %v", c.in, c.want, got)
		}
	}

}
