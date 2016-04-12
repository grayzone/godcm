package dcmdata

import (
	"testing"
)

func TestNewDcmSequenceOfItems(t *testing.T) {
	cases := []struct {
		in_1 DcmTag
		in_2 uint32
		in_3 bool
		want *DcmSequenceOfItems
	}{
		{*NewDcmTag(), 0, true, &DcmSequenceOfItems{DcmElement: *NewDcmElement(*NewDcmTag(), 0), lastItemComplete: true, readAsUN_: true, itemList: nil}},
	}
	for _, c := range cases {
		got := NewDcmSequenceOfItems(c.in_1, c.in_2, c.in_3)
		if *got != *c.want {
			t.Errorf("NewDcmSequenceOfItems(%v,%v,%v) == want %v got %v", c.in_1, c.in_2, c.in_3, c.want, got)
		}
	}
}
