package dcmdata

import (
	"log"

	"github.com/grayzone/godcm/ofstd"
)

type DcmItem struct {
	DcmObject
	elementList *DcmList
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

/** This function takes care of group length and padding elements
 *  in the current element list according to what is specified in
 *  glenc and padenc. If required, this function does the following
 *  two things:
 *    a) it calculates the group length of all groups which are
 *       contained in this item and sets the calculated values
 *       in the corresponding group length elements and
 *    b) it inserts a corresponding padding element (or, in case
 *       of sequences: padding elements) with a corresponding correct
 *       size into the element list.
 *  @param glenc          Encoding type for group length; specifies what shall
 *                        be done with group length tags.
 *  @param padenc         Encoding type for padding; specifies what shall be
 *                        done with padding tags.
 *  @param xfer           The transfer syntax that shall be used.
 *  @param enctype        Encoding type for sequences; specifies how sequences
 *                        will be handled.
 *  @param padlen         The length up to which the dataset shall be padded,
 *                        if padding is desired.
 *  @param subPadlen      For sequences (ie sub elements), the length up to
 *                        which item shall be padded, if padding is desired.
 *  @param instanceLength Number of extra bytes added to the item/dataset
 *                        length used when computing the padding; this
 *                        parameter is for instance used to pass the length
 *                        of the file meta header from the DcmFileFormat to
 *                        the DcmDataset object.
 *  @return status, EC_Normal if successful, an error code otherwise
 */
func (item *DcmItem) ComputeGroupLengthAndPadding(glenc E_GrpLenEncoding,
	padenc E_PaddingEncoding,
	xfer E_TransferSyntax,
	enctype E_EncodingType,
	padlen uint32,
	subPadlen uint32,
	instanceLength uint32) ofstd.OFCondition {
	/* if certain conditions are met, this is considered to be an illegal call. */
	if ((padenc == EPD_withPadding) && (padlen%2 == 1 || subPadlen%2 == 1)) || ((glenc == EGL_recalcGL || glenc == EGL_withGL || padenc == EPD_withPadding) && xfer == EXS_Unknown) {
		return EC_IllegalCall
	}
	/* if the caller specified that group length tags and padding */
	/* tags are not supposed to be changed, there is nothing to do. */
	if (glenc == EGL_noChange) && (padenc == EPD_noChange) {
		return ofstd.EC_Normal
	}
	/* if we get to this point, we need to do something. First of all, set the error indicator to normal. */
	l_error := ofstd.EC_Normal
	if item.elementList.Empty() != true {
		/* initialize some variables */
		xferSyn := NewDcmXfer(xfer)
		var do *DcmObject
		var beginning bool = true
		var lastGrp uint16 = 0x0000
		var actGrp uint16
		var actGLElem *DcmUnsignedLong
		//	var paddingGL *DcmUnsignedLong
		//	var grplen uint32
		//	var sublen uint32
		var groupLengthExceeded = false

		/* determine the current seek mode and set the list pointer to the first element */
		var seekmode E_ListPos = ELP_next
		item.elementList.Seek(ELP_first)

		/* start a loop: we want to go through all elements as long as everything is okay */
		for item.elementList.Seek(seekmode) != nil && l_error.Good() {
			seekmode = ELP_next
			do = item.elementList.Get(ELP_atpos)
			/* if the current element is a sequence, compute group length and padding for the sub sequence */
			if do.GetVR() == EVR_SQ {
				// add size of sequence header
				tmplen := instanceLength + xferSyn.SizeofTagHeader(EVR_SQ)
				// call computeGroupLengthAndPadding for all contained items
				var sq DcmSequenceOfItems
				sq.DcmObject = *do
				l_error = sq.ComputeGroupLengthAndPadding(glenc, padenc, xfer, enctype, subPadlen, subPadlen, tmplen)
			}
			/* if everything is ok so far */
			if l_error.Good() {
				/* in case one of the following two conditions is met */
				/*  (i) the caller specified that we want to add or remove group length elements and the current */
				/*      element's tag shows that it is a group length element (tag's element number equals 0x0000) */
				/*  (ii) the caller specified that we want to add or remove padding elements and the current */
				/*      element's tag shows that it is a padding element (tag is (0xfffc,0xfffc) */
				/* then we want to delete the current (group length or padding) element */
				if ((glenc == EGL_withGL || glenc == EGL_withoutGL) && do.GetETag() == 0x0000) || ((padenc != EPD_noChange) && (do.GetTag().DcmTagKey == DCM_DataSetTrailingPadding)) {
					seekmode = ELP_atpos // remove advances 1 element forward -> make next seek() work
				} else if (glenc == EGL_withGL) || (glenc == EGL_recalcGL) {
					/* if the above mentioned conditions are not met but the caller specified that we want to add group */
					/* length tags for every group or that we want to recalculate values for existing group length tags */

					/* we need to determine the current element's group number */
					actGrp = do.GetGTag()

					/* and if the group number is different from the last remembered group number or */
					/* if this id the very first element that is treated then we've found a new group */
					if (actGrp != lastGrp) || beginning { // new Group found
						/* set beginning to false in order to specify that the */
						/* very first element has already been treated */
						beginning = false

						/* if the current element is a group length element and its data type */
						/* is not UL replace this element with one that has a UL datatype since */
						/* group length elements are supposed to have this data type */
						if (do.GetETag() == 0x0000) && (do.Ident() != EVR_UL) {
							item.elementList.Remove()
							vr := NewDcmVR(EVR_UL)
							tagUL := NewDcmTagWithGEV(actGrp, 0x0000, *vr)
							dUL := NewDcmUnsignedLong(*tagUL, 0)

							item.elementList.Insert(&dUL.DcmObject, ELP_prev)
							do = &dUL.DcmObject
							log.Println("DcmItem: Group Length with VR other than UL found, corrected")
						} else if glenc == EGL_withGL {
							/* if the above mentioned condition is not met but the caller specified */
							/* that we want to add group length elements, we need to add such an element */

							// Create GroupLength element
							vr := NewDcmVR(EVR_UL)
							tagUL := NewDcmTagWithGEV(actGrp, 0x0000, *vr)
							// insert new GroupLength element
							dUL := NewDcmUnsignedLong(*tagUL, 0)
							item.elementList.Insert(&dUL.DcmObject, ELP_prev)
							do = &dUL.DcmObject
						}
						/* in case we want to add padding elements and the current element is a */
						/* padding element we want to remember the padding element so that the */
						/* group length of this element can be stored later */
						if (padenc == EPD_withPadding) && (actGrp == 0xFFFC) {
							//				paddingGL = NewDcmUnsignedLong(do.GetTag(), 0)
						}
						/* if actGLElem contains a valid pointer it was set in one of the last iterations */
						/* to the group lenght element of the last group. We need to write the current computed */
						/* group length value to this element. Exception: If group length exceeds maximum possible */
						/* value, than remove group length element instead of setting it */
						if actGLElem != nil {
							if groupLengthExceeded != true {
								// do not use putUint32() in order to make sure that the resulting VM is really 1
								//				actGLElem.

							}

						}

					}

				}

			}

		}

	}

	return ofstd.EC_Normal

}
