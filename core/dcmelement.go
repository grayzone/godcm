package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

// DcmElement indentified the data element tag.
type DcmElement struct {
	Tag    DcmTag
	Name   string
	VR     string
	Length int64
	Value  []byte
}

// GetValueString convert value to string according to VR
func (e DcmElement) GetValueString() string {
	buf := bytes.NewBuffer(e.Value)
	var result string
	switch e.VR {
	case "FL", "FD", "OD", "OF":
		var f float64
		binary.Read(buf, binary.LittleEndian, &f)
		result = fmt.Sprintf("%f", f)
	case "", "OL", "SL", "SS", "UL":
		var i int32
		binary.Read(buf, binary.LittleEndian, &i)
		result = fmt.Sprintf("%d", i)
	case "US", "US or SS":
		var i uint16
		binary.Read(buf, binary.LittleEndian, &i)
		result = fmt.Sprintf("%d", i)
	case "AE", "AS", "CS", "DA", "DS", "DT", "IS", "LO", "LT", "PN", "ST", "UI", "UT", "TM", "SH":
		result = string(bytes.Trim(e.Value, "\x00"))
	default:
		result = fmt.Sprintf("%x", e.Value)
	}
	return result
}

// String convert to string value
func (e DcmElement) String() string {
	return fmt.Sprintf("Tag:%s; VR:%s; Length:%d; Value:%s", e.Tag, e.VR, e.Length, e.GetValueString())
}

// ReadDcmTag is to read group and element
func (e *DcmElement) ReadDcmTag(s *DcmFileStream) error {
	var err error
	e.Tag.Group, err = s.ReadUINT16()
	if err != nil {
		return err
	}
	e.Tag.Element, err = s.ReadUINT16()
	if err != nil {
		return err
	}
	return nil
}

// ReadDcmVR is to read vr
func (e *DcmElement) ReadDcmVR(s *DcmFileStream) error {
	var err error
	e.VR, err = s.ReadString(2)
	return err
}

// ReadValueLengthWithExplicitVR gets the value length of the dicom element with explicit VR.
func (e *DcmElement) ReadValueLengthWithExplicitVR(s *DcmFileStream) error {
	switch e.VR {
	case "OB", "OD", "OF", "OL", "OW", "SQ", "UC", "UR", "UT", "UN":
		// skip the reserved 2 bytes
		_, err := s.Skip(2)
		if err != nil {
			return err
		}
		err = e.ReadValueLengthUint32(s)
		if err != nil {
			return err
		}
	default:
		// read value length
		err := e.ReadValueLengthUint16(s)
		if err != nil {
			return err
		}
	}
	return nil
}

// ReadValueLengthWithImplicitVR gets the value length of the dicom element with implicit VR.
func (e *DcmElement) ReadValueLengthWithImplicitVR(s *DcmFileStream) error {
	return e.ReadValueLengthUint32(s)
}

// ReadValueLengthUint16 read 2 bytes value length
func (e *DcmElement) ReadValueLengthUint16(s *DcmFileStream) error {
	l, err := s.ReadUINT16()
	if err != nil {
		return err
	}
	e.Length = int64(l)
	return nil
}

// ReadValueLengthUint32 read 4 bytes value length
func (e *DcmElement) ReadValueLengthUint32(s *DcmFileStream) error {
	l, err := s.ReadUINT32()
	if err != nil {
		return err
	}
	e.Length = int64(l)
	return nil

}

// ReadValue get or skip the element value.
func (e *DcmElement) ReadValue(s *DcmFileStream, isReadValue bool, isReadPixel bool) error {
	var err error
	if !isReadPixel {
		if e.Tag.Group == 0x7fe0 {
			_, err = s.Skip(e.Length)
			return err
		}
	}
	if isReadValue {
		// read element value
		e.Value, err = s.Read(e.Length)
	} else {
		_, err = s.Skip(e.Length)
	}
	return err
}

// ReadDcmElement read one dicom element.
func (e *DcmElement) ReadDcmElement(s *DcmFileStream, isExplicitVR bool, isReadValue bool, isReadPixel bool) error {
	if isExplicitVR {
		return e.ReadDcmElementWithExplicitVR(s, isReadValue, isReadPixel)
	}
	return e.ReadDcmElementWithImplicitVR(s, isReadValue, isReadPixel)
}

// ReadDcmElementWithExplicitVR read the data element with explicit VR.
func (e *DcmElement) ReadDcmElementWithExplicitVR(s *DcmFileStream, isReadValue bool, isReadPixel bool) error {
	// read dicom tag
	err := e.ReadDcmTag(s)
	if err != nil {
		return err
	}

	// read VR
	err = e.ReadDcmVR(s)
	if err != nil {
		return err
	}

	//read the value length
	err = e.ReadValueLengthWithExplicitVR(s)
	if err != nil {
		return err
	}

	err = e.ReadValue(s, isReadValue, isReadPixel)
	if err != nil {
		return err
	}

	log.Println(e.String())

	return nil
}

// ReadDcmElementWithImplicitVR read the data element with implicit VR.
func (e *DcmElement) ReadDcmElementWithImplicitVR(s *DcmFileStream, isReadValue bool, isReadPixel bool) error {

	// read dciom tag
	err := e.ReadDcmTag(s)
	if err != nil {
		return err
	}
	// get VR from Dicom Element registry
	err = FindDcmElmentByTag(e)
	if err != nil {
		log.Println(err.Error())
	}

	// read the value length
	err = e.ReadValueLengthWithImplicitVR(s)
	if err != nil {
		return err
	}

	err = e.ReadValue(s, isReadValue, isReadPixel)
	if err != nil {
		return err
	}

	log.Println(e)

	/*
		if elem.Tag.Group != 0x7fe0 {
					log.Println(elem)
		}
	*/
	return nil

}
