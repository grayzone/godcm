package image

import (
	"bytes"
	"encoding/binary"
	"errors"
	_ "log" // for debug
	"os"
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

	PixelData []byte
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
	if image.PixelRepresentation == 0 {
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

func (image DcmImage) rescaleWindowLevel(pixel int16) uint8 {
	var value float64
	if (image.WindowCenter == 0.0) && (image.WindowWidth == 0.0) {
		var slope float64
		if image.minValue != image.maxValue {
			slope = 255.0 / float64(image.maxValue-image.minValue)
		} else {
			slope = 1.0
		}
		value = float64(pixel-image.minValue) * slope
	} else {

		shift := image.WindowCenter - image.WindowWidth/2.0
		slope := 255.0 / image.WindowWidth
		value = (float64(pixel) - shift) * slope
	}
	var result uint8
	if value < 0 {
		result = 0
	} else if value > 255 {
		result = 255
	} else {
		result = uint8(value)
	}
	return result
}

func (image DcmImage) convertTo8Bit() []uint8 {
	var result []uint8
	count := image.Rows * image.Columns
	image.findPixelExtremeValue()

	for i := uint32(0); i < count; i++ {
		p := binary.LittleEndian.Uint16(image.PixelData[2*i : 2*i+2])

		pixel := image.clipHighBits(int16(p))
		pixel = image.rescalePixel(pixel)

		b := image.rescaleWindowLevel(pixel)
		result = append(result, b)
	}
	return result
}

func (image DcmImage) findPixelExtremeValue() {
	// skip to find the max/min value if window level is not 0
	if (image.WindowCenter != 0.0) || (image.WindowWidth != 0.0) {
		return
	}
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
}
