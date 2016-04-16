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

	for range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 131, 14} {
		tag, _ := reader.FileStream.ReadDcmTag()
		log.Println(tag)

		vr, _ := reader.FileStream.ReadDcmVR()
		log.Println(vr)

		if vr == "OB" {
			reader.FileStream.Skip(2)
			l, _ := reader.FileStream.ReadUINT32()
			log.Println(l)
			s, _ := reader.FileStream.ReadString(int64(l))
			log.Println(s)
		} else {
			l, _ := reader.FileStream.ReadUINT16()
			log.Println(l)
			s, _ := reader.FileStream.ReadString(int64(l))
			log.Println(s)
		}

	}

	/*
		for !reader.FileStream.Eos() {

		}
	*/

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
	log.Println(b, string(b))
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
