package dcmdata

import (
	"testing"
)

func TestNewDcmItem(t *testing.T) {
	cases := []struct {
		in_1 DcmTag
		in_2 uint32
		want *DcmItem
	}{
		{*NewDcmTag(), 0, &DcmItem{DcmObject: *NewDcmObject(*NewDcmTag(), 0), elementList: nil, lastElementComplete: true, fStartPosition: 0}},
	}
	for _, c := range cases {
		got := NewDcmItem(c.in_1, c.in_2)
		if *got != *c.want {
			t.Errorf("NewDcmItem(%v,%v) == want %v got %v", c.in_1, c.in_2, c.want, got)
		}
	}
}
