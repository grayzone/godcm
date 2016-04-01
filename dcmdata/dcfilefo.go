package dcmdata

type DcmFileFormat struct {
	DcmSequenceOfItems
	fileReadMode E_FileReadMode
}

func (fo *DcmFileFormat) SetReadMode(readMode E_FileReadMode) {
	fo.fileReadMode = readMode
}

func (fo *DcmFileFormat) GetReadMode() E_FileReadMode {
	return fo.fileReadMode
}

func (fo *DcmFileFormat) Read(instream []byte) {
}

func (fo *DcmFileFormat) GetDataset() DcmDataset {
	var ds DcmDataset
	return ds

}

func (fo *DcmFileFormat) LoadFile(filename string, readXfer E_TransferSyntax, groupLength E_GrpLenEncoding, maxReadLength uint32, readMode E_FileReadMode) {
	if readMode == ERM_dataset {
		//	fo.GetDataset().

	}
}
