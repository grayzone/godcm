package dcmdata

import (
	"testing"
)

func TestNewDcmMetaInfo(t *testing.T) {
	cases := []struct {
		want *DcmMetaInfo
	}{
		{&DcmMetaInfo{filePreamble: "", preambleUsed: false, fPreambleTransferState: ERW_init, xfer: META_HEADER_DEFAULT_TRANSFERSYNTAX}},
	}
	for _, c := range cases {
		got := NewDcmMetaInfo()
		if (got.filePreamble != c.want.filePreamble) || (got.preambleUsed != c.want.preambleUsed) || (got.fPreambleTransferState != c.want.fPreambleTransferState) || (got.xfer != c.want.xfer) {
			t.Errorf("NewDcmMetaInfo() == want %v got %v", c.want, got)
		}
	}

}
