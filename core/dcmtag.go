package core

import "fmt"

// DcmTag contains the group and element of the dicom element.
type DcmTag struct {
	Group   uint16
	Element uint16
}

// String is to convert tag to string
func (t DcmTag) String() string {
	return fmt.Sprintf("0x%04x%04x", t.Group, t.Element)
}
