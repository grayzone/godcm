package core

import "testing"

func TestDcmFileStreamOpen(t *testing.T) {
	var s DcmFileStream
	err := s.Open()
	if err == nil {
		t.Error("DcmFileStream.Open() failed")
	}
}

func TestDcmFileStreamClose(t *testing.T) {
	var s DcmFileStream
	err := s.Close()
	if err != nil {
		t.Error("DcmFileStream.Close() failed")
	}
}
