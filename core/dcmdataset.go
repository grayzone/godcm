package core

import (
	"errors"
	_ "log" // for debug
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
		elem.isReadValue = isReadValue
		elem.isReadPixel = isReadPixel

		err := elem.ReadDcmElement(stream)
		if err != nil {
			return err
		}
		//		log.Println(elem)
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

// SOPInstanceUID get the SOP Instance UID of the dicom image
func (dataset DcmDataset) SOPInstanceUID() string {
	return dataset.getElementValue(DCMSOPInstanceUID)
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

// WindowWidth gets the window width of the dicom image
func (dataset DcmDataset) WindowWidth() string {
	return dataset.getElementValue(DCMWindowWidth)
}

// NumberOfFrames gets the frame number
func (dataset DcmDataset) NumberOfFrames() string {
	return dataset.getElementValue(DCMNumberOfFrames)
}

/*
	PixelWidth    float64
	PixelHeight   float64
*/

// BitsAllocated gets the bits allocated value
func (dataset DcmDataset) BitsAllocated() string {
	return dataset.getElementValue(DCMBitsAllocated)
}

// BitsStored gets the bits stored
func (dataset DcmDataset) BitsStored() string {
	return dataset.getElementValue(DCMBitsStored)
}

// HighBit gets the high bit
func (dataset DcmDataset) HighBit() string {
	return dataset.getElementValue(DCMHighBit)
}

// PhotometricInterpretation gets photometric interpretation
func (dataset DcmDataset) PhotometricInterpretation() string {
	return dataset.getElementValue(DCMPhotometricInterpretation)
}

// SamplesPerPixel gets samples per pixel
func (dataset DcmDataset) SamplesPerPixel() string {
	return dataset.getElementValue(DCMSamplesPerPixel)
}

// PixelRepresentation gets pixel representation
// unsigned (0) or signed (1), the default is unsigned
func (dataset DcmDataset) PixelRepresentation() string {
	return dataset.getElementValue(DCMPixelRepresentation)
}

// PlanarConfiguration gets planar configuration
// the default is interlaced, 0 meaning the channels are interlaced which is the common way of serializing color pixels or 1 meaning its separated
func (dataset DcmDataset) PlanarConfiguration() string {
	return dataset.getElementValue(DCMPlanarConfiguration)
}

// PixelData get the pixel data of the dicom image.
func (dataset DcmDataset) PixelData() []byte {
	var elem DcmElement
	elem.Tag = DCMPixelData
	err := dataset.FindElement(&elem)
	if err != nil {
		return nil
	}
	if elem.Squence != nil {
		var maxLengthIndex int
		var maxLength int64
		for i, v := range elem.Squence.Item {
			if maxLength < v.Length {
				maxLength = v.Length
				maxLengthIndex = i
			}
		}
		return elem.Squence.Item[maxLengthIndex].Value
	}
	return elem.Value
}
