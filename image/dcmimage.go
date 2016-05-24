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

	/*
		var data []uint32
		var index int
		for i := uint32(0); i < image.Columns; i++ {
			for j := uint32(0); j < image.Rows; j++ {
				data = append(data, uint8(image.PixelData[index]))
				index++
			}
		}
	*/
	//	data := image.convertTo8Bit()

	//	binary.Write(buf, binary.LittleEndian, data)
	var result []int16
	count := image.Rows * image.Columns
	for i := uint32(0); i < count; i++ {
		var p uint16
		//		buf := bytes.NewBuffer(image.PixelData[2*i : 2*i+2])
		p = binary.LittleEndian.Uint16(image.PixelData[2*i : 2*i+2])

		result = append(result, int16(p))
	}
	image.PixelRepresentation = 1
	if image.PixelRepresentation == 0 {
		var nMask int32
		nMask = 0xffff << (image.HighBit + 1)
		//	sign = 1 << image.HighBit

		for i := uint32(0); i < count; i++ {
			p := int32(result[i])
			p &= ^nMask
			result[i] = int16(p)
		}
	} else {
		var nSignBit int32
		nSignBit = 1 << image.HighBit
		var nMask int32
		nMask = 0xffff << (image.HighBit + 1)
		for i := uint32(0); i < count; i++ {
			p := int32(result[i])
			if (p & nSignBit) != 0 {
				p |= nMask
			} else {
				p &= ^nMask
			}
			result[i] = int16(p)
		}

	}

	shift := image.WindowCenter - image.WindowWidth/2.0
	slope := 255.0 / image.WindowWidth
	var data []uint8
	for i := uint32(0); i < count; i++ {
		value := (float64(result[i]) - shift) * slope
		if value < 0 {
			value = 0
		} else if value > 255 {
			value = 255
		}

		data = append(data, uint8(value))
		//	data = append(data, uint8(result[i]))
	}
	binary.Write(buf, binary.LittleEndian, data)
	f.Write(buf.Bytes())

	return nil
}

/*
func (image DcmImage) convertTo8Bit() []uint8 {
	var result []uint8
	count := image.Rows * image.Columns

	if image.HighBit < 15 {
		var nMask int32
		if image.PixelRepresentation == 0 {
			nMask = 0xffff << (image.HighBit + 1)

			for i := uint32(0); i < count; i++ {
				var p int32
				buf := bytes.NewBuffer(image.PixelData[2*i : 2*i+1])
				binary.Read(buf, binary.LittleEndian, &p)
				p &= ^nMask
				result = append(result, uint8(p))
			}

		} else {
			nSignBit := int32(1 << image.HighBit)
			nMask = 0xffff << (image.HighBit + 1)

			for i := uint32(0); i < count; i++ {
				var p int32
				buf := bytes.NewBuffer(image.PixelData[2*i : 2*i+1])
				binary.Read(buf, binary.LittleEndian, &p)
				if (p & nSignBit) != 0 {
					p |= nMask
				} else {
					p &= ^nMask
				}
				result = append(result, uint8(p))
			}

		}

	}

	if (image.RescaleSlope != 1.0) || (image.RescaleIntercept != 0.0) {
		for i := uint32(0); i < count; i++ {
			var p int32
			buf := bytes.NewBuffer(image.PixelData[2*i : 2*i+1])
			binary.Read(buf, binary.LittleEndian, &p)
			fvalue := float64(p)*image.RescaleSlope + image.RescaleIntercept

			result = append(result, uint8(fvalue))
		}
	}

	if (image.WindowCenter != 0.0) || (image.WindowWidth != 0.0) {
		shift := image.WindowCenter - image.WindowWidth/2.0
		slope := 255.0 / image.WindowWidth
		for i := uint32(0); i < count; i++ {
			var p int32
			buf := bytes.NewBuffer(image.PixelData[2*i : 2*i+1])
			binary.Read(buf, binary.LittleEndian, &p)
			value := (float64(p) - shift) * slope
			if value < 0 {
				value = 0
			} else if value > 255 {
				value = 255
			}
			result = append(result, uint8(value))
		}
	} else {
		min, max := image.findPixelExtremeValue()
		var slope float64
		if min != max {
			slope = 255.0 / float64(max-min)
		} else {
			slope = 1.0
		}
		for i := uint32(0); i < count; i++ {
			var p int32
			buf := bytes.NewBuffer(image.PixelData[2*i : 2*i+1])
			binary.Read(buf, binary.LittleEndian, &p)
			value := float64(p-min) * slope
			if value < 0 {
				value = 0
			} else if value > 255 {
				value = 255
			}
			result = append(result, uint8(value))
		}
	}
	return result
}

func (image DcmImage) findPixelExtremeValue() (int32, int32) {
	var min, max int32
	count := image.Columns * image.Rows
	for i := uint32(0); i < count; i++ {
		var p int32
		buf := bytes.NewBuffer(image.PixelData[2*i : 2*i+1])
		binary.Read(buf, binary.LittleEndian, &p)
		if i == 0 {
			min = p
			max = p
		}
		if p < min {
			min = p
		}
		if p > max {
			max = p
		}
	}
	return min, max
}
*/
