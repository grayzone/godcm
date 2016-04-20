package main

import (
	"log"
	"os"

	"github.com/grayzone/godcm/core"
)

func ReadDicom() {
	_, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	filename := "./test/data/CT-MONO2-16-ankle"
	var reader core.DcmReader
	err = reader.ReadFile(filename)
	if err != nil {
		log.Println(err.Error())
	}
}

func main() {
	ReadDicom()
}
