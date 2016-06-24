package dcmimage

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

// WriteBMP write pixel data to BMP file
func (di DcmImage) WriteBMP(filename string, bits uint16, frame int) error {
	if di.IsCompressed {
		err := errors.New("not support compressed format")
		return err
	}
	switch bits {
	case 8:
		if di.IsMonochrome() == false {
			err := errors.New("not support color image in 8 bits")
			return err
		}
	case 24:
	case 32:
	default:
		err := errors.New("not support BMP bits")
		return err
	}

	pixelData, err := di.getPixelDataOfFrame(frame)
	if err != nil {
		return err
	}
	pixel := di.convertTo8Bit(pixelData)
	d := di.convertToImageData(pixel)

	//	log.Println("pixel data length", len(d))

	var fileHeader BitmapFileHeader
	fileHeader.bfType[0] = 'B'
	fileHeader.bfType[1] = 'M'
	fileHeader.bfSize = 54 + uint32(di.Rows*di.Columns)
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
	infoHeader.biWidth = di.Columns
	infoHeader.biHeight = di.Rows
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

	rgbplane := bits / 8

	gap := rgbplane * uint16((4-(di.Columns&0x3))&0x3)

	for i := di.Rows; i > uint32(0); i-- {
		for j := uint32(0); j < di.Columns; j++ {
			r := d[3*di.Columns*(i-1)+3*j : 3*di.Columns*(i-1)+3*j+1]
			g := d[3*di.Columns*(i-1)+3*j+1 : 3*di.Columns*(i-1)+3*j+2]
			b := d[3*di.Columns*(i-1)+3*j+2 : 3*di.Columns*(i-1)+3*j+3]

			switch bits {
			case 8:
				binary.Write(buf, binary.LittleEndian, r)
			case 24:
				binary.Write(buf, binary.LittleEndian, b)
				binary.Write(buf, binary.LittleEndian, g)
				binary.Write(buf, binary.LittleEndian, r)
			case 32:
				p := uint32(r[0])<<16 | uint32(g[0])<<8 | uint32(b[0])
				binary.Write(buf, binary.LittleEndian, p)
			}

		}

		if bits != 32 {
			for i := uint16(0); i < gap; i++ {
				binary.Write(buf, binary.LittleEndian, uint8(0))
			}
		}
	}
	f.Write(buf.Bytes())
	return nil
}
