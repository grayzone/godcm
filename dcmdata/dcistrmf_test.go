package dcmdata

import (
	"github.com/grayzone/godcm/ofstd"
	"os"
	"testing"
)

func gettestdatafolder() string {
	cur, err := os.Getwd()
	if err != nil {
		return ""
	}
	result := cur + "/../test/data/"
	return result
}

func TestNewDcmFileProducer(t *testing.T) {
	cases := []struct {
		in_1 string
		in_2 int64
		want *DcmFileProducer
	}{
		{"", 0, &DcmFileProducer{status_: ofstd.MakeOFCondition(OFM_dcmdata, 18, ofstd.OF_error, "open : The system cannot find the file specified."), size_: 0}},
		{gettestdatafolder() + "gh220.dcm", 0, &DcmFileProducer{status_: ofstd.EC_Normal, size_: 454}},
		{gettestdatafolder() + "gh223.dcm", 0, &DcmFileProducer{status_: ofstd.EC_Normal, size_: 702}},
		{gettestdatafolder() + "gh133.dcm", 0, &DcmFileProducer{status_: ofstd.EC_Normal, size_: 2980438}},
	}
	for _, c := range cases {
		got := NewDcmFileProducer(c.in_1, c.in_2)

		if (got.size_ != c.want.size_) || (got.status_ != c.want.status_) {
			t.Errorf("NewDcmFileProducer(%v,%v) == want %v got %v", c.in_1, c.in_2, c.want, got)
		}
	}
}
