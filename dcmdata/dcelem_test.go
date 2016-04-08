package dcmdata

import (
	"testing"
)

func TestNewDcmElement(t *testing.T) {
	cases := []struct {
		in_1 DcmTag
		in_2 uint32
		want *DcmElement
	}{
		{*NewDcmTag(), 0, &DcmElement{fByteOrder: GLocalByteOrder}},
	}
	for _, c := range cases {
		got := NewDcmElement(c.in_1, c.in_2)
		if got.fByteOrder != c.want.fByteOrder {
			t.Errorf("NewDcmElement() == want %v got %v", c.want, got)
		}
	}
}
