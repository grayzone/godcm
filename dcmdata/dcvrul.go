package dcmdata

import "github.com/grayzone/godcm/ofstd"

/** a class representing the DICOM value representation 'Unsigned Long' (UL)
 */
type DcmUnsignedLong struct {
	DcmElement
}

func NewDcmUnsignedLong(tag DcmTag, l uint32) *DcmUnsignedLong {
	var d DcmUnsignedLong
	elem := *NewDcmElement(tag, l)
	d.DcmElement = elem
	return &d
}

/** set element value to given integer array data
 *  @param uintVals unsigned integer data to be set
 *  @param numUints number of integer values to be set
 *  @return status, EC_Normal if successful, an error code otherwise
 */
func (ul *DcmUnsignedLong) PutUint32Array(uintVals *uint32, numUints uint32) ofstd.OFCondition {
	ul.errorFlag = ofstd.EC_Normal
	if numUints > 0 {
		/* check for valid data */
		if uintVals != nil {
			//ul.errorFlag =
		}

	}
	return ofstd.EC_Normal

}
