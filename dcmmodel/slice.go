package dcmmodel

import "github.com/grayzone/godcm/core"

type Slice struct {
	SOPInstanceUID       string `orm:"unique;column(sopinstanceuid)"`
	SOPClassUID          string `orm:"column(sopclassuid)"`
	SpecificCharacterSet string `orm:"column(specificcharacterset)"`
	InstanceNumber       string `orm:"column(instancenumber)"`
	PatientOrientation   string `orm:"column(patientorientation)"`
	ContentDate          string `orm:"column(contentdate)"`
	ContentTime          string `orm:"column(contenttime)"`

	//pixel
	SamplesPerPixel           string `orm:"column(samplesperpixel)"`
	PhotometricInterpretation string `orm:"column(photometricinterpretation)"`
	Rows                      string `orm:"column(rows)"`
	Columns                   string `orm:"column(columns)"`
	BitsAllocated             string `orm:"column(bitsallocated)"`
	BitsStored                string `orm:"column(bitsstored)"`
	HighBit                   string `orm:"column(highbit)"`
	PixelRepresentation       string `orm:"column(pixelrepresentation)"`
	PlanarConfiguration       string `orm:"column(planarconfiguration)"`
	WindowCenter              string `orm:"column(windowcenter)"`
	WindowWidth               string `orm:"column(windowwidth)"`
	NumberOfFrames            string `orm:"column(numberofframes)"`
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
