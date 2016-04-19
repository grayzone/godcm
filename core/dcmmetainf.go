package core

import (
	"log"
	"os"
)

// DcmMetaInfo is to store DICOM meta data.
type DcmMetaInfo struct {
	Preamble []byte // length: 128
	Prefix   []byte // length: 4
	Elements []DcmElement
}

/*
// String convert to string value
func (meta DcmMetaInfo) String() string {
	var result string
	result += fmt.Sprintf("%x", meta.Preamble) + ";"
	result += string(meta.Prefix) + ";"
	for _, v := range meta.Elements {
		result += v.String()
	}
	return result
}
*/

// NewDcmMetaInfo to initialize the struct with all the tags used for dicom file meta information
func NewDcmMetaInfo() *DcmMetaInfo {
	var meta DcmMetaInfo

	meta.Elements = append(meta.Elements, DcmElement{Tag: FileMetaInformationGroupLength})
	meta.Elements = append(meta.Elements, DcmElement{Tag: FileMetaInformationVersion})
	meta.Elements = append(meta.Elements, DcmElement{Tag: MediaStorageSOPClassUID})
	meta.Elements = append(meta.Elements, DcmElement{Tag: MediaStorageSOPInstanceUID})
	meta.Elements = append(meta.Elements, DcmElement{Tag: TransferSyntaxUID})
	meta.Elements = append(meta.Elements, DcmElement{Tag: ImplementationClassUID})
	meta.Elements = append(meta.Elements, DcmElement{Tag: ImplementationVersionName})
	meta.Elements = append(meta.Elements, DcmElement{Tag: SourceApplicationEntityTitle})
	meta.Elements = append(meta.Elements, DcmElement{Tag: SendingApplicationEntityTitle})
	meta.Elements = append(meta.Elements, DcmElement{Tag: ReceivingApplicationEntityTitle})
	meta.Elements = append(meta.Elements, DcmElement{Tag: PrivateInformationCreatorUID})
	meta.Elements = append(meta.Elements, DcmElement{Tag: PrivateInformation})

	return &meta
}

// FindDcmElement find the element by tag
func (meta *DcmMetaInfo) FindDcmElement(tag DcmTag) (*DcmElement, error) {
	for i, v := range meta.Elements {
		if v.Tag == tag {
			return &meta.Elements[i], nil
		}
	}
	return nil, nil
}

// GetTransferSyntaxUID return the transfer syntax string of the DICOM file.
func (meta *DcmMetaInfo) GetTransferSyntaxUID() (string, error) {
	elem, err := meta.FindDcmElement(TransferSyntaxUID)
	if err != nil {
		return "", err
	}
	return elem.Value, nil
}

// Read meta information from file stream
func (meta *DcmMetaInfo) Read(stream *DcmFileStream) error {
	// turn to the beginning of the file
	_, err := stream.FileHandler.Seek(0, os.SEEK_SET)
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

	// read dicom meta datasets
	for !stream.Eos() {
		group, err := stream.ReadUINT16()
		if err != nil {
			return err
		}
		if group != 0x0002 {
			err = stream.Putback(2)
			return err
		}
		element, err := stream.ReadUINT16()
		if err != nil {
			return err
		}
		vr, err := stream.ReadDcmVR()
		if err != nil {
			return err
		}
		length, err := stream.ReadDcmElementValueLength(vr)
		if err != nil {
			return err
		}
		value, err := stream.ReadString(length)
		if err != nil {
			return err
		}
		elem, err := meta.FindDcmElement(DcmTag{group, element})
		if err != nil {
			log.Println(err.Error)
			continue
		}
		elem.VR = vr
		elem.Length = length
		elem.Value = value
	}

	return nil
}

// CheckTransferSyntaxUID is to identify the Transfer Syntax used to encode the following data set.
func (meta *DcmMetaInfo) CheckTransferSyntaxUID() error {
	uid, err := meta.GetTransferSyntaxUID()
	if err != nil {
		return err
	}
	log.Println(uid)

	return nil

}
