package image

import (
	"bytes"
	"encoding/binary"
	"errors"
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
	biWidth         int32
	biHeight        int32
	biPlanes        uint16
	biBitCount      uint16
	biCompression   uint32
	biSizeImage     uint32
	biXPelsPerMeter int32
	biYPelsPerMeter int32
	biClrUsed       uint32
	biClrImportant  uint32
}

// DcmImage provide the "DICOM image toolkit"
type DcmImage struct {
	Rows          int32
	Columns       int32
	PixelWidth    float64
	PixelHeight   float64
	BitsAllocated uint16
	BitsStored    uint16
	HighBit       uint16
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
	fileHeader.bfSize = 54 + uint32(image.Rows*image.Columns*3)
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
	infoHeader.biSizeImage = 0
	infoHeader.biXPelsPerMeter = 0
	infoHeader.biYPelsPerMeter = 0
	infoHeader.biClrUsed = 0
	infoHeader.biClrImportant = 0

	f, _ := os.Create(filename)
	defer f.Close()

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, fileHeader.bfType)
	binary.Write(buf, binary.LittleEndian, fileHeader.bfSize)
	binary.Write(buf, binary.LittleEndian, fileHeader.bfReserved1)
	binary.Write(buf, binary.LittleEndian, fileHeader.bfReserved2)
	binary.Write(buf, binary.LittleEndian, fileHeader.bfOffBits)

	binary.Write(buf, binary.LittleEndian, infoHeader.bitSize)
	binary.Write(buf, binary.LittleEndian, infoHeader.biWidth)
	binary.Write(buf, binary.LittleEndian, infoHeader.biHeight)
	binary.Write(buf, binary.LittleEndian, infoHeader.biPlanes)
	binary.Write(buf, binary.LittleEndian, infoHeader.biBitCount)
	binary.Write(buf, binary.LittleEndian, infoHeader.biCompression)

	binary.Write(buf, binary.LittleEndian, infoHeader.biSizeImage)

	binary.Write(buf, binary.LittleEndian, infoHeader.biXPelsPerMeter)
	binary.Write(buf, binary.LittleEndian, infoHeader.biYPelsPerMeter)
	binary.Write(buf, binary.LittleEndian, infoHeader.biClrUsed)
	binary.Write(buf, binary.LittleEndian, infoHeader.biClrImportant)

	if palette != nil {
		binary.Write(buf, binary.LittleEndian, palette)
	}

	var data []uint8

	for i := int32(0); i < image.Columns; i++ {
		for j := int32(0); j < image.Rows; j++ {
			data = append(data, 16)
		}
	}
	binary.Write(buf, binary.LittleEndian, data)

	f.Write(buf.Bytes())

	return nil
}
