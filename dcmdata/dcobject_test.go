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

func TestDcmObjectTransferState(t *testing.T) {
	cases := []struct {
		in   DcmObject
		want E_TransferState
	}{
		{DcmObject{tag: *NewDcmTag()}, ERW_init},
	}
	for _, c := range cases {
		got := c.in.TransferState()
		if got != c.want {
			t.Errorf("TransferState() == want %v got %v", c.want, got)
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
