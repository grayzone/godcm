package dcmimage

import (
	"errors"
	"image"
	"image/color"
	"image/jpeg"
	"log" // for debug
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
	log.Println("pixel data length:", len(pixelData))
	d := di.convertTo8Bit(pixelData)

	if !di.IsMonochrome() {
		m := image.NewRGBA(image.Rect(int(di.Columns), int(di.Rows), 0, 0))
		var index int
		for y := 0; y < int(di.Rows); y++ {
			for x := 0; x < int(di.Columns); x++ {
				r := d[index]
				index++
				g := d[index]
				index++
				b := d[index]
				index++
				c := color.RGBA{r, g, b, 0}
				m.Set(x, y, c)
			}
		}
		return jpeg.Encode(outfile, m, nil)
	}

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
	/*
		o := new(jpeg.Options)
		o.Quality = 100
	*/
	return jpeg.Encode(outfile, m, nil)

}
