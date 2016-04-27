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
	err = elem.ReadDcmTagElemment(stream)
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

	err = elem.ReadValue(stream, true, false)
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
