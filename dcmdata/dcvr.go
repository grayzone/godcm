package dcmdata

/** Global flag to enable/disable the generation of VR=UN
 */
var DcmEnableUnknownVRGeneration = true

/** Global flag to enable/disable the generation of VR=UT
 */
var DcmEnableUnlimitedTextVRGeneration = true

/** Global flag to enable/disable the automatic re-conversion of defined
 *  length UN elements read in an explicit VR transfer syntax, if the real
 *  VR is defined in the data dictionary.
 */
var DcmEnableUnknownVRConversion = false /* default OFFalse */

/*
** VR property table
 */

const DCMVR_PROP_NONE = 0x00
const DCMVR_PROP_NONSTANDARD = 0x01
const DCMVR_PROP_INTERNAL = 0x02
const DCMVR_PROP_EXTENDEDLENGTHENCODING = 0x04
const DCMVR_PROP_ISASTRING = 0x08

type dcmdataInterface interface {
	SetVR()
	GetEVR()
}

type DcmEVR int

const (

	/// application entity title
	EVR_AE DcmEVR = iota

	/// age string
	EVR_AS

	/// attribute tag
	EVR_AT

	/// code string
	EVR_CS

	/// date string
	EVR_DA

	/// decimal string
	EVR_DS

	/// date time string
	EVR_DT

	/// float single-precision
	EVR_FL

	/// float double-precision
	EVR_FD

	/// integer string
	EVR_IS

	/// long string
	EVR_LO

	/// long text
	EVR_LT

	/// other byte
	EVR_OB

	/// other float
	EVR_OF

	/// other word
	EVR_OW

	/// person name
	EVR_PN

	/// short string
	EVR_SH

	/// signed long
	EVR_SL

	/// sequence of items
	EVR_SQ

	/// signed short
	EVR_SS

	/// short text
	EVR_ST

	/// time string
	EVR_TM

	/// unique identifier
	EVR_UI

	/// unsigned long
	EVR_UL

	/// unsigned short
	EVR_US

	/// unlimited text
	EVR_UT

	/// OB or OW depending on context
	EVR_ox

	/// SS or US depending on context
	EVR_xs

	/// US, SS or OW depending on context, used for LUT Data (thus the name)
	EVR_lt

	/// na="not applicable", for data which has no VR
	EVR_na

	/// up="unsigned pointer", used internally for DICOMDIR suppor
	EVR_up

	/// used internally for items
	EVR_item

	/// used internally for meta info datasets
	EVR_metainfo

	/// used internally for datasets
	EVR_dataset

	/// used internally for DICOM files
	EVR_fileFormat

	/// used internally for DICOMDIR objects
	EVR_dicomDir

	/// used internally for DICOMDIR records
	EVR_dirRecord

	/// used internally for pixel sequences in a compressed image
	EVR_pixelSQ

	/// used internally for pixel items in a compressed image
	EVR_pixelItem

	/// used internally for elements with unknown VR (encoded with 4-byte length field in explicit VR)
	EVR_UNKNOWN

	/// unknown value representation
	EVR_UN

	/// used internally for uncompressed pixeld data
	EVR_PixelData

	/// used internally for overlay data
	EVR_OverlayData

	/// used internally for elements with unknown VR with 2-byte length field in explicit VR
	EVR_UNKNOWN2B
)

func (evr DcmEVR) String() string {
	var result string

	for i := 0; i < DcmVRDict_DIM; i++ {
		if evr == dcmVRDict[i].vr {
			result = dcmVRDict[i].vrstr
		}
	}
	return result

}

type dcmVREntry struct {
	vr             DcmEVR // Enumeration Value of Value representation
	vrstr          string
	vrName         string // Name of Value representation
	fValWidth      uint   // Length of minimal unit, used for swapping
	propertyFlags  int    // Normal, internal, non-standard vr
	minValueLength uint32 // Minimum length of a single value (bytes)
	maxValueLength uint32 // Maximum length of a single value (bytes)
}

