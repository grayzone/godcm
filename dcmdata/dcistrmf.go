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
	err := result.Open(filename)
	if err != nil {
		result.status_ = ofstd.MakeOFCondition(OFM_dcmdata, 18, ofstd.OF_error, err.Error())
		return &result
	}
	result.size_, _ = result.file_.Seek(0, os.SEEK_END)
	_, err = result.file_.Seek(offset, os.SEEK_SET)
	if err != nil {
		result.status_ = ofstd.MakeOFCondition(OFM_dcmdata, 18, ofstd.OF_error, err.Error())
		return &result
	}
	return &result
}

func (p *DcmFileProducer) Open(filename string) error {
	var err error
	p.file_, err = os.Open(filename)
	return err
}

func (p *DcmFileProducer) Close() error {
	return p.file_.Close()
}

func (p *DcmFileProducer) Good() bool {
	return p.status_.Good()
}

func (p *DcmFileProducer) Status() ofstd.OFCondition {
	return p.status_
}

func (p *DcmFileProducer) Eos() bool {
	if p.file_ == nil {
		return true
	}
	size, _ := p.file_.Seek(0, os.SEEK_CUR)
	return size == p.size_
}
