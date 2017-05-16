package models

import "github.com/grayzone/godcm/core"

type Patient struct {
	PatientName      string
	PatientID        string
	PatientBirthDate string
	PatientSex       string
	Study            []Study
}

func (this *Patient) Parse(dataset core.DcmDataset) {
	this.PatientName = dataset.GetElementValue(core.DCMPatientName)
	this.PatientID = dataset.GetElementValue(core.DCMPatientID)
	this.PatientBirthDate = dataset.GetElementValue(core.DCMPatientBirthDate)
	this.PatientSex = dataset.GetElementValue(core.DCMPatientSex)

	for i := range this.Study {
		this.Study[i].Parse(dataset)
	}
}
