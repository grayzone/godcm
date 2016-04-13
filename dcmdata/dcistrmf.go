package dcmdata

import (
	"github.com/grayzone/godcm/ofstd"
	"os"
)

type DcmFileProducer struct {
	file_   *os.File
	status_ ofstd.OFCondition
	size_   int64
}

type DcmInputFileStream struct {
	DcmInputStream
	producer_ DcmFileProducer
	filename  string
}

func NewDcmFileProducer(filename string, offset int64) *DcmFileProducer {
	var result DcmFileProducer
	result.status_ = ofstd.EC_Normal
	var err error
	result.file_, err = os.Open(filename)
	defer result.file_.Close()
	if err != nil {
		result.status_ = ofstd.MakeOFCondition(OFM_dcmdata, 18, ofstd.OF_error, err.Error())
		return &result
	}
	result.size_, _ = result.file_.Seek(offset, os.SEEK_END)
	_, err = result.file_.Seek(offset, os.SEEK_SET)
	if err != nil {
		result.status_ = ofstd.MakeOFCondition(OFM_dcmdata, 18, ofstd.OF_error, err.Error())
		return &result
	}
	return &result
}
