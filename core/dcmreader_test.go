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
		{gettestdatafolder() + "GH178.dcm", DcmDataSet{}},
		{gettestdatafolder() + "xr_chest.dcm", DcmDataSet{}},
		{gettestdatafolder() + "xr_chicken2.dcm", DcmDataSet{}},
		{gettestdatafolder() + "CT-MONO2-16-ankle", DcmDataSet{}},
		{gettestdatafolder() + "T14/IM-0001-0001.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH184.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH064.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH133.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH179A.dcm", DcmDataSet{}},
		{gettestdatafolder() + "CT1_J2KI", DcmDataSet{}},
		{gettestdatafolder() + "GH223.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH195.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_ELE_undefinded_length.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", DcmDataSet{}},
		{gettestdatafolder() + "MR-MONO2-8-16x-heart.dcm", DcmDataSet{}},

		/*
			{gettestdatafolder() + "GH179B.dcm", DcmDataSet{}}, // incomplete file
		*/
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
		{gettestdatafolder() + "GH178.dcm", DcmDataSet{}},
		{gettestdatafolder() + "xr_chest.dcm", DcmDataSet{}},
		{gettestdatafolder() + "xr_chicken2.dcm", DcmDataSet{}},
		{gettestdatafolder() + "CT-MONO2-16-ankle", DcmDataSet{}},
		{gettestdatafolder() + "T14/IM-0001-0001.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH184.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH064.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH133.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH179A.dcm", DcmDataSet{}},
		{gettestdatafolder() + "CT1_J2KI", DcmDataSet{}},
		{gettestdatafolder() + "GH223.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH195.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", DcmDataSet{}},
		{gettestdatafolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_ELE_undefinded_length.dcm", DcmDataSet{}},
		{gettestdatafolder() + "MR-MONO2-8-16x-heart.dcm", DcmDataSet{}},

		/*
			{gettestdatafolder() + "GH179B.dcm", DcmDataSet{}}, // incomplete file
		*/
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

func TestGetPatientID(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{gettestdatafolder() + "GH178.dcm", "55555"},
		{gettestdatafolder() + "xr_chest.dcm", "234"},
		{gettestdatafolder() + "xr_chicken2.dcm", "CHICKEN"},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		err := reader.ReadFile(c.in)
		if err != nil {
			t.Errorf("DcmReader.ReadFile(): %s", err.Error())
		}
		got, err := reader.GetPatientID()
		if err != nil {
			t.Errorf("GetPatientID() %s, error : %s", c.in, err.Error())
		}
		if got != c.want {
			t.Errorf("GetPatientID() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestGetPatientIDFailed(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{gettestdatafolder() + "GH220.dcm", ""},
		{gettestdatafolder() + "CT-MONO2-16-ankle", ""},
		{gettestdatafolder() + "GH223.dcm", ""},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		err := reader.ReadFile(c.in)
		got, err := reader.GetPatientID()
		if err == nil || got != c.want {
			t.Errorf("GetPatientID() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestGetPatientName(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{gettestdatafolder() + "T14/IM-0001-0001.dcm", "WRIX"},
		{gettestdatafolder() + "MR-MONO2-8-16x-heart.dcm", "Anonymized"},
		{gettestdatafolder() + "CT-MONO2-16-ankle", "Anonymized"},
		{gettestdatafolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", "CompressedSamples^CT1"},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		err := reader.ReadFile(c.in)
		if err != nil {
			t.Errorf("DcmReader.ReadFile(): %s", err.Error())
		}
		got, err := reader.GetPatientName()
		if err != nil {
			t.Errorf("GetPatientName() %s, error : %s", c.in, err.Error())
		}
		if got != c.want {
			t.Errorf("GetPatientName() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestGetPatientNameFailed(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{gettestdatafolder() + "GH223.dcm", ""},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		err := reader.ReadFile(c.in)
		got, err := reader.GetPatientName()
		if err == nil || got != c.want {
			t.Errorf("GetPatientName() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}
