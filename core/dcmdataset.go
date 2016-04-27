package core

import (
	"errors"
	_ "log"
)

// DcmDataSet is to contain the DICOM dataset from file.
type DcmDataSet struct {
	Elements []DcmElement
}

func (dataset *DcmDataSet) Read(stream *DcmFileStream, isExplicitVR bool, byteOrder EByteOrder, isReadValue bool, isReadPixel bool) error {
	for !stream.Eos() {
		//	for range [12]int{} {
		var elem DcmElement
		elem.isExplicitVR = isExplicitVR
		elem.byteOrder = byteOrder
		err := elem.ReadDcmElement(stream, isReadValue, isReadPixel)
		if err != nil {
			return err
		}
		//	log.Println(elem)
		dataset.Elements = append(dataset.Elements, elem)
	}
	return nil
}

// FindElement find the element information from the data set.
func (dataset DcmDataSet) FindElement(e *DcmElement) error {
	for _, v := range dataset.Elements {
		if e.Tag == v.Tag {
			*e = v
			return nil
		}
	}
	str := "not find the tag '" + e.Tag.String() + "' in the data set"
	return errors.New(str)
}
