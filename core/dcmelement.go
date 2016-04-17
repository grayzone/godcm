package core

// DcmElement indentified the data element tag.
type DcmElement struct {
	Tag    DcmTag
	VR     string
	Length int64
	Value  string
	IsMeta bool
}
