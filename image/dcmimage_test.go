package image

import "testing"

func TestWrite8BMP(t *testing.T) {
	filename := "8.bmp"

	var di DcmImage

	di.Columns = 512
	di.Rows = 512
	di.BitsAllocated = 8
	di.BitsStored = 12
	di.HighBit = 11
	err := di.WriteBMP(filename, 8, 0)
	if err != nil {
		t.Errorf("WriteBMP() %s", err.Error())
	}
}

func TestWrite16BMP(t *testing.T) {
	filename := "16.bmp"

	var di DcmImage

	di.Columns = 512
	di.Rows = 512
	di.BitsAllocated = 16
	di.BitsStored = 12
	di.HighBit = 11
	err := di.WriteBMP(filename, 16, 0)
	if err == nil {
		t.Errorf("WriteBMP() %s", err.Error())
	}
}

func TestWrite24BMP(t *testing.T) {
	filename := "24.bmp"

	var di DcmImage

	di.Columns = 64
	di.Rows = 64
	di.BitsAllocated = 24
	di.BitsStored = 12
	di.HighBit = 11
	err := di.WriteBMP(filename, 24, 0)
	if err != nil {
		t.Errorf("WriteBMP() %s", err.Error())
	}
}

func TestWrite32BMP(t *testing.T) {
	filename := "32.bmp"

	var di DcmImage

	di.Columns = 64
	di.Rows = 64
	di.BitsAllocated = 32
	di.BitsStored = 12
	di.HighBit = 11
	err := di.WriteBMP(filename, 32, 0)
	if err != nil {
		t.Errorf("WriteBMP() %s", err.Error())
	}
}
