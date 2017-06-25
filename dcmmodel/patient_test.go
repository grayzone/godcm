package dcmmodel

import (
	"fmt"
	"testing"

	"github.com/grayzone/godcm/core"
	"github.com/grayzone/godcm/util"
)

func TestPatientParse(t *testing.T) {
	cases := []struct {
		in   string
		want Patient
	}{
		{util.GetTestDataFolder() + "GH220.dcm", Patient{}},
		{util.GetTestDataFolder() + "GH178.dcm", Patient{}},
		{util.GetTestDataFolder() + "xr_chest.dcm", Patient{}},
		{util.GetTestDataFolder() + "xr_chicken2.dcm", Patient{}},
		{util.GetTestDataFolder() + "CT-MONO2-16-ankle", Patient{}},
		{util.GetTestDataFolder() + "T14/IM-0001-0001.dcm", Patient{}},
		{util.GetTestDataFolder() + "GH184.dcm", Patient{}},
		{util.GetTestDataFolder() + "GH064.dcm", Patient{}},
		{util.GetTestDataFolder() + "GH133.dcm", Patient{}},
		{util.GetTestDataFolder() + "GH179A.dcm", Patient{}},
		{util.GetTestDataFolder() + "CT1_J2KI", Patient{}},
		{util.GetTestDataFolder() + "GH223.dcm", Patient{}},
		{util.GetTestDataFolder() + "GH195.dcm", Patient{}},
		{util.GetTestDataFolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_ELE_undefinded_length.dcm", Patient{}},
		{util.GetTestDataFolder() + "GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm", Patient{}},
		{util.GetTestDataFolder() + "MR-MONO2-8-16x-heart.dcm", Patient{}},

		/*
			{util.GetTestDataFolder() + "GH179B.dcm", DcmDataset{}}, // incomplete file
		*/
	}
	for _, c := range cases {
		var reader core.DcmReader
		reader.IsReadValue = true
		err := reader.ReadFile(c.in)
		if err != nil {
			t.Errorf("DcmReader.ReadFile(): %s", err.Error())
			return
		}
		var p Patient
		p.Parse(reader.Dataset)
		fmt.Println(p)
	}

}
