package dcmdata

import (
	"testing"

	"github.com/grayzone/godcm/ofstd"
)

func TestDcmTagSetVR(t *testing.T) {
	cases := []struct {
		in        DcmVR
		want      DcmVR
		errorFlag ofstd.OFCondition
	}{
		{DcmVR{DcmEVR(EVR_UNKNOWN)}, DcmVR{DcmEVR(EVR_UNKNOWN)}, EC_InvalidTag},
		{DcmVR{DcmEVR(EVR_AE)}, DcmVR{DcmEVR(EVR_AE)}, ofstd.EC_Normal},
	}
	for _, c := range cases {
		var tag DcmTag
		got := tag.SetVR(c.in)
		if (got != c.want) || (tag.errorFlag.Text() != c.errorFlag.Text()) {
			t.Errorf("SetVR(%q) == %q with error flag: %q, want %q and %q", c.in.vr, got.vr, tag.errorFlag.Text(), c.want.vr, c.errorFlag.Text())
		}
	}

}

func TestGetVR(t *testing.T) {
	cases := []struct {
		in   DcmTag
		want DcmVR
	}{
		{DcmTag{DcmVR: DcmVR{DcmEVR(EVR_UNKNOWN)}}, DcmVR{DcmEVR(EVR_UNKNOWN)}},
		{DcmTag{DcmVR: DcmVR{DcmEVR(EVR_AE)}}, DcmVR{DcmEVR(EVR_AE)}},
	}
	for _, c := range cases {
		got := c.in.GetVR()
		if got != c.want {
			t.Errorf("%q GetVR() == %q , want %q ", c.in.GetEVR(), got.vr, c.want.vr)
		}
	}
}

func TestGetEVR(t *testing.T) {
	cases := []struct {
		in   DcmTag
		want DcmEVR
	}{
		{DcmTag{DcmVR: DcmVR{DcmEVR(EVR_UNKNOWN)}}, DcmEVR(EVR_UNKNOWN)},
		{DcmTag{DcmVR: DcmVR{DcmEVR(EVR_AE)}}, DcmEVR(EVR_AE)},
	}
	for _, c := range cases {
		got := c.in.GetEVR()
		if got != c.want {
			t.Errorf("%q GetEVR()== %q , want %q ", c.in.GetEVR(), got, c.want)
		}
	}
}

func TestDcmTagGetVRName(t *testing.T) {
	cases := []struct {
		in   DcmTag
		want string
	}{
		{DcmTag{DcmVR: DcmVR{DcmEVR(EVR_UNKNOWN)}}, "??"},
		{DcmTag{DcmVR: DcmVR{DcmEVR(EVR_AE)}}, "AE"},
	}
	for _, c := range cases {
		got := c.in.GetVRName()
		if got != c.want {
			t.Errorf("%q GetVRName()== %q , want %q ", c.in.GetEVR(), got, c.want)
		}
	}
}

func TestGetGTag(t *testing.T) {
	cases := []struct {
		in   DcmTag
		want uint16
	}{
		{DcmTag{DcmTagKey: DcmTagKey{0x0010, 0x0010}}, 0x0010},
		{DcmTag{DcmTagKey: DcmTagKey{0xFFFF, 0x0010}}, 0xFFFF},
	}
	for _, c := range cases {
		got := c.in.GetGTag()
		if got != c.want {
			t.Errorf("%q GetGTag()== 0x%04x, want 0x%04x ", c.in.GetEVR(), got, c.want)
		}
	}
}

func TestGetETag(t *testing.T) {
	cases := []struct {
		in   DcmTag
		want uint16
	}{
		{DcmTag{DcmTagKey: DcmTagKey{0x0010, 0x001F}}, 0x001F},
		{DcmTag{DcmTagKey: DcmTagKey{0xFFFF, 0x0010}}, 0x0010},
	}
	for _, c := range cases {
		got := c.in.GetETag()
		if got != c.want {
			t.Errorf("%q GetETag()== 0x%04x, want 0x%04x ", c.in.GetEVR(), got, c.want)
		}
	}
}

func TestGetXTag(t *testing.T) {
	cases := []struct {
		in   DcmTag
		want DcmTagKey
	}{
		{DcmTag{DcmTagKey: DcmTagKey{0x0010, 0x001F}}, DcmTagKey{0x0010, 0x001F}},
		{DcmTag{DcmTagKey: DcmTagKey{0xFFFF, 0x0010}}, DcmTagKey{0xFFFF, 0x0010}},
	}
	for _, c := range cases {
		got := c.in.GetXTag()
		if !c.want.Equal(got) {
			t.Errorf("%s GetXTag()== %q, want %q ", c.in.DcmTagKey.ToString(), got.ToString(), c.want.ToString())
		}
	}
}

func TestGetTagName(t *testing.T) {
	cases := []struct {
		in   DcmTag
		want string
	}{
		{DcmTag{DcmTagKey: DcmTagKey{0x0010, 0x001F}}, "Unknown Tag & Data"},
		{DcmTag{DcmTagKey: DcmTagKey{0xFFFF, 0x0010}}, "Unknown Tag & Data"},
	}
	for _, c := range cases {
		got := c.in.GetTagName()
		if got != c.want {
			t.Errorf("GetTagName()== %q, want %q ", got, c.want)
		}
	}
}
