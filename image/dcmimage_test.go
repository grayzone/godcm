package image

import (
	"os"
	"strconv"
	"testing"

	"github.com/grayzone/godcm/core"
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
		in string
	}{
		{"MR-MONO2-8-16x-heart.dcm"},
		{"US-MONO2-8-8x-execho.dcm"},
		{"xr_tspine.dcm"},
		{"xr_chest.dcm"},
		{"IM0.dcm"},
		{"image_09-12-2013_4.dcm"},
		{"CT-MONO2-16-ankle"},
		{"xr_chicken2.dcm"},
		{"T23/IM-0001-0001.dcm"},
		{"IM-0001-0010.dcm"},
		{"GH195.dcm"},
		{"GH064.dcm"},
		{"GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm"},
		{"GH177_D_CLUNIE_CT1_IVRLE_BigEndian_ELE_undefinded_length.dcm"},
	}
	for _, c := range cases {
		var reader core.DcmReader
		reader.IsReadPixel = true
		reader.IsReadValue = true
		filepath := gettestdatafolder() + c.in
		err := reader.ReadFile(filepath)
		isCompressed, err := reader.IsCompressed()
		if err != nil {
			t.Errorf("WriteBMP() %s", err.Error())
		}

		pixeldata := reader.Dataset.PixelData()

		var img DcmImage

		img.IsCompressed = isCompressed

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

		num, _ = strconv.ParseUint(reader.Dataset.PixelRepresentation(), 10, 16)
		img.PixelRepresentation = uint16(num.(uint64))

		img.PhotometricInterpretation = reader.Dataset.PhotometricInterpretation()

		img.PixelData = pixeldata

		//		sop := reader.Dataset.SOPInstanceUID()
		err = img.WriteBMP(c.in+".bmp", 8, 0)
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
