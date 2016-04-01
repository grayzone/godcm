package dcmdata

type DcmMetaInfo struct {
	DcmItem

	filePreamble           string           // buffer for 132 byte DICOM file preamble
	preambleUsed           bool             // true if the preamble was read from stream
	fPreambleTransferState E_TransferState  // transfer state of the preamble
	xfer                   E_TransferSyntax // transfer syntax in which the meta-header was read
}

func (mi *DcmMetaInfo) Read(instream []byte) {
}
