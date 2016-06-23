package dcmimage

import (
	"encoding/binary"
	"errors"
	"image"
	"image/color"
	_ "log" // for debug
)

var (
	MAXBITS uint16 = 32
)

// DcmImage provide the "DICOM image toolkit"
type DcmImage struct {
	Rows                      uint32
	Columns                   uint32
	PixelWidth                float64
	PixelHeight               float64
	BitsAllocated             uint16
	BitsStored                uint16
	HighBit                   uint16
	PhotometricInterpretation string
	SamplesPerPixel           uint16
	PixelRepresentation       uint16
	//	PlanarConfiguration       uint16
	RescaleIntercept float64
	RescaleSlope     float64
	WindowCenter     float64
	WindowWidth      float64

	IsReverse    bool
	IsCompressed bool
	IsBigEndian  bool

	//	RescaleType          string
	//	PresentationLUTShape string

	minValue int16
	maxValue int16

	//	AbsMinimum float64
	//	AbsMaximum float64

	high float64
	low  float64

	NumberOfFrames int
	PixelData      []byte
}

func maxval(bits uint16, pos uint32) uint32 {
	if bits > MAXBITS {
		return uint32(1<<bits - 1)
	}
	return uint32(1<<bits - pos)
}

func (di DcmImage) clipHighBits(pixel int16) int16 {
	if di.HighBit > 15 {
		return pixel
	}

	nMask := 0xffff << (di.HighBit + 1)
	if di.PixelRepresentation != 0 {
		nSignBit := 1 << di.HighBit
		if (pixel & int16(nSignBit)) != 0 {
			pixel |= int16(nMask)
			return pixel
		}
	}
	pixel &= ^int16(nMask)
	return pixel
}

func (di DcmImage) rescalePixel(pixel int16) int16 {
	if di.RescaleSlope == 1.0 && di.RescaleIntercept == 0.0 {
		return pixel
	}

	if di.RescaleSlope == 0 && di.RescaleIntercept == 0 {
		return pixel
	}

	return int16(float64(pixel)*di.RescaleSlope + di.RescaleIntercept)
}

func (di DcmImage) nowindow(pixel int16) uint8 {
	gradient := (di.high - di.low) / float64(di.maxValue-di.minValue)
	value := float64(pixel-di.minValue) * gradient
	if value > di.high {
		return uint8(di.high)
	}
	if value < 0 {
		return 0
	}
	return uint8(value)
}

func (di DcmImage) window(pixel int16) uint8 {
	min := di.WindowCenter - di.WindowWidth/2.0 + 0.5
	max := di.WindowCenter + di.WindowWidth/2.0 + 0.5
	gradient := (di.high - di.low) / di.WindowWidth
	value := float64(pixel)
	if value < min {
		return 0
	} else if value > max {
		return uint8(di.high)
	}
	value = (value - min) * gradient

	if value > di.high {
		return uint8(di.high)
	}
	if value < 0 {
		return 0
	}
	return uint8(value)

}

func (di DcmImage) rescaleWindowLevel(pixel int16) uint8 {
	if (di.WindowCenter == 0.0) && (di.WindowWidth == 0.0) {
		return di.nowindow(pixel)
	}
	return di.window(pixel)
}

func (di DcmImage) byteTouint8(pixelData []byte) []uint8 {
	var result []uint8
	for i := range pixelData {
		b := uint8(pixelData[i])
		if di.IsReverse {
			b = uint8(di.high) - b
		}
		result = append(result, b)
	}
	return result
}

func (di DcmImage) int16Touint8(pixelData []byte) []uint8 {
	var result []uint8
	count := di.Columns * di.Rows
	for i := uint32(0); i < count; i++ {
		b := pixelData[2*i : 2*i+2]
		var pixel int16
		if di.IsBigEndian {
			pixel = int16(binary.BigEndian.Uint16(b))
		} else {
			pixel = int16(binary.LittleEndian.Uint16(b))
		}
		pixel = di.clipHighBits(pixel)
		pixel = di.rescalePixel(pixel)
		p := di.rescaleWindowLevel(pixel)
		if di.IsReverse {
			p = uint8(di.high) - p
		}
		result = append(result, p)
	}
	return result
}

func (di DcmImage) convertTo8Bit(pixel []byte) []uint8 {
	di.determinReverse()
	if di.BitsAllocated <= 8 {
		return di.byteTouint8(pixel)
	}
	di.determineMinMax()
	return di.int16Touint8(pixel)
}

/*
func (di *DcmImage) findAbsMaxMinValue() {
	if di.PixelRepresentation == 1 {
		di.AbsMinimum = -float64(maxval(di.BitsStored-1, 0))
		di.AbsMaximum = float64(maxval(di.BitsStored-1, 1))
		return
	}
	di.AbsMinimum = 0
	di.AbsMaximum = float64(maxval(di.BitsStored, 1))
}
*/

/*
func (di DcmImage) getAbsMaxRange() float64 {
	di.findAbsMaxMinValue()
	return di.AbsMaximum - di.AbsMinimum + 1
}


func (di *DcmImage) determinHighLow() {
	if di.PixelRepresentation == 0 {
		di.high = 0
		di.low = 255
		return
	}
	di.high = 255
	di.low = 0
}
*/

// IsMonochrome check whether image is monochrome or not.
func (di DcmImage) IsMonochrome() bool {
	return di.PhotometricInterpretation == "MONOCHROME1" || di.PhotometricInterpretation == "MONOCHROME2"
}

func (di *DcmImage) determinReverse() {
	if di.PhotometricInterpretation == "MONOCHROME1" {
		di.IsReverse = true
	}
}

func (di *DcmImage) determineMinMax() {

	di.high = float64(maxval(8, 1))
	di.low = 0

	// skip to find the max/min value if window level is not 0

	if (di.WindowCenter != 0.0) || (di.WindowWidth != 0.0) {
		return
	}

	//	di.findAbsMaxMinValue()
	count := di.Columns * di.Rows
	for i := uint32(0); i < count; i++ {
		var pixel int16
		if di.BitsAllocated > 8 {
			b := di.PixelData[2*i : 2*i+2]
			if di.IsBigEndian {
				pixel = int16(binary.BigEndian.Uint16(b))
			} else {
				pixel = int16(binary.LittleEndian.Uint16(b))
			}
		} else {
			pixel = int16(di.PixelData[i])
		}

		if i == 0 {
			di.minValue = pixel
			di.maxValue = pixel
		}
		if pixel < di.minValue {
			di.minValue = pixel
		}
		if pixel > di.maxValue {
			di.maxValue = pixel
		}
	}

	//	log.Println("min", di.minValue, "max", di.maxValue)
}

func (di DcmImage) getPixelDataOfFrame(frame int) ([]byte, error) {
	size := int(di.Columns * di.Rows * uint32(di.SamplesPerPixel))
	if size == 0 {
		err := errors.New("getPixelDataOfFrame : SamplesPerPixel is zero")
		return nil, err
	}
	num := len(di.PixelData) / size

	if frame > num {
		err := errors.New("getPixelDataOfFrame : out of range")
		return nil, err
	}
	return di.PixelData[size*frame : size*frame+size], nil
}

func (di DcmImage) convertToImage(frame int) (image.Image, error) {
	if di.IsCompressed {
		err := errors.New("not supported compressed format")
		return nil, err
	}
	pixelData, err := di.getPixelDataOfFrame(frame)
	if err != nil {
		return nil, err
	}
	//	log.Println("pixel data length:", len(pixelData))
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
		return m, nil
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
	return m, nil
}
