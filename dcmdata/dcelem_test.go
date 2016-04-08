package dcmdata

import (
	"github.com/grayzone/godcm/ofstd"
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

func TestDcmElementGetString(t *testing.T) {
	cases := []struct {
		in   string
		want ofstd.OFCondition
	}{
		{"", EC_IllegalCall},
	}
	for _, c := range cases {
		var e DcmElement
		got := e.GetString(c.in)
		if got != c.want {
			t.Errorf("GetString() == want %v got %v", c.want, got)
		}
	}
}
