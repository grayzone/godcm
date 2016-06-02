package dcmimage

import (
	"encoding/binary"
	"log" // for debug
	"strings"
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

	PixelData []byte
	Data      []uint8
}

/*
func maxval(bits uint16, pos uint32) uint32 {
	if bits > MAXBITS {
		return uint32(1<<bits - 1)
	}
	return uint32(1<<bits - pos)
}
*/

func (di DcmImage) clipHighBits(pixel int16) int16 {
	if di.HighBit > 15 {
		return pixel
	}
	/*
		if di.BitsAllocated == di.BitsStored {
			return pixel
		}
	*/
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

	return di.getRescaling(pixel).(int16)
}

func (di DcmImage) getRescaling(value interface{}) interface{} {
	switch value.(type) {
	case int16:
		return int16(float64(value.(int16))*di.RescaleSlope + di.RescaleIntercept)
	case float64:
		return value.(float64)*di.RescaleSlope + di.RescaleIntercept
	}
	return 0
}

/*
func (di *DcmImage) checkRescaling() {

	if di.RescaleSlope < 0 {
		var tmp interface{}
		tmp = di.minValue
		di.minValue = di.getRescaling(di.maxValue).(int16)
		di.maxValue = di.getRescaling(tmp).(int16)

		tmp = di.AbsMinimum
		di.AbsMinimum = di.getRescaling(di.AbsMaximum).(float64)
		di.AbsMaximum = di.getRescaling(tmp).(float64)
		return
	}
	di.minValue = di.getRescaling(di.minValue).(int16)
	di.maxValue = di.getRescaling(di.maxValue).(int16)

	di.AbsMinimum = di.getRescaling(di.AbsMinimum).(float64)
	di.AbsMaximum = di.getRescaling(di.AbsMaximum).(float64)
}
*/
func (di DcmImage) nowindow(pixel int16) uint8 {
	//	var outrange float64
	//	outrange = di.high - di.low + 1

	//	ocnt := di.getAbsMaxRange()
	gradient := 255.0 / float64(di.maxValue-di.minValue)
	value := float64(pixel-di.minValue) * gradient
	if value > 255 {
		return 255
	}
	if value < 0 {
		return 0
	}
	return uint8(value)
}

func (di DcmImage) window(pixel int16) uint8 {
	min := di.WindowCenter - di.WindowWidth/2.0 + 0.5
	max := di.WindowCenter + di.WindowWidth/2.0 + 0.5
	slope := 255.0 / di.WindowWidth
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

func (di DcmImage) rescaleWindowLevel(pixel int16) uint8 {
	if (di.WindowCenter == 0.0) && (di.WindowWidth == 0.0) {
		return di.nowindow(pixel)
	}
	return di.window(pixel)
}

func (di *DcmImage) byteTouint8() {
	di.Data = nil
	for i := range di.PixelData {
		b := uint8(di.PixelData[i])
		if di.IsReverse {
			b = 255 - b
		}
		di.Data = append(di.Data, b)
	}
}

func (di *DcmImage) int16Touint8() {
	di.Data = nil
	count := di.Columns * di.Rows

	for i := uint32(0); i < count; i++ {
		b := di.PixelData[2*i : 2*i+2]
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
			p = 255 - p
		}
		di.Data = append(di.Data, p)
	}
}

func (di *DcmImage) convertTo8Bit() {
	di.determinReverse()
	if di.BitsAllocated <= 8 {
		di.byteTouint8()
		return
	}
	di.determineMinMax()
	di.int16Touint8()

	/*

		var result []uint8

		rgbplane := bits / 8

		gap := rgbplane * uint16((4-(di.Columns&0x3))&0x3)
		for i := di.Rows; i > uint32(0); i-- {
			for j := uint32(0); j < di.Columns; j++ {
				var pixel int16

				if di.BitsAllocated > 8 {
					b := di.PixelData[2*di.Columns*i-2*di.Columns+2*j : 2*di.Columns*i-2*di.Columns+2*j+2]
					if di.IsBigEndian {
						pixel = int16(binary.BigEndian.Uint16(b))
					} else {
						pixel = int16(binary.LittleEndian.Uint16(b))
					}
				} else {
					pixel = int16(di.PixelData[di.Columns*i-di.Columns+j])
				}

				pixel = di.clipHighBits(pixel)
				pixel = di.rescalePixel(pixel)

				b := di.rescaleWindowLevel(pixel)
				//b := di.nowindow(pixel)
				if di.IsReverse {
					b = 255 - b
				}

				for i := uint16(0); i < rgbplane; i++ {
					result = append(result, uint8(b))
				}

			}
			if bits != 32 {
				for i := uint16(0); i < gap; i++ {
					result = append(result, uint8(0))
				}
			}
		}
		return result

	*/
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

func (di *DcmImage) determinReverse() {
	if strings.ToUpper(di.PhotometricInterpretation) == "MONOCHROME1" {
		di.IsReverse = true
	}
}

func (di *DcmImage) determineMinMax() {
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

	log.Println("min", di.minValue, "max", di.maxValue)
	//	di.checkRescaling()
	//	log.Println("min2", di.minValue, "max2", di.maxValue)
}
