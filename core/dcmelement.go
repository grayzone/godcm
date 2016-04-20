package core

// DcmElement indentified the data element tag.
type DcmElement struct {
	Tag    DcmTag
	Name   string
	VR     string
	Length int64
	Value  string
}

// String convert to string value
/*
func (e DcmElement) String() string {
	return fmt.Sprintf("%s;%s;%d;%s", e.Tag, e.VR, e.Length, e.Value)
}
*/
