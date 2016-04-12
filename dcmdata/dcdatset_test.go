package dcmdata

import (
	"github.com/grayzone/godcm/ofstd"
	"testing"
)

func TestNewDcmDataset(t *testing.T) {
	cases := []struct {
		want *DcmDataset
	}{
		{&DcmDataset{DcmItem: *NewDcmItem(DCM_ItemTag, DCM_UndefinedLength), OriginalXfer: EXS_Unknown, CurrentXfer: EXS_LittleEndianExplicit}},
	}
	for _, c := range cases {
		got := NewDcmDataset()
		if *got != *c.want {
			t.Errorf("NewDcmDataset() == want %v got %v", c.want, got)
		}
	}

}

func TestDcdatsetLoadFile(t *testing.T) {
	cases := []struct {
		in_0 DcmDataset
		in_1 string
		in_2 E_TransferSyntax
		in_3 E_GrpLenEncoding
		in_4 uint32
		want ofstd.OFCondition
	}{
		{DcmDataset{}, "", EXS_Unknown, EGL_noChange, DCM_MaxReadLength, ofstd.EC_Normal},
	}
	for _, c := range cases {
		got := c.in_0.LoadFile(c.in_1, c.in_2, c.in_3, c.in_4)

		if got != c.want {
			t.Errorf("LoadFile() == want %v got %v", c.want, got)
		}
	}
}
