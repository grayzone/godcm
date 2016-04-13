package dcmdata

import (
	"github.com/grayzone/godcm/ofstd"
)

type DcmProducer interface {
	Good() bool
	Status() ofstd.OFCondition
	Eos() bool
	Avail() int64
	Read(buflen int64) ([]byte, int64)
	Skip(skiplen int64) int64
	Putback(num int64)
}

/** pure virtual abstract base class for input stream factories,
 *  i.e. objects that can create a new input stream
 */

type DcmInputStreamFactory struct {
}

type DcmInputStream struct {
	tell int64
	mark int64
}

func NewDcmInputStream() *DcmInputStream {
	var result DcmInputStream
	return &result
}
