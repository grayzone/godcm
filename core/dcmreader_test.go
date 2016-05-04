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

func TestDcmReaderReadFileDICOMWithoutReadValue(t *testing.T) {
	cases := []struct {
		in   string
		want DcmDataset
	}{
		{gettestdatafolder() + "GH220.dcm", DcmDataset{}},
		{gettestdatafolder() + "GH178.dcm", DcmDataset{}},
		{gettestdatafolder() + "xr_chest.dcm", DcmDataset{}},
		{gettestdatafolder() + "xr_chicken2.dcm", DcmDataset{}},
		{gettestdatafolder() + "CT-MONO2-16-ankle", DcmDataset{}},
		{gettestdatafolder() + "T14/IM-0001-0001.dcm", DcmDataset{}},
		{gettestdatafolder() + "GH184.dcm", DcmDataset{}},
		{gettestdatafolder() + "GH064.dcm", DcmDataset{}},
		{gettestdatafolder() + "GH133.dcm", DcmDataset{}},
		{gettestdatafolder() + "GH179A.dcm", DcmDataset{}},
		{gettestdatafolder() + "CT1_J2KI", DcmDataset{}},
		{gettestdatafolder() + "GH223.dcm", DcmDataset{}},
		{gettestdatafolder() + "GH195.dcm", DcmDataset{}},
		{gettestdatafolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_ELE_undefinded_length.dcm", DcmDataset{}},
		{gettestdatafolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", DcmDataset{}},
		{gettestdatafolder() + "MR-MONO2-8-16x-heart.dcm", DcmDataset{}},

		/*
			{gettestdatafolder() + "GH179B.dcm", DcmDataset{}}, // incomplete file
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
		want DcmDataset
	}{
		{gettestdatafolder() + "GH220.dcm", DcmDataset{}},
		{gettestdatafolder() + "GH178.dcm", DcmDataset{}},
		{gettestdatafolder() + "xr_chest.dcm", DcmDataset{}},
		{gettestdatafolder() + "xr_chicken2.dcm", DcmDataset{}},
		{gettestdatafolder() + "CT-MONO2-16-ankle", DcmDataset{}},
		{gettestdatafolder() + "T14/IM-0001-0001.dcm", DcmDataset{}},
		{gettestdatafolder() + "GH184.dcm", DcmDataset{}},
		{gettestdatafolder() + "GH064.dcm", DcmDataset{}},
		{gettestdatafolder() + "GH133.dcm", DcmDataset{}},
		{gettestdatafolder() + "GH179A.dcm", DcmDataset{}},
		{gettestdatafolder() + "CT1_J2KI", DcmDataset{}},
		{gettestdatafolder() + "GH223.dcm", DcmDataset{}},
		{gettestdatafolder() + "GH195.dcm", DcmDataset{}},
		{gettestdatafolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", DcmDataset{}},
		{gettestdatafolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_ELE_undefinded_length.dcm", DcmDataset{}},
		{gettestdatafolder() + "MR-MONO2-8-16x-heart.dcm", DcmDataset{}},

		/*
			{gettestdatafolder() + "GH179B.dcm", DcmDataset{}}, // incomplete file
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

func TestFileMetaInformationGroupLength(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{gettestdatafolder() + "T14/IM-0001-0001.dcm", "194"},
		{gettestdatafolder() + "MR-MONO2-8-16x-heart.dcm", "132"},
		{gettestdatafolder() + "CT-MONO2-16-ankle", "192"},
		{gettestdatafolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", "212"},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.ReadFile(c.in)

		got := reader.Meta.FileMetaInformationGroupLength()
		if got != c.want {
			t.Errorf("FileMetaInformationGroupLength() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestFileMetaInformationVersion(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{gettestdatafolder() + "T14/IM-0001-0001.dcm", "0001"},
		{gettestdatafolder() + "MR-MONO2-8-16x-heart.dcm", "0001"},
		{gettestdatafolder() + "CT-MONO2-16-ankle", "0001"},
		{gettestdatafolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", "0001"},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.ReadFile(c.in)

		got := reader.Meta.FileMetaInformationVersion()
		if got != c.want {
			t.Errorf("FileMetaInformationVersion() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestMediaStorageSOPClassUID(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{gettestdatafolder() + "T14/IM-0001-0001.dcm", "1.2.840.10008.5.1.4.1.1.4"},
		{gettestdatafolder() + "MR-MONO2-8-16x-heart.dcm", "1.2.840.10008.5.1.4.1.1.4"},
		{gettestdatafolder() + "CT-MONO2-16-ankle", "1.2.840.10008.5.1.4.1.1.7"},
		{gettestdatafolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", "1.2.840.10008.5.1.4.1.1.2"},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.ReadFile(c.in)

		got := reader.Meta.MediaStorageSOPClassUID()
		if got != c.want {
			t.Errorf("MediaStorageSOPClassUID() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestMediaStorageSOPInstanceUID(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{gettestdatafolder() + "T14/IM-0001-0001.dcm", "1.3.12.2.1107.5.2.5.11090.5.0.582504825601085"},
		{gettestdatafolder() + "MR-MONO2-8-16x-heart.dcm", "999.999.2.19960619.163000.1.103"},
		{gettestdatafolder() + "CT-MONO2-16-ankle", "1.2.840.113619.2.1.2411.1031152382.365.1.736169244"},
		{gettestdatafolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", "1.2.276.0.7230010.3.1.4.1787205428.2345.1071048146.1"},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.ReadFile(c.in)

		got := reader.Meta.MediaStorageSOPInstanceUID()
		if got != c.want {
			t.Errorf("MediaStorageSOPInstanceUID() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestTransferSyntaxUID(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{gettestdatafolder() + "T14/IM-0001-0001.dcm", "1.2.840.10008.1.2.4.91"},
		{gettestdatafolder() + "MR-MONO2-8-16x-heart.dcm", "1.2.840.10008.1.2.1"},
		{gettestdatafolder() + "CT-MONO2-16-ankle", "1.2.840.10008.1.2"},
		{gettestdatafolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", "1.2.840.10008.1.2.2"},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.ReadFile(c.in)

		got := reader.Meta.TransferSyntaxUID()
		if got != c.want {
			t.Errorf("TransferSyntaxUID() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestImplementationClassUID(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{gettestdatafolder() + "T14/IM-0001-0001.dcm", "1.3.6.1.4.1.19291.2.1"},
		{gettestdatafolder() + "MR-MONO2-8-16x-heart.dcm", "999.999"},
		{gettestdatafolder() + "CT-MONO2-16-ankle", "1.2.840.113619.6.5"},
		{gettestdatafolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", "1.2.276.0.7230010.3.0.3.5.4"},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.ReadFile(c.in)

		got := reader.Meta.ImplementationClassUID()
		if got != c.want {
			t.Errorf("ImplementationClassUID() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestImplementationVersionName(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{gettestdatafolder() + "T14/IM-0001-0001.dcm", "OSIRIX001"},
		{gettestdatafolder() + "MR-MONO2-8-16x-heart.dcm", ""},
		{gettestdatafolder() + "CT-MONO2-16-ankle", "1_2_5"},
		{gettestdatafolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", "OFFIS_DCMTK_354"},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.ReadFile(c.in)

		got := reader.Meta.ImplementationVersionName()
		if got != c.want {
			t.Errorf("ImplementationVersionName() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestSourceApplicationEntityTitle(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{gettestdatafolder() + "T14/IM-0001-0001.dcm", "OsiriX"},
		{gettestdatafolder() + "MR-MONO2-8-16x-heart.dcm", ""},
		{gettestdatafolder() + "CT-MONO2-16-ankle", "CTN_STORAGE"},
		{gettestdatafolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", "CLUNIE1"},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.ReadFile(c.in)

		got := reader.Meta.SourceApplicationEntityTitle()
		if got != c.want {
			t.Errorf("SourceApplicationEntityTitle() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestSendingApplicationEntityTitle(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{gettestdatafolder() + "T14/IM-0001-0001.dcm", ""},
		{gettestdatafolder() + "MR-MONO2-8-16x-heart.dcm", ""},
		{gettestdatafolder() + "CT-MONO2-16-ankle", ""},
		{gettestdatafolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", ""},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.ReadFile(c.in)

		got := reader.Meta.SendingApplicationEntityTitle()
		if got != c.want {
			t.Errorf("SendingApplicationEntityTitle() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestReceivingApplicationEntityTitle(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{gettestdatafolder() + "T14/IM-0001-0001.dcm", ""},
		{gettestdatafolder() + "MR-MONO2-8-16x-heart.dcm", ""},
		{gettestdatafolder() + "CT-MONO2-16-ankle", ""},
		{gettestdatafolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", ""},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.ReadFile(c.in)

		got := reader.Meta.ReceivingApplicationEntityTitle()
		if got != c.want {
			t.Errorf("ReceivingApplicationEntityTitle() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestPrivateInformationCreatorUID(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{gettestdatafolder() + "T14/IM-0001-0001.dcm", ""},
		{gettestdatafolder() + "MR-MONO2-8-16x-heart.dcm", ""},
		{gettestdatafolder() + "CT-MONO2-16-ankle", ""},
		{gettestdatafolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", ""},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.ReadFile(c.in)

		got := reader.Meta.PrivateInformationCreatorUID()
		if got != c.want {
			t.Errorf("PrivateInformationCreatorUID() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestPrivateInformation(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{gettestdatafolder() + "T14/IM-0001-0001.dcm", ""},
		{gettestdatafolder() + "MR-MONO2-8-16x-heart.dcm", ""},
		{gettestdatafolder() + "CT-MONO2-16-ankle", ""},
		{gettestdatafolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", ""},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.ReadFile(c.in)

		got := reader.Meta.PrivateInformation()
		if got != c.want {
			t.Errorf("PrivateInformation() %s, want '%v' got '%v'", c.in, c.want, got)
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
		{gettestdatafolder() + "GH223.dcm", ""},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		err := reader.ReadFile(c.in)
		if err != nil {
			t.Errorf("DcmReader.ReadFile(): %s", err.Error())
		}
		got := reader.Dataset.PatientName()
		if got != c.want {
			t.Errorf("GetPatientName() %s, want '%v' got '%v'", c.in, c.want, got)
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
		{gettestdatafolder() + "GH220.dcm", ""},
		{gettestdatafolder() + "CT-MONO2-16-ankle", ""},
		{gettestdatafolder() + "GH223.dcm", ""},
		{gettestdatafolder() + "MR-MONO2-8-16x-heart.dcm", ""},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		err := reader.ReadFile(c.in)
		if err != nil {
			t.Errorf("DcmReader.ReadFile(): %s", err.Error())
		}
		got := reader.Dataset.PatientID()
		if got != c.want {
			t.Errorf("GetPatientID() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestGetModality(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{gettestdatafolder() + "GH178.dcm", "CT"},
		{gettestdatafolder() + "xr_chest.dcm", "CR"},
		{gettestdatafolder() + "xr_chicken2.dcm", "CR"},
		{gettestdatafolder() + "MR-MONO2-8-16x-heart.dcm", "MR"},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.ReadFile(c.in)
		got := reader.Dataset.Modality()
		if got != c.want {
			t.Errorf("GetModality() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}
