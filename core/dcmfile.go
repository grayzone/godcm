package core

// DcmFile contains all the elements from file
type DcmFile struct {
	FileMetaInfo DcmMetaInfo
	FileDataset  DcmDataset
}
