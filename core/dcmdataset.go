package core

import (
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
		//		log.Println(elem)
		dataset.Elements = append(dataset.Elements, elem)
	}
	return nil
}
