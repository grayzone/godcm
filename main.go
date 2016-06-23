package main

import (
	"log"
	"os"

	"strconv"

	"github.com/grayzone/godcm/core"
	"github.com/grayzone/godcm/dcmimage"
)

var folder = "./test/data/"

func readdicmfile(filename string, isReadValue bool) {
	var reader core.DcmReader
	reader.IsReadValue = isReadValue
	err := reader.ReadFile(folder + filename)
	if err != nil {
		log.Println(err.Error())
	}

	for _, v := range reader.Meta.Elements {
		log.Println(v.String())
	}
	for i := range reader.Dataset.Elements {
		log.Println(reader.Dataset.Elements[i].String())
	}

}

func getimageinfo(filename string) dcmimage.DcmImage {
	var reader core.DcmReader
	reader.IsReadPixel = true
	reader.IsReadValue = true
	err := reader.ReadFile(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	isCompressed, err := reader.IsCompressed()
	if err != nil {
		log.Println(err.Error())
	}

	isBigEndian, err := reader.IsBigEndian()
	if err != nil {
		log.Println(err.Error())
	}

	pixeldata := reader.Dataset.PixelData()

	var img dcmimage.DcmImage

	img.IsCompressed = isCompressed
	img.IsBigEndian = isBigEndian

	var num interface{}

	num, _ = strconv.ParseUint(reader.Dataset.BitsAllocated(), 10, 16)
	img.BitsAllocated = uint16(num.(uint64))

	num, _ = strconv.ParseUint(reader.Dataset.BitsStored(), 10, 16)
	img.BitsStored = uint16(num.(uint64))

	num, _ = strconv.ParseUint(reader.Dataset.Columns(), 10, 32)
	img.Columns = uint32(num.(uint64))

	num, _ = strconv.ParseUint(reader.Dataset.Rows(), 10, 32)
	img.Rows = uint32(num.(uint64))

	num, _ = strconv.ParseUint(reader.Dataset.HighBit(), 10, 16)
	img.HighBit = uint16(num.(uint64))

	num, _ = strconv.ParseFloat(reader.Dataset.WindowCenter(), 64)
	img.WindowCenter = num.(float64)

	num, _ = strconv.ParseFloat(reader.Dataset.WindowWidth(), 64)
	img.WindowWidth = num.(float64)

	num, _ = strconv.ParseFloat(reader.Dataset.RescaleIntercept(), 64)
	img.RescaleIntercept = num.(float64)

	num, _ = strconv.ParseFloat(reader.Dataset.RescaleSlope(), 64)
	img.RescaleSlope = num.(float64)

	num, _ = strconv.ParseUint(reader.Dataset.PixelRepresentation(), 10, 16)
	img.PixelRepresentation = uint16(num.(uint64))

	img.PhotometricInterpretation = reader.Dataset.PhotometricInterpretation()

	num, _ = strconv.ParseUint(reader.Dataset.NumberOfFrames(), 10, 64)
	img.NumberOfFrames = int(num.(uint64))

	num, _ = strconv.ParseUint(reader.Dataset.SamplesPerPixel(), 10, 16)
	img.SamplesPerPixel = uint16(num.(uint64))

	img.PixelData = pixeldata

	return img

}

func convert2bmp(filename string, bits uint16) {
	img := getimageinfo(folder + filename)

	frame := img.NumberOfFrames
	for i := 0; i < frame; i++ {
		newfile := filename + "_" + strconv.FormatUint(uint64(i), 10) + ".bmp"
		err := img.WriteBMP(newfile, bits, i)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func convert2png(filename string) {
	img := getimageinfo(folder + filename)

	frame := img.NumberOfFrames
	for i := 0; i < frame; i++ {
		newfile := filename + "_" + strconv.FormatUint(uint64(i), 10) + ".png"
		err := img.ConvertToPNG(newfile, i)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func convert2jpg(filename string) {
	img := getimageinfo(folder + filename)

	frame := img.NumberOfFrames
	for i := 0; i < frame; i++ {
		newfile := filename + "_" + strconv.FormatUint(uint64(i), 10) + ".jpg"
		err := img.ConvertToJPG(newfile, i)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

var testfile = []string{
	"US-RGB-8-esopecho.dcm",
	"MR-MONO2-8-16x-heart.dcm",
	"xr_chest.dcm",
	"GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm",
	"US-MONO2-8-8x-execho.dcm",
	"xr_tspine.dcm",
	"IM0.dcm",
	"image_09-12-2013_4.dcm",
	"CT-MONO2-16-ankle",
	"xr_chicken2.dcm",
	"T23/IM-0001-0001.dcm",
	"IM-0001-0010.dcm",
	"GH195.dcm",
	"GH064.dcm",
	"GH177_D_CLUNIE_CT1_IVRLE_BigEndian_ELE_undefinded_length.dcm",
}

func testParseDcm() {

	var index int
	var isReadValue bool
	switch len(os.Args) {
	case 1:
		readdicmfile(testfile[0], true)
	case 2:
		index, _ = strconv.Atoi(os.Args[1])
		readdicmfile(testfile[index], isReadValue)
	case 3:
		index, _ = strconv.Atoi(os.Args[1])
		isReadValue, _ = strconv.ParseBool(os.Args[2])
		readdicmfile(testfile[index], isReadValue)
	}
}

func testdcm2bmp() {
	var index int
	switch len(os.Args) {
	case 1:
		convert2bmp(testfile[0], 32)
	case 2:
		index, _ = strconv.Atoi(os.Args[1])
		convert2bmp(testfile[index], 8)
	case 3:
		index, _ = strconv.Atoi(os.Args[1])
		bits, _ := strconv.Atoi(os.Args[2])
		convert2bmp(testfile[index], uint16(bits))
	}
}

func testdcm2png() {
	var index int
	switch len(os.Args) {
	case 1:
		convert2png(testfile[0])
	case 2:
		index, _ = strconv.Atoi(os.Args[1])
		convert2png(testfile[index])
	}
}

func testdcm2jpg() {
	var index int
	switch len(os.Args) {
	case 1:
		convert2jpg(testfile[0])
	case 2:
		index, _ = strconv.Atoi(os.Args[1])
		convert2jpg(testfile[index])
	}
}

/*


func convert2png16(filename string) {
	img := getimageinfo(filename)
	err := img.ConvertToPNG16("test16.png")
	if err != nil {
		log.Println(err.Error())
	}
}


func convert2jpg16(filename string) {
	img := getimageinfo(filename)
	err := img.ConvertToJPG16("test16.jpg")
	if err != nil {
		log.Println(err.Error())
	}
}


func testdcm2png16() {
	var index int
	switch len(os.Args) {
	case 1:
		convert2png16(folder + testfile[0])
	case 2:
		index, _ = strconv.Atoi(os.Args[1])
		convert2png16(folder + testfile[index])
	}
}


func testdcm2jpg16() {
	var index int
	switch len(os.Args) {
	case 1:
		convert2jpg16(folder + testfile[0])
	case 2:
		index, _ = strconv.Atoi(os.Args[1])
		convert2jpg16(folder + testfile[index])
	}
}
*/
func main() {
	//	testParseDcm()
	testdcm2bmp()
	// testdcm2png()
	//	testdcm2jpg()
}