var dcmVRDict = []dcmVREntry{
	{EVR_AE, "EVR_AE", "AE", 1, DCMVR_PROP_ISASTRING, 0, 16},
	{EVR_AS, "EVR_AS", "AS", 1, DCMVR_PROP_ISASTRING, 4, 4},
	{EVR_AT, "EVR_AT", "AT", 2, DCMVR_PROP_NONE, 4, 4},
	{EVR_CS, "EVR_CS", "CS", 1, DCMVR_PROP_ISASTRING, 0, 16},
	{EVR_DA, "EVR_DA", "DA", 1, DCMVR_PROP_ISASTRING, 8, 10},
	{EVR_DS, "EVR_DS", "DS", 1, DCMVR_PROP_ISASTRING, 0, 16},
	{EVR_DT, "EVR_DT", "DT", 1, DCMVR_PROP_ISASTRING, 0, 26},
	{EVR_FL, "EVR_FL", "FL", 4, DCMVR_PROP_NONE, 4, 4},
	{EVR_FD, "EVR_FD", "FD", 8, DCMVR_PROP_NONE, 8, 8},
	{EVR_IS, "EVR_IS", "IS", 1, DCMVR_PROP_ISASTRING, 0, 12},
	{EVR_LO, "EVR_LO", "LO", 1, DCMVR_PROP_ISASTRING, 0, 64},
	{EVR_LT, "EVR_LT", "LT", 1, DCMVR_PROP_ISASTRING, 0, 10240},
	{EVR_OB, "EVR_OB", "OB", 4, DCMVR_PROP_EXTENDEDLENGTHENCODING, 0, DCM_UndefinedLength},
	{EVR_OF, "EVR_OF", "OF", 4, DCMVR_PROP_EXTENDEDLENGTHENCODING, 0, DCM_UndefinedLength},
	{EVR_OW, "EVR_OW", "OW", 4, DCMVR_PROP_EXTENDEDLENGTHENCODING, 0, DCM_UndefinedLength},
	{EVR_PN, "EVR_PN", "PN", 1, DCMVR_PROP_ISASTRING, 0, 64},
	{EVR_SH, "EVR_SH", "SH", 1, DCMVR_PROP_ISASTRING, 0, 16},
	{EVR_SL, "EVR_SL", "SL", 4, DCMVR_PROP_NONE, 4, 4},
	{EVR_SQ, "EVR_SQ", "SQ", 0, DCMVR_PROP_EXTENDEDLENGTHENCODING, 0, DCM_UndefinedLength},
	{EVR_SS, "EVR_SS", "SS", 4, DCMVR_PROP_NONE, 2, 2},
	{EVR_ST, "EVR_ST", "ST", 1, DCMVR_PROP_ISASTRING, 0, 1024},
	{EVR_TM, "EVR_TM", "TM", 1, DCMVR_PROP_ISASTRING, 0, 16},
	{EVR_UI, "EVR_UI", "UI", 1, DCMVR_PROP_ISASTRING, 0, 64},
	{EVR_UL, "EVR_UL", "UL", 4, DCMVR_PROP_NONE, 4, 4},
	{EVR_US, "EVR_US", "US", 4, DCMVR_PROP_NONE, 2, 2},
	{EVR_UT, "EVR_UT", "UT", 1, DCMVR_PROP_ISASTRING | DCMVR_PROP_EXTENDEDLENGTHENCODING, 0, DCM_UndefinedLength},
	{EVR_ox, "EVR_ox", "ox", 4, DCMVR_PROP_NONSTANDARD | DCMVR_PROP_EXTENDEDLENGTHENCODING, 0, DCM_UndefinedLength},
	{EVR_xs, "EVR_xs", "xs", 4, DCMVR_PROP_NONSTANDARD, 2, 2},
	{EVR_lt, "EVR_lt", "lt", 4, DCMVR_PROP_NONSTANDARD | DCMVR_PROP_EXTENDEDLENGTHENCODING, 0, DCM_UndefinedLength},
	{EVR_na, "EVR_na", "na", 0, DCMVR_PROP_NONSTANDARD, 0, 0},
	{EVR_up, "EVR_up", "up", 4, DCMVR_PROP_NONSTANDARD, 4, 4},
	/* unique prefixes have been "invented" for the following internal VRs */
	{EVR_item, "EVR_item", "it_EVR_item", 0, DCMVR_PROP_NONSTANDARD | DCMVR_PROP_INTERNAL, 0, 0},
	{EVR_metainfo, "EVR_metainfo", "mi_EVR_metainfo", 0, DCMVR_PROP_NONSTANDARD | DCMVR_PROP_INTERNAL, 0, 0},
	{EVR_dataset, "EVR_dataset", "ds_EVR_dataset", 0, DCMVR_PROP_NONSTANDARD | DCMVR_PROP_INTERNAL, 0, 0},
	{EVR_fileFormat, "EVR_fileFormat", "ff_EVR_fileFormat", 0, DCMVR_PROP_NONSTANDARD | DCMVR_PROP_INTERNAL, 0, 0},
	{EVR_dicomDir, "EVR_dicomDir", "dd_EVR_dicomDir", 0, DCMVR_PROP_NONSTANDARD | DCMVR_PROP_INTERNAL, 0, 0},
	{EVR_dirRecord, "EVR_dirRecord", "dr_EVR_dirRecord", 0, DCMVR_PROP_NONSTANDARD | DCMVR_PROP_INTERNAL, 0, 0},

	{EVR_pixelSQ, "EVR_pixelSQ", "ps_EVR_pixelSQ", 4, DCMVR_PROP_NONSTANDARD | DCMVR_PROP_INTERNAL, 0, DCM_UndefinedLength},
	/* Moved from internal use to non standard only: necessary to distinguish from "normal" OB */
	{EVR_pixelItem, "EVR_pixelItem", "pi", 4, DCMVR_PROP_NONSTANDARD, 0, DCM_UndefinedLength},
	/* EVR_UNKNOWN (i.e. "future" VRs) should be mapped to UN or OB */
	{EVR_UNKNOWN, "EVR_UNKNOWN", "??", 4, DCMVR_PROP_NONSTANDARD | DCMVR_PROP_INTERNAL | DCMVR_PROP_EXTENDEDLENGTHENCODING, 0, DCM_UndefinedLength},

	/* Unknown Value Representation - Supplement 14 */
	{EVR_UN, "EVR_UN", "UN", 4, DCMVR_PROP_EXTENDEDLENGTHENCODING, 0, DCM_UndefinedLength},

	/* Pixel Data - only used in ident() */
	{EVR_PixelData, "EVR_PixelData", "PixelData", 0, DCMVR_PROP_INTERNAL, 0, DCM_UndefinedLength},
	/* Overlay Data - only used in ident() */
	{EVR_OverlayData, "EVR_OverlayData", "OverlayData", 0, DCMVR_PROP_INTERNAL, 0, DCM_UndefinedLength},
	/* illegal VRs, we assume no extended length coding */
	{EVR_UNKNOWN2B, "EVR_UNKNOWN2B", "??", 4, DCMVR_PROP_NONSTANDARD | DCMVR_PROP_INTERNAL, 0, DCM_UndefinedLength},
}

