package core

import (
	"errors"
	_ "log" // for debug
	"strings"
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

func (dataset DcmDataset) GetElementValue(tag DcmTag) string {
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
	return dataset.GetElementValue(DCMPatientID)
}

// PatientName get the patient name from the dicom file
func (dataset DcmDataset) PatientName() string {
	return dataset.GetElementValue(DCMPatientName)
}

// Modality get the modality of the dicom image
func (dataset DcmDataset) Modality() string {
	return dataset.GetElementValue(DCMModality)
}

// StudyInstanceUID get the Study Instance UID of the dicom image
func (dataset DcmDataset) StudyInstanceUID() string {
	return dataset.GetElementValue(DCMStudyInstanceUID)
}

func (dataset DcmDataset) StudyDate() string {
	return dataset.GetElementValue(DCMStudyDate)
}

// SOPInstanceUID get the SOP Instance UID of the dicom image
func (dataset DcmDataset) SOPInstanceUID() string {
	return dataset.GetElementValue(DCMSOPInstanceUID)
}

// Rows get the rows of the dicom image
func (dataset DcmDataset) Rows() string {
	return dataset.GetElementValue(DCMRows)
}

// Columns get the columns of the dicom image
func (dataset DcmDataset) Columns() string {
	return dataset.GetElementValue(DCMColumns)
}

// WindowCenter gets the window center of the dicom image
func (dataset DcmDataset) WindowCenter() string {
	return dataset.GetElementValue(DCMWindowCenter)
}

// WindowWidth gets the window width of the dicom image
func (dataset DcmDataset) WindowWidth() string {
	return dataset.GetElementValue(DCMWindowWidth)
}

// NumberOfFrames gets the frame number
func (dataset DcmDataset) NumberOfFrames() string {
	result := dataset.GetElementValue(DCMNumberOfFrames)
	if len(result) == 0 {
		result = "1"
	}
	return result
}

/*
	PixelWidth    float64
	PixelHeight   float64
*/

// BitsAllocated gets the bits allocated value
func (dataset DcmDataset) BitsAllocated() string {
	return dataset.GetElementValue(DCMBitsAllocated)
}

// BitsStored gets the bits stored
func (dataset DcmDataset) BitsStored() string {
	return dataset.GetElementValue(DCMBitsStored)
}

// HighBit gets the high bit
func (dataset DcmDataset) HighBit() string {
	return dataset.GetElementValue(DCMHighBit)
}

// PhotometricInterpretation gets photometric interpretation
func (dataset DcmDataset) PhotometricInterpretation() string {
	s := dataset.GetElementValue(DCMPhotometricInterpretation)
	return strings.ToUpper(s)
}

// SamplesPerPixel gets samples per pixel
func (dataset DcmDataset) SamplesPerPixel() string {
	return dataset.GetElementValue(DCMSamplesPerPixel)
}

// PixelRepresentation gets pixel representation
// unsigned (0) or signed (1), the default is unsigned
func (dataset DcmDataset) PixelRepresentation() string {
	return dataset.GetElementValue(DCMPixelRepresentation)
}

// PlanarConfiguration gets planar configuration
// the default is interlaced, 0 meaning the channels are interlaced which is the common way of serializing color pixels or 1 meaning its separated
func (dataset DcmDataset) PlanarConfiguration() string {
	return dataset.GetElementValue(DCMPlanarConfiguration)
}

// RescaleIntercept gets rescale intercept
func (dataset DcmDataset) RescaleIntercept() string {
	return dataset.GetElementValue(DCMRescaleIntercept)
}

// RescaleSlope gets rescale slope
func (dataset DcmDataset) RescaleSlope() string {
	return dataset.GetElementValue(DCMRescaleSlope)
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
