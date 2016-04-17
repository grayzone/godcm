package core

import (
	"errors"
	"log"
)

// DICOM3FILEIDENTIFIER is the DiCOM index in the file header.
const DICOM3FILEIDENTIFIER = "DICM"

// DcmReader is to read DICOM file
type DcmReader struct {
	Dataset    DcmDataset
	FileStream DcmFileStream
}

// ReadFile is to read dicom file.
func (reader *DcmReader) ReadFile(filename string) error {
	reader.FileStream.FileName = filename
	err := reader.FileStream.Open()
	if err != nil {
		return err
	}
	defer reader.FileStream.Close()
	isDCM3, err := reader.IsDicom3()
	if !isDCM3 {
		return err
	}

	//read dicom file meta information
	var dcmfile DcmFile
	err = dcmfile.FileMetaInfo.Read(&reader.FileStream)
	if err != nil {
		return err
	}

	log.Println(dcmfile.FileMetaInfo)

	return nil
}

// IsDicom3 is to check the file is supported by DICOM 3.0 or not.
func (reader *DcmReader) IsDicom3() (bool, error) {
	_, err := reader.FileStream.Skip(128)
	if err != nil {
		return false, err
	}
	b, err := reader.FileStream.Read(int64(len(DICOM3FILEIDENTIFIER)))
	if err != nil {
		return false, err
	}
	if string(b) == DICOM3FILEIDENTIFIER {
		return true, nil
	}
	return false, errors.New("Only supprot DICOM 3.0.")
}

// ReadOneDcmElement is to parse one dicom element
func (reader *DcmReader) ReadOneDcmElement() (DcmElement, error) {
	var elem DcmElement

	return elem, nil

}
