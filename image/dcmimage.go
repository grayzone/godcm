package image

import (
	"bytes"
	"encoding/binary"
	"errors"
	"log" // for debug
	"os"
)

var (
	MAXBITS uint16 = 32
)

// BitmapFileHeader is for BMP file header
type BitmapFileHeader struct {
	bfType      [2]byte // must always be 'BM'
	bfSize      uint32
	bfReserved1 uint16 // reserved, should be '0'
	bfReserved2 uint16 // reserved, should be '0'
	bfOffBits   uint32
}

// BitmapInfoHeader is for BMP info header
type BitmapInfoHeader struct {
	bitSize         uint32
	biWidth         uint32
	biHeight        uint32
	biPlanes        uint16
	biBitCount      uint16
	biCompression   uint32
	biSizeImage     uint32
	biXPelsPerMeter uint32
	biYPelsPerMeter uint32
	biClrUsed       uint32
	biClrImportant  uint32
}

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
	PlanarConfiguration       uint16
	RescaleIntercept          float64
	RescaleSlope              float64
	WindowCenter              float64
	WindowWidth               float64

	RescaleType          string
	PresentationLUTShape string

	minValue int16
	maxValue int16

	AbsMinimum float64
	AbsMaximum float64

	high float64
	low  float64

	PixelData []byte
}

func maxval(bits uint16, pos uint32) uint32 {
	if bits > MAXBITS {
		return uint32(1<<bits - 1)
	}
	return uint32(1<<bits - pos)
}

// WriteBMP write pixel data to BMP file
func (image DcmImage) WriteBMP(filename string, bits uint16, frame int) error {
	switch bits {
	case 8:
	case 24:
	case 32:
	default:
		err := errors.New("not supported BMP format")
		return err
	}
	var fileHeader BitmapFileHeader
	fileHeader.bfType[0] = 'B'
	fileHeader.bfType[1] = 'M'
	fileHeader.bfSize = 54 + uint32(image.Rows*image.Columns)
	fileHeader.bfReserved1 = 0
	fileHeader.bfReserved2 = 0
	fileHeader.bfOffBits = 54

	var palette *[256]uint32
	if bits == 8 {
		palette = new([256]uint32)
		fileHeader.bfSize += 1024
		fileHeader.bfOffBits += 1024
		for i := uint32(0); i < 256; i++ {
			palette[i] = uint32((i << 16) | (i << 8) | i)
		}
	}

	var infoHeader BitmapInfoHeader
	infoHeader.bitSize = 40
	infoHeader.biWidth = image.Columns
	infoHeader.biHeight = image.Rows
	infoHeader.biPlanes = 1
	infoHeader.biBitCount = bits
	infoHeader.biCompression = 0
	infoHeader.biSizeImage = uint32((uint32(infoHeader.biWidth)*uint32(infoHeader.biBitCount) + 31) / 32 * 4 * infoHeader.biHeight)
	infoHeader.biXPelsPerMeter = 0
	infoHeader.biYPelsPerMeter = 0
	infoHeader.biClrUsed = 0
	infoHeader.biClrImportant = 0

	f, _ := os.Create(filename)
	defer f.Close()

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, fileHeader)
	binary.Write(buf, binary.LittleEndian, infoHeader)

	if palette != nil {
		binary.Write(buf, binary.LittleEndian, palette)
	}

	data := image.convertTo8Bit()

	binary.Write(buf, binary.LittleEndian, data)
	f.Write(buf.Bytes())
	return nil
}

func (image DcmImage) clipHighBits(pixel int16) int16 {
	if image.HighBit > 15 {
		return pixel
	}
	nMask := 0xffff << (image.HighBit + 1)
	if image.PixelRepresentation != 0 {
		nSignBit := 1 << image.HighBit
		if (pixel & int16(nSignBit)) != 0 {
			pixel |= int16(nMask)
			return pixel
		}
	}
	pixel &= ^int16(nMask)
	return pixel
}

func (image DcmImage) rescalePixel(pixel int16) int16 {
	if image.RescaleSlope == 1.0 && image.RescaleIntercept == 0.0 {
		return pixel
	}
	return image.getRescaling(pixel).(int16)
}

func (image DcmImage) getRescaling(value interface{}) interface{} {
	switch value.(type) {
	case int16:
		return int16(float64(value.(int16))*image.RescaleSlope + image.RescaleIntercept)
	case float64:
		return value.(float64)*image.RescaleSlope + image.RescaleIntercept
	}
	return 0
}

