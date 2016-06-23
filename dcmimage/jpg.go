package dcmimage

import (
	"image/jpeg"
	_ "log" // for debug
	"os"
)

// ConvertToJPG convert dicom file to jpg file
func (di DcmImage) ConvertToJPG(filepath string, frame int) error {
	m, err := di.convertToImage(frame)
	if err != nil {
		return err
	}
	outfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer outfile.Close()
	/*
		o := new(jpeg.Options)
		o.Quality = 100
	*/
	return jpeg.Encode(outfile, m, nil)

}
