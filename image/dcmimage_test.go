package image

import (
	"github.com/grayzone/godcm/core"
	"os"
	"strconv"
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

func TestWrite8BMP(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{gettestdatafolder() + "IM0.dcm", "1"},
	}
	for _, c := range cases {
		var reader core.DcmReader
		reader.IsReadPixel = true
		reader.IsReadValue = true
		err := reader.ReadFile(c.in)
		pixeldata := reader.Dataset.PixelData()

		var img DcmImage

		var num interface{}

		num, _ = strconv.ParseUint(reader.Dataset.BitsAllocated(), 10, 16)
		img.BitsAllocated = uint16(num.(uint64))

		num, _ = strconv.ParseUint(reader.Dataset.Columns(), 10, 32)
		img.Columns = uint32(num.(uint64))

		num, _ = strconv.ParseUint(reader.Dataset.Rows(), 10, 32)
		img.Rows = uint32(num.(uint64))

		num, _ = strconv.ParseUint(reader.Dataset.HighBit(), 10, 16)
		img.HighBit = uint16(num.(uint64))

		num, _ = strconv.ParseFloat(reader.Dataset.WindowCenter(), 64)
		img.WindowCenter = num.(float64)

		num, _ = strconv.ParseFloat(reader.Dataset.WindowWidth(), 64)
		img.WindowWidth = num.(float64)

		num, _ = strconv.ParseFloat(reader.Dataset.RescaleIntercept(), 64)
		img.RescaleIntercept = num.(float64)

		num, _ = strconv.ParseFloat(reader.Dataset.RescaleSlope(), 64)
		img.RescaleSlope = num.(float64)

		img.PixelData = pixeldata

		err = img.WriteBMP("test.bmp", 8, 0)
		if err != nil {
			t.Errorf("WriteBMP() %s", err.Error())
		}
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
