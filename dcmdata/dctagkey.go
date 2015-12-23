package dcmdata

import "fmt"

/// macro for an "undefined" attribute tag that is never used in DICOM
var DCM_UndefinedTagKey = DcmTagKey{group: 0xffff, element: 0xffff}

type DcmTagKey struct {
	group   uint16
	element uint16
}

func (dtk *DcmTagKey) init() {
	dtk.group = 0xffff
	dtk.element = 0xffff
}

func NewDcmTagKey() *DcmTagKey {
	return &DcmTagKey{group: 0xffff, element: 0xffff}
}

/** set value to given tag key
 *  @param key attribute tag to copy
 */
func (dtk *DcmTagKey) SetByKey(key DcmTagKey) {
	dtk.group = key.group
	dtk.element = key.element
}

/** set value to given group and element
 *  @param g group
 *  @param e element
 */
func (dtk *DcmTagKey) SetByValue(group uint16, element uint16) {
	dtk.group = group
	dtk.element = element
}

/** set group to given number
 *  @param g group
 */
func (dtk *DcmTagKey) SetGroup(group uint16) {
	dtk.group = group
}

/** set element to given number
 *  @param e element
 */
func (dtk *DcmTagKey) SetElement(element uint16) {
	dtk.element = element
}

/** returns group number
 *  @return returns the group number of the tag key
 */
func (dtk *DcmTagKey) GetGroup() uint16 {
	return dtk.group
}

/** returns element number
 *  @return Returns the element number of the tag key
 */
func (dtk *DcmTagKey) GetElement() uint16 {
	return dtk.element
}

func (dtk *DcmTagKey) Equal(key DcmTagKey) bool {
	return (dtk.group == key.GetGroup()) && (dtk.element == key.GetElement())
}

/** returns true, if group is valid (permitted in DICOM files).
 *  Referring to the standard, groups 1, 3, 5, 7 and 0xFFFF are illegal.
 *  @return returns OFTrue if tag key has a valid group number.
 */
func (dtk *DcmTagKey) HasValidGroup() bool {
	if ((dtk.group & 1) != 0) && ((dtk.group <= 7) || (dtk.group == 0xFFFF)) {
		return false
	}
	return true
}

/** checks whether the tag key is a valid group length element.
 *  Also calls hasValidGroup().
 *  @return returns OFTrue if tag key is a valid group length element
 */
func (dtk *DcmTagKey) IsGroupLength() bool {
	return (dtk.element == 0) && dtk.HasValidGroup()
}

/** returns true if the tag key is private, ie. whether it has an odd group
 *  number. Also hasValidGroup() is called.
 *  @return returns OFTrue if group is private and valid.
 */
func (dtk *DcmTagKey) IsPrivate() bool {
	return ((dtk.group & 1) != 0) && dtk.HasValidGroup()
}

/** returns true, if tag is a private reservation tag
 *  of the form (gggg,00xx) with gggg being odd and
 *  xx in the range of 10 and FF.
 *  @return returns OFTrue if tag key is a private reservation key
 */
func (dtk *DcmTagKey) IsPrivateReservation() bool {
	return dtk.IsPrivate() && dtk.element >= 0x10 && dtk.element <= 0xFF
}

/** generate a simple hash code for this attribute tag,
 *  used for fast look-up in the DICOM dictionary
 *  @return hash code for this tag
 */
func (dtk *DcmTagKey) Hash() uint32 {
	return ((uint32(int(dtk.GetGroup())<<16) & 0xffff0000) | (uint32(int(dtk.GetElement()) & 0xffff)))
}

/** convert tag key to string having the form "(gggg,eeee)".
 *  @return the string representation of this tag key
 */
func (dtk *DcmTagKey) ToString() string {
	var result string
	if dtk.group == 0xFFFF && dtk.element == 0xFFFF {
		result = "(????,????)"
	} else {
		result = fmt.Sprintf("(%04x,%04x)", dtk.group, dtk.element)
	}
	return result
}

/** returns true if a data element with the given tag key can
 *  be digitally signed, false otherwise
 *  @return true if signable, false otherwise
 */
func (dtk *DcmTagKey) IsSignableTag() bool {
	//no group length tags (element number of 0000)
	if dtk.element == 0x0000 {
		return false
	}
	// no Length to End Tag
	if (dtk.group == 0x0008) && (dtk.element == 0x0001) {
		return false
	}

	//no tags with group number less than 0008
	if dtk.group < 0x0008 {
		return false
	}

	//no tags from group FFFA (digital signatures sequence)
	if dtk.group == 0xFFFA {
		return false
	}

	// no MAC Parameters sequence
	if (dtk.group == 0x4ffe) && (dtk.element == 0x0001) {
		return false
	}

	//no Data Set trailing Padding
	if (dtk.group == 0xfffc) && (dtk.element == 0xfffc) {
		return false
	}

	//no Sequence or Item Delimitation Tag
	if (dtk.group == 0xfffe) && ((dtk.element == 0xe00d) || (dtk.element == 0xe0dd)) {
		return false
	}
	return true
}
