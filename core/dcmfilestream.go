package core

import "os"

// DcmFileStream is to read binary file to bytes.
type DcmFileStream struct {
	FileName    string
	FileHandler *os.File
}

// Open is to open a file
func (s *DcmFileStream) Open() error {
	var err error
	s.FileHandler, err = os.Open(s.FileName)
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
func (s *DcmFileStream) Skip(length int64) error {
	_, err := s.FileHandler.Seek(length, os.SEEK_SET)
	return err
}

// Read is to read bytes by given length
func (s *DcmFileStream) Read(length int64) ([]byte, error) {
	if length == 0 {
		return []byte{}, nil
	}
	b := make([]byte, length)
	_, err := s.FileHandler.Read(b)
	return b, err
}
