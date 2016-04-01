package dcmdata

type DcmFileFormat struct {
	DcmSequenceOfItems
	fileReadMode E_FileReadMode
}

func (fo *DcmFileFormat) SetReadMode(readMode E_FileReadMode) {
	fo.fileReadMode = readMode
}

func (fo *DcmFileFormat) GetReadMode() {
	return fo.fileReadMode
}

func (fo *DcmFileFormat) Read(instream []byte) {
}
