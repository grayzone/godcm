package core

import "os"

// DcmMetaInfo is to store DICOM meta data.
type DcmMetaInfo struct {
	Preamble []byte // length: 128
	Prefix   []byte // length: 4
	Elements []DcmElement
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
		var elem DcmElement
		elem.Tag.Group = group
		elem.Tag.Element = element
		elem.VR = vr
		elem.Length = length
		elem.Value = value
		elem.IsMeta = true
		meta.Elements = append(meta.Elements, elem)
	}
	return nil
}
