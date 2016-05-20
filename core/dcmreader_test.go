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

func TestPatientName(t *testing.T) {
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
			t.Errorf("PatientName() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestPatientID(t *testing.T) {
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
			t.Errorf("PatientID() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestModality(t *testing.T) {
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
			t.Errorf("Modality() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestRows(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{gettestdatafolder() + "GH178.dcm", "512"},
		{gettestdatafolder() + "xr_chest.dcm", "2048"},
		{gettestdatafolder() + "xr_chicken2.dcm", "3015"},
		{gettestdatafolder() + "MR-MONO2-8-16x-heart.dcm", "256"},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.ReadFile(c.in)
		got := reader.Dataset.Rows()
		if got != c.want {
			t.Errorf("Rows() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestColumns(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{gettestdatafolder() + "GH178.dcm", "512"},
		{gettestdatafolder() + "xr_chest.dcm", "2495"},
		{gettestdatafolder() + "xr_chicken2.dcm", "2505"},
		{gettestdatafolder() + "MR-MONO2-8-16x-heart.dcm", "256"},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.ReadFile(c.in)
		got := reader.Dataset.Columns()
		if got != c.want {
			t.Errorf("Columns() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestWindowCenter(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{gettestdatafolder() + "GH178.dcm", "-2"},
		{gettestdatafolder() + "xr_chest.dcm", "2.04750000E+03"},
		{gettestdatafolder() + "xr_chicken2.dcm", ""},
		{gettestdatafolder() + "MR-MONO2-8-16x-heart.dcm", ""},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.ReadFile(c.in)
		got := reader.Dataset.WindowCenter()
		if got != c.want {
			t.Errorf("WindowCenter() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestWindowWidth(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{gettestdatafolder() + "GH178.dcm", "1"},
		{gettestdatafolder() + "xr_chest.dcm", "4.09500000E+03"},
		{gettestdatafolder() + "xr_chicken2.dcm", ""},
		{gettestdatafolder() + "MR-MONO2-8-16x-heart.dcm", ""},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.ReadFile(c.in)
		got := reader.Dataset.WindowWidth()
		if got != c.want {
			t.Errorf("WindowWidth() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestSOPInstanceUID(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{gettestdatafolder() + "GH178.dcm", "1.2.840.113619.2.284.3.17442826.413.1388989542.640"},
		{gettestdatafolder() + "xr_chest.dcm", "1.3.51.0.7.99.2155959091.28444.877621460.2"},
		{gettestdatafolder() + "xr_chicken2.dcm", "1.2.392.200036.9125.4.0.219104458.164963328.1055409964"},
		{gettestdatafolder() + "MR-MONO2-8-16x-heart.dcm", "999.999.2.19960619.163000.1.103"},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.ReadFile(c.in)
		got := reader.Dataset.SOPInstanceUID()
		if got != c.want {
			t.Errorf("SOPInstanceUID() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestPixelData(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{gettestdatafolder() + "GH178.dcm", 0},
		{gettestdatafolder() + "xr_chest.dcm", 10219520},
		{gettestdatafolder() + "xr_chicken2.dcm", 659012},
		{gettestdatafolder() + "MR-MONO2-8-16x-heart.dcm", 1048576},
		{gettestdatafolder() + "IM-0001-0010.dcm", 103518},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.IsReadPixel = true
		reader.ReadFile(c.in)
		p := reader.Dataset.PixelData()
		var got int
		if p != nil {
			got = len(p)
		}
		if got != c.want {
			t.Errorf("SOPInstanceUID() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestNumberOfFrames(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{gettestdatafolder() + "GH178.dcm", ""},
		{gettestdatafolder() + "xr_chest.dcm", ""},
		{gettestdatafolder() + "xr_chicken2.dcm", ""},
		{gettestdatafolder() + "MR-MONO2-8-16x-heart.dcm", "16"},
		{gettestdatafolder() + "IM-0001-0010.dcm", ""},
		{gettestdatafolder() + "T23/IM-0001-0001.dcm", ""},
		{gettestdatafolder() + "T14/IM-0001-0001.dcm", ""},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.IsReadPixel = true
		reader.ReadFile(c.in)
		got := reader.Dataset.NumberOfFrames()
		if got != c.want {
			t.Errorf("NumberOfFrames() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestBitsAllocated(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{gettestdatafolder() + "GH178.dcm", "16"},
		{gettestdatafolder() + "xr_chest.dcm", "16"},
		{gettestdatafolder() + "xr_chicken2.dcm", "16"},
		{gettestdatafolder() + "MR-MONO2-8-16x-heart.dcm", "8"},
		{gettestdatafolder() + "IM-0001-0010.dcm", "16"},
		{gettestdatafolder() + "T23/IM-0001-0001.dcm", "16"},
		{gettestdatafolder() + "T14/IM-0001-0001.dcm", "16"},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.IsReadPixel = true
		reader.ReadFile(c.in)
		got := reader.Dataset.BitsAllocated()
		if got != c.want {
			t.Errorf("BitsAllocated() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestBitsStored(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{gettestdatafolder() + "GH178.dcm", "16"},
		{gettestdatafolder() + "xr_chest.dcm", "12"},
		{gettestdatafolder() + "xr_chicken2.dcm", "12"},
		{gettestdatafolder() + "MR-MONO2-8-16x-heart.dcm", "8"},
		{gettestdatafolder() + "IM-0001-0010.dcm", "12"},
		{gettestdatafolder() + "T23/IM-0001-0001.dcm", "12"},
		{gettestdatafolder() + "T14/IM-0001-0001.dcm", "12"},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.IsReadPixel = true
		reader.ReadFile(c.in)
		got := reader.Dataset.BitsStored()
		if got != c.want {
			t.Errorf("BitsStored() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestHighBit(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{gettestdatafolder() + "GH178.dcm", "15"},
		{gettestdatafolder() + "xr_chest.dcm", "11"},
		{gettestdatafolder() + "xr_chicken2.dcm", "11"},
		{gettestdatafolder() + "MR-MONO2-8-16x-heart.dcm", "7"},
		{gettestdatafolder() + "IM-0001-0010.dcm", "11"},
		{gettestdatafolder() + "T23/IM-0001-0001.dcm", "11"},
		{gettestdatafolder() + "T14/IM-0001-0001.dcm", "11"},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.IsReadPixel = true
		reader.ReadFile(c.in)
		got := reader.Dataset.HighBit()
		if got != c.want {
			t.Errorf("HighBit() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}
