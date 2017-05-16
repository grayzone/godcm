package core

import (
	"testing"

	"github.com/grayzone/godcm/util"
)

func TestDcmReaderReadFileNONDICOM(t *testing.T) {
	cases := []struct {
		in   string
		want DcmDataset
	}{
		{"", DcmDataset{}},
		{util.GetTestDataFolder() + "minimumdict.xml", DcmDataset{}},
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
		{util.GetTestDataFolder() + "GH220.dcm", DcmDataset{}},
		{util.GetTestDataFolder() + "GH178.dcm", DcmDataset{}},
		{util.GetTestDataFolder() + "xr_chest.dcm", DcmDataset{}},
		{util.GetTestDataFolder() + "xr_chicken2.dcm", DcmDataset{}},
		{util.GetTestDataFolder() + "CT-MONO2-16-ankle", DcmDataset{}},
		{util.GetTestDataFolder() + "T14/IM-0001-0001.dcm", DcmDataset{}},
		{util.GetTestDataFolder() + "GH184.dcm", DcmDataset{}},
		{util.GetTestDataFolder() + "GH064.dcm", DcmDataset{}},
		{util.GetTestDataFolder() + "GH133.dcm", DcmDataset{}},
		{util.GetTestDataFolder() + "GH179A.dcm", DcmDataset{}},
		{util.GetTestDataFolder() + "CT1_J2KI", DcmDataset{}},
		{util.GetTestDataFolder() + "GH223.dcm", DcmDataset{}},
		{util.GetTestDataFolder() + "GH195.dcm", DcmDataset{}},
		{util.GetTestDataFolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_ELE_undefinded_length.dcm", DcmDataset{}},
		{util.GetTestDataFolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", DcmDataset{}},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", DcmDataset{}},

		/*
			{util.GetTestDataFolder() + "GH179B.dcm", DcmDataset{}}, // incomplete file
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
		{util.GetTestDataFolder() + "GH220.dcm", DcmDataset{}},
		{util.GetTestDataFolder() + "GH178.dcm", DcmDataset{}},
		{util.GetTestDataFolder() + "xr_chest.dcm", DcmDataset{}},
		{util.GetTestDataFolder() + "xr_chicken2.dcm", DcmDataset{}},
		{util.GetTestDataFolder() + "CT-MONO2-16-ankle", DcmDataset{}},
		{util.GetTestDataFolder() + "T14/IM-0001-0001.dcm", DcmDataset{}},
		{util.GetTestDataFolder() + "GH184.dcm", DcmDataset{}},
		{util.GetTestDataFolder() + "GH064.dcm", DcmDataset{}},
		{util.GetTestDataFolder() + "GH133.dcm", DcmDataset{}},
		{util.GetTestDataFolder() + "GH179A.dcm", DcmDataset{}},
		{util.GetTestDataFolder() + "CT1_J2KI", DcmDataset{}},
		{util.GetTestDataFolder() + "GH223.dcm", DcmDataset{}},
		{util.GetTestDataFolder() + "GH195.dcm", DcmDataset{}},
		{util.GetTestDataFolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", DcmDataset{}},
		{util.GetTestDataFolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_ELE_undefinded_length.dcm", DcmDataset{}},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", DcmDataset{}},

		/*
			{util.GetTestDataFolder() + "GH179B.dcm", DcmDataset{}}, // incomplete file
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
		{util.GetTestDataFolder() + "T14/IM-0001-0001.dcm", "194"},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", "132"},
		{util.GetTestDataFolder() + "CT-MONO2-16-ankle", "192"},
		{util.GetTestDataFolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", "212"},
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
		{util.GetTestDataFolder() + "T14/IM-0001-0001.dcm", "0001"},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", "0001"},
		{util.GetTestDataFolder() + "CT-MONO2-16-ankle", "0001"},
		{util.GetTestDataFolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", "0001"},
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
		{util.GetTestDataFolder() + "T14/IM-0001-0001.dcm", "1.2.840.10008.5.1.4.1.1.4"},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", "1.2.840.10008.5.1.4.1.1.4"},
		{util.GetTestDataFolder() + "CT-MONO2-16-ankle", "1.2.840.10008.5.1.4.1.1.7"},
		{util.GetTestDataFolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", "1.2.840.10008.5.1.4.1.1.2"},
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
		{util.GetTestDataFolder() + "T14/IM-0001-0001.dcm", "1.3.12.2.1107.5.2.5.11090.5.0.582504825601085"},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", "999.999.2.19960619.163000.1.103"},
		{util.GetTestDataFolder() + "CT-MONO2-16-ankle", "1.2.840.113619.2.1.2411.1031152382.365.1.736169244"},
		{util.GetTestDataFolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", "1.2.276.0.7230010.3.1.4.1787205428.2345.1071048146.1"},
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
		{util.GetTestDataFolder() + "T14/IM-0001-0001.dcm", "1.2.840.10008.1.2.4.91"},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", "1.2.840.10008.1.2.1"},
		{util.GetTestDataFolder() + "CT-MONO2-16-ankle", "1.2.840.10008.1.2"},
		{util.GetTestDataFolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", "1.2.840.10008.1.2.2"},
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
		{util.GetTestDataFolder() + "T14/IM-0001-0001.dcm", "1.3.6.1.4.1.19291.2.1"},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", "999.999"},
		{util.GetTestDataFolder() + "CT-MONO2-16-ankle", "1.2.840.113619.6.5"},
		{util.GetTestDataFolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", "1.2.276.0.7230010.3.0.3.5.4"},
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
		{util.GetTestDataFolder() + "T14/IM-0001-0001.dcm", "OSIRIX001"},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", ""},
		{util.GetTestDataFolder() + "CT-MONO2-16-ankle", "1_2_5"},
		{util.GetTestDataFolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", "OFFIS_DCMTK_354"},
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
		{util.GetTestDataFolder() + "T14/IM-0001-0001.dcm", "OsiriX"},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", ""},
		{util.GetTestDataFolder() + "CT-MONO2-16-ankle", "CTN_STORAGE"},
		{util.GetTestDataFolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", "CLUNIE1"},
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
		{util.GetTestDataFolder() + "T14/IM-0001-0001.dcm", ""},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", ""},
		{util.GetTestDataFolder() + "CT-MONO2-16-ankle", ""},
		{util.GetTestDataFolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", ""},
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
		{util.GetTestDataFolder() + "T14/IM-0001-0001.dcm", ""},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", ""},
		{util.GetTestDataFolder() + "CT-MONO2-16-ankle", ""},
		{util.GetTestDataFolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", ""},
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
		{util.GetTestDataFolder() + "T14/IM-0001-0001.dcm", ""},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", ""},
		{util.GetTestDataFolder() + "CT-MONO2-16-ankle", ""},
		{util.GetTestDataFolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", ""},
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
		{util.GetTestDataFolder() + "T14/IM-0001-0001.dcm", ""},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", ""},
		{util.GetTestDataFolder() + "CT-MONO2-16-ankle", ""},
		{util.GetTestDataFolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", ""},
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
		{util.GetTestDataFolder() + "T14/IM-0001-0001.dcm", "WRIX"},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", "Anonymized"},
		{util.GetTestDataFolder() + "CT-MONO2-16-ankle", "Anonymized"},
		{util.GetTestDataFolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", "CompressedSamples^CT1"},
		{util.GetTestDataFolder() + "GH223.dcm", ""},
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
		{util.GetTestDataFolder() + "GH178.dcm", "55555"},
		{util.GetTestDataFolder() + "xr_chest.dcm", "234"},
		{util.GetTestDataFolder() + "xr_chicken2.dcm", "CHICKEN"},
		{util.GetTestDataFolder() + "GH220.dcm", ""},
		{util.GetTestDataFolder() + "CT-MONO2-16-ankle", ""},
		{util.GetTestDataFolder() + "GH223.dcm", ""},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", ""},
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
		{util.GetTestDataFolder() + "GH178.dcm", "CT"},
		{util.GetTestDataFolder() + "xr_chest.dcm", "CR"},
		{util.GetTestDataFolder() + "xr_chicken2.dcm", "CR"},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", "MR"},
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
		{util.GetTestDataFolder() + "GH178.dcm", "512"},
		{util.GetTestDataFolder() + "xr_chest.dcm", "2048"},
		{util.GetTestDataFolder() + "xr_chicken2.dcm", "3015"},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", "256"},
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
		{util.GetTestDataFolder() + "GH178.dcm", "512"},
		{util.GetTestDataFolder() + "xr_chest.dcm", "2495"},
		{util.GetTestDataFolder() + "xr_chicken2.dcm", "2505"},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", "256"},
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
		{util.GetTestDataFolder() + "GH178.dcm", "-2"},
		{util.GetTestDataFolder() + "xr_chest.dcm", "2.04750000E+03"},
		{util.GetTestDataFolder() + "xr_chicken2.dcm", ""},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", ""},
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
		{util.GetTestDataFolder() + "GH178.dcm", "1"},
		{util.GetTestDataFolder() + "xr_chest.dcm", "4.09500000E+03"},
		{util.GetTestDataFolder() + "xr_chicken2.dcm", ""},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", ""},
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
		{util.GetTestDataFolder() + "GH178.dcm", "1.2.840.113619.2.284.3.17442826.413.1388989542.640"},
		{util.GetTestDataFolder() + "xr_chest.dcm", "1.3.51.0.7.99.2155959091.28444.877621460.2"},
		{util.GetTestDataFolder() + "xr_chicken2.dcm", "1.2.392.200036.9125.4.0.219104458.164963328.1055409964"},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", "999.999.2.19960619.163000.1.103"},
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
		{util.GetTestDataFolder() + "GH178.dcm", 0},
		{util.GetTestDataFolder() + "xr_chest.dcm", 10219520},
		{util.GetTestDataFolder() + "xr_chicken2.dcm", 659012},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", 1048576},
		{util.GetTestDataFolder() + "IM-0001-0010.dcm", 103518},
		{util.GetTestDataFolder() + "IM0.dcm", 524288},
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
		{util.GetTestDataFolder() + "GH178.dcm", "1"},
		{util.GetTestDataFolder() + "xr_chest.dcm", "1"},
		{util.GetTestDataFolder() + "xr_chicken2.dcm", "1"},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", "16"},
		{util.GetTestDataFolder() + "IM-0001-0010.dcm", "1"},
		{util.GetTestDataFolder() + "T23/IM-0001-0001.dcm", "1"},
		{util.GetTestDataFolder() + "T14/IM-0001-0001.dcm", "1"},
		{util.GetTestDataFolder() + "IM0.dcm", "1"},
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
		{util.GetTestDataFolder() + "GH178.dcm", "16"},
		{util.GetTestDataFolder() + "xr_chest.dcm", "16"},
		{util.GetTestDataFolder() + "xr_chicken2.dcm", "16"},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", "8"},
		{util.GetTestDataFolder() + "IM-0001-0010.dcm", "16"},
		{util.GetTestDataFolder() + "T23/IM-0001-0001.dcm", "16"},
		{util.GetTestDataFolder() + "T14/IM-0001-0001.dcm", "16"},
		{util.GetTestDataFolder() + "IM0.dcm", "16"},
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
		{util.GetTestDataFolder() + "GH178.dcm", "16"},
		{util.GetTestDataFolder() + "xr_chest.dcm", "12"},
		{util.GetTestDataFolder() + "xr_chicken2.dcm", "12"},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", "8"},
		{util.GetTestDataFolder() + "IM-0001-0010.dcm", "12"},
		{util.GetTestDataFolder() + "T23/IM-0001-0001.dcm", "12"},
		{util.GetTestDataFolder() + "T14/IM-0001-0001.dcm", "12"},
		{util.GetTestDataFolder() + "IM0.dcm", "13"},
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
		{util.GetTestDataFolder() + "GH178.dcm", "15"},
		{util.GetTestDataFolder() + "xr_chest.dcm", "11"},
		{util.GetTestDataFolder() + "xr_chicken2.dcm", "11"},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", "7"},
		{util.GetTestDataFolder() + "IM-0001-0010.dcm", "11"},
		{util.GetTestDataFolder() + "T23/IM-0001-0001.dcm", "11"},
		{util.GetTestDataFolder() + "T14/IM-0001-0001.dcm", "11"},
		{util.GetTestDataFolder() + "IM0.dcm", "12"},
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

func TestPhotometricInterpretation(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{util.GetTestDataFolder() + "GH178.dcm", "MONOCHROME2"},
		{util.GetTestDataFolder() + "xr_chest.dcm", "MONOCHROME1"},
		{util.GetTestDataFolder() + "xr_chicken2.dcm", "MONOCHROME1"},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", "MONOCHROME2"},
		{util.GetTestDataFolder() + "IM-0001-0010.dcm", "MONOCHROME2"},
		{util.GetTestDataFolder() + "T23/IM-0001-0001.dcm", "MONOCHROME2"},
		{util.GetTestDataFolder() + "T14/IM-0001-0001.dcm", "MONOCHROME2"},
		{util.GetTestDataFolder() + "IM0.dcm", "MONOCHROME2"},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.IsReadPixel = true
		reader.ReadFile(c.in)
		got := reader.Dataset.PhotometricInterpretation()
		if got != c.want {
			t.Errorf("PhotometricInterpretation() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestSamplesPerPixel(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{util.GetTestDataFolder() + "GH178.dcm", "1"},
		{util.GetTestDataFolder() + "xr_chest.dcm", "1"},
		{util.GetTestDataFolder() + "xr_chicken2.dcm", "1"},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", "1"},
		{util.GetTestDataFolder() + "IM-0001-0010.dcm", "1"},
		{util.GetTestDataFolder() + "T23/IM-0001-0001.dcm", "1"},
		{util.GetTestDataFolder() + "T14/IM-0001-0001.dcm", "1"},
		{util.GetTestDataFolder() + "IM0.dcm", "1"},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.IsReadPixel = true
		reader.ReadFile(c.in)
		got := reader.Dataset.SamplesPerPixel()
		if got != c.want {
			t.Errorf("SamplesPerPixel() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestPixelRepresentation(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{util.GetTestDataFolder() + "GH178.dcm", "1"},
		{util.GetTestDataFolder() + "xr_chest.dcm", "0"},
		{util.GetTestDataFolder() + "xr_chicken2.dcm", "0"},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", "0"},
		{util.GetTestDataFolder() + "IM-0001-0010.dcm", "0"},
		{util.GetTestDataFolder() + "T23/IM-0001-0001.dcm", "0"},
		{util.GetTestDataFolder() + "T14/IM-0001-0001.dcm", "0"},
		{util.GetTestDataFolder() + "IM0.dcm", "1"},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.IsReadPixel = true
		reader.ReadFile(c.in)
		got := reader.Dataset.PixelRepresentation()
		if got != c.want {
			t.Errorf("PixelRepresentation() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestPlanarConfiguration(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{util.GetTestDataFolder() + "GH178.dcm", ""},
		{util.GetTestDataFolder() + "xr_chest.dcm", ""},
		{util.GetTestDataFolder() + "xr_chicken2.dcm", ""},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", ""},
		{util.GetTestDataFolder() + "IM-0001-0010.dcm", ""},
		{util.GetTestDataFolder() + "T23/IM-0001-0001.dcm", ""},
		{util.GetTestDataFolder() + "T14/IM-0001-0001.dcm", ""},
		{util.GetTestDataFolder() + "IM0.dcm", ""},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.IsReadPixel = true
		reader.ReadFile(c.in)
		got := reader.Dataset.PlanarConfiguration()
		if got != c.want {
			t.Errorf("PlanarConfiguration() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestRescaleIntercept(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{util.GetTestDataFolder() + "GH178.dcm", "-1024"},
		{util.GetTestDataFolder() + "xr_chest.dcm", "0.00000000E+00"},
		{util.GetTestDataFolder() + "xr_chicken2.dcm", "0.000000"},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", ""},
		{util.GetTestDataFolder() + "IM-0001-0010.dcm", ""},
		{util.GetTestDataFolder() + "T23/IM-0001-0001.dcm", ""},
		{util.GetTestDataFolder() + "T14/IM-0001-0001.dcm", ""},
		{util.GetTestDataFolder() + "IM0.dcm", "0"},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.IsReadPixel = true
		reader.ReadFile(c.in)
		got := reader.Dataset.RescaleIntercept()
		if got != c.want {
			t.Errorf("RescaleIntercept() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestRescaleSlope(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{util.GetTestDataFolder() + "GH178.dcm", "1"},
		{util.GetTestDataFolder() + "xr_chest.dcm", "1.00000000E+00"},
		{util.GetTestDataFolder() + "xr_chicken2.dcm", "1.000000"},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", ""},
		{util.GetTestDataFolder() + "IM-0001-0010.dcm", ""},
		{util.GetTestDataFolder() + "T23/IM-0001-0001.dcm", ""},
		{util.GetTestDataFolder() + "T14/IM-0001-0001.dcm", ""},
		{util.GetTestDataFolder() + "IM0.dcm", "1"},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.IsReadPixel = true
		reader.ReadFile(c.in)
		got := reader.Dataset.RescaleSlope()
		if got != c.want {
			t.Errorf("RescaleSlope() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestIsCompressed(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{util.GetTestDataFolder() + "GH178.dcm", false},
		{util.GetTestDataFolder() + "xr_chest.dcm", false},
		{util.GetTestDataFolder() + "xr_chicken2.dcm", true},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", false},
		{util.GetTestDataFolder() + "IM-0001-0010.dcm", true},
		{util.GetTestDataFolder() + "T23/IM-0001-0001.dcm", true},
		{util.GetTestDataFolder() + "T14/IM-0001-0001.dcm", true},
		{util.GetTestDataFolder() + "IM0.dcm", false},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.IsReadPixel = false
		reader.ReadFile(c.in)
		got, _ := reader.IsCompressed()
		if got != c.want {
			t.Errorf("IsCompressed() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}

func TestIsBigEndian(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{util.GetTestDataFolder() + "GH178.dcm", false},
		{util.GetTestDataFolder() + "xr_chest.dcm", false},
		{util.GetTestDataFolder() + "xr_chicken2.dcm", false},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", false},
		{util.GetTestDataFolder() + "IM-0001-0010.dcm", false},
		{util.GetTestDataFolder() + "T23/IM-0001-0001.dcm", false},
		{util.GetTestDataFolder() + "T14/IM-0001-0001.dcm", false},
		{util.GetTestDataFolder() + "IM0.dcm", false},
		{util.GetTestDataFolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", true},
	}
	for _, c := range cases {
		var reader DcmReader
		reader.IsReadValue = true
		reader.IsReadPixel = false
		reader.ReadFile(c.in)
		got, _ := reader.IsBigEndian()
		if got != c.want {
			t.Errorf("IsBigEndian() %s, want '%v' got '%v'", c.in, c.want, got)
		}
	}
}
