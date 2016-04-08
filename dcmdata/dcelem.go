package dcmdata

import (
	"C"

	"github.com/grayzone/godcm/ofstd"
)

type DcmElement struct {
	DcmObject
	/// current byte order of attribute value in memory
	fByteOrder E_ByteOrder

	/// required information to load value later
	fLoadValue DcmInputStreamFactory

	/// value of the element
	fValue uint16
}

/** get a pointer to the element value of the current element as type string.
 *  Requires element to be of corresponding VR, otherwise an error is returned.
 *  This method does not copy, but returns a pointer to the element value,
 *  which remains under control of this object and is valid only until the next
 *  read, write or put operation.
 *  @param val pointer to value returned in this parameter upon success
 *  @return EC_Normal upon success, an error code otherwise
 */
/** constructor.
 *  Create new element from given tag and length.
 *  @param tag DICOM tag for the new element
 *  @param len value length for the new element
 */
func NewDcmElement(tag DcmTag, l uint32) *DcmElement {
	var result DcmElement
	obj := NewDcmObject(tag, l)
	result.DcmObject = *obj
	result.fByteOrder = GLocalByteOrder
	return &result

}

/** get a pointer to the element value of the current element as type string.
 *  Requires element to be of corresponding VR, otherwise an error is returned.
 *  This method does not copy, but returns a pointer to the element value,
 *  which remains under control of this object and is valid only until the next
 *  read, write or put operation.
 *  @param val pointer to value returned in this parameter upon success
 *  @return EC_Normal upon success, an error code otherwise
 */
func (e *DcmElement) GetString(val string) ofstd.OFCondition { // for strings
	e.errorFlag = EC_IllegalCall
	return e.errorFlag
}

/** calculate the value length (without attribute tag, VR and length field)
 *  of this DICOM element when encoded with the given transfer syntax and
 *  the given encoding type for sequences. Never returns undefined length.
 *  @param xfer transfer syntax for length calculation
 *  @param enctype sequence encoding type for length calculation
 *  @return value length of DICOM element
 */
func (e *DcmElement) GetLength(xfer E_TransferSyntax, enctype E_EncodingType) uint32 {

	return e.length
}

/** calculate the length of this DICOM element when encoded with the
 *  given transfer syntax and the given encoding type for sequences.
 *  For elements, the length includes the length of the tag, length field,
 *  VR field and the value itself, for items and sequences it returns
 *  the length of the complete item or sequence including delimitation tags
 *  if applicable. Never returns undefined length.
 *  @param xfer transfer syntax for length calculation
 *  @param enctype sequence encoding type for length calculation
 *  @return length of DICOM element
 */
func (e *DcmElement) CalcElementLength(xfer E_TransferSyntax, enctype E_EncodingType) uint32 {
	xferSyn := NewDcmXfer(xfer)
	vr := e.GetVR()
	if (vr == EVR_UNKNOWN2B) || (vr == EVR_na) {
		vr = EVR_UN
	}

	headerLength := xferSyn.SizeofTagHeader(vr)

	elemLength := e.GetLength(xfer, enctype)
	if ofstd.Check32BitAddOverflow(headerLength, elemLength) {
		return DCM_UndefinedLength
	} else {
		return headerLength + elemLength
	}
}
