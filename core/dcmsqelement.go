package core

import (
	_ "log"
)

// DcmSQElement contain a SQ Data Element
type DcmSQElement struct {
	Item []DcmElement
}

func (sq DcmSQElement) String() string {
	var result string
	for _, v := range sq.Item {
		result += v.String()
		result += " | "
	}
	return result
}

// Read the items in an SQ data element
func (sq *DcmSQElement) Read(stream *DcmFileStream, length int64, isExplicitVR bool, isReadValue bool) error {
	if isExplicitVR {
		return sq.ReadItemsWithExplicitVR(stream, length, isReadValue)
	}
	return sq.ReadItemsWithImplicitVR(stream)
}

func readItemWithUndefinedLength(e *DcmElement, s *DcmFileStream, isReadValue bool) error {
	isFoundDelimTag := false
	for !isFoundDelimTag {
		//read 2 bytes to check the group of the delim tag.
		bg, err := s.Read(2)
		if err != nil {
			return err
		}
		if bg[0] == 0xFE && bg[1] == 0xFF {
			// continue to check the element of the delim tag.
			be, err := s.Read(2)
			if err != nil {
				return err
			}
			if be[0] == 0x0D && be[1] == 0xE0 {
				// find the delim tag
				_, err = s.Skip(4)
				if err != nil {
					return err
				}
				isFoundDelimTag = true
			}
			if isReadValue {
				e.Value = append(e.Value, bg...)
				e.Value = append(e.Value, be...)
			}
		} else {
			if isReadValue {
				e.Value = append(e.Value, bg...)
			}
		}
	}
	return nil
}

// ReadItemsWithExplicitVR the items in an SQ data element with explicit VR
func (sq *DcmSQElement) ReadItemsWithExplicitVR(stream *DcmFileStream, length int64, isReadValue bool) error {
	startPos := stream.Position
	var delta int64
	for !stream.Eos() && delta < length {
		var elem DcmElement
		elem.isReadValue = isReadValue
		err := elem.ReadDcmTag(stream)
		//		log.Println("SQ :", elem, delta)

		if err != nil {
			return err
		}
		if elem.Tag == DCMItem {
			// read item length
			err = elem.ReadValueLengthUint32(stream)
			if err != nil {
				return err
			}

			if elem.Length == 0xFFFFFFFF {
				err = readItemWithUndefinedLength(&elem, stream, isReadValue)
				if err != nil {
					return err
				}
			} else {
				// read item value
				err = elem.ReadValue(stream)
				if err != nil {
					return err
				}
			}
			//			log.Println(elem)
			sq.Item = append(sq.Item, elem)
		}
		if elem.Tag == DCMSequenceDelimitationItem {
			// read item length
			err = elem.ReadValueLengthUint32(stream)
			if err != nil {
				return err
			}
			//			log.Println(elem)
			sq.Item = append(sq.Item, elem)
			break
		}
		delta = stream.Position - startPos
	}
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
