package dcmdata

import (
	"log"

	"github.com/grayzone/godcm/ofstd"
)

type DcmItem struct {
	DcmObject
	elementList         *DcmList
	lastElementComplete bool
	fStartPosition      int
}

func NewDcmItem(tag DcmTag, len uint32) *DcmItem {
	return &DcmItem{*NewDcmObject(tag, len), nil, true, 0}
}

/** Virtual object copying. This method can be used for DcmObject
 *  and derived classes to get a deep copy of an object. Internally
 *  the assignment operator is called if the given DcmObject parameter
 *  is of the same type as "this" object instance. If not, an error
 *  is returned. This function permits copying an object by value
 *  in a virtual way which therefore is different to just calling the
 *  assignment operator of DcmElement which could result in slicing
 *  the object.
 *  @param rhs - [in] The instance to copy from. Has to be of the same
 *                class type as "this" object
 *  @return EC_Normal if copying was successful, error otherwise
 */
func (item *DcmItem) CopyFrom(rhs DcmObject) ofstd.OFCondition {
	if &rhs != &item.DcmObject {
		if rhs.Ident() != item.Ident() {
			return EC_IllegalCall
		}
		item.DcmObject = rhs
	}
	return ofstd.EC_Normal
}

/** calculate the value length (without attribute tag, VR and length field)
 *  of this DICOM element when encoded with the given transfer syntax and
 *  the given encoding type for sequences.
 *  If length encodig is set to be explicit and the item content is larger
 *  than the available 32-bit length field, then undefined length is
 *  returned. If "dcmWriteOversizedSeqsAndItemsUndefined" is disabled,
 *  also the internal DcmObject errorFlag is set to
 *  EC_SeqOrItemContentOverflow.
 *  @param xfer transfer syntax for length calculation
 *  @param enctype sequence encoding type for length calculation
 *  @return value length of DICOM element
 */
func (item *DcmItem) GetLength(xfer E_TransferSyntax, enctype E_EncodingType) uint32 {
	var itemlen uint32
	if item.elementList.Empty() != true {
		item.elementList.Seek(ELP_first)
		for item.elementList.Seek(ELP_next) != nil {
			do := item.elementList.Get(ELP_atpos)
			sublen := do.CalcElementLength(xfer, enctype)
			/* explicit length: be sure that total size of contained elements fits into item's
			   32 Bit length field. If not, switch encoding automatically to undefined
			   length for this item. Nevertheless, any contained elements will be
			   written with explicit length if possible.
			*/
			if (enctype == EET_ExplicitLength) && (ofstd.Check32BitAddOverflow(sublen, itemlen)) {
				if DcmWriteOversizedSeqsAndItemsUndefined {
					log.Println("DcmItem: Explicit length of item exceeds 32-Bit length field, trying to encode with undefined length")
				} else {
					log.Println("DcmItem: Explicit length of item exceeds 32-Bit length field, aborting write")
					item.errorFlag = EC_SeqOrItemContentOverflow
				}
				return DCM_UndefinedLength
			} else {
				itemlen = itemlen + sublen
			}
		}
	}
	return itemlen
}

/** calculate the length of this DICOM element when encoded with the
 *  given transfer syntax and the given encoding type for sequences.
 *  For elements, the length includes the length of the tag, length field,
 *  VR field and the value itself, for items and sequences it returns
 *  the length of the complete item or sequence including delimitation tags
 *  if applicable.
 *  If length encodig is set to be explicit and the total item size is
 *  larger than the available 32-bit length field, then undefined length
 *  is returned. If "dcmWriteOversizedSeqsAndItemsImplicit" is disabled,
 *  also the internal DcmObject errorFlag is set to EC_SeqOrItemContentOverflow
 *  in case the item content (excluding tag header etc.) is already too
 *  large.
 *  @param xfer transfer syntax for length calculation
 *  @param enctype sequence encoding type for length calculation
 *  @return length of DICOM element
 */
func (item *DcmItem) CalcElementLength(xfer E_TransferSyntax, enctype E_EncodingType) uint32 {
	xferSyn := NewDcmXfer(xfer)
	/* Length of item's start header */
	headersize := xferSyn.SizeofTagHeader(item.GetVR())
	/* Length of item's content, i.e. contained elements */
	itemlen := item.GetLength(xfer, enctype)
	/* Since the item's total length can exceed the maximum length of 32 bit, it is
	 * always necessary to check for overflows. The approach taken is not elegant
	 * but should work...
	 */
	if (itemlen == DCM_UndefinedLength) || ofstd.Check32BitAddOverflow(itemlen, headersize) {
		return DCM_UndefinedLength
	}
	itemlen = itemlen + xferSyn.SizeofTagHeader(item.GetVR())
	if enctype == EET_UndefinedLength { // add bytes for closing item tag marker if necessary
		if ofstd.Check32BitAddOverflow(itemlen, 8) {
			return DCM_UndefinedLength
		} else {
			itemlen = itemlen + 8
		}
	}

	return itemlen
}

func (item *DcmItem) Clear() ofstd.OFCondition {
	item.errorFlag = ofstd.EC_Normal
	if item.elementList != nil {
		item.elementList.DeleteAllElements()
	}
	item.setLengthField(0)
	return item.errorFlag
}

func (item *DcmItem) ComputeGroupLengthAndPadding(glenc E_GrpLenEncoding, padenc E_PaddingEncoding, xfer E_TransferSyntax, enctype E_EncodingType, padlen uint32, subPadlen uint32, instanceLength uint32) ofstd.OFCondition {
	if (padenc == EPD_withPadding && ((padlen%2) != 0 || (subPadlen%2) != 0)) || ((glenc == EGL_recalcGL || glenc == EGL_withGL || padenc == EPD_withPadding) && xfer == EXS_Unknown) {
		return EC_IllegalCall
	}
	if glenc == EGL_noChange && padenc == EPD_noChange {
		return ofstd.EC_Normal
	}

	err := ofstd.EC_Normal
	if item.elementList.Empty() {
		return err
	}
	xferSyn := NewDcmXfer(xfer)
	seekmode := ELP_next
	item.elementList.Seek(ELP_first)
	for err.Good() && (item.elementList.Seek(seekmode) != nil) {
		seekmode = ELP_next
		d := item.elementList.Get(ELP_atpos)
		if d.GetVR() == EVR_SQ {
			templen := instanceLength + xferSyn.SizeofTagHeader(EVR_SQ)
			t := NewDcmItem(d.tag, 0)
			t.ComputeGroupLengthAndPadding(glenc, padenc, xfer, enctype, subPadlen, subPadlen, templen)
		}
		if !err.Good() {
			continue
		}
		if ((glenc == EGL_withGL || glenc == EGL_withoutGL) && d.GetETag() == 0x0000) || (padenc != EPD_noChange && d.GetTag().DcmTagKey == DCM_DataSetTrailingPadding) {
			item.elementList.Remove()
			seekmode = ELP_atpos
		} else if glenc == EGL_withGL || glenc == EGL_recalcGL {

		}
	}

	return err
}
