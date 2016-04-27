package core

import (
	"errors"
	"os"
	"strings"
)

// DcmFileStream is to read binary file to bytes.
type DcmFileStream struct {
	FileName    string
	fileHandler *os.File
	Size        int64
	Position    int64
}

// Open is to open a file
func (s *DcmFileStream) Open() error {
	var err error
	s.fileHandler, err = os.Open(s.FileName)
	if err != nil {
		return err
	}
	s.Size, err = s.fileHandler.Seek(0, os.SEEK_END)
	if err != nil {
		return err
	}
	_, err = s.fileHandler.Seek(0, os.SEEK_SET)
	return err
}

// Close is to close a file
func (s *DcmFileStream) Close() error {
	if s.fileHandler != nil {
		return s.fileHandler.Close()
	}
	return nil
}

// Skip the bytes by given length
func (s *DcmFileStream) Skip(skiplength int64) (int64, error) {
	var result int64
	if s.fileHandler == nil {
		return result, errors.New("The file is not opened yet.")
	}
	if skiplength == 0 {
		return result, nil
	}
	pos, _ := s.fileHandler.Seek(0, os.SEEK_CUR)

	if s.Size-pos < skiplength {
		result = s.Size - pos
	} else {
		result = skiplength
	}
	_, err := s.fileHandler.Seek(skiplength, os.SEEK_CUR)

	s.Position += result
	return result, err
}

// SeekToBegin set the handler to the beginning of the file.
func (s *DcmFileStream) SeekToBegin() error {
	_, err := s.fileHandler.Seek(0, os.SEEK_SET)

	s.Position = 0
	return err
}

// Putback the bytes by given length
func (s *DcmFileStream) Putback(num int64) error {
	if num == 0 {
		return nil
	}
	pos, _ := s.fileHandler.Seek(0, os.SEEK_CUR)
	if num > pos {
		return errors.New("Parser failure: Putback operation failed")
	}
	_, err := s.fileHandler.Seek(-num, os.SEEK_CUR)
	if err != nil {
		return err
	}
	s.Position -= num
	return nil
}

// Eos is to check the end of the DICOM file.
func (s *DcmFileStream) Eos() bool {
	if s.fileHandler == nil {
		return true
	}
	size, _ := s.fileHandler.Seek(0, os.SEEK_CUR)
	//	log.Println(size, s.Size)
	return size == s.Size

}

// Read is to read bytes by given length
func (s *DcmFileStream) Read(length int64) ([]byte, error) {
	if length == 0 {
		return []byte{}, nil
	}
	b := make([]byte, length)
	_, err := s.fileHandler.Read(b)

	s.Position += length
	return b, err
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
