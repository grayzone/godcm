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

func TestEmpty(t *testing.T) {
	cases := []struct {
		want bool
	}{
		{true},
	}

	for _, c := range cases {
		var l DcmList
		//		t.Log(l)
		got := l.Empty()
		if got != c.want {
			t.Errorf("Empty(), want '%v' got '%v'", c.want, got)
		}
	}

}
