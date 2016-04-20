package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
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
	fmt.Println(e.VR)
	buf := bytes.NewBuffer(e.Value)
	var result string
	switch e.VR {
	case "FL", "FD", "OD", "OF":
		var f float64
		binary.Read(buf, binary.LittleEndian, &f)
		result = fmt.Sprintf("%f", f)
	case "OL", "SH", "SL", "SS", "UL", "US":
		var i int
		binary.Read(buf, binary.LittleEndian, &i)
		result = fmt.Sprintf("%d", i)
	case "AE", "AS", "CS", "DA", "DS", "DT", "IS", "LO", "LT", "PN", "ST", "UI", "UT":
		binary.Read(buf, binary.LittleEndian, &result)
	default:
		result = fmt.Sprintf("%x", e.Value)
	}
	return result
}

// String convert to string value
func (e DcmElement) String() string {
	return fmt.Sprintf("Tag:%s; VR:%s; Length:%d; Value:%s", e.Tag, e.VR, e.Length, e.GetValueString())
}
