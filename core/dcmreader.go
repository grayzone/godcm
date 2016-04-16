package core

import "errors"
import "log"

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

	return nil
}

// IsDicom3 is to check the file is supported by DICOM 3.0 or not.
func (reader *DcmReader) IsDicom3() (bool, error) {
	err := reader.FileStream.Skip(128)
	if err != nil {
		return false, err
	}
	b, err := reader.FileStream.Read(int64(len(DICOM3FILEIDENTIFIER)))
	if err != nil {
		return false, err
	}
	log.Println(b, string(b))
	if string(b) == DICOM3FILEIDENTIFIER {
		return true, nil
	}
	return false, errors.New("Only supprot DICOM 3.0.")
}
