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

/*
func (di DcmImage) window16(pixel int16) uint16 {

	if (di.WindowCenter == 0.0) && (di.WindowWidth == 0.0) {
		return di.nowindow16(pixel)
	}

	min := di.WindowCenter - di.WindowWidth/2.0 + 0.5
	max := di.WindowCenter + di.WindowWidth/2.0 + 0.5
	slope := 65535.0 / di.WindowWidth
	value := float64(pixel)
	if value < min {
		return 0
	} else if value > max {
		return 65535
	}
	value = (value - min) * slope

	if value > 65535 {
		return 65535
	}
	if value < 0 {
		return 0
	}
	return uint16(value)
}

func (di DcmImage) nowindow16(pixel int16) uint16 {
	gradient := 65535.0 / float64(di.maxValue-di.minValue)
	value := float64(int16(pixel)-di.minValue) * gradient
	if value > 65535 {
		return 65535
	}
	if value < 0 {
		return 0
	}
	return uint16(value)
}

func (di DcmImage) convertToPNG16(filepath string) error {
	if di.IsCompressed {
		err := errors.New("not supported compressed format")
		return err
	}
	outfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer outfile.Close()

	//	di.convertTo8Bit()

	m := image.NewGray16(image.Rect(int(di.Columns), int(di.Rows), 0, 0))
	di.determinReverse()
	di.determineMinMax()
	for y := di.Rows; y > uint32(0); y-- {
		for x := uint32(0); x < di.Columns; x++ {
			var pixel uint16
			b := di.PixelData[2*di.Columns*y-2*di.Columns+2*x : 2*di.Columns*y-2*di.Columns+2*x+2]
			if di.IsBigEndian {
				pixel = binary.BigEndian.Uint16(b)
			} else {
				pixel = binary.LittleEndian.Uint16(b)
			}

			//		pixel = uint16(di.clipHighBits(int16(pixel)))
			var p int16
			p = di.rescalePixel(int16(pixel))
			pixel = di.window16(p)
			if di.IsReverse {
				pixel = 65535 - pixel
			}
			c := color.Gray16{pixel}
			m.SetGray16(int(x), int(y), c)
		}
	}
	return png.Encode(outfile, m)
}
*/
