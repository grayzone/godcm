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
		want DcmDataSet
	}{
		{"", DcmDataSet{}},
		{gettestdatafolder() + "minimumdict.xml", DcmDataSet{}},
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

func TestDcmReaderReadFileDICOMWithoutReadValue(t *testing.T) {
	cases := []struct {
		in   string
		want DcmDataSet
	}{
		{gettestdatafolder() + "GH220.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH064.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH133.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH178.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH179A.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH184.dcm", DcmDataSet{}},
		{gettestdatafolder() + "CT1_J2KI", DcmDataSet{}},
		{gettestdatafolder() + "CT-MONO2-16-ankle", DcmDataSet{}},
		{gettestdatafolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_ELE_undefinded_length.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH179B.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH195.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH223.dcm", DcmDataSet{}},
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

func TestDcmReaderReadFileDICOM(t *testing.T) {
	cases := []struct {
		in   string
		want DcmDataSet
	}{
		{gettestdatafolder() + "GH220.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH064.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH133.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH178.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH179A.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH184.dcm", DcmDataSet{}},
		{gettestdatafolder() + "CT1_J2KI", DcmDataSet{}},
		{gettestdatafolder() + "CT-MONO2-16-ankle", DcmDataSet{}},
		{gettestdatafolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_ELE_undefinded_length.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH179B.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH195.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH223.dcm", DcmDataSet{}},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		err := reader.ReadFile(c.in)
		if err != nil {
			t.Errorf("DcmReader.ReadFile(): %s", err.Error())
			return
		}
	}
}
