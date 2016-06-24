package dcmimage

import (
	"image/png"
	"os"
)

// ConvertToPNG convert dicom file to png file.
func (di DcmImage) ConvertToPNG(filepath string, frame int) error {
	m, err := di.convertToImage(frame)
	if err != nil {
		return err
	}
	outfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer outfile.Close()
	return png.Encode(outfile, m)
}
