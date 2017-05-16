package models

import "github.com/grayzone/godcm/core"

type Study struct {
	StudyInstanceUID       string
	StudyDate              string
	StudyTime              string
	ReferringPhysicianName string
	StudyID                string
	AccessionNumber        string
	Series                 []Series
}

func (this *Study) Parse(dataset core.DcmDataset) {
	this.StudyInstanceUID = dataset.GetElementValue(core.DCMStudyInstanceUID)
	this.StudyDate = dataset.GetElementValue(core.DCMStudyDate)
	this.StudyTime = dataset.GetElementValue(core.DCMStudyTime)
	this.ReferringPhysicianName = dataset.GetElementValue(core.DCMReferringPhysicianName)
	this.StudyID = dataset.GetElementValue(core.DCMStudyID)
	this.AccessionNumber = dataset.GetElementValue(core.DCMAccessionNumber)

	for i := range this.Series {
		this.Series[i].Parse(dataset)
	}
}
