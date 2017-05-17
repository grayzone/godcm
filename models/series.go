package models

import "github.com/grayzone/godcm/core"

type Series struct {
	SeriesInstanceUID string
	SeriesNumber      string
	Modality          string
	Laterality        string
	Slice             []Slice
}

func (this *Series) Parse(dataset core.DcmDataset) {
	this.SeriesInstanceUID = dataset.GetElementValue(core.DCMSeriesInstanceUID)
	this.SeriesNumber = dataset.GetElementValue(core.DCMSeriesNumber)
	this.Modality = dataset.GetElementValue(core.DCMModality)
	this.Laterality = dataset.GetElementValue(core.DCMLaterality)

	var s Slice
	s.Parse(dataset)
	this.Slice = append(this.Slice, s)
}