func (image *DcmImage) checkRescaling() {

	if image.RescaleSlope < 0 {
		var tmp interface{}
		tmp = image.minValue
		image.minValue = image.getRescaling(image.maxValue).(int16)
		image.maxValue = image.getRescaling(tmp).(int16)

		tmp = image.AbsMinimum
		image.AbsMinimum = image.getRescaling(image.AbsMaximum).(float64)
		image.AbsMaximum = image.getRescaling(tmp).(float64)
		return
	}
	image.minValue = image.getRescaling(image.minValue).(int16)
	image.maxValue = image.getRescaling(image.maxValue).(int16)

	image.AbsMinimum = image.getRescaling(image.AbsMinimum).(float64)
	image.AbsMaximum = image.getRescaling(image.AbsMaximum).(float64)
}

func (image DcmImage) nowindow(pixel int16) uint8 {
	//	var outrange float64
	//	outrange = image.high - image.low + 1

	//	ocnt := image.getAbsMaxRange()
	gradient := 255.0 / float64(image.maxValue-image.minValue)
	value := float64(pixel-image.minValue) * gradient
	if value > 255 {
		return 255
	}
	if value < 0 {
		return 0
	}
	return uint8(value)
}

func (image DcmImage) window(pixel int16) uint8 {

	/*
		width1 := image.WindowWidth - 1
		outrange := image.high - image.low
		var offset float64
		var gradient float64
		if width1 != 0 {
			offset = (image.high - (image.WindowCenter-0.5)/width1 + 0.5) * outrange
			gradient = outrange / width1
		}
		value := offset + float64(pixel)*gradient
	*/

	min := image.WindowCenter - image.WindowWidth/2.0 + 0.5
	max := image.WindowCenter + image.WindowWidth/2.0 + 0.5
	slope := 255.0 / image.WindowWidth
	value := float64(pixel)
	if value < min {
		return 0
	} else if value > max {
		return 255
	}
	value = (value - min) * slope

	if value > 255 {
		return 255
	}
	if value < 0 {
		return 0
	}
	return uint8(value)

}

func (image DcmImage) rescaleWindowLevel(pixel int16) uint8 {
	if (image.WindowCenter == 0.0) && (image.WindowWidth == 0.0) {
		return image.nowindow(pixel)
	}
	return image.window(pixel)
}

func (image DcmImage) convertTo8Bit() []uint8 {
	image.determineMinMax()
	image.determinHighLow()
	var result []uint8
	gap := (4 - (image.Columns & 0x3)) & 0x3
	for i := image.Rows; i > uint32(0); i-- {
		for j := uint32(0); j < image.Columns; j++ {
			p := binary.LittleEndian.Uint16(image.PixelData[2*image.Columns*i-2*image.Columns+2*j : 2*image.Columns*i-2*image.Columns+2*j+2])

			pixel := image.clipHighBits(int16(p))
			pixel = image.rescalePixel(pixel)

			b := image.rescaleWindowLevel(pixel)
			//b := image.nowindow(pixel)
			result = append(result, uint8(b))
		}
		for i := uint32(0); i < gap; i++ {
			result = append(result, uint8(0))
		}
	}
	return result
}

func (image *DcmImage) findAbsMaxMinValue() {
	if image.PixelRepresentation == 1 {
		image.AbsMinimum = -float64(maxval(image.BitsStored-1, 0))
		image.AbsMaximum = float64(maxval(image.BitsStored-1, 1))
		return
	}
	image.AbsMinimum = 0
	image.AbsMaximum = float64(maxval(image.BitsStored, 1))
}

func (image DcmImage) getAbsMaxRange() float64 {
	image.findAbsMaxMinValue()
	return image.AbsMaximum - image.AbsMinimum + 1
}

func (image *DcmImage) determinHighLow() {
	if image.PixelRepresentation == 0 {
		image.high = 0
		image.low = 255
		return
	}
	image.high = 255
	image.low = 0
}

func (image *DcmImage) determineMinMax() {
	// skip to find the max/min value if window level is not 0
	/*
		if (image.WindowCenter != 0.0) || (image.WindowWidth != 0.0) {
			return
		}
	*/
	image.findAbsMaxMinValue()
	count := image.Columns * image.Rows
	for i := uint32(0); i < count; i++ {
		p := int16(binary.LittleEndian.Uint16(image.PixelData[2*i : 2*i+2]))
		if i == 0 {
			image.minValue = p
			image.maxValue = p
		}
		if p < image.minValue {
			image.minValue = p
		}
		if p > image.maxValue {
			image.maxValue = p
		}
	}

	log.Println("min1", image.minValue, "max1", image.maxValue, "absmin", image.AbsMinimum, "absmax", image.AbsMaximum)
	image.checkRescaling()
	log.Println("min2", image.minValue, "max2", image.maxValue)
}
