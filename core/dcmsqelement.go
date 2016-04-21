package core

// DcmSQElement contain an SQ Data Element
type DcmSQElement struct {
	Item []DcmElement
}

// Read the items in an SQ data element
func (sq *DcmSQElement) Read(stream *DcmFileStream, isExplicitVR bool) error {
	if isExplicitVR {
		return sq.ReadItemsWithExplicitVR(stream)
	}
	return sq.ReadItemsWithImplicitVR(stream)
}

// ReadItemsWithExplicitVR the items in an SQ data element with explicit VR
func (sq *DcmSQElement) ReadItemsWithExplicitVR(stream *DcmFileStream) error {
	return nil
}

/*

var (
	DCMItem                     = DcmTagKey{0xfffe, 0xe000}
	DCMItemDelimitationItem     = DcmTagKey{0xfffe, 0xe00d}
	DCMSequenceDelimitationItem = DcmTagKey{0xfffe, 0xe0dd}
)


*/

// ReadItemsWithImplicitVR the items in an SQ data element with implicit VR
func (sq *DcmSQElement) ReadItemsWithImplicitVR(stream *DcmFileStream) error {

	/*
		for !stream.Eos() {
			var elem DcmElement
			err := elem.ReadDcmTag(stream)
			if err != nil {
				return err
			}
			if elem.Tag == DCMItem {
				// read item length
			}

			// Sequence Delimitation Item
			if elem.Tag.Group == 0xFFFE && elem.Tag.Element == 0xE0DD {
				// read item length
				err = elem.ReadValueLengthWithImplicitVR(stream)
				if err != nil {
					return err
				}
				break
			}

			sq.Item = appemd(sq.Item, elem)
		}


	*/
	return nil
}
