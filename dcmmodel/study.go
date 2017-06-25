package dcmmodel

import "github.com/grayzone/godcm/core"

type Study struct {
	StudyInstanceUID       string `orm:"unique;column(studyinstanceuid)"`
	StudyDate              string `orm:"column(studydate)"`
	StudyTime              string `orm:"column(studytime)"`
	ReferringPhysicianName string `orm:"column(referringphysicianname)"`
	StudyID                string `orm:"column(studyid)"`
	AccessionNumber        string `orm:"column(accessionnumber)"`
	//	Series                 []Series `orm:"-"`
}

func (this *Study) Parse(dataset core.DcmDataset) {
	this.StudyInstanceUID = dataset.GetElementValue(core.DCMStudyInstanceUID)
	this.StudyDate = dataset.GetElementValue(core.DCMStudyDate)
	this.StudyTime = dataset.GetElementValue(core.DCMStudyTime)
	this.ReferringPhysicianName = dataset.GetElementValue(core.DCMReferringPhysicianName)
	this.StudyID = dataset.GetElementValue(core.DCMStudyID)
	this.AccessionNumber = dataset.GetElementValue(core.DCMAccessionNumber)

	/*
		var s Series
		s.Parse(dataset)
		this.Series = append(this.Series, s)
	*/
}