var DcmVRDict_DIM = len(dcmVRDict) // sizeof(DcmVRDict) / sizeof(DcmVREntry);

type DcmVR struct {
	vr DcmEVR
}

func NewDcmVR(v DcmEVR) *DcmVR {
	return &DcmVR{v}
}

/** assign new VR value by name
 *  @param vrName symbolic name of value representation
 */
func (v *DcmVR) SetByName(name string) {
	v.vr = EVR_UNKNOWN
	isfound := false
	for i := 0; i < DcmVRDict_DIM && !isfound; i++ {
		if name == dcmVRDict[i].vrName {
			isfound = true
			v.vr = dcmVRDict[i].vr
			break
		}
	}
	if len(name) == 2 {
		if name[0] == '?' && name[1] == '?' {
			v.vr = EVR_UNKNOWN2B
		}
		if !isfound && ((name[0] < 'A') || (name[0] > 'Z') || (name[1] < 'A') || (name[1] > 'Z')) {
			v.vr = EVR_UNKNOWN2B
		}
	}

}

/** assign new VR value
 *  @param evr enumerated VR value
 */
func (v *DcmVR) SetVR(evr DcmEVR) {
	if int(evr) >= 0 && int(evr) < DcmVRDict_DIM {
		v.vr = evr
	} else {
		v.vr = EVR_UNKNOWN
	}
}

/** get enumerated VR managed by this object
 *  @return enumerated VR
 */
func (v *DcmVR) GetEVR() DcmEVR {
	return v.vr
}

/** get enumerated standard VR managed by this object.
 *  If this object manages a non-standard, internal VR such as EVR_ox,
 *  this method returns the enumerated VR to which the internal VR will
 *  be mapped when writing the DICOM object.
 *  @return enumerated VR
 */
func (v *DcmVR) GetValidEVR() DcmEVR {
	result := DcmEVR(EVR_UNKNOWN)
	if v.IsStandard() {
		result = v.vr
	} else {
		switch v.vr {
		case DcmEVR(EVR_up):
			result = DcmEVR(EVR_UL)
		case DcmEVR(EVR_xs):
			result = DcmEVR(EVR_US)
		case DcmEVR(EVR_lt):
			result = DcmEVR(EVR_OW)
		case DcmEVR(EVR_ox), DcmEVR(EVR_pixelSQ):
			result = DcmEVR(EVR_OB)
		default:
			result = DcmEVR(EVR_UN) /* handle as Unknown VR (Supplement 14) */
		}
	}

	/*
	 ** If the generation of UN is not globally enabled then use OB instead.
	 ** We may not want to generate UN if other software cannot handle it.
	 */
	if result == DcmEVR(EVR_UN) {
		if !DcmEnableUnknownVRGeneration {
			result = DcmEVR(EVR_OB) /* handle UT as if OB */
		}
	}

	/*
	 ** If the generation of UT is not globally enabled then use OB instead.
	 ** We may not want to generate UT if other software cannot handle it.
	 */
	if result == DcmEVR(EVR_UT) {
		if !DcmEnableUnlimitedTextVRGeneration {
			result = DcmEVR(EVR_OB) /* handle UT as if OB */
		}
	}

	return result
}

