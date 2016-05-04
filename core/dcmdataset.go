package core

import (
	"errors"
	_ "log"
)

// DcmDataset is to contain the DICOM dataset from file
type DcmDataset struct {
	Elements []DcmElement
}

func (dataset *DcmDataset) Read(stream *DcmFileStream, isExplicitVR bool, byteOrder EByteOrder, isReadValue bool, isReadPixel bool) error {
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

// FindElement find the element information from the data set
func (dataset DcmDataset) FindElement(e *DcmElement) error {
	for _, v := range dataset.Elements {
		if e.Tag == v.Tag {
			*e = v
			return nil
		}
	}
	str := "not find the tag '" + e.Tag.String() + "' in the data set"
	return errors.New(str)
}

// PatientID get the patient ID from the dicom file
func (dataset DcmDataset) PatientID() string {
	var elem DcmElement
	elem.Tag = DCMPatientID
	err := dataset.FindElement(&elem)
	if err != nil {
		return ""
	}
	return elem.GetValueString()
}

// PatientName get the patient name from the dicom file
func (dataset DcmDataset) PatientName() string {
	var elem DcmElement
	elem.Tag = DCMPatientName
	err := dataset.FindElement(&elem)
	if err != nil {
		return ""
	}
	return elem.GetValueString()
}

// Modality get the modality of the dicom image
func (dataset DcmDataset) Modality() string {
	var elem DcmElement
	elem.Tag = DCMModality
	err := dataset.FindElement(&elem)
	if err != nil {
		return ""
	}
	return elem.GetValueString()
}

// Rows get the rows of the dicom image
func (dataset DcmDataset) Rows() string {
	var elem DcmElement
	elem.Tag = DCMRows
	err := dataset.FindElement(&elem)
	if err != nil {
		return ""
	}
	return elem.GetValueString()
}

// Columns get the columns of the dicom image
func (dataset DcmDataset) Columns() string {
	var elem DcmElement
	elem.Tag = DCMColumns
	err := dataset.FindElement(&elem)
	if err != nil {
		return ""
	}
	return elem.GetValueString()
}
