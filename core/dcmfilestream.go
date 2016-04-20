package core

import (
	"bytes"
	"encoding/binary"
	"errors"
	"log"
	"os"
	"strings"
)

// DcmFileStream is to read binary file to bytes.
type DcmFileStream struct {
	FileName    string
	FileHandler *os.File
	Size        int64
	Tell        int64
}

// Open is to open a file
func (s *DcmFileStream) Open() error {
	var err error
	s.FileHandler, err = os.Open(s.FileName)
	if err != nil {
		return err
	}
	s.Size, err = s.FileHandler.Seek(0, os.SEEK_END)
	if err != nil {
		return err
	}
	_, err = s.FileHandler.Seek(0, os.SEEK_SET)
	return err
}

// Close is to close a file
func (s *DcmFileStream) Close() error {
	if s.FileHandler != nil {
		return s.FileHandler.Close()
	}
	return nil
}

// Skip the bytes by given length
func (s *DcmFileStream) Skip(skiplength int64) (int64, error) {
	var result int64
	if s.FileHandler == nil {
		return result, errors.New("The file is not opened yet.")
	}
	if skiplength == 0 {
		return result, nil
	}
	pos, _ := s.FileHandler.Seek(0, os.SEEK_CUR)

	if s.Size-pos < skiplength {
		result = s.Size - pos
	} else {
		result = skiplength
	}
	_, err := s.FileHandler.Seek(skiplength, os.SEEK_CUR)
	s.Tell += result
	return result, err
}

// Putback the bytes by given length
func (s *DcmFileStream) Putback(num int64) error {
	if num == 0 {
		return nil
	}
	pos, _ := s.FileHandler.Seek(0, os.SEEK_CUR)
	if num > pos {
		return errors.New("Parser failure: Putback operation failed")
	}
	_, err := s.FileHandler.Seek(-num, os.SEEK_CUR)
	return err
}

// Eos is to check the end of the DICOM file.
func (s *DcmFileStream) Eos() bool {
	if s.FileHandler == nil {
		return true
	}
	size, _ := s.FileHandler.Seek(0, os.SEEK_CUR)
	return size == s.Size

}

// Read is to read bytes by given length
func (s *DcmFileStream) Read(length int64) ([]byte, error) {
	if length == 0 {
		return []byte{}, nil
	}
	b := make([]byte, length)
	pos, err := s.FileHandler.Read(b)
	s.Tell += int64(pos)
	return b, err
}

// ReadUINT16 is to read a uint16 value from the file.
func (s *DcmFileStream) ReadUINT16() (uint16, error) {
	v, err := s.Read(2)
	if err != nil {
		return 0, err
	}
	var result uint16
	buf := bytes.NewReader(v)
	err = binary.Read(buf, binary.LittleEndian, &result)
	return result, err
}

// ReadUINT32 is to read a uint32 value from the file.
func (s *DcmFileStream) ReadUINT32() (uint32, error) {
	v, err := s.Read(4)
	if err != nil {
		return 0, err
	}
	var result uint32
	buf := bytes.NewReader(v)
	err = binary.Read(buf, binary.LittleEndian, &result)
	return result, err
}

// ReadString is to read a string from the file.
func (s *DcmFileStream) ReadString(slen int64) (string, error) {
	v, err := s.Read(slen)
	if err != nil {
		return "", err
	}
	str := strings.TrimRight(string(v), "\x00")
	return str, nil
}

// ReadDcmTag is to read group and element
func (s *DcmFileStream) ReadDcmTag() (DcmTag, error) {
	var t DcmTag
	var err error
	t.Group, err = s.ReadUINT16()
	if err != nil {
		return t, err
	}
	t.Element, err = s.ReadUINT16()
	if err != nil {
		return t, err
	}
	return t, nil
}

// ReadDcmVR is to read vr
func (s *DcmFileStream) ReadDcmVR() (string, error) {
	return s.ReadString(2)
}

// ReadValueLengthWithExplicitVR gets the value length of the dicom element with explicit VR.
func (s *DcmFileStream) ReadValueLengthWithExplicitVR(vr string) (int64, error) {
	var result int64
	switch vr {
	case "OB", "OD", "OF", "OL", "OW", "SQ", "UC", "UR", "UT", "UN":
		// skip the reserved 2 bytes
		_, err := s.Skip(2)
		if err != nil {
			return 0, err
		}
		l, err := s.ReadUINT32()
		if err != nil {
			return 0, err
		}
		result = int64(l)
	default:
		// read value length
		l, err := s.ReadUINT16()
		if err != nil {
			return 0, err
		}
		result = int64(l)
	}
	return result, nil
}

// ReadValueLengthWithImplicitVR gets the value length of the dicom element with implicit VR.
func (s *DcmFileStream) ReadValueLengthWithImplicitVR(vr string) (int64, error) {
	var result int64
	l, err := s.ReadUINT32()
	if err != nil {
		return 0, err
	}
	result = int64(l)
	return result, nil
}

// ReadDcmElement read one dicom element.
func (s *DcmFileStream) ReadDcmElement(isExplicitVR bool) (DcmElement, error) {
	if isExplicitVR {
		return s.ReadDcmElementWithExplicitVR()
	}
	return s.ReadDcmElementWithImplicitVR()
}

// ReadDcmElementWithExplicitVR read the data element with explicit VR.
func (s *DcmFileStream) ReadDcmElementWithExplicitVR() (DcmElement, error) {
	var elem DcmElement
	var err error
	// read dicom tag
	elem.Tag, err = s.ReadDcmTag()
	if err != nil {
		return elem, err
	}

	// read VR
	elem.VR, err = s.ReadDcmVR()
	if err != nil {
		return elem, err
	}

	//read the value length
	elem.Length, err = s.ReadValueLengthWithExplicitVR(elem.VR)
	if err != nil {
		return elem, err
	}

	// read element value
	elem.Value, err = s.ReadString(elem.Length)

	return elem, err
}

// ReadDcmElementWithImplicitVR read the data element with implicit VR.
func (s *DcmFileStream) ReadDcmElementWithImplicitVR() (DcmElement, error) {
	var elem DcmElement
	var err error

	// read dciom tag
	elem.Tag, err = s.ReadDcmTag()
	if err != nil {
		return elem, err
	}
	// get VR from Dicom Element registry
	err = FindDcmElmentByTag(&elem)
	if err != nil {
		log.Println(err.Error())

	}

	// read the value length
	elem.Length, err = s.ReadValueLengthWithImplicitVR(elem.VR)
	if err != nil {
		return elem, err
	}

	// read element value
	elem.Value, err = s.ReadString(elem.Length)

	if elem.Tag.Group != 0x7fe0 {
		log.Println(elem)
	}

	return elem, err

}
