package dcmdata

var META_HEADER_DEFAULT_TRANSFERSYNTAX = EXS_LittleEndianExplicit

type DcmMetaInfo struct {
	DcmItem

	filePreamble           string           // buffer for 132 byte DICOM file preamble
	preambleUsed           bool             // true if the preamble was read from stream
	fPreambleTransferState E_TransferState  // transfer state of the preamble
	xfer                   E_TransferSyntax // transfer syntax in which the meta-header was read
}

func NewDcmMetaInfo() *DcmMetaInfo {
	var dcmitem DcmItem
	dcmitem.DcmObject = *NewDcmObject(DCM_ItemTag, 0)
	mi := DcmMetaInfo{DcmItem: dcmitem, preambleUsed: false, fPreambleTransferState: ERW_init, xfer: META_HEADER_DEFAULT_TRANSFERSYNTAX}

	return &mi
}

func (mi *DcmMetaInfo) Read(instream []byte) {
}
