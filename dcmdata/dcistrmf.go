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

func (p *DcmFileProducer) Avail() int64 {
	if p.file_ == nil {
		return 0
	}
	size, _ := p.file_.Seek(0, os.SEEK_CUR)
	return p.size_ - size
}

func (p *DcmFileProducer) Read(buflen int64) ([]byte, int64) {
	var result int64
	if !p.Good() || (p.file_ == nil) || (buflen == 0) {
		return nil, result
	}
	buf := make([]byte, buflen)
	r, _ := p.file_.Read(buf)
	result = int64(r)
	return buf, result
}

func (p *DcmFileProducer) Skip(skiplen int64) int64 {
	var result int64
	if !p.Good() || (p.file_ == nil) || (skiplen == 0) {
		return result
	}
	pos, _ := p.file_.Seek(0, os.SEEK_CUR)
	if p.size_-pos < skiplen {
		result = p.size_ - pos
	} else {
		result = skiplen
	}
	_, err := p.file_.Seek(result, os.SEEK_CUR)
	if err != nil {
		p.status_ = ofstd.MakeOFCondition(OFM_dcmdata, 18, ofstd.OF_error, err.Error())
	}
	return result
}

func (p *DcmFileProducer) Putback(num int64) {
	if !p.Good() || (p.file_ == nil) || (num == 0) {
		return
	}
	pos, _ := p.file_.Seek(0, os.SEEK_CUR)
	if num > pos {
		p.status_ = EC_PutbackFailed // tried to putback before start of file
		return
	}
	_, err := p.file_.Seek(-num, os.SEEK_CUR)
	if err != nil {
		p.status_ = ofstd.MakeOFCondition(OFM_dcmdata, 18, ofstd.OF_error, err.Error())
	}
}
