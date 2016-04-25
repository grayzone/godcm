package main

import (
	"log"

	"github.com/grayzone/godcm/core"
)

func readdicmfile(filename string) {
	var reader core.DcmReader
	reader.IsReadValue = true
	err := reader.ReadFile(filename)
	if err != nil {
		log.Println(err.Error())
	}
}

// ReadDicomWithImplicitVR read the dicom file with implicit VR for testing.
func ReadDicomWithImplicitVR() {
	readdicmfile("./test/data/CT-MONO2-16-ankle")
}

// ReadDicomWithExplicitVR read the dicom file with explicit VR for testing.
func ReadDicomWithExplicitVR() {
	readdicmfile("./test/data/GH178.dcm")
}

func main() {
	//	ReadDicomWithImplicitVR()
	ReadDicomWithExplicitVR()
}
