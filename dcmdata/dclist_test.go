package dcmdata

import "testing"

func TestNewDcmList(t *testing.T) {
	cases := []struct {
		want DcmList
	}{
		{DcmList{nil, nil, nil, 0}},
	}

	for _, c := range cases {
		got := NewDcmList()

		if *got != c.want {
			t.Errorf("NewDcmList(), want '%v' got '%v'", c.want, got)
		}
	}

}

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

func TestDcmListCard(t *testing.T) {
	cases := []struct {
		in   DcmList
		want uint32
	}{
		{DcmList{}, 0},
	}

	for _, c := range cases {
		got := c.in.Card()
		if got != c.want {
			t.Errorf("Card(), want '%v' got '%v'", c.want, got)
		}
	}

}

func TestDcmListAppend(t *testing.T) {
	cases := []struct {
		in_0   *DcmList
		in_1   *DcmObject
		want_1 *DcmObject
		want_2 uint32
	}{
		{&DcmList{}, nil, nil, 0},
	}

	for _, c := range cases {
		got_1 := c.in_0.Append(c.in_1)
		got_2 := c.in_0.Card()
		if (got_1 != c.want_1) || (got_2 != c.want_2) {
			t.Errorf("%v Append(%v), want '%v' | %v got '%v' | %v ", c.in_0, c.in_1, c.want_1, c.want_2, got_1, got_2)
		}
	}

}

func TestDcmListPrepend(t *testing.T) {
	cases := []struct {
		in_0   *DcmList
		in_1   *DcmObject
		want_1 *DcmObject
		want_2 uint32
	}{
		{&DcmList{}, nil, nil, 0},
	}

	for _, c := range cases {
		got_1 := c.in_0.Prepend(c.in_1)
		got_2 := c.in_0.Card()
		if (got_1 != c.want_1) || (got_2 != c.want_2) {
			t.Errorf("%v Prepend(%v), want '%v' | %v got '%v' | %v ", c.in_0, c.in_1, c.want_1, c.want_2, got_1, got_2)
		}
	}

}

func TestDcmListInsert(t *testing.T) {
	cases := []struct {
		in_0   *DcmList
		in_1   *DcmObject
		in_2   E_ListPos
		want_1 *DcmObject
		want_2 uint32
	}{
		{&DcmList{}, nil, ELP_atpos, nil, 0},
	}

	for _, c := range cases {
		got_1 := c.in_0.Insert(c.in_1, c.in_2)
		got_2 := c.in_0.Card()
		if (got_1 != c.want_1) || (got_2 != c.want_2) {
			t.Errorf("%v Insert(%v,%v), want '%v' | %v got '%v' | %v ", c.in_0, c.in_1, c.in_2, c.want_1, c.want_2, got_1, got_2)
		}
	}
}

func TestDcmListRemove(t *testing.T) {
	cases := []struct {
		in     *DcmList
		want_1 *DcmObject
		want_2 uint32
	}{
		{&DcmList{}, nil, 0},
	}

	for _, c := range cases {
		got_1 := c.in.Remove()
		got_2 := c.in.Card()
		if (got_1 != c.want_1) || (got_2 != c.want_2) {
			t.Errorf("%v Remove(), want '%v' | %v got '%v' | %v ", c.in, c.want_1, c.want_2, got_1, got_2)
		}
	}

}

func TestDcmListGet(t *testing.T) {
	cases := []struct {
		in_0 *DcmList
		in_1 E_ListPos
		want *DcmObject
	}{
		{&DcmList{}, ELP_atpos, nil},
	}

	for _, c := range cases {
		got := c.in_0.Get(c.in_1)
		if got != c.want {
			t.Errorf("%v Get(%v), want '%v' got '%v' ", c.in_0, c.in_1, c.want, got)
		}
	}
}

func TestDcmListSeek(t *testing.T) {
	cases := []struct {
		in_0 *DcmList
		in_1 E_ListPos
		want *DcmObject
	}{
		{&DcmList{}, ELP_first, nil},
		{&DcmList{}, ELP_last, nil},
		{&DcmList{}, ELP_prev, nil},
		{&DcmList{}, ELP_next, nil},
	}

	for _, c := range cases {
		got := c.in_0.Seek(c.in_1)
		if got != c.want {
			t.Errorf("%v Seek(%v), want '%v' got '%v' ", c.in_0, c.in_1, c.want, got)
		}
	}
}

func TestDcmListSeekTo(t *testing.T) {
	cases := []struct {
		in_0 *DcmList
		in_1 uint32
		want *DcmObject
	}{
		{&DcmList{}, 0, nil},
	}

	for _, c := range cases {
		got := c.in_0.Seek_to(c.in_1)
		if got != c.want {
			t.Errorf("%v Seek_to(%v), want '%v' got '%v' ", c.in_0, c.in_1, c.want, got)
		}
	}

}
