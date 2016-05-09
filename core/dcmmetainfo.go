package core

import (
	"errors"
	_ "log"
)

// DcmMetaInfo is to store DICOM meta data.
type DcmMetaInfo struct {
	Preamble        []byte // length: 128
	Prefix          []byte // length: 4
	Elements        []DcmElement
	isEndofMetaInfo bool
}

// FindDcmElement find the element by tag
func (meta DcmMetaInfo) FindDcmElement(elem *DcmElement) error {
	for _, v := range meta.Elements {
		if v.Tag == elem.Tag {
			*elem = v
			return nil
		}
	}
	err := "Not find the tag '" + elem.Tag.String() + "' in meta information."
	return errors.New(err)
}

// GetTransferSyntaxUID return the transfer syntax string of the DICOM file.
func (meta DcmMetaInfo) GetTransferSyntaxUID() (string, error) {
	var elem DcmElement
	elem.Tag = DCMTransferSyntaxUID
	err := meta.FindDcmElement(&elem)
	if err != nil {
		return "", err
	}
	return elem.GetValueString(), nil
}

// ReadOneElement read one DICOM element in meta information.
func (meta *DcmMetaInfo) ReadOneElement(stream *DcmFileStream) error {
	var elem DcmElement
	elem.isReadValue = true
	var err error
	err = elem.ReadDcmTagGroup(stream)
	if err != nil {
		return err
	}
	if elem.Tag.Group != 0x0002 {
		stream.Putback(2)
		meta.isEndofMetaInfo = true
		return nil
	}
	err = elem.ReadDcmTagElement(stream)
	if err != nil {
		return err
	}

	err = elem.ReadDcmVR(stream)
	if err != nil {
		return err
	}

	err = elem.ReadValueLengthWithExplicitVR(stream)
	if err != nil {
		return err
	}

	err = elem.ReadValue(stream)
	if err != nil {
		return err
	}
	//	log.Println(elem)
	meta.Elements = append(meta.Elements, elem)
	return nil
}

// Read meta information from file stream
func (meta *DcmMetaInfo) Read(stream *DcmFileStream) error {
	// turn to the beginning of the file
	err := stream.SeekToBegin()
	if err != nil {
		return err
	}
	// read the preamble
	meta.Preamble, err = stream.Read(128)
	if err != nil {
		return err
	}
	//read the prefix
	meta.Prefix, err = stream.Read(4)
	if err != nil {
		return err
	}
	// read dicom meta datasets
	for !stream.Eos() && !meta.isEndofMetaInfo {
		err = meta.ReadOneElement(stream)
		if err != nil {
			return err
		}
	}
	return nil
}

// IsExplicitVR is to check if the tag is Explicit VR structure
func (meta DcmMetaInfo) IsExplicitVR() (bool, error) {
	uid, err := meta.GetTransferSyntaxUID()
	if err != nil {
		return false, err
	}
	var xfer DcmXfer
	xfer.XferID = uid
	err = xfer.GetDcmXferByID()
	if err != nil {
		return false, err
	}
	return xfer.IsExplicitVR(), nil
}

// GetByteOrder get the byte orber of the file
func (meta DcmMetaInfo) GetByteOrder() (EByteOrder, error) {
	uid, err := meta.GetTransferSyntaxUID()
	if err != nil {
		return EBOunknown, err
	}
	var xfer DcmXfer
	xfer.XferID = uid
	err = xfer.GetDcmXferByID()
	if err != nil {
		return EBOunknown, err
	}
	return xfer.ByteOrder, nil
}

// FindElement get the element info by given tag
func (meta DcmMetaInfo) FindElement(e *DcmElement) error {
	for _, v := range meta.Elements {
		if e.Tag == v.Tag {
			*e = v
			return nil
		}
	}
	str := "not find the tag '" + e.Tag.String() + "' in the data set"
	return errors.New(str)
}

func (meta DcmMetaInfo) getElementValue(tag DcmTag) string {
	var elem DcmElement
	elem.Tag = tag
	err := meta.FindElement(&elem)
	if err != nil {
		return ""
	}
	return elem.GetValueString()
}

// FileMetaInformationGroupLength gets meta information group length
func (meta DcmMetaInfo) FileMetaInformationGroupLength() string {
	return meta.getElementValue(DCMFileMetaInformationGroupLength)
}

// FileMetaInformationVersion gets meta information version
func (meta DcmMetaInfo) FileMetaInformationVersion() string {
	return meta.getElementValue(DCMFileMetaInformationVersion)
}

// MediaStorageSOPClassUID gets media storage SOP Class UID
func (meta DcmMetaInfo) MediaStorageSOPClassUID() string {
	return meta.getElementValue(DCMMediaStorageSOPClassUID)
}

// MediaStorageSOPInstanceUID gets media storage SOP Instance UID
func (meta DcmMetaInfo) MediaStorageSOPInstanceUID() string {
	return meta.getElementValue(DCMMediaStorageSOPInstanceUID)
}

// TransferSyntaxUID gets Transfer Syntax UID
func (meta DcmMetaInfo) TransferSyntaxUID() string {
	return meta.getElementValue(DCMTransferSyntaxUID)
}

// ImplementationClassUID gets Implementation Class UID
func (meta DcmMetaInfo) ImplementationClassUID() string {
	return meta.getElementValue(DCMImplementationClassUID)
}

// ImplementationVersionName gets implementation version name
func (meta DcmMetaInfo) ImplementationVersionName() string {
	return meta.getElementValue(DCMImplementationVersionName)
}

// SourceApplicationEntityTitle gets source application entity title
func (meta DcmMetaInfo) SourceApplicationEntityTitle() string {
	return meta.getElementValue(DCMSourceApplicationEntityTitle)
}

// SendingApplicationEntityTitle gets sending application entity title
func (meta DcmMetaInfo) SendingApplicationEntityTitle() string {
	return meta.getElementValue(DCMSendingApplicationEntityTitle)
}

// ReceivingApplicationEntityTitle gets receiving application entity title
func (meta DcmMetaInfo) ReceivingApplicationEntityTitle() string {
	return meta.getElementValue(DCMReceivingApplicationEntityTitle)
}

// PrivateInformationCreatorUID gets private information createor UID
func (meta DcmMetaInfo) PrivateInformationCreatorUID() string {
	return meta.getElementValue(DCMPrivateInformationCreatorUID)
}

// PrivateInformation gets private information
func (meta DcmMetaInfo) PrivateInformation() string {
	return meta.getElementValue(DCMPrivateInformation)
}
