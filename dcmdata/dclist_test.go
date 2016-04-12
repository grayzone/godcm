package dcmdata

import "testing"

func TestNewDcmListNode(t *testing.T) {
	cases := []struct {
		in   *DcmObject
		want *DcmListNode
	}{
		{nil, &DcmListNode{nil, nil, nil}},
	}

	for _, c := range cases {
		got := NewDcmListNode(c.in)

		if *got != *c.want {
			t.Errorf("NewDcmListNode(), want '%v' got '%v'", c.want, got)
		}
	}
}

func TestDcmListNodeValue(t *testing.T) {
	cases := []struct {
		in   *DcmListNode
		want *DcmObject
	}{
		{&DcmListNode{}, nil},
	}

	for _, c := range cases {
		got := c.in.Value()

		if got != c.want {
			t.Errorf("Value(), want '%v' got '%v'", c.want, got)
		}
	}
}

func TestEmpty(t *testing.T) {
	cases := []struct {
		in   DcmList
		want bool
	}{
		{DcmList{}, true},
	}

	for _, c := range cases {

		got := c.in.Empty()
		if got != c.want {
			t.Errorf("Empty(), want '%v' got '%v'", c.want, got)
		}
	}

}

func TestDcmListValid(t *testing.T) {
	cases := []struct {
		in   DcmList
		want bool
	}{
		{DcmList{}, false},
	}

	for _, c := range cases {
		got := c.in.Valid()
		if got != c.want {
			t.Errorf("Valid(), want '%v' got '%v'", c.want, got)
		}
	}

}
