package dcmdata

import "github.com/grayzone/godcm/ofstd"

const DcmTag_ERROR_TagName = "Unknown Tag & Data"

type DcmTag struct {
	DcmTagKey
	DcmVR                            /// VR of this attribute tag
	tagName        string            /// name of this attribute tag, remains NULL unless getTagName() is called
	privateCreator string            /// private creator code, remains NULL unless setPrivateCreator() is called
	errorFlag      ofstd.OFCondition /// current error code, EC_Normal if a valid VR for the tag is known
}

var (
	DCM_ItemTag = DcmTag{DcmTagKey: DCM_Item, DcmVR: *NewDcmVR(EVR_na)}

	DCM_ItemDelimitationItemTag     = DcmTag{DcmTagKey: DCM_ItemDelimitationItem, DcmVR: *NewDcmVR(EVR_na)}
	DCM_SequenceDelimitationItemTag = DcmTag{DcmTagKey: DCM_SequenceDelimitationItem, DcmVR: *NewDcmVR(EVR_na)}
	DCM_InternalUseTag              = DcmTag{DcmTagKey: DcmTagKey{0xfffe, 0xfffe}, DcmVR: *NewDcmVR(EVR_UNKNOWN)}
)

func NewDcmTag() *DcmTag {
	var tag DcmTag
	tag.vr = EVR_UNKNOWN
	tag.errorFlag = EC_InvalidTag
	return &tag
}

/** constructor.
 *  Initializes group/element and VR from given parameters.
 *  No dictionary lookup needed/performed.
 *  @param g tag group
 *  @param e tag element
 *  @param avr VR
 */
func NewDcmTagWithGEV(g uint16, e uint16, avr DcmVR) *DcmTag {
	var tag DcmTag

	tag.group = g
	tag.element = e
	tag.DcmVR = avr

	tag.errorFlag = ofstd.EC_Normal
	return &tag

}

/// set specific VR
func (tag *DcmTag) SetVR(avr DcmVR) DcmVR {
	tag.DcmVR = avr
	if tag.GetEVR() == EVR_UNKNOWN {
		tag.errorFlag = EC_InvalidTag
	} else {
		tag.errorFlag = ofstd.EC_Normal
	}
	return tag.GetVR()
}

/// returns VR object by value
func (tag *DcmTag) GetVR() DcmVR {
	return tag.DcmVR
}

/// returns VR code
func (tag *DcmTag) GetEVR() DcmEVR {
	return tag.DcmVR.GetEVR()
}

/// returns name of VR
func (tag *DcmTag) GetVRName() string {
	return tag.DcmVR.GetVRName()
}

/** returns tag group
 *  @return tag group
 */
func (tag *DcmTag) GetGTag() uint16 {
	return tag.DcmTagKey.group
}

/** returns tag element
 *  @return tag element
 */
func (tag *DcmTag) GetETag() uint16 {
	return tag.DcmTagKey.element
}

/** returns a copy of the tag key by value
 *  @return copy of tag key, by value
 */
func (tag *DcmTag) GetXTag() DcmTagKey {
	return tag.DcmTagKey
}

/** returns name of attribute tag.
 *  If name has not been accessed before, a dictionary lookup
 *  under consideration of the current private creator code
 *  is performed.  If no attribute name is found, a default
 *  name is used.  Never returns NULL.
 *  @return attribute tag name, never NULL.
 */
func (tag *DcmTag) GetTagName() string {

	result := DcmTag_ERROR_TagName
	// to be continued

	return result
}

/** returns true if a data element with the given tag and VR
 *  can be digitally signed, false otherwise
 *  @return true if signable, false otherwise
 */
func (tag *DcmTag) IsSignable() bool {
	result := tag.IsSignableTag()
	if result {
		result = !tag.IsUnknownVR()
	}
	return result
}

/** returns true if the VR used for writing is "UN"
 */
func (tag *DcmTag) IsUnknownVR() bool {
	result := false
	switch tag.GetValidEVR() {
	case EVR_UNKNOWN, EVR_UNKNOWN2B, EVR_UN:
		result = true
	default:
		result = false
	}
	return result
}

/** convert the given string to a DICOM tag value
 *  @param name name or tag of the attribute to be searched for.
 *    If the name of the attribute is given the spelling has to be consistent
 *    with the spelling used in the data dictionary (e.g. "PatientName").
 *    If the tag values are used the format is "gggg,eeee" (i.e. two hexa-
 *    decimal numbers separated by a comma).
 *  @param value variable in which the resulting tag value is stored.
 *    If this functions fails to find the specified tag, this variable
 *    remains unchanged.
 *  @return status, EC_Normal upon success, an error code otherwise
 */
