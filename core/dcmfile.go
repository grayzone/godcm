package core

// DcmFile contains all the elements from file
type DcmFile struct {
	FileMetaInfo DcmMetaInfo
	FileDataSet  DcmDataSet
}

// NewDcmFile is to create an instance of DcmFile
func NewDcmFile() *DcmFile {
	var f DcmFile
	f.FileMetaInfo = *NewDcmMetaInfo()

	return &f
}
