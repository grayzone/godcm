package main

import (
	"log"
	"os"

	"strconv"

	"github.com/grayzone/godcm/core"
	"github.com/grayzone/godcm/image"
)

func readdicmfile(filename string, isReadValue bool) {
	var reader core.DcmReader
	reader.IsReadValue = isReadValue
	err := reader.ReadFile(filename)
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

func convert2bmp(filename string) {
	var reader core.DcmReader
	reader.IsReadPixel = true
	reader.IsReadValue = true
	err := reader.ReadFile(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	pixeldata := reader.Dataset.PixelData()

	var img image.DcmImage

	var num uint64

	num, _ = strconv.ParseUint(reader.Dataset.BitsAllocated(), 10, 16)
	img.BitsAllocated = uint16(num)

	num, _ = strconv.ParseUint(reader.Dataset.Columns(), 10, 32)
	img.Columns = uint32(num)

	num, _ = strconv.ParseUint(reader.Dataset.Rows(), 10, 32)
	img.Rows = uint32(num)

	num, _ = strconv.ParseUint(reader.Dataset.Rows(), 10, 16)
	img.HighBit = uint16(num)

	img.PixelData = pixeldata

	err = img.WriteBMP("test.bmp", 8, 0)
	if err != nil {
		log.Println(err.Error())
	}
}

var testfile = []string{
	"./test/data/IM0.dcm",
	"./test/data/image_09-12-2013_4.dcm",
	"./test/data/IM-0001-0010.dcm",
	"./test/data/CT-MONO2-16-ankle",
	"./test/data/GH195.dcm",
	"./test/data/GH064.dcm",
	"./test/data/GH177_D_CLUNIE_CT1_IVRLE_BigEndian_undefined_length.dcm",
	"./test/data/GH177_D_CLUNIE_CT1_IVRLE_BigEndian_ELE_undefinded_length.dcm",
	"./test/data/MR-MONO2-8-16x-heart.dcm",
	"./test/data/xr_chicken2.dcm",
	"./test/data/T23/IM-0001-0001.dcm",
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
	convert2bmp(testfile[0])
}

func main() {
	//	testParseDcm()
	testdcm2bmp()
}
