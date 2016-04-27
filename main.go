package main

import (
	"log"
	"os"

	"strconv"

	"github.com/grayzone/godcm/core"
)

func readdicmfile(filename string, isReadValue bool) {
	var reader core.DcmReader
	reader.IsReadValue = isReadValue
	err := reader.ReadFile(filename)
	if err != nil {
		log.Println(err.Error())
	}
}

var testfile = []string{
	"./test/data/CT-MONO2-16-ankle",
	"./test/data/GH195.dcm",
	"./test/data/GH064.dcm",
}

func main() {
	var index int
	var isReadValue bool
	switch len(os.Args) {
	case 1:
		readdicmfile(testfile[1], true)
	case 2:
		index, _ = strconv.Atoi(os.Args[1])
		readdicmfile(testfile[index], isReadValue)
	case 3:
		index, _ = strconv.Atoi(os.Args[1])
		isReadValue, _ = strconv.ParseBool(os.Args[2])
		readdicmfile(testfile[index], isReadValue)
	}
}
