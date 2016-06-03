package dcmimage

import (
	"errors"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

// ConvertToJPG convert dicom file to jpg file
func (di DcmImage) ConvertToJPG(filepath string, frame int) error {
	if di.IsCompressed {
		err := errors.New("not supported compressed format")
		return err
	}
	outfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer outfile.Close()

	pixelData, err := di.getPixelDataOfFrame(frame)
	if err != nil {
		return err
	}
	d := di.convertTo8Bit(pixelData)

	m := image.NewGray(image.Rect(int(di.Columns), int(di.Rows), 0, 0))
	var index int
	for y := 0; y < int(di.Rows); y++ {
		for x := 0; x < int(di.Columns); x++ {
			r := d[index]
			index++
			c := color.Gray{r}
			m.SetGray(x, y, c)
		}
	}
	return jpeg.Encode(outfile, m, nil)
}
