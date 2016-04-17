package core

import (
	"os"
	"testing"
)

func gettestdatafolder() string {
	cur, err := os.Getwd()
	if err != nil {
		return ""
	}
	result := cur + "/../test/data/"
	return result
}

func TestDcmReaderReadFileNONDICOM(t *testing.T) {
	cases := []struct {
		in   string
		want DcmDataset
	}{
		{"", DcmDataset{}},
		{gettestdatafolder() + "minimumdict.xml", DcmDataset{}},
	}
	for _, c := range cases {
		var reader DcmReader
		err := reader.ReadFile(c.in)
		if err == nil {
			t.Errorf("DcmReader.ReadFile(): %s", err.Error())
			return
		}
	}
}

func TestDcmReaderReadFileDICOM(t *testing.T) {
	cases := []struct {
		in   string
		want DcmDataset
	}{
		{gettestdatafolder() + "GH220.dcm", DcmDataset{}},
		{gettestdatafolder() + "GH064.dcm", DcmDataset{}},
		{gettestdatafolder() + "GH133.dcm", DcmDataset{}},
		{gettestdatafolder() + "GH178.dcm", DcmDataset{}},
		{gettestdatafolder() + "GH179A.dcm", DcmDataset{}},
		{gettestdatafolder() + "GH184.dcm", DcmDataset{}},
	}
	for _, c := range cases {
		var reader DcmReader
		err := reader.ReadFile(c.in)
		if err != nil {
			t.Errorf("DcmReader.ReadFile(): %s", err.Error())
			return
		}
	}
}
