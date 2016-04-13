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
		{gettestdatafolder() + "GH220.dcm", 0, &DcmFileProducer{status_: ofstd.EC_Normal, size_: 454}},
		{gettestdatafolder() + "GH223.dcm", 0, &DcmFileProducer{status_: ofstd.EC_Normal, size_: 702}},
		{gettestdatafolder() + "GH133.dcm", 0, &DcmFileProducer{status_: ofstd.EC_Normal, size_: 2980438}},
	}
	for _, c := range cases {
		got := NewDcmFileProducer(c.in_1, c.in_2)
		defer got.Close()
		defer c.want.Close()

		if (got.size_ != c.want.size_) || (got.status_.Status() != c.want.status_.Status()) {
			t.Errorf("NewDcmFileProducer(%v,%v) == want %v got %v", c.in_1, c.in_2, c.want, got)
		}
	}
}

func TestDcmFileProducerGood(t *testing.T) {
	cases := []struct {
		in   *DcmFileProducer
		want bool
	}{
		{NewDcmFileProducer("", 0), false},
		{NewDcmFileProducer(gettestdatafolder()+"GH220.dcm", 0), true},
		{NewDcmFileProducer(gettestdatafolder()+"GH223.dcm", 0), true},
		{NewDcmFileProducer(gettestdatafolder()+"GH133.dcm", 0), true},
	}
	for _, c := range cases {
		got := c.in.Good()
		defer c.in.Close()

		if got != c.want {
			t.Errorf(" %v Good() == want %v got %v", c.in, c.want, got)
		}
	}

}

func TestDcmFileProducerStatus(t *testing.T) {
	cases := []struct {
		in   *DcmFileProducer
		want ofstd.OFCondition
	}{
		{NewDcmFileProducer("", 0), ofstd.MakeOFCondition(OFM_dcmdata, 18, ofstd.OF_error, "open : The system cannot find the file specified.")},
		{NewDcmFileProducer(gettestdatafolder()+"GH220.dcm", 0), ofstd.EC_Normal},
		{NewDcmFileProducer(gettestdatafolder()+"GH223.dcm", 0), ofstd.EC_Normal},
		{NewDcmFileProducer(gettestdatafolder()+"GH133.dcm", 0), ofstd.EC_Normal},
	}
	for _, c := range cases {
		got := c.in.Status()
		defer c.in.Close()

		if got.Status() != c.want.Status() {
			t.Errorf(" %v Status() == want %v got %v", c.in, c.want, got)
		}
	}

}

func TestDcmFileProducerEos(t *testing.T) {
	cases := []struct {
		in   *DcmFileProducer
		want bool
	}{
		{NewDcmFileProducer("", 0), true},
		{NewDcmFileProducer(gettestdatafolder()+"GH220.dcm", 0), false},
		{NewDcmFileProducer(gettestdatafolder()+"GH223.dcm", 0), false},
		{NewDcmFileProducer(gettestdatafolder()+"GH133.dcm", 0), false},
		{NewDcmFileProducer(gettestdatafolder()+"GH220.dcm", 454), true},
		{NewDcmFileProducer(gettestdatafolder()+"GH223.dcm", 702), true},
		{NewDcmFileProducer(gettestdatafolder()+"GH133.dcm", 2980438), true},
	}
	for _, c := range cases {
		got := c.in.Eos()
		defer c.in.Close()

		if got != c.want {
			t.Errorf(" %v Eos() == want %v got %v", c.in, c.want, got)
		}
	}

}
