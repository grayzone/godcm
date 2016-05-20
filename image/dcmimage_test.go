package image

import (
	"testing"
)

func TestWriteBMP(t *testing.T) {
	filename := "test.bmp"

	var di DcmImage

	di.Columns = 64
	di.Rows = 64
	di.BitsAllocated = 16
	di.BitsStored = 12
	di.HighBit = 11
	err := di.WriteBMP(filename, 8, 0)
	if err != nil {
		t.Errorf("WriteBMP() %s", err.Error())
	}
}
