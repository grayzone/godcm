package dcmdata

import "testing"

func TestSetByName(t *testing.T) {
	cases := []struct {
		in   string
		want DcmEVR
	}{
		{"", EVR_UNKNOWN},
		{" ", EVR_UNKNOWN},
		{".", EVR_UNKNOWN},
		{"AA", EVR_UNKNOWN},
		{"ZZ", EVR_UNKNOWN},
		{"XX", EVR_UNKNOWN},
		{"YY", EVR_UNKNOWN},
		{"aa", EVR_UNKNOWN2B},
		{"bb", EVR_UNKNOWN2B},
		{"cc", EVR_UNKNOWN2B},
		{"??", EVR_UNKNOWN2B},
		{"?A", EVR_UNKNOWN2B},
		{"a?", EVR_UNKNOWN2B},
		{"AE", EVR_AE},
		{"AS", EVR_AS},
		{"AT", EVR_AT},
		{"CS", EVR_CS},
		{"DA", EVR_DA},
		{"DS", EVR_DS},
		{"DT", EVR_DT},
		{"FL", EVR_FL},
		{"FD", EVR_FD},
		{"IS", EVR_IS},
		{"LO", EVR_LO},
		{"LT", EVR_LT},
		{"OB", EVR_OB},
		{"OF", EVR_OF},
		{"OW", EVR_OW},
		{"PN", EVR_PN},
		{"SH", EVR_SH},
		{"SL", EVR_SL},
		{"SQ", EVR_SQ},
		{"SS", EVR_SS},
		{"ST", EVR_ST},
		{"TM", EVR_TM},
		{"UI", EVR_UI},
		{"UL", EVR_UL},
		{"US", EVR_US},
		{"UT", EVR_UT},
		{"ox", EVR_ox},
		{"xs", EVR_xs},
		{"lt", EVR_lt},
		{"na", EVR_na},
		{"up", EVR_up},
		{"it_EVR_item", EVR_item},
		{"mi_EVR_metainfo", EVR_metainfo},
		{"ds_EVR_dataset", EVR_dataset},
		{"ff_EVR_fileFormat", EVR_fileFormat},
		{"dd_EVR_dicomDir", EVR_dicomDir},
		{"dr_EVR_dirRecord", EVR_dirRecord},
		{"ps_EVR_pixelSQ", EVR_pixelSQ},
		{"pi", EVR_pixelItem},
		{"PixelData", EVR_PixelData},
		{"OverlayData", EVR_OverlayData},
	}
	for _, c := range cases {
		var v DcmVR
		v.SetByName(c.in)

		got := v.GetEVR()
		if got != c.want {
			t.Errorf("SetByName(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestSetVR(t *testing.T) {
	for i := 0; i < DcmVRDict_DIM+100; i++ {
		var v DcmVR
		v.SetVR(DcmEVR(i))
		if i >= 0 && i < DcmVRDict_DIM {
			if v.GetEVR() != DcmEVR(i) {
				t.Errorf("excepted %q, got %q", DcmEVR(i), v.GetEVR())
			}
		} else {
			if v.GetEVR() != EVR_UNKNOWN {
				t.Errorf("excepted %q, got %q", DcmEVR(EVR_UNKNOWN), v.GetEVR())
			}
		}

	}
}

func TestIsStandard(t *testing.T) {
	for i := 0; i < DcmVRDict_DIM+100; i++ {
		var v DcmVR
		v.SetVR(DcmEVR(i))
		if i >= 0 && i < 26 {
			if v.IsStandard() != true {
				t.Errorf("%d, %q, excepted true, got %v", i, v.GetEVR(), v.IsStandard())
			}
		} else {
			if v.GetEVR() == DcmEVR(EVR_UN) || v.GetEVR() == DcmEVR(EVR_PixelData) || v.GetEVR() == DcmEVR(EVR_OverlayData) {
				if v.IsStandard() != true {
					t.Errorf("%d, %q, excepted true, got %v", i, v.GetEVR(), v.IsStandard())
				}
			} else if v.IsStandard() != false {
				t.Errorf("%d, %q, excepted false, got %v", i, v.GetEVR(), v.IsStandard())
			}
		}
	}
}

func TestGetValidEVR(t *testing.T) {
	cases1 := []struct {
		in   DcmEVR
		want DcmEVR
	}{
		{EVR_AE, EVR_AE},
		{EVR_AS, EVR_AS},
		{EVR_AT, EVR_AT},
		{EVR_CS, EVR_CS},
		{EVR_DA, EVR_DA},
		{EVR_DS, EVR_DS},
		{EVR_DT, EVR_DT},
		{EVR_FL, EVR_FL},
		{EVR_FD, EVR_FD},
		{EVR_IS, EVR_IS},
		{EVR_LO, EVR_LO},
		{EVR_LT, EVR_LT},
		{EVR_OB, EVR_OB},
		{EVR_OF, EVR_OF},
		{EVR_OW, EVR_OW},
		{EVR_PN, EVR_PN},
		{EVR_SH, EVR_SH},
		{EVR_SL, EVR_SL},
		{EVR_SQ, EVR_SQ},
		{EVR_SS, EVR_SS},
		{EVR_ST, EVR_ST},
		{EVR_TM, EVR_TM},
		{EVR_UI, EVR_UI},
		{EVR_UL, EVR_UL},
		{EVR_US, EVR_US},
		{EVR_UT, EVR_UT},
		{EVR_UN, EVR_UN},
		{EVR_PixelData, EVR_PixelData},
		{EVR_OverlayData, EVR_OverlayData},

		{EVR_ox, EVR_OB},
		{EVR_xs, EVR_US},
		{EVR_lt, EVR_OW},
		{EVR_na, EVR_UN},
		{EVR_up, EVR_UL},
		{EVR_item, EVR_UN},
		{EVR_metainfo, EVR_UN},
		{EVR_dataset, EVR_UN},
		{EVR_fileFormat, EVR_UN},
		{EVR_dicomDir, EVR_UN},
		{EVR_dirRecord, EVR_UN},
		{EVR_pixelSQ, EVR_OB},
		{EVR_pixelItem, EVR_UN},
		{EVR_UNKNOWN, EVR_UN},
		{EVR_UNKNOWN2B, EVR_UN},
	}
	cases2 := []struct {
		in   DcmEVR
		want DcmEVR
	}{
		{EVR_AE, EVR_AE},
		{EVR_AS, EVR_AS},
		{EVR_AT, EVR_AT},
		{EVR_CS, EVR_CS},
		{EVR_DA, EVR_DA},
		{EVR_DS, EVR_DS},
		{EVR_DT, EVR_DT},
		{EVR_FL, EVR_FL},
		{EVR_FD, EVR_FD},
		{EVR_IS, EVR_IS},
		{EVR_LO, EVR_LO},
		{EVR_LT, EVR_LT},
		{EVR_OB, EVR_OB},
		{EVR_OF, EVR_OF},
		{EVR_OW, EVR_OW},
		{EVR_PN, EVR_PN},
		{EVR_SH, EVR_SH},
		{EVR_SL, EVR_SL},
		{EVR_SQ, EVR_SQ},
		{EVR_SS, EVR_SS},
		{EVR_ST, EVR_ST},
		{EVR_TM, EVR_TM},
		{EVR_UI, EVR_UI},
		{EVR_UL, EVR_UL},
		{EVR_US, EVR_US},
		{EVR_UT, EVR_UT},
		{EVR_UN, EVR_OB},
		{EVR_PixelData, EVR_PixelData},
		{EVR_OverlayData, EVR_OverlayData},

		{EVR_ox, EVR_OB},
		{EVR_xs, EVR_US},
		{EVR_lt, EVR_OW},
		{EVR_na, EVR_OB},
		{EVR_up, EVR_UL},
		{EVR_item, EVR_OB},
		{EVR_metainfo, EVR_OB},
		{EVR_dataset, EVR_OB},
		{EVR_fileFormat, EVR_OB},
		{EVR_dicomDir, EVR_OB},
		{EVR_dirRecord, EVR_OB},
		{EVR_pixelSQ, EVR_OB},
		{EVR_pixelItem, EVR_OB},
		{EVR_UNKNOWN, EVR_OB},
		{EVR_UNKNOWN2B, EVR_OB},
	}
	cases3 := []struct {
		in   DcmEVR
		want DcmEVR
	}{
		{EVR_AE, EVR_AE},
		{EVR_AS, EVR_AS},
		{EVR_AT, EVR_AT},
		{EVR_CS, EVR_CS},
		{EVR_DA, EVR_DA},
		{EVR_DS, EVR_DS},
		{EVR_DT, EVR_DT},
		{EVR_FL, EVR_FL},
		{EVR_FD, EVR_FD},
		{EVR_IS, EVR_IS},
		{EVR_LO, EVR_LO},
		{EVR_LT, EVR_LT},
		{EVR_OB, EVR_OB},
		{EVR_OF, EVR_OF},
		{EVR_OW, EVR_OW},
		{EVR_PN, EVR_PN},
		{EVR_SH, EVR_SH},
		{EVR_SL, EVR_SL},
		{EVR_SQ, EVR_SQ},
		{EVR_SS, EVR_SS},
		{EVR_ST, EVR_ST},
		{EVR_TM, EVR_TM},
		{EVR_UI, EVR_UI},
		{EVR_UL, EVR_UL},
		{EVR_US, EVR_US},
		{EVR_UT, EVR_OB},
		{EVR_UN, EVR_OB},
		{EVR_PixelData, EVR_PixelData},
		{EVR_OverlayData, EVR_OverlayData},

		{EVR_ox, EVR_OB},
		{EVR_xs, EVR_US},
		{EVR_lt, EVR_OW},
		{EVR_na, EVR_OB},
		{EVR_up, EVR_UL},
		{EVR_item, EVR_OB},
		{EVR_metainfo, EVR_OB},
		{EVR_dataset, EVR_OB},
		{EVR_fileFormat, EVR_OB},
		{EVR_dicomDir, EVR_OB},
		{EVR_dirRecord, EVR_OB},
		{EVR_pixelSQ, EVR_OB},
		{EVR_pixelItem, EVR_OB},
		{EVR_UNKNOWN, EVR_OB},
		{EVR_UNKNOWN2B, EVR_OB},
	}
	DcmEnableUnknownVRGeneration = true
	for _, c := range cases1 {
		var v DcmVR
		v.SetVR(c.in)

		got := v.GetValidEVR()
		if got != c.want {
			t.Errorf("GetValidEVR(%q) == %q, want %q", c.in, got, c.want)
		}
	}
	DcmEnableUnknownVRGeneration = false
	for _, c := range cases2 {
		var v DcmVR
		v.SetVR(c.in)

		got := v.GetValidEVR()
		if got != c.want {
			t.Errorf("GetValidEVR(%q) == %q, want %q", c.in, got, c.want)
		}
	}

	DcmEnableUnlimitedTextVRGeneration = false
	for _, c := range cases3 {
		var v DcmVR
		v.SetVR(c.in)

		got := v.GetValidEVR()
		if got != c.want {
			t.Errorf("GetValidEVR(%q) == %q, want %q", c.in, got, c.want)
		}
	}
	DcmEnableUnknownVRGeneration = true
	DcmEnableUnlimitedTextVRGeneration = true

}

func TestGetVRName(t *testing.T) {
	cases := []struct {
		in   DcmEVR
		want string
	}{
		{EVR_AE, "AE"},
		{EVR_AS, "AS"},
		{EVR_AT, "AT"},
		{EVR_CS, "CS"},
		{EVR_DA, "DA"},
		{EVR_DS, "DS"},
		{EVR_DT, "DT"},
		{EVR_FL, "FL"},
		{EVR_FD, "FD"},
		{EVR_IS, "IS"},
		{EVR_LO, "LO"},
		{EVR_LT, "LT"},
		{EVR_OB, "OB"},
		{EVR_OF, "OF"},
		{EVR_OW, "OW"},
		{EVR_PN, "PN"},
		{EVR_SH, "SH"},
		{EVR_SL, "SL"},
		{EVR_SQ, "SQ"},
		{EVR_SS, "SS"},
		{EVR_ST, "ST"},
		{EVR_TM, "TM"},
		{EVR_UI, "UI"},
		{EVR_UL, "UL"},
		{EVR_US, "US"},
		{EVR_UT, "UT"},
		{EVR_ox, "ox"},
		{EVR_xs, "xs"},
		{EVR_lt, "lt"},
		{EVR_na, "na"},
		{EVR_up, "up"},
		{EVR_item, "it_EVR_item"},
		{EVR_metainfo, "mi_EVR_metainfo"},
		{EVR_dataset, "ds_EVR_dataset"},
		{EVR_fileFormat, "ff_EVR_fileFormat"},
		{EVR_dicomDir, "dd_EVR_dicomDir"},
		{EVR_dirRecord, "dr_EVR_dirRecord"},
		{EVR_pixelSQ, "ps_EVR_pixelSQ"},
		{EVR_pixelItem, "pi"},
		{EVR_UNKNOWN, "??"},
		{EVR_UN, "UN"},
		{EVR_PixelData, "PixelData"},
		{EVR_OverlayData, "OverlayData"},
		{EVR_UNKNOWN2B, "??"},
	}
	for _, c := range cases {
		var v DcmVR
		v.SetVR(c.in)
		got := v.GetVRName()
		if got != c.want {
			t.Errorf("GetVRName(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestGetValidVRName(t *testing.T) {
	cases1 := []struct {
		in   DcmEVR
		want string
	}{
		{EVR_AE, "AE"},
		{EVR_AS, "AS"},
		{EVR_AT, "AT"},
		{EVR_CS, "CS"},
		{EVR_DA, "DA"},
		{EVR_DS, "DS"},
		{EVR_DT, "DT"},
		{EVR_FL, "FL"},
		{EVR_FD, "FD"},
		{EVR_IS, "IS"},
		{EVR_LO, "LO"},
		{EVR_LT, "LT"},
		{EVR_OB, "OB"},
		{EVR_OF, "OF"},
		{EVR_OW, "OW"},
		{EVR_PN, "PN"},
		{EVR_SH, "SH"},
		{EVR_SL, "SL"},
		{EVR_SQ, "SQ"},
		{EVR_SS, "SS"},
		{EVR_ST, "ST"},
		{EVR_TM, "TM"},
		{EVR_UI, "UI"},
		{EVR_UL, "UL"},
		{EVR_US, "US"},
		{EVR_UT, "UT"},
		{EVR_UN, "UN"},
		{EVR_PixelData, "PixelData"},
		{EVR_OverlayData, "OverlayData"},

		{EVR_ox, "OB"},
		{EVR_xs, "US"},
		{EVR_lt, "OW"},
		{EVR_na, "UN"},
		{EVR_up, "UL"},
		{EVR_item, "UN"},
		{EVR_metainfo, "UN"},
		{EVR_dataset, "UN"},
		{EVR_fileFormat, "UN"},
		{EVR_dicomDir, "UN"},
		{EVR_dirRecord, "UN"},
		{EVR_pixelSQ, "OB"},
		{EVR_pixelItem, "UN"},
		{EVR_UNKNOWN, "UN"},
		{EVR_UNKNOWN2B, "UN"},
	}
	cases2 := []struct {
		in   DcmEVR
		want string
	}{
		{EVR_AE, "AE"},
		{EVR_AS, "AS"},
		{EVR_AT, "AT"},
		{EVR_CS, "CS"},
		{EVR_DA, "DA"},
		{EVR_DS, "DS"},
		{EVR_DT, "DT"},
		{EVR_FL, "FL"},
		{EVR_FD, "FD"},
		{EVR_IS, "IS"},
		{EVR_LO, "LO"},
		{EVR_LT, "LT"},
		{EVR_OB, "OB"},
		{EVR_OF, "OF"},
		{EVR_OW, "OW"},
		{EVR_PN, "PN"},
		{EVR_SH, "SH"},
		{EVR_SL, "SL"},
		{EVR_SQ, "SQ"},
		{EVR_SS, "SS"},
		{EVR_ST, "ST"},
		{EVR_TM, "TM"},
		{EVR_UI, "UI"},
		{EVR_UL, "UL"},
		{EVR_US, "US"},
		{EVR_UT, "UT"},
		{EVR_UN, "OB"},
		{EVR_PixelData, "PixelData"},
		{EVR_OverlayData, "OverlayData"},

		{EVR_ox, "OB"},
		{EVR_xs, "US"},
		{EVR_lt, "OW"},
		{EVR_na, "OB"},
		{EVR_up, "UL"},
		{EVR_item, "OB"},
		{EVR_metainfo, "OB"},
		{EVR_dataset, "OB"},
		{EVR_fileFormat, "OB"},
		{EVR_dicomDir, "OB"},
		{EVR_dirRecord, "OB"},
		{EVR_pixelSQ, "OB"},
		{EVR_pixelItem, "OB"},
		{EVR_UNKNOWN, "OB"},
		{EVR_UNKNOWN2B, "OB"},
	}
	cases3 := []struct {
		in   DcmEVR
		want string
	}{
		{EVR_AE, "AE"},
		{EVR_AS, "AS"},
		{EVR_AT, "AT"},
		{EVR_CS, "CS"},
		{EVR_DA, "DA"},
		{EVR_DS, "DS"},
		{EVR_DT, "DT"},
		{EVR_FL, "FL"},
		{EVR_FD, "FD"},
		{EVR_IS, "IS"},
		{EVR_LO, "LO"},
		{EVR_LT, "LT"},
		{EVR_OB, "OB"},
		{EVR_OF, "OF"},
		{EVR_OW, "OW"},
		{EVR_PN, "PN"},
		{EVR_SH, "SH"},
		{EVR_SL, "SL"},
		{EVR_SQ, "SQ"},
		{EVR_SS, "SS"},
		{EVR_ST, "ST"},
		{EVR_TM, "TM"},
		{EVR_UI, "UI"},
		{EVR_UL, "UL"},
		{EVR_US, "US"},
		{EVR_UT, "OB"},
		{EVR_UN, "OB"},
		{EVR_PixelData, "PixelData"},
		{EVR_OverlayData, "OverlayData"},

		{EVR_ox, "OB"},
		{EVR_xs, "US"},
		{EVR_lt, "OW"},
		{EVR_na, "OB"},
		{EVR_up, "UL"},
		{EVR_item, "OB"},
		{EVR_metainfo, "OB"},
		{EVR_dataset, "OB"},
		{EVR_fileFormat, "OB"},
		{EVR_dicomDir, "OB"},
		{EVR_dirRecord, "OB"},
		{EVR_pixelSQ, "OB"},
		{EVR_pixelItem, "OB"},
		{EVR_UNKNOWN, "OB"},
		{EVR_UNKNOWN2B, "OB"},
	}
	DcmEnableUnknownVRGeneration = true
	for _, c := range cases1 {
		var v DcmVR
		v.SetVR(c.in)
		got := v.GetValidVRName()
		if got != c.want {
			t.Errorf("GetValidVRName(%q) == %q, want %q", c.in, got, c.want)
		}
	}
	DcmEnableUnknownVRGeneration = false
	for _, c := range cases2 {
		var v DcmVR
		v.SetVR(c.in)
		got := v.GetValidVRName()
		if got != c.want {
			t.Errorf("GetValidVRName(%q) == %q, want %q", c.in, got, c.want)
		}
	}
	DcmEnableUnlimitedTextVRGeneration = false
	for _, c := range cases3 {
		var v DcmVR
		v.SetVR(c.in)
		got := v.GetValidVRName()
		if got != c.want {
			t.Errorf("GetValidVRName(%q) == %q, want %q", c.in, got, c.want)
		}
	}

	DcmEnableUnknownVRGeneration = true
	DcmEnableUnlimitedTextVRGeneration = true
}

func TestGetValueWidth(t *testing.T) {
	cases := []struct {
		in   DcmEVR
		want uint
	}{
		{EVR_AE, 1},
		{EVR_AS, 1},
		{EVR_AT, 2},
		{EVR_CS, 1},
		{EVR_DA, 1},
		{EVR_DS, 1},
		{EVR_DT, 1},
		{EVR_FL, 4},
		{EVR_FD, 8},
		{EVR_IS, 1},
		{EVR_LO, 1},
		{EVR_LT, 1},
		{EVR_OB, 4},
		{EVR_OF, 4},
		{EVR_OW, 4},
		{EVR_PN, 1},
		{EVR_SH, 1},
		{EVR_SL, 4},
		{EVR_SQ, 0},
		{EVR_SS, 4},
		{EVR_ST, 1},
		{EVR_TM, 1},
		{EVR_UI, 1},
		{EVR_UL, 4},
		{EVR_US, 4},
		{EVR_UT, 1},
		{EVR_ox, 4},
		{EVR_xs, 4},
		{EVR_lt, 4},
		{EVR_na, 0},
		{EVR_up, 4},
		{EVR_item, 0},
		{EVR_metainfo, 0},
		{EVR_dataset, 0},
		{EVR_fileFormat, 0},
		{EVR_dicomDir, 0},
		{EVR_dirRecord, 0},
		{EVR_pixelSQ, 4},
		{EVR_pixelItem, 4},
		{EVR_UNKNOWN, 4},
		{EVR_UN, 4},
		{EVR_PixelData, 0},
		{EVR_OverlayData, 0},
		{EVR_UNKNOWN2B, 4},
	}
	for _, c := range cases {
		var v DcmVR
		v.SetVR(c.in)
		got := v.GetValueWidth()
		if got != c.want {
			t.Errorf("GetValueWidth(%q) == %d, want %d", c.in, got, c.want)
		}
	}

}

func TestIsForInternalUseOnly(t *testing.T) {
	cases := []struct {
		in   DcmEVR
		want bool
	}{
		{EVR_AE, false},
		{EVR_AS, false},
		{EVR_AT, false},
		{EVR_CS, false},
		{EVR_DA, false},
		{EVR_DS, false},
		{EVR_DT, false},
		{EVR_FL, false},
		{EVR_FD, false},
		{EVR_IS, false},
		{EVR_LO, false},
		{EVR_LT, false},
		{EVR_OB, false},
		{EVR_OF, false},
		{EVR_OW, false},
		{EVR_PN, false},
		{EVR_SH, false},
		{EVR_SL, false},
		{EVR_SQ, false},
		{EVR_SS, false},
		{EVR_ST, false},
		{EVR_TM, false},
		{EVR_UI, false},
		{EVR_UL, false},
		{EVR_US, false},
		{EVR_UT, false},
		{EVR_ox, false},
		{EVR_xs, false},
		{EVR_lt, false},
		{EVR_na, false},
		{EVR_up, false},
		{EVR_item, true},
		{EVR_metainfo, true},
		{EVR_dataset, true},
		{EVR_fileFormat, true},
		{EVR_dicomDir, true},
		{EVR_dirRecord, true},
		{EVR_pixelSQ, true},
		{EVR_pixelItem, false},
		{EVR_UNKNOWN, true},
		{EVR_UN, false},
		{EVR_PixelData, true},
		{EVR_OverlayData, true},
		{EVR_UNKNOWN2B, true},
	}
	for _, c := range cases {
		var v DcmVR
		v.SetVR(c.in)
		got := v.IsForInternalUseOnly()
		if got != c.want {
			t.Errorf("IsForInternalUseOnly(%q) == %v, want %v", c.in, got, c.want)
		}
	}

}

func TestIsaString(t *testing.T) {
	cases := []struct {
		in   DcmEVR
		want bool
	}{
		{EVR_AE, true},
		{EVR_AS, true},
		{EVR_AT, false},
		{EVR_CS, true},
		{EVR_DA, true},
		{EVR_DS, true},
		{EVR_DT, true},
		{EVR_FL, false},
		{EVR_FD, false},
		{EVR_IS, true},
		{EVR_LO, true},
		{EVR_LT, true},
		{EVR_OB, false},
		{EVR_OF, false},
		{EVR_OW, false},
		{EVR_PN, true},
		{EVR_SH, true},
		{EVR_SL, false},
		{EVR_SQ, false},
		{EVR_SS, false},
		{EVR_ST, true},
		{EVR_TM, true},
		{EVR_UI, true},
		{EVR_UL, false},
		{EVR_US, false},
		{EVR_UT, true},
		{EVR_ox, false},
		{EVR_xs, false},
		{EVR_lt, false},
		{EVR_na, false},
		{EVR_up, false},
		{EVR_item, false},
		{EVR_metainfo, false},
		{EVR_dataset, false},
		{EVR_fileFormat, false},
		{EVR_dicomDir, false},
		{EVR_dirRecord, false},
		{EVR_pixelSQ, false},
		{EVR_pixelItem, false},
		{EVR_UNKNOWN, false},
		{EVR_UN, false},
		{EVR_PixelData, false},
		{EVR_OverlayData, false},
		{EVR_UNKNOWN2B, false},
	}
	for _, c := range cases {
		var v DcmVR
		v.SetVR(c.in)
		got := v.IsaString()
		if got != c.want {
			t.Errorf("IsaString(%q) == %v, want %v", c.in, got, c.want)
		}
	}

}

func TestUsesExtendedLengthEncoding(t *testing.T) {
	cases := []struct {
		in   DcmEVR
		want bool
	}{
		{EVR_AE, false},
		{EVR_AS, false},
		{EVR_AT, false},
		{EVR_CS, false},
		{EVR_DA, false},
		{EVR_DS, false},
		{EVR_DT, false},
		{EVR_FL, false},
		{EVR_FD, false},
		{EVR_IS, false},
		{EVR_LO, false},
		{EVR_LT, false},
		{EVR_OB, true},
		{EVR_OF, true},
		{EVR_OW, true},
		{EVR_PN, false},
		{EVR_SH, false},
		{EVR_SL, false},
		{EVR_SQ, true},
		{EVR_SS, false},
		{EVR_ST, false},
		{EVR_TM, false},
		{EVR_UI, false},
		{EVR_UL, false},
		{EVR_US, false},
		{EVR_UT, true},
		{EVR_ox, true},
		{EVR_xs, false},
		{EVR_lt, true},
		{EVR_na, false},
		{EVR_up, false},
		{EVR_item, false},
		{EVR_metainfo, false},
		{EVR_dataset, false},
		{EVR_fileFormat, false},
		{EVR_dicomDir, false},
		{EVR_dirRecord, false},
		{EVR_pixelSQ, false},
		{EVR_pixelItem, false},
		{EVR_UNKNOWN, true},
		{EVR_UN, true},
		{EVR_PixelData, false},
		{EVR_OverlayData, false},
		{EVR_UNKNOWN2B, false},
	}
	for _, c := range cases {
		var v DcmVR
		v.SetVR(c.in)
		got := v.UsesExtendedLengthEncoding()
		if got != c.want {
			t.Errorf("UsesExtendedLengthEncoding(%q) == %v, want %v", c.in, got, c.want)
		}
	}

}

func TestGetMinValueLength(t *testing.T) {
	cases := []struct {
		in   DcmEVR
		want uint32
	}{
		{EVR_AE, 0},
		{EVR_AS, 4},
		{EVR_AT, 4},
		{EVR_CS, 0},
		{EVR_DA, 8},
		{EVR_DS, 0},
		{EVR_DT, 0},
		{EVR_FL, 4},
		{EVR_FD, 8},
		{EVR_IS, 0},
		{EVR_LO, 0},
		{EVR_LT, 0},
		{EVR_OB, 0},
		{EVR_OF, 0},
		{EVR_OW, 0},
		{EVR_PN, 0},
		{EVR_SH, 0},
		{EVR_SL, 4},
		{EVR_SQ, 0},
		{EVR_SS, 2},
		{EVR_ST, 0},
		{EVR_TM, 0},
		{EVR_UI, 0},
		{EVR_UL, 4},
		{EVR_US, 2},
		{EVR_UT, 0},
		{EVR_ox, 0},
		{EVR_xs, 2},
		{EVR_lt, 0},
		{EVR_na, 0},
		{EVR_up, 4},
		{EVR_item, 0},
		{EVR_metainfo, 0},
		{EVR_dataset, 0},
		{EVR_fileFormat, 0},
		{EVR_dicomDir, 0},
		{EVR_dirRecord, 0},
		{EVR_pixelSQ, 0},
		{EVR_pixelItem, 0},
		{EVR_UNKNOWN, 0},
		{EVR_UN, 0},
		{EVR_PixelData, 0},
		{EVR_OverlayData, 0},
		{EVR_UNKNOWN2B, 0},
	}

	for _, c := range cases {
		var v DcmVR
		v.SetVR(c.in)
		got := v.GetMinValueLength()
		if got != c.want {
			t.Errorf("GetMinValueLength(%q) == %v, want %v", c.in, got, c.want)
		}
	}

}

func TestGetMaxValueLength(t *testing.T) {
	cases := []struct {
		in   DcmEVR
		want uint32
	}{
		{EVR_AE, 16},
		{EVR_AS, 4},
		{EVR_AT, 4},
		{EVR_CS, 16},
		{EVR_DA, 10},
		{EVR_DS, 16},
		{EVR_DT, 26},
		{EVR_FL, 4},
		{EVR_FD, 8},
		{EVR_IS, 12},
		{EVR_LO, 64},
		{EVR_LT, 10240},
		{EVR_OB, DCM_UndefinedLength},
		{EVR_OF, DCM_UndefinedLength},
		{EVR_OW, DCM_UndefinedLength},
		{EVR_PN, 64},
		{EVR_SH, 16},
		{EVR_SL, 4},
		{EVR_SQ, DCM_UndefinedLength},
		{EVR_SS, 2},
		{EVR_ST, 1024},
		{EVR_TM, 16},
		{EVR_UI, 64},
		{EVR_UL, 4},
		{EVR_US, 2},
		{EVR_UT, DCM_UndefinedLength},
		{EVR_ox, DCM_UndefinedLength},
		{EVR_xs, 2},
		{EVR_lt, DCM_UndefinedLength},
		{EVR_na, 0},
		{EVR_up, 4},
		{EVR_item, 0},
		{EVR_metainfo, 0},
		{EVR_dataset, 0},
		{EVR_fileFormat, 0},
		{EVR_dicomDir, 0},
		{EVR_dirRecord, 0},
		{EVR_pixelSQ, DCM_UndefinedLength},
		{EVR_pixelItem, DCM_UndefinedLength},
		{EVR_UNKNOWN, DCM_UndefinedLength},
		{EVR_UN, DCM_UndefinedLength},
		{EVR_PixelData, DCM_UndefinedLength},
		{EVR_OverlayData, DCM_UndefinedLength},
		{EVR_UNKNOWN2B, DCM_UndefinedLength},
	}

	for _, c := range cases {
		var v DcmVR
		v.SetVR(c.in)
		got := v.GetMaxValueLength()
		if got != c.want {
			t.Errorf("GetMaxValueLength(%q) == %v, want %v", c.in, got, c.want)
		}
	}

}

func TestIsEquivalent(t *testing.T) {
	cases := []struct {
		base DcmEVR
		in   DcmEVR
		want bool
	}{
		{EVR_AE, EVR_AE, true},
		{EVR_AE, EVR_AS, false},
		{EVR_AS, EVR_AS, true},
		{EVR_AT, EVR_AT, true},
		{EVR_CS, EVR_CS, true},

		{EVR_ox, EVR_OB, true},
		{EVR_ox, EVR_OW, true},
		{EVR_ox, EVR_UN, false},

		{EVR_lt, EVR_OW, true},
		{EVR_lt, EVR_US, true},
		{EVR_lt, EVR_SS, true},

		{EVR_OB, EVR_ox, true},

		{EVR_OW, EVR_ox, true},
		{EVR_OW, EVR_lt, true},

		{EVR_up, EVR_UL, true},

		{EVR_UL, EVR_up, true},

		{EVR_xs, EVR_SS, true},
		{EVR_xs, EVR_US, true},

		{EVR_SS, EVR_xs, true},
		{EVR_SS, EVR_lt, true},
		{EVR_US, EVR_xs, true},
		{EVR_US, EVR_lt, true},
		{EVR_US, EVR_OW, false},
	}

	for _, c := range cases {
		var v1 DcmVR
		v1.SetVR(c.base)
		var v2 DcmVR
		v2.SetVR(c.in)

		got := v1.IsEquivalent(v2)
		if got != c.want {
			t.Errorf("%q IsEquivalent(%q) == %v, want %v", c.base, c.in, got, c.want)
		}
	}
}

func TestVRToString(t *testing.T) {
	cases := []struct {
		in   DcmEVR
		want string
	}{
		{EVR_AE, "EVR_AE"},
		{EVR_AS, "EVR_AS"},
		{EVR_AT, "EVR_AT"},
		{EVR_CS, "EVR_CS"},
		/*
			{EVR_DA, "EVR_DA", "DA", 1, DCMVR_PROP_ISASTRING, 8, 10},
			{EVR_DS, "EVR_DS", "DS", 1, DCMVR_PROP_ISASTRING, 0, 16},
			{EVR_DT, "EVR_DT", "DT", 1, DCMVR_PROP_ISASTRING, 0, 26},
			{EVR_FL, "EVR_FL", "FL", 4, DCMVR_PROP_NONE, 4, 4},
			{EVR_FD, "EVR_FD", "FD", 8, DCMVR_PROP_NONE, 8, 8},
			{EVR_IS, "EVR_IS", "IS", 1, DCMVR_PROP_ISASTRING, 0, 12},
			{EVR_LO, "EVR_LO", "LO", 1, DCMVR_PROP_ISASTRING, 0, 64},
			{EVR_LT, "EVR_LT", "LT", 1, DCMVR_PROP_ISASTRING, 0, 10240},
			{EVR_OB, "EVR_OB", "OB", 4, DCMVR_PROP_EXTENDEDLENGTHENCODING, 0, DCM_UndefinedLength},
			{EVR_OF, "EVR_OF", "OF", 4, DCMVR_PROP_EXTENDEDLENGTHENCODING, 0, DCM_UndefinedLength},
			{EVR_OW, "EVR_OW", "OW", 4, DCMVR_PROP_EXTENDEDLENGTHENCODING, 0, DCM_UndefinedLength},
			{EVR_PN, "EVR_PN", "PN", 1, DCMVR_PROP_ISASTRING, 0, 64},
			{EVR_SH, "EVR_SH", "SH", 1, DCMVR_PROP_ISASTRING, 0, 16},
			{EVR_SL, "EVR_SL", "SL", 4, DCMVR_PROP_NONE, 4, 4},
			{EVR_SQ, "EVR_SQ", "SQ", 0, DCMVR_PROP_EXTENDEDLENGTHENCODING, 0, DCM_UndefinedLength},
			{EVR_SS, "EVR_SS", "SS", 4, DCMVR_PROP_NONE, 2, 2},
			{EVR_ST, "EVR_ST", "ST", 1, DCMVR_PROP_ISASTRING, 0, 1024},
			{EVR_TM, "EVR_TM", "TM", 1, DCMVR_PROP_ISASTRING, 0, 16},
			{EVR_UI, "EVR_UI", "UI", 1, DCMVR_PROP_ISASTRING, 0, 64},
			{EVR_UL, "EVR_UL", "UL", 4, DCMVR_PROP_NONE, 4, 4},
			{EVR_US, "EVR_US", "US", 4, DCMVR_PROP_NONE, 2, 2},
			{EVR_UT, "EVR_UT", "UT", 1, DCMVR_PROP_ISASTRING | DCMVR_PROP_EXTENDEDLENGTHENCODING, 0, DCM_UndefinedLength},
			{EVR_ox, "EVR_ox", "ox", 4, DCMVR_PROP_NONSTANDARD | DCMVR_PROP_EXTENDEDLENGTHENCODING, 0, DCM_UndefinedLength},
			{EVR_xs, "EVR_xs", "xs", 4, DCMVR_PROP_NONSTANDARD, 2, 2},
			{EVR_lt, "EVR_lt", "lt", 4, DCMVR_PROP_NONSTANDARD | DCMVR_PROP_EXTENDEDLENGTHENCODING, 0, DCM_UndefinedLength},
			{EVR_na, "EVR_na", "na", 0, DCMVR_PROP_NONSTANDARD, 0, 0},
			{EVR_up, "EVR_up", "up", 4, DCMVR_PROP_NONSTANDARD, 4, 4},

			{EVR_item, "EVR_item", "it_EVR_item", 0, DCMVR_PROP_NONSTANDARD | DCMVR_PROP_INTERNAL, 0, 0},
			{EVR_metainfo, "EVR_metainfo", "mi_EVR_metainfo", 0, DCMVR_PROP_NONSTANDARD | DCMVR_PROP_INTERNAL, 0, 0},
			{EVR_dataset, "EVR_dataset", "ds_EVR_dataset", 0, DCMVR_PROP_NONSTANDARD | DCMVR_PROP_INTERNAL, 0, 0},
			{EVR_fileFormat, "EVR_fileFormat", "ff_EVR_fileFormat", 0, DCMVR_PROP_NONSTANDARD | DCMVR_PROP_INTERNAL, 0, 0},
			{EVR_dicomDir, "EVR_dicomDir", "dd_EVR_dicomDir", 0, DCMVR_PROP_NONSTANDARD | DCMVR_PROP_INTERNAL, 0, 0},
			{EVR_dirRecord, "EVR_dirRecord", "dr_EVR_dirRecord", 0, DCMVR_PROP_NONSTANDARD | DCMVR_PROP_INTERNAL, 0, 0},

			{EVR_pixelSQ, "EVR_pixelSQ", "ps_EVR_pixelSQ", 4, DCMVR_PROP_NONSTANDARD | DCMVR_PROP_INTERNAL, 0, DCM_UndefinedLength},

			{EVR_pixelItem, "EVR_pixelItem", "pi", 4, DCMVR_PROP_NONSTANDARD, 0, DCM_UndefinedLength},

			{EVR_UNKNOWN, "EVR_UNKNOWN", "??", 4, DCMVR_PROP_NONSTANDARD | DCMVR_PROP_INTERNAL | DCMVR_PROP_EXTENDEDLENGTHENCODING, 0, DCM_UndefinedLength},

			{EVR_UN, "EVR_UN", "UN", 4, DCMVR_PROP_EXTENDEDLENGTHENCODING, 0, DCM_UndefinedLength},

			{EVR_PixelData, "EVR_PixelData", "PixelData", 0, DCMVR_PROP_INTERNAL, 0, DCM_UndefinedLength},

			{EVR_OverlayData, "EVR_OverlayData", "OverlayData", 0, DCMVR_PROP_INTERNAL, 0, DCM_UndefinedLength},

			{EVR_UNKNOWN2B, "EVR_UNKNOWN2B", "??", 4, DCMVR_PROP_NONSTANDARD | DCMVR_PROP_INTERNAL, 0, DCM_UndefinedLength},

		*/
	}

	for _, c := range cases {
		got := c.in.String()
		if got != c.want {
			t.Errorf("String(%q) == %v, want %v", c.in, got, c.want)
		}
	}

}
