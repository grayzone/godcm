package models

import "github.com/grayzone/godcm/core"

type Slice struct {
	SOPInstanceUID       string
	SOPClassUID          string
	SpecificCharacterSet string
	InstanceNumber       string
	PatientOrientation   string
	ContentDate          string
	ContentTime          string

	//pixel
	SamplesPerPixel           string
	PhotometricInterpretation string
	Rows                      string
	Columns                   string
	BitsAllocated             string
	BitsStored                string
	HighBit                   string
	PixelRepresentation       string
	PlanarConfiguration       string
	WindowCenter              string
	WindowWidth               string
	NumberOfFrames            string
}

func (this *Slice) Parse(dataset core.DcmDataset) {
	this.SOPInstanceUID = dataset.GetElementValue(core.DCMSOPInstanceUID)
	this.SOPClassUID = dataset.GetElementValue(core.DCMSOPClassUID)
	this.SpecificCharacterSet = dataset.GetElementValue(core.DCMSpecificCharacterSet)
	this.InstanceNumber = dataset.GetElementValue(core.DCMInstanceNumber)
	this.PatientOrientation = dataset.GetElementValue(core.DCMPatientOrientation)
	this.ContentDate = dataset.GetElementValue(core.DCMContentDate)
	this.ContentTime = dataset.GetElementValue(core.DCMContentTime)

	this.SamplesPerPixel = dataset.GetElementValue(core.DCMSamplesPerPixel)
	this.PhotometricInterpretation = dataset.GetElementValue(core.DCMPhotometricInterpretation)
	this.Rows = dataset.GetElementValue(core.DCMRows)
	this.Columns = dataset.GetElementValue(core.DCMColumns)
	this.BitsAllocated = dataset.GetElementValue(core.DCMBitsAllocated)
	this.BitsStored = dataset.GetElementValue(core.DCMBitsStored)
	this.HighBit = dataset.GetElementValue(core.DCMHighBit)
	this.PixelRepresentation = dataset.GetElementValue(core.DCMPixelPresentation)
	this.PlanarConfiguration = dataset.GetElementValue(core.DCMPlanarConfiguration)
	this.WindowCenter = dataset.GetElementValue(core.DCMWindowCenter)
	this.WindowWidth = dataset.GetElementValue(core.DCMWindowWidth)
	this.NumberOfFrames = dataset.GetElementValue(core.DCMNumberOfFrames)
}