/** returns true if VR is a standard DICOM VR
 *  @return true if VR is a standard DICOM VR
 */
func (v *DcmVR) IsStandard() bool {
	if (dcmVRDict[v.vr].propertyFlags & DCMVR_PROP_NONSTANDARD) == DCMVR_PROP_NONSTANDARD {
		return false
	}
	return true
}

/** get symbolic VR name for this object
 *  @return VR name string, never NULL
 */
func (v *DcmVR) GetVRName() string {
	return dcmVRDict[v.vr].vrName
}

/** get symbolic standard VR name for this object
 *  If this object manages a non-standard, internal VR such as EVR_ox,
 *  this method returns the name of the VR to which the internal VR will
 *  be mapped when writing the DICOM object.
 *  @return VR name string, never NULL
 */
func (v *DcmVR) GetValidVRName() string {
	var r DcmVR
	r.SetVR(v.GetValidEVR())
	return r.GetVRName()
}

/** compute the size for non-empty values of this VR.
 *  For fixed size VRs such as OW, US, SL, the method returns the size
 *  of each value, in bytes.  For variable length VRs (strings), it returns 1.
 *  For internal VRs it returns 0.
 *  @return size of values of this VR
 */
func (v *DcmVR) GetValueWidth() uint {
	return dcmVRDict[v.vr].fValWidth
}

/** returns true if VR is for internal use only
 *  @return true if VR is for internal use only
 */
func (v *DcmVR) IsForInternalUseOnly() bool {
	if (dcmVRDict[v.vr].propertyFlags & DCMVR_PROP_INTERNAL) == DCMVR_PROP_INTERNAL {
		return true
	}
	return false
}

/** returns true if VR represents a string
 *  @return true if VR represents a string
 */
func (v *DcmVR) IsaString() bool {
	if (dcmVRDict[v.vr].propertyFlags & DCMVR_PROP_ISASTRING) == DCMVR_PROP_ISASTRING {
		return true
	}
	return false
}

/** returns true if VR uses an extended length encoding for explicit transfer syntaxes
 *  @return true if VR uses an extended length encoding for explicit transfer syntaxes
 */
func (v *DcmVR) UsesExtendedLengthEncoding() bool {
	if (dcmVRDict[v.vr].propertyFlags & DCMVR_PROP_EXTENDEDLENGTHENCODING) == DCMVR_PROP_EXTENDEDLENGTHENCODING {
		return true
	}
	return false
}

/** check if VRs are equivalent
 *  VRs are considered equivalent if equal or if one of them is an internal VR
 *  and the other one is a possible standard VR to which the internal one maps.
 *  @param avr VR to compare with
 *  @return true if VRs are equivalent, false otherwise
 */
func (v *DcmVR) IsEquivalent(avr DcmVR) bool {
	evr := avr.GetEVR()
	if v.vr == avr.GetEVR() {
		return true
	}
	result := false
	switch v.vr {
	case EVR_ox:
		result = (evr == EVR_OB || evr == EVR_OW)
	case EVR_lt:
		result = (evr == EVR_OW || evr == EVR_US || evr == EVR_SS)
	case EVR_OB:
		result = (evr == EVR_ox)
	case EVR_OW:
		result = (evr == EVR_ox || evr == EVR_lt)
	case EVR_up:
		result = (evr == EVR_UL)
	case EVR_UL:
		result = (evr == EVR_up)
	case EVR_xs:
		result = (evr == EVR_SS || evr == EVR_US)
	case EVR_SS, EVR_US:
		result = (evr == EVR_xs || evr == EVR_lt)
	default:
	}

	return result
}

/** return minimum length of a value with this VR (in bytes), assuming single byte characters
 *  @return minimum length of a value
 */
func (v *DcmVR) GetMinValueLength() uint32 {
	return dcmVRDict[v.vr].minValueLength
}

/** return maximum length of a value with this VR (in bytes), assuming single byte characters
 *  @return maximum length of a value
 */
func (v *DcmVR) GetMaxValueLength() uint32 {
	return dcmVRDict[v.vr].maxValueLength
}
