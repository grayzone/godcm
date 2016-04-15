package dcmdata

import (
	"os"
	"reflect"
	"testing"

	"github.com/grayzone/godcm/ofstd"
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

func TestDcmFileProducerAvail(t *testing.T) {
	cases := []struct {
		in   *DcmFileProducer
		want int64
	}{
		{NewDcmFileProducer("", 0), 0},
		{NewDcmFileProducer(gettestdatafolder()+"GH220.dcm", 0), 454},
		{NewDcmFileProducer(gettestdatafolder()+"GH223.dcm", 0), 702},
		{NewDcmFileProducer(gettestdatafolder()+"GH133.dcm", 0), 2980438},
		{NewDcmFileProducer(gettestdatafolder()+"GH220.dcm", 10), 444},
		{NewDcmFileProducer(gettestdatafolder()+"GH223.dcm", 100), 602},
		{NewDcmFileProducer(gettestdatafolder()+"GH133.dcm", 1000), 2979438},
	}
	for _, c := range cases {
		got := c.in.Avail()
		defer c.in.Close()

		if got != c.want {
			t.Errorf(" %v Avail() == want %v got %v", c.in, c.want, got)
		}
	}

}

func TestDcmFileProducerRead(t *testing.T) {
	cases := []struct {
		in_0   *DcmFileProducer
		in_1   int64
		want_1 []byte
		want_2 int64
	}{
		{NewDcmFileProducer("", 0), 0, nil, 0},
		{NewDcmFileProducer(gettestdatafolder()+"GH184.dcm", 0), 1, []byte{0x52}, 1},
		{NewDcmFileProducer(gettestdatafolder()+"GH184.dcm", 0), 2, []byte{0x52, 0x5f}, 2},
		{NewDcmFileProducer(gettestdatafolder()+"GH184.dcm", 0), 3, []byte{0x52, 0x5f, 0x31}, 3},
		{NewDcmFileProducer(gettestdatafolder()+"GH184.dcm", 0), 4, []byte{0x52, 0x5f, 0x31, 0x46}, 4},
	}
	for _, c := range cases {
		got_1, got_2 := c.in_0.Read(c.in_1)
		defer c.in_0.Close()

		if !reflect.DeepEqual(got_1, c.want_1) || got_2 != c.want_2 {
			t.Errorf(" %v Read(%v) == want %x | %v got %x | %v", c.in_0, c.in_1, c.want_1, c.want_2, got_1, got_2)
		}
	}

}

func TestDcmFileProducerSkip(t *testing.T) {
	cases := []struct {
		in_0 *DcmFileProducer
		in_1 int64
		want int64
	}{
		{NewDcmFileProducer("", 0), 0, 0},
		{NewDcmFileProducer(gettestdatafolder()+"GH220.dcm", 0), 1, 1},
		{NewDcmFileProducer(gettestdatafolder()+"GH220.dcm", 0), 454, 454},
		{NewDcmFileProducer(gettestdatafolder()+"GH220.dcm", 0), 453, 453},
		{NewDcmFileProducer(gettestdatafolder()+"GH220.dcm", 0), 455, 454},
	}
	for _, c := range cases {
		got := c.in_0.Skip(c.in_1)
		defer c.in_0.Close()

		if got != c.want {
			t.Errorf(" %v Skip(%v) == want %v got %v ", c.in_0, c.in_1, c.want, got)
		}
	}

}

func TestDcmFileProducerPutback(t *testing.T) {
	cases := []struct {
		in_0 *DcmFileProducer
		in_1 int64
		want int64
	}{
		{NewDcmFileProducer("", 0), 0, 0},
		{NewDcmFileProducer(gettestdatafolder()+"GH220.dcm", 100), 1, 99},
		{NewDcmFileProducer(gettestdatafolder()+"GH220.dcm", 0), 454, 0},
		{NewDcmFileProducer(gettestdatafolder()+"GH220.dcm", 100), 453, 100},
		{NewDcmFileProducer(gettestdatafolder()+"GH220.dcm", 10), 455, 10},
		{NewDcmFileProducer(gettestdatafolder()+"GH220.dcm", 100), 10, 90},
		{NewDcmFileProducer(gettestdatafolder()+"GH220.dcm", 10), 9, 1},
	}
	for _, c := range cases {
		c.in_0.Putback(c.in_1)
		got, _ := c.in_0.file_.Seek(0, os.SEEK_CUR)
		defer c.in_0.Close()

		if got != c.want {
			t.Errorf(" %v Putback(%v) == want %v got %v ", c.in_0, c.in_1, c.want, got)
		}
	}

}

func TestNewDcmInputFileStream(t *testing.T) {
	cases := []struct {
		in_1 string
		in_2 int64
		want *DcmInputFileStream
	}{
		{"", 0, &DcmInputFileStream{status: ofstd.MakeOFCondition(OFM_dcmdata, 18, ofstd.OF_error, "open : The system cannot find the file specified."), size: 0}},
		{gettestdatafolder() + "GH220.dcm", 0, &DcmInputFileStream{status: ofstd.EC_Normal, size: 454}},
		{gettestdatafolder() + "GH223.dcm", 0, &DcmInputFileStream{status: ofstd.EC_Normal, size: 702}},
		{gettestdatafolder() + "GH133.dcm", 0, &DcmInputFileStream{status: ofstd.EC_Normal, size: 2980438}},
	}
	for _, c := range cases {
		got := NewDcmInputFileStream(c.in_1, c.in_2)
		defer got.Close()
		defer c.want.Close()

		if (got.size != c.want.size) || (got.status.Status() != c.want.status.Status()) {
			t.Errorf("NewDcmInputFileStream(%v,%v) == want %v got %v", c.in_1, c.in_2, c.want, got)
		}
	}
}

func TestDcmInputFileStreamGood(t *testing.T) {
	cases := []struct {
		in   *DcmInputFileStream
		want bool
	}{
		{NewDcmInputFileStream("", 0), false},
		{NewDcmInputFileStream(gettestdatafolder()+"GH220.dcm", 0), true},
		{NewDcmInputFileStream(gettestdatafolder()+"GH223.dcm", 0), true},
		{NewDcmInputFileStream(gettestdatafolder()+"GH133.dcm", 0), true},
	}
	for _, c := range cases {
		got := c.in.Good()
		defer c.in.Close()

		if got != c.want {
			t.Errorf(" %v Good() == want %v got %v", c.in, c.want, got)
		}
	}

}

func TestDcmInputFileStreamStatus(t *testing.T) {
	cases := []struct {
		in   *DcmInputFileStream
		want ofstd.OFCondition
	}{
		{NewDcmInputFileStream("", 0), ofstd.MakeOFCondition(OFM_dcmdata, 18, ofstd.OF_error, "open : The system cannot find the file specified.")},
		{NewDcmInputFileStream(gettestdatafolder()+"GH220.dcm", 0), ofstd.EC_Normal},
		{NewDcmInputFileStream(gettestdatafolder()+"GH223.dcm", 0), ofstd.EC_Normal},
		{NewDcmInputFileStream(gettestdatafolder()+"GH133.dcm", 0), ofstd.EC_Normal},
	}
	for _, c := range cases {
		got := c.in.Status()
		defer c.in.Close()

		if got.Status() != c.want.Status() {
			t.Errorf(" %v Status() == want %v got %v", c.in, c.want, got)
		}
	}

}

func TestDcmInputFileStreamEos(t *testing.T) {
	cases := []struct {
		in   *DcmInputFileStream
		want bool
	}{
		{NewDcmInputFileStream("", 0), true},
		{NewDcmInputFileStream(gettestdatafolder()+"GH220.dcm", 0), false},
		{NewDcmInputFileStream(gettestdatafolder()+"GH223.dcm", 0), false},
		{NewDcmInputFileStream(gettestdatafolder()+"GH133.dcm", 0), false},
		{NewDcmInputFileStream(gettestdatafolder()+"GH220.dcm", 454), true},
		{NewDcmInputFileStream(gettestdatafolder()+"GH223.dcm", 702), true},
		{NewDcmInputFileStream(gettestdatafolder()+"GH133.dcm", 2980438), true},
	}
	for _, c := range cases {
		got := c.in.Eos()
		defer c.in.Close()

		if got != c.want {
			t.Errorf(" %v Eos() == want %v got %v", c.in, c.want, got)
		}
	}

}

func TestDcmInputFileStreamAvail(t *testing.T) {
	cases := []struct {
		in   *DcmInputFileStream
		want int64
	}{
		{NewDcmInputFileStream("", 0), 0},
		{NewDcmInputFileStream(gettestdatafolder()+"GH220.dcm", 0), 454},
		{NewDcmInputFileStream(gettestdatafolder()+"GH223.dcm", 0), 702},
		{NewDcmInputFileStream(gettestdatafolder()+"GH133.dcm", 0), 2980438},
		{NewDcmInputFileStream(gettestdatafolder()+"GH220.dcm", 10), 444},
		{NewDcmInputFileStream(gettestdatafolder()+"GH223.dcm", 100), 602},
		{NewDcmInputFileStream(gettestdatafolder()+"GH133.dcm", 1000), 2979438},
	}
	for _, c := range cases {
		got := c.in.Avail()
		defer c.in.Close()

		if got != c.want {
			t.Errorf(" %v Avail() == want %v got %v", c.in, c.want, got)
		}
	}

}

func TestDcmInputFileStreamRead(t *testing.T) {
	cases := []struct {
		in_0   *DcmInputFileStream
		in_1   int64
		want_1 []byte
		want_2 int64
	}{
		{NewDcmInputFileStream("", 0), 0, nil, 0},
		{NewDcmInputFileStream(gettestdatafolder()+"GH184.dcm", 0), 1, []byte{0x52}, 1},
		{NewDcmInputFileStream(gettestdatafolder()+"GH184.dcm", 0), 2, []byte{0x52, 0x5f}, 2},
		{NewDcmInputFileStream(gettestdatafolder()+"GH184.dcm", 0), 3, []byte{0x52, 0x5f, 0x31}, 3},
		{NewDcmInputFileStream(gettestdatafolder()+"GH184.dcm", 0), 4, []byte{0x52, 0x5f, 0x31, 0x46}, 4},
	}
	for _, c := range cases {
		got_1, got_2 := c.in_0.Read(c.in_1)
		defer c.in_0.Close()

		if !reflect.DeepEqual(got_1, c.want_1) || got_2 != c.want_2 {
			t.Errorf(" %v Read(%v) == want %x | %v got %x | %v", c.in_0, c.in_1, c.want_1, c.want_2, got_1, got_2)
		}
	}

}

func TestDcmInputFileStreamSkip(t *testing.T) {
	cases := []struct {
		in_0 *DcmInputFileStream
		in_1 int64
		want int64
	}{
		{NewDcmInputFileStream("", 0), 0, 0},
		{NewDcmInputFileStream(gettestdatafolder()+"GH220.dcm", 0), 1, 1},
		{NewDcmInputFileStream(gettestdatafolder()+"GH220.dcm", 0), 454, 454},
		{NewDcmInputFileStream(gettestdatafolder()+"GH220.dcm", 0), 453, 453},
		{NewDcmInputFileStream(gettestdatafolder()+"GH220.dcm", 0), 455, 454},
	}
	for _, c := range cases {
		got := c.in_0.Skip(c.in_1)
		defer c.in_0.Close()

		if got != c.want {
			t.Errorf(" %v Skip(%v) == want %v got %v ", c.in_0, c.in_1, c.want, got)
		}
	}

}

func TestDcmInputFileStreamPutback1(t *testing.T) {
	cases := []struct {
		in_0 *DcmInputFileStream
		in_1 int64
		want int64
	}{
		{NewDcmInputFileStream("", 0), 0, 0},
		{NewDcmInputFileStream(gettestdatafolder()+"GH220.dcm", 100), 1, 99},
		{NewDcmInputFileStream(gettestdatafolder()+"GH220.dcm", 0), 454, 0},
		{NewDcmInputFileStream(gettestdatafolder()+"GH220.dcm", 100), 453, 100},
		{NewDcmInputFileStream(gettestdatafolder()+"GH220.dcm", 10), 455, 10},
		{NewDcmInputFileStream(gettestdatafolder()+"GH220.dcm", 100), 10, 90},
		{NewDcmInputFileStream(gettestdatafolder()+"GH220.dcm", 10), 9, 1},
	}
	for _, c := range cases {
		c.in_0.Putback(c.in_1)
		got, _ := c.in_0.file.Seek(0, os.SEEK_CUR)
		defer c.in_0.Close()

		if got != c.want {
			t.Errorf(" %v Putback(%v) == want %v got %v ", c.in_0, c.in_1, c.want, got)
		}
	}

}

func TestDcmInputFileStreamPutback2(t *testing.T) {
	cases := []struct {
		in_0    *DcmInputFileStream
		in_skip int64
		want_1  int64
		want_2  int64
	}{
		{NewDcmInputFileStream("", 0), 0, 0, 0},
		{NewDcmInputFileStream(gettestdatafolder()+"GH220.dcm", 100), 1, 1, 100},
		{NewDcmInputFileStream(gettestdatafolder()+"GH220.dcm", 0), 454, 454, 0},
		{NewDcmInputFileStream(gettestdatafolder()+"GH220.dcm", 100), 453, 354, 100},
		{NewDcmInputFileStream(gettestdatafolder()+"GH220.dcm", 10), 455, 444, 10},
		{NewDcmInputFileStream(gettestdatafolder()+"GH220.dcm", 100), 10, 10, 100},
		{NewDcmInputFileStream(gettestdatafolder()+"GH220.dcm", 10), 9, 9, 10},
	}
	for _, c := range cases {
		c.in_0.Mark()
		c.in_0.Skip(c.in_skip)
		got_1 := c.in_0.Tell()
		c.in_0.Putback()
		got_2, _ := c.in_0.file.Seek(0, os.SEEK_CUR)
		defer c.in_0.Close()

		if got_1 != c.want_1 || got_2 != c.want_2 {
			t.Errorf(" %v Putback() == want %v | %v got %v | %v", c.in_0, c.want_1, c.want_2, got_1, got_2)
		}
	}

}
