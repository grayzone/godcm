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
