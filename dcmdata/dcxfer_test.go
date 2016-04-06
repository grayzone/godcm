package dcmdata

import "testing"

func TestFindMachineTransferSyntax(t *testing.T) {
	cases := []struct {
		want E_ByteOrder
	}{
		{EBO_LittleEndian},
	}
	for _, c := range cases {
		got := FindMachineTransferSyntax()
		if got != c.want {
			t.Errorf("FindMachineTransferSyntax(), want %q got %q", c.want, got)
		}
	}

}

func TestNewDcmXfer(t *testing.T) {
	cases := []struct {
		in   E_TransferSyntax
		want DcmXfer
	}{
		{EXS_LittleEndianImplicit,
			DcmXfer{UID_LittleEndianImplicitTransferSyntax,
				"Little Endian Implicit",
				EXS_LittleEndianImplicit,
				EBO_LittleEndian,
				EVT_Implicit,
				EJE_NotEncapsulated,
				0, 0,
				false,
				ESC_none}},
	}
	for _, c := range cases {
		got := NewDcmXfer(c.in)
		if *got != c.want {
			t.Errorf("NewDcmXfer(%d), want %q got %q", c.in, c.want.xferName, got.xferName)
		}
	}

}

func TestDcmxferToString(t *testing.T) {
	cases := []struct {
		in   E_ByteOrder
		want string
	}{
		{EBO_unknown, "EBO_unknown"},
		{EBO_LittleEndian, "EBO_LittleEndian"},
		{EBO_BigEndian, "EBO_BigEndian"},
	}
	for _, c := range cases {
		got := c.in.String()
		if got != c.want {
			t.Errorf("String(%d), want %q got %q", c.in, c.want, got)
		}
	}

}

func TestIsExplicitVR(t *testing.T) {
	cases := []struct {
		in   E_TransferSyntax
		want bool
	}{
		{EXS_LittleEndianImplicit, false},
		{EXS_BigEndianImplicit, false},
		{EXS_LittleEndianExplicit, true},
		{EXS_BigEndianExplicit, true},
		{EXS_JPEGProcess1TransferSyntax, true},
		{EXS_JPEGProcess2_4TransferSyntax, true},
		{EXS_JPEGProcess3_5TransferSyntax, true},
		{EXS_JPEGProcess6_8TransferSyntax, true},
		{EXS_JPEGProcess7_9TransferSyntax, true},
		{EXS_JPEGProcess10_12TransferSyntax, true},
		{EXS_JPEGProcess11_13TransferSyntax, true},
		{EXS_JPEGProcess14TransferSyntax, true},
		{EXS_JPEGProcess15TransferSyntax, true},
		{EXS_JPEGProcess16_18TransferSyntax, true},
		{EXS_JPEGProcess17_19TransferSyntax, true},
		{EXS_JPEGProcess20_22TransferSyntax, true},
		{EXS_JPEGProcess21_23TransferSyntax, true},
		{EXS_JPEGProcess24_26TransferSyntax, true},
		{EXS_JPEGProcess25_27TransferSyntax, true},
		{EXS_JPEGProcess28TransferSyntax, true},
		{EXS_JPEGProcess29TransferSyntax, true},
		{EXS_JPEGProcess14SV1TransferSyntax, true},
		{EXS_RLELossless, true},
		{EXS_JPEGLSLossless, true},
		{EXS_JPEGLSLossy, true},
		{EXS_DeflatedLittleEndianExplicit, true},
		{EXS_JPEG2000osslessOnly, true},
		{EXS_JPEG2000, true},
		{EXS_MPEG2MainProfileAtMainLevel, true},
		{EXS_MPEG2MainProfileAtHighLevel, true},
		{EXS_JPEG2000MulticomponentLosslessOnly, true},
		{EXS_JPEG2000Multicomponent, true},
		{EXS_JPIPReferenced, true},
		{EXS_JPIPReferencedDeflate, true},
	}
	for _, c := range cases {
		got := NewDcmXfer(c.in).IsExplicitVR()
		if got != c.want {
			t.Errorf("IsExplicitVR(%v), want %v got %v", c.in, c.want, got)
		}
	}

}
