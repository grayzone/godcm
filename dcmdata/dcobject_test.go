package dcmdata

import (
	"testing"

	"github.com/grayzone/godcm/ofstd"
)

func TestNewDcmObject(t *testing.T) {
	cases := []struct {
		in_tag    DcmTag
		in_length uint32
		want      *DcmObject
	}{
		{DcmTag{}, 1, &DcmObject{length: 1, fTransferState: ERW_init, fTransferredBytes: 0, errorFlag: ofstd.EC_Normal}},
	}
	for _, c := range cases {
		got := NewDcmObject(c.in_tag, c.in_length)
		if (got.length != 1) || (got.fTransferState != ERW_init) || (got.fTransferredBytes != 0) || (got.errorFlag != ofstd.EC_Normal) {
			t.Errorf("NewDcmTag() == want %v got %v", c.want, got)
		}
	}
}

func TestDcmObjectGetVR(t *testing.T) {
	cases := []struct {
		in   DcmObject
		want DcmEVR
	}{
		{DcmObject{tag: *NewDcmTag()}, EVR_UNKNOWN},
	}
	for _, c := range cases {
		got := c.in.GetVR()
		if got.String() != c.want.String() {
			t.Errorf("GetVR() == want %v got %v", c.want, got)
		}
	}
}

func TestDcmObjectIsaString(t *testing.T) {
	cases := []struct {
		in   DcmObject
		want bool
	}{
		{DcmObject{tag: *NewDcmTag()}, false},
	}
	for _, c := range cases {
		got := c.in.IsaString()
		if got != c.want {
			t.Errorf("IsaString() == want %v got %v", c.want, got)
		}
	}

}

func TestDcmObjectSetTransferState(t *testing.T) {
	cases := []struct {
		in_0 DcmObject
		in_1 E_TransferState
		want E_TransferState
	}{
		{DcmObject{tag: *NewDcmTag()}, ERW_ready, ERW_ready},
	}
	for _, c := range cases {
		c.in_0.SetTransferState(c.in_1)
		got := c.in_0.GetTransferState()
		if got != c.want {
			t.Errorf("%v SetTransferState(%v) == want %v got %v", c.in_0, c.in_1, c.want, got)
		}
	}
}

func TestDcmObjectGetTransferState(t *testing.T) {
	cases := []struct {
		in   DcmObject
		want E_TransferState
	}{
		{DcmObject{tag: *NewDcmTag()}, ERW_init},
	}
	for _, c := range cases {
		got := c.in.GetTransferState()
		if got != c.want {
			t.Errorf("GetTransferState() == want %v got %v", c.want, got)
		}
	}
}

func TestDcmObjectTransferInit(t *testing.T) {
	cases := []struct {
		in     DcmObject
		want_1 E_TransferState
		want_2 uint32
	}{
		{DcmObject{tag: *NewDcmTag()}, ERW_init, 0},
	}
	for _, c := range cases {
		c.in.TransferInit()
		if (c.in.fTransferState != c.want_1) || (c.in.fTransferredBytes != c.want_2) {
			t.Errorf("TransferInit() == want %v|%v got %v|%v", c.want_1, c.want_2, c.in.fTransferState, c.in.fTransferredBytes)
		}
	}
}

func TestDcmObjectTransferEnd(t *testing.T) {
	cases := []struct {
		in   DcmObject
		want E_TransferState
	}{
		{DcmObject{tag: *NewDcmTag()}, ERW_notInitialized},
	}
	for _, c := range cases {
		c.in.TransferEnd()
		if c.in.fTransferState != c.want {
			t.Errorf("TransferEnd() == want %v got %v", c.want, c.in.fTransferState)
		}
	}
}

func TestDcmObjectContainsUnknownVR(t *testing.T) {
	cases := []struct {
		in   DcmObject
		want bool
	}{
		{DcmObject{tag: *NewDcmTag()}, true},
	}
	for _, c := range cases {
		got := c.in.ContainsUnknownVR()
		if got != c.want {
			t.Errorf("ContainsUnknownVR() == want %v got %v", c.want, got)
		}
	}
}

func TestDcmObjectGetGTag(t *testing.T) {
	cases := []struct {
		in   DcmObject
		want uint16
	}{
		{DcmObject{tag: *NewDcmTag()}, 0xffff},
	}
	for _, c := range cases {
		got := c.in.GetGTag()
		if got != c.want {
			t.Errorf("GetGTag() == want %v got %v", c.want, got)
		}
	}
}

func TestDcmObjectGetETag(t *testing.T) {
	cases := []struct {
		in   DcmObject
		want uint16
	}{
		{DcmObject{tag: *NewDcmTag()}, 0xffff},
	}
	for _, c := range cases {
		got := c.in.GetETag()
		if got != c.want {
			t.Errorf("GetETag() == want %v got %v", c.want, got)
		}
	}
}

func TestDcmObjectGetTag(t *testing.T) {
	cases := []struct {
		in   DcmObject
		want DcmTag
	}{
		{DcmObject{tag: *NewDcmTag()}, DcmTag{DcmTagKey: DcmTagKey{0xffff, 0xffff}}},
	}
	for _, c := range cases {
		got := c.in.GetTag()
		if got.DcmTagKey != c.want.DcmTagKey {
			t.Errorf("GetTag() == want %v got %v", c.want, got)
		}
	}
}

func TestDcmObjectIsEmpty(t *testing.T) {
	cases := []struct {
		in   bool
		want bool
	}{
		{true, true},
		{false, true},
	}
	for _, c := range cases {
		var obj DcmObject
		got := obj.IsEmpty(c.in)
		if got != c.want {
			t.Errorf("IsEmpty() == want %v got %v", c.want, got)
		}
	}
}

func TestDcmObjectIsLeaf(t *testing.T) {
	cases := []struct {
		in   DcmObject
		want bool
	}{
		{DcmObject{tag: *NewDcmTag()}, false},
	}
	for _, c := range cases {
		got := c.in.IsLeaf()
		if got != c.want {
			t.Errorf("IsLeaf() == want %v got %v", c.want, got)
		}
	}
}

func TestDcmObjectGetLengthField(t *testing.T) {
	cases := []struct {
		in   DcmObject
		want uint32
	}{
		{DcmObject{tag: *NewDcmTag()}, 0},
	}
	for _, c := range cases {
		got := c.in.GetLengthField()
		if got != c.want {
			t.Errorf("GetLengthField() == want %v got %v", c.want, got)
		}
	}
}

func TestDcmObjectCalcElementLength(t *testing.T) {
	cases := []struct {
		in_1 E_TransferSyntax
		in_2 E_EncodingType
		want uint32
	}{
		{EXS_Unknown, EET_ExplicitLength, 0},
	}
	for _, c := range cases {
		var obj DcmObject
		got := obj.CalcElementLength(c.in_1, c.in_2)
		if got != c.want {
			t.Errorf("CalcElementLength() == want %v got %v", c.want, got)
		}
	}
}

func TestDcmObjectIdent(t *testing.T) {
	cases := []struct {
		in   DcmObject
		want DcmEVR
	}{
		{DcmObject{tag: *NewDcmTag()}, EVR_UNKNOWN},
	}
	for _, c := range cases {
		got := c.in.Ident()
		if got != c.want {
			t.Errorf("Ident() == want %v got %v", c.want, got)
		}
	}
}
