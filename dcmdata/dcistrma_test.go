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
