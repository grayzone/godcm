package dcmdata

import "github.com/grayzone/godcm/ofstd"

/** class representing a DICOM Sequence of Items (SQ).
 *  This class is derived from class DcmElement (and not from DcmObject) despite the fact
 *  that sequences have no value field as such, they maintain a list of items. However,
 *  all APIs in class DcmItem and class DcmDataset accept DcmElements.
 *  This is ugly and causes some DcmElement API methods to be useless with DcmSequence.
 */
type DcmSequenceOfItems struct {
	DcmElement
	/** flag used during suspended I/O. Indicates whether the last item
	 *  was completely or only partially read/written during the last call
	 *  to read/write.
	 */
	lastItemComplete bool

	/** used during reading. Contains the position in the stream where
	 *  the sequence started (needed for calculating the remaining number of
	 *  bytes available for a fixed-length sequence).
	 */
	fStartPosition ofstd.Offile_off_t

	/** true if this sequence has been instantiated while reading an UN element
	 *  with undefined length
	 */
	readAsUN_ bool

	/// the list of items maintained by this sequence object
	itemList *DcmList
}

func NewDcmSequenceOfItems(tag DcmTag, l uint32, readAsUN bool) *DcmSequenceOfItems {
	var sq DcmSequenceOfItems
	sq.DcmElement = *NewDcmElement(tag, l)
	sq.lastItemComplete = true
	sq.readAsUN_ = readAsUN
	return &sq
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
func (sq *DcmSequenceOfItems) ComputeGroupLengthAndPadding(glenc E_GrpLenEncoding, padenc E_PaddingEncoding, xfer E_TransferSyntax, enctype E_EncodingType, padlen uint32, subPadlen uint32, instanceLength uint32) ofstd.OFCondition {
	l_error := ofstd.EC_Normal
	if sq.itemList.Empty() != true {
		sq.itemList.Seek(ELP_first)
		for l_error.Good() && (sq.itemList.Seek(ELP_next) != nil) {
			var do DcmItem
			do.DcmObject = *sq.itemList.Get(ELP_atpos)
			l_error = do.ComputeGroupLengthAndPadding(glenc, padenc, xfer, enctype, padlen, subPadlen, instanceLength)
		}
	}
	return l_error
}
