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

func (dataset DcmDataset) getElementValue(tag DcmTag) string {
	var elem DcmElement
	elem.Tag = tag
	err := dataset.FindElement(&elem)
	if err != nil {
		return ""
	}
	return elem.GetValueString()
}

// PatientID get the patient ID from the dicom file
func (dataset DcmDataset) PatientID() string {
	return dataset.getElementValue(DCMPatientID)
}

// PatientName get the patient name from the dicom file
func (dataset DcmDataset) PatientName() string {
	return dataset.getElementValue(DCMPatientName)
}

// Modality get the modality of the dicom image
func (dataset DcmDataset) Modality() string {
	return dataset.getElementValue(DCMModality)
}

// Rows get the rows of the dicom image
func (dataset DcmDataset) Rows() string {
	return dataset.getElementValue(DCMRows)
}

// Columns get the columns of the dicom image
func (dataset DcmDataset) Columns() string {
	return dataset.getElementValue(DCMColumns)
}

// WindowCenter gets the window center of the dicom image
func (dataset DcmDataset) WindowCenter() string {
	return dataset.getElementValue(DCMWindowCenter)
}

// WindowWidth gets tge window width of the dicom image
func (dataset DcmDataset) WindowWidth() string {
	return dataset.getElementValue(DCMWindowWidth)
}

// PixelData get the pixel data of the dicom image.
func (dataset DcmDataset) PixelData() []byte {

	return nil
}
