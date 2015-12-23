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
