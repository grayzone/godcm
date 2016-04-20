package core

import "testing"

func TestNewDcmXfer(t *testing.T) {
	cases := []struct {
		in   ETransferSyntax
		want DcmXfer
	}{
		{EXSLittleEndianImplicit,
			DcmXfer{UIDLittleEndianImplicitTransferSyntax,
				"Little Endian Implicit",
				EXSLittleEndianImplicit,
				EBOLittleEndian,
				EVTImplicit,
				EJENotEncapsulated,
				0, 0,
				false,
				ESCnone}},
	}
	for _, c := range cases {
		got := NewDcmXfer(c.in)
		if *got != c.want {
			t.Errorf("NewDcmXfer(%d), want %q got %q", c.in, c.want.XferName, got.XferName)
		}
	}

}

func TestDcmxferToString(t *testing.T) {
	cases := []struct {
		in   EByteOrder
		want string
	}{
		{EBOunknown, "Unknown"},
		{EBOLittleEndian, "LittleEndian"},
		{EBOBigEndian, "BigEndian"},
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
		in   ETransferSyntax
		want bool
	}{
		{EXSLittleEndianImplicit, false},
		{EXSBigEndianImplicit, false},
		{EXSLittleEndianExplicit, true},
		{EXSBigEndianExplicit, true},
		{EXSJPEGProcess1TransferSyntax, true},
		{EXSJPEGProcess24TransferSyntax, true},
		{EXSJPEGProcess35TransferSyntax, true},
		{EXSJPEGProcess68TransferSyntax, true},
		{EXSJPEGProcess79TransferSyntax, true},
		{EXSJPEGProcess1012TransferSyntax, true},
		{EXSJPEGProcess1113TransferSyntax, true},
		{EXSJPEGProcess14TransferSyntax, true},
		{EXSJPEGProcess15TransferSyntax, true},
		{EXSJPEGProcess1618TransferSyntax, true},
		{EXSJPEGProcess1719TransferSyntax, true},
		{EXSJPEGProcess2022TransferSyntax, true},
		{EXSJPEGProcess2123TransferSyntax, true},
		{EXSJPEGProcess2426TransferSyntax, true},
		{EXSJPEGProcess2527TransferSyntax, true},
		{EXSJPEGProcess28TransferSyntax, true},
		{EXSJPEGProcess29TransferSyntax, true},
		{EXSJPEGProcess14SV1TransferSyntax, true},
		{EXSRLELossless, true},
		{EXSJPEGLSLossless, true},
		{EXSJPEGLSLossy, true},
		{EXSDeflatedLittleEndianExplicit, true},
		{EXSJPEG2000osslessOnly, true},
		{EXSJPEG2000, true},
		{EXSMPEG2MainProfileAtMainLevel, true},
		{EXSMPEG2MainProfileAtHighLevel, true},
		{EXSJPEG2000MulticomponentLosslessOnly, true},
		{EXSJPEG2000Multicomponent, true},
		{EXSJPIPReferenced, true},
		{EXSJPIPReferencedDeflate, true},
	}
	for _, c := range cases {
		got := NewDcmXfer(c.in).IsExplicitVR()
		if got != c.want {
			t.Errorf("IsExplicitVR(%v), want %v got %v", c.in, c.want, got)
		}
	}

}
