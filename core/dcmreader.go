package core

import (
	"errors"
	"log"
	"strconv"

	"github.com/grayzone/godcm/dcmimage"
)

// DICOM3FILEIDENTIFIER is the DiCOM index in the file header.
const DICOM3FILEIDENTIFIER = "DICM"

// DcmReader is to read DICOM file
type DcmReader struct {
	fs          DcmFileStream
	Meta        DcmMetaInfo
	Dataset     DcmDataset
	IsReadValue bool
	IsReadPixel bool
}

// ReadFile is to read dicom file.
func (reader *DcmReader) ReadFile(filename string) error {
	reader.fs.FileName = filename
	err := reader.fs.Open()
	if err != nil {
		return err
	}
	defer reader.fs.Close()
	isDCM3, err := reader.IsDicom3()
	if !isDCM3 {
		return err
	}

	//read dicom file meta information
	err = reader.Meta.Read(&reader.fs)
	if err != nil {
		return err
	}

	isExplicitVR, err := reader.Meta.IsExplicitVR()
	if err != nil {
		return err
	}

	byteOrder, err := reader.Meta.GetByteOrder()
	if err != nil {
		return err
	}
	// read dicom dataset
	err = reader.Dataset.Read(&reader.fs, isExplicitVR, byteOrder, reader.IsReadValue, reader.IsReadPixel)
	if err != nil {
		return err
	}

	return nil
}

func (reader *DcmReader) GetImageInfo() dcmimage.DcmImage {
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

func (reader DcmReader) Convert2PNG(filepath string) error {
	img := reader.GetImageInfo()
	frame := img.NumberOfFrames
	for i := 0; i < frame; i++ {
		var newfile string
		if frame == 1 {
			newfile = filepath + ".png"
		} else {
			newfile = filepath + "_" + strconv.FormatUint(uint64(i), 10) + ".png"
		}
		err := img.ConvertToPNG(newfile, i)
		if err != nil {
			log.Println(err.Error())
		}
	}
	return nil
}

// IsDicom3 is to check the file is supported by DICOM 3.0 or not.
func (reader DcmReader) IsDicom3() (bool, error) {
	_, err := reader.fs.Skip(128)
	if err != nil {
		return false, err
	}
	b, err := reader.fs.Read(int64(len(DICOM3FILEIDENTIFIER)))
	if err != nil {
		return false, err
	}
	if string(b) != DICOM3FILEIDENTIFIER {
		return false, errors.New("Only supprot DICOM 3.0.")
	}
	reader.fs.Putback(132)
	return true, nil
}

// IsCompressed check whether pixel data only exist in compressed format
func (reader DcmReader) IsCompressed() (bool, error) {
	var xfer DcmXfer
	xfer.XferID = reader.Meta.TransferSyntaxUID()
	err := xfer.GetDcmXferByID()
	if err != nil {
		return false, err
	}
	if xfer.IsCompressed() {
		log.Println("XferName: ", xfer.XferName)
	}
	return xfer.IsCompressed(), nil
}

// IsBigEndian check whether pixel data need to be swapped
func (reader DcmReader) IsBigEndian() (bool, error) {
	var xfer DcmXfer
	xfer.XferID = reader.Meta.TransferSyntaxUID()
	err := xfer.GetDcmXferByID()
	if err != nil {
		return false, err
	}
	/*
		if xfer.IsBigEndian() {
			log.Println("ByteOrder: ", xfer.ByteOrder.String())
		}
	*/
	return xfer.IsBigEndian(), nil
}
