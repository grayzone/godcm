package dcmdata

import "fmt"

/// macro for an "undefined" attribute tag that is never used in DICOM
var DCM_UndefinedTagKey = DcmTagKey{Group: 0xffff, Element: 0xffff}

type DcmTagKey struct {
	Group   uint16
	Element uint16
}

func NewDcmTagKey() *DcmTagKey {
	return &DcmTagKey{Group: 0xffff, Element: 0xffff}
}

func (dtk *DcmTagKey) Set(g, e uint16) {
	dtk.Group, dtk.Element = g, e
}

func (dtk *DcmTagKey) Equal(key DcmTagKey) bool {
	return (dtk.Group == key.Group) && (dtk.Element == key.Element)
}

/** returns true, if group is valid (permitted in DICOM files).
 *  Referring to the standard, groups 1, 3, 5, 7 and 0xFFFF are illegal.
 *  @return returns OFTrue if tag key has a valid group number.
 */
func (dtk *DcmTagKey) HasValidGroup() bool {
	if ((dtk.Group & 1) != 0) && ((dtk.Group <= 7) || (dtk.Group == 0xFFFF)) {
		return false
	}
	return true
}

/** checks whether the tag key is a valid group length element.
 *  Also calls hasValidGroup().
 *  @return returns OFTrue if tag key is a valid group length element
 */
func (dtk *DcmTagKey) IsGroupLength() bool {
	return (dtk.Element == 0) && dtk.HasValidGroup()
}

/** returns true if the tag key is private, ie. whether it has an odd group
 *  number. Also hasValidGroup() is called.
 *  @return returns OFTrue if group is private and valid.
 */
func (dtk *DcmTagKey) IsPrivate() bool {
	return ((dtk.Group & 1) != 0) && dtk.HasValidGroup()
}

/** returns true, if tag is a private reservation tag
 *  of the form (gggg,00xx) with gggg being odd and
 *  xx in the range of 10 and FF.
 *  @return returns OFTrue if tag key is a private reservation key
 */
func (dtk *DcmTagKey) IsPrivateReservation() bool {
	return dtk.IsPrivate() && dtk.Element >= 0x10 && dtk.Element <= 0xFF
}

/** generate a simple hash code for this attribute tag,
 *  used for fast look-up in the DICOM dictionary
 *  @return hash code for this tag
 */
func (dtk *DcmTagKey) Hash() uint32 {
	return ((uint32(int(dtk.Group)<<16) & 0xffff0000) | (uint32(int(dtk.Element) & 0xffff)))
}

/** convert tag key to string having the form "(gggg,eeee)".
 *  @return the string representation of this tag key
 */
func (dtk *DcmTagKey) ToString() string {
	var result string
	if dtk.Group == 0xFFFF && dtk.Element == 0xFFFF {
		result = "(????,????)"
	} else {
		result = fmt.Sprintf("(%04x,%04x)", dtk.Group, dtk.Element)
	}
	return result
}

/** returns true if a data element with the given tag key can
 *  be digitally signed, false otherwise
 *  @return true if signable, false otherwise
 */
func (dtk *DcmTagKey) IsSignableTag() bool {
	//no group length tags (element number of 0000)
	if dtk.Element == 0x0000 {
		return false
	}
	// no Length to End Tag
	if (dtk.Group == 0x0008) && (dtk.Element == 0x0001) {
		return false
	}

	//no tags with group number less than 0008
	if dtk.Group < 0x0008 {
		return false
	}

	//no tags from group FFFA (digital signatures sequence)
	if dtk.Group == 0xFFFA {
		return false
	}

	// no MAC Parameters sequence
	if (dtk.Group == 0x4ffe) && (dtk.Element == 0x0001) {
		return false
	}

	//no Data Set trailing Padding
	if (dtk.Group == 0xfffc) && (dtk.Element == 0xfffc) {
		return false
	}

	//no Sequence or Item Delimitation Tag
	if (dtk.Group == 0xfffe) && ((dtk.Element == 0xe00d) || (dtk.Element == 0xe0dd)) {
		return false
	}
	return true
}
