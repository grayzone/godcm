package core

import (
	"errors"
)

// DICOM3FILEIDENTIFIER is the DiCOM index in the file header.
const DICOM3FILEIDENTIFIER = "DICM"

// DcmReader is to read DICOM file
type DcmReader struct {
	fs          DcmFileStream
	Meta        DcmMetaInfo
	Dataset     DcmDataset
	IsReadValue bool
	IsReadPixel bool
}

// ReadFile is to read dicom file.
func (reader *DcmReader) ReadFile(filename string) error {
	reader.fs.FileName = filename
	err := reader.fs.Open()
	if err != nil {
		return err
	}
	defer reader.fs.Close()
	isDCM3, err := reader.IsDicom3()
	if !isDCM3 {
		return err
	}

	//read dicom file meta information
	err = reader.Meta.Read(&reader.fs)
	if err != nil {
		return err
	}

	isExplicitVR, err := reader.Meta.IsExplicitVR()
	if err != nil {
		return err
	}

	byteOrder, err := reader.Meta.GetByteOrder()
	if err != nil {
		return err
	}
	// read dicom dataset
	err = reader.Dataset.Read(&reader.fs, isExplicitVR, byteOrder, reader.IsReadValue, reader.IsReadPixel)
	if err != nil {
		return err
	}

	return nil
}

// IsDicom3 is to check the file is supported by DICOM 3.0 or not.
func (reader *DcmReader) IsDicom3() (bool, error) {
	_, err := reader.fs.Skip(128)
	if err != nil {
		return false, err
	}
	b, err := reader.fs.Read(int64(len(DICOM3FILEIDENTIFIER)))
	if err != nil {
		return false, err
	}
	if string(b) != DICOM3FILEIDENTIFIER {
		return false, errors.New("Only supprot DICOM 3.0.")
	}
	reader.fs.Putback(132)
	return true, nil
}
