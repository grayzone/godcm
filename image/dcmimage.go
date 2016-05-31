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

	PixelData []byte
}

func maxval(bits uint16, pos uint32) uint32 {
	if bits > MAXBITS {
		return uint32(1<<bits - 1)
	}
	return uint32(1<<bits - pos)
}

func expandSign(value uint16, signBit uint16, signMask uint16) uint16 {
	if (value & signBit) == 0 {
		return value
	}
	return value | signMask
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
	return int16(float64(pixel)*image.RescaleSlope + image.RescaleIntercept)
}

/*
func (image DcmImage) covertPixelToInputRepresentation(pixel uint16) uint16 {
	var mask uint16
	for i := uint16(0); i < image.BitsStored; i++ {
		mask |= 1 << i
	}
	sign := uint16(1 << (image.BitsStored - 1))
	var smask uint16
	for i := image.BitsStored; i < 16; i++ {
		smask |= uint16(1 << i)
	}
	shift := image.HighBit + 1 - image.BitsStored

	if shift == 0 {
		return expandSign(pixel&uint16(mask), sign, smask)
	}
	return expandSign((pixel>>shift)&uint16(mask), sign, smask)
}

*/

func (image DcmImage) countRescaling(value interface{}) interface{} {
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
		image.minValue = image.countRescaling(image.maxValue).(int16)
		image.maxValue = image.countRescaling(tmp).(int16)

		tmp = image.AbsMinimum
		image.AbsMinimum = image.countRescaling(image.AbsMaximum).(float64)
		image.AbsMaximum = image.countRescaling(tmp).(float64)
		return
	}
	image.minValue = image.countRescaling(image.minValue).(int16)
	image.maxValue = image.countRescaling(image.maxValue).(int16)

	image.AbsMinimum = image.countRescaling(image.AbsMinimum).(float64)
	image.AbsMaximum = image.countRescaling(image.AbsMaximum).(float64)
}

func (image DcmImage) nowindow(pixel int16) uint8 {
	var outrange float64
	if image.PixelRepresentation == 1 {
		outrange = float64(maxval(8, 0))
	} else {
		outrange = 0 - 255 + 1
	}

	ocnt := image.getAbsMaxRange()
	gradient := outrange / ocnt
	value := float64(pixel-image.minValue) * gradient
	return uint8(value)
}

func (image DcmImage) window(pixel int16) uint8 {
	shift := image.WindowCenter - image.WindowWidth/2.0
	slope := 255.0 / image.WindowWidth
	value := (float64(pixel) - shift) * slope
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
	var result []uint8
	gap := (4 - (image.Columns & 0x3)) & 0x3
	for i := image.Rows; i > uint32(0); i-- {
		for j := uint32(0); j < image.Columns; j++ {
			p := binary.LittleEndian.Uint16(image.PixelData[2*image.Columns*i-2*image.Columns+2*j : 2*image.Columns*i-2*image.Columns+2*j+2])

			//		pixel := image.covertPixelToInputRepresentation(int16(p))
			pixel := image.clipHighBits(int16(p))
			pixel = image.rescalePixel(pixel)

			//	b := image.rescaleWindowLevel(pixel)
			b := image.nowindow(pixel)
			result = append(result, uint8(b))
		}
		for i := uint32(0); i < gap; i++ {
			result = append(result, uint8(0))
		}
	}

	/*
		var data []uint16
		var data1 []uint16
		for i := uint32(0); i < image.Rows*image.Columns; i++ {
			p := binary.LittleEndian.Uint16(image.PixelData[2*i : 2*i+2])
			pixel := image.covertPixelToInputRepresentation(p)
			data = append(data, pixel)
			data1 = append(data1, p)
		}
		f, _ := os.Create("test.bin")
		defer f.Close()
		buf := new(bytes.Buffer)
		binary.Write(buf, binary.LittleEndian, data1)
		f.Write(buf.Bytes())
	*/
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

func (image *DcmImage) determineMinMax() {
	// skip to find the max/min value if window level is not 0

	if (image.WindowCenter != 0.0) || (image.WindowWidth != 0.0) {
		return
	}

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
