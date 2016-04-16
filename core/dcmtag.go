package core

import "fmt"

// DcmTag contains the group and element of the dicom element.
type DcmTag struct {
	Group   uint16
	Element uint16
}

// ToString is to convert tag to string
func (t *DcmTag) ToString() string {
	return fmt.Sprintf("0x%x%x", t.Group, t.Element)
}
