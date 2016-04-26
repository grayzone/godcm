package core

import (
	"errors"
	"log"
)

// DICOM3FILEIDENTIFIER is the DiCOM index in the file header.
const DICOM3FILEIDENTIFIER = "DICM"

// DcmReader is to read DICOM file
type DcmReader struct {
	DcmFileStream
	Meta        DcmMetaInfo
	Dataset     DcmDataSet
	IsReadValue bool
	IsReadPixel bool
}

// NewDcmReader create a new instance of DcmReader.
func NewDcmReader() *DcmReader {
	var reader DcmReader
	reader.Meta = *NewDcmMetaInfo()
	return &reader
}

// ReadFile is to read dicom file.
func (reader *DcmReader) ReadFile(filename string) error {
	reader.FileName = filename
	err := reader.Open()
	if err != nil {
		return err
	}
	defer reader.Close()
	isDCM3, err := reader.IsDicom3()
	if !isDCM3 {
		return err
	}

	//read dicom file meta information
	err = reader.Meta.Read(&reader.DcmFileStream)
	log.Println(reader.Meta.Elements, err)
	if err != nil {
		return err
	}

	isExplicitVR, err := reader.Meta.IsExplicitVR()
	if err != nil {
		return err
	}

	// read dicom dataset
	err = reader.Dataset.Read(&reader.DcmFileStream, isExplicitVR, reader.IsReadValue, reader.IsReadPixel)
	if err != nil {
		return err
	}

	//	log.Println(dcmfile.FileDataSet.Elements)

	return nil
}

// IsDicom3 is to check the file is supported by DICOM 3.0 or not.
func (reader *DcmReader) IsDicom3() (bool, error) {
	_, err := reader.Skip(128)
	if err != nil {
		return false, err
	}
	b, err := reader.Read(int64(len(DICOM3FILEIDENTIFIER)))
	if err != nil {
		return false, err
	}
	if string(b) == DICOM3FILEIDENTIFIER {
		reader.Putback(132)
		return true, nil
	}
	return false, errors.New("Only supprot DICOM 3.0.")
}
