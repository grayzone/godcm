package dcmmodel

import "github.com/grayzone/godcm/core"

type Patient struct {
	PatientName      string  `orm:"column(patientname)"`
	PatientID        string  `orm:"column(patientid)"`
	PatientBirthDate string  `orm:"column(patientbirthdate)"`
	PatientSex       string  `orm:"column(patientsex)"`
	Study            []Study `orm:"-"`
}

func (this *Patient) Parse(dataset core.DcmDataset) {
	this.PatientName = dataset.GetElementValue(core.DCMPatientName)
	this.PatientID = dataset.GetElementValue(core.DCMPatientID)
	this.PatientBirthDate = dataset.GetElementValue(core.DCMPatientBirthDate)
	this.PatientSex = dataset.GetElementValue(core.DCMPatientSex)

	var s Study
	s.Parse(dataset)
	this.Study = append(this.Study, s)
}
