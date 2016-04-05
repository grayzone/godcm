package dcmdata

import "github.com/grayzone/godcm/ofstd"

// Maxinum number of read bytes for a Value Element
const DCM_MaxReadLength = 4096

// Maximum length of tag and length in a DICOM element
const DCM_TagInfoLength = 12

// Optimum line length if not all data printed
const DCM_OptPrintLineLength = 70

// Optimum value length if not all data printed
const DCM_OptPrintValueLength = 40

// Optimum attribute name length (for tree output)
const DCM_OptPrintAttributeNameLength = 35

/** This flags defines whether automatic correction should be applied to input
 *  data (e.g. stripping of padding blanks, removal of blanks in UIDs, etc).
 *  Default is enabled.
 */
var DcmEnableAutomaticInputDataCorrection = true /* default OFTrue */

/** This flag defines the handling of illegal odd-length attributes: If flag is
 *  true, odd lengths are respected (i.e. an odd number of bytes is read from
 *  the input stream.) After successful reading, padding to even number of bytes
 *  is enforced by adding a zero pad byte if dcmEnableAutomaticInputDataCorrection
 *  is true. Otherwise the odd number of bytes remains as read.
 *
 *  If flag is false, old (pre DCMTK 3.5.2) behaviour applies: The length field
 *  implicitly incremented and an even number of bytes is read from the stream.
 */
var DcmAcceptOddAttributeLength = true /* default OFTrue */

/** This flag defines how UN attributes with undefined length are treated
 *  by the parser when reading. The default is to expect the content of the
 *  UN element (up to and including the sequence delimitation item)
 *  to be encoded in Implicit VR Little Endian, as described in CP 246.
 *  DCMTK expects the attribute to be encoded like a DICOM sequence, i.e.
 *  the content of each item is parsed as a DICOM dataset.
 *  If the flag is disabled old (pre DCMTK 3.5.4) behaviour applies: The
 *  attribute is treated as if it was an Explicit VR SQ element.
 *
 *  Note that the flag only affects the read behaviour but not the write
 *  behaviour - DCMTK will never write UN elements with undefined length.
 */
var DcmEnableCP246Support = true /* default OFTrue */

/** DCMTK releases up to 3.5.3 created a non-conforming byte stream
 *  as input to the MAC algorithm when creating or verifying digital signatures
 *  including compressed pixel data (i.e. signatures including attribute
 *  (7FE0,0010) in an encapsulated transfer syntax). This has been fixed
 *  in DCMTK 3.5.4, but this flag allows to revert to the old behavior
 *  in order to create or verify signatures that are compatible with older
 *  releases. Default is "off" (OFFalse).
 */
var DcmEnableOldSignatureFormat = false /* default OFFalse */

/** This flag defines whether the transfer syntax for uncompressed datasets
 *  is detected automatically.  The automatic detection has been introduced
 *  since there are (incorrectly encoded) DICOM dataset stored with a
 *  different transfer syntax than specified in the meta header.
 */
var DcmAutoDetectDatasetXfer = false /* default OFFalse */

/** This flag defines how non-standard VRs are treated by the parser when
 *  reading. The default is to treat data element with non-standard VR as
 *  unknown. If this flag is enabled, the parser will try to read the data
 *  element with Implicit VR Little Endian transfer syntax.
 */
var DcmAcceptUnexpectedImplicitEncoding = false /* default OFFalse */

/** This flag indicates, whether private attributes with implicit transfer
 *  syntax having a maximum length should be handled as sequences (ignoring
 *  any dictionary entries for that tag). This can happen, if for example
 *  a private creator element is illegally inserted with VR SQ
 *  (undefined length and implicit coding). The parser usually would then
 *  try to parse the element with VR=LO (private creator) with maximum
 *  length, which would lead to an error. The default behaviour is to
 *  rely on the dictionary.
 */
var DcmReadImplPrivAttribMaxLengthAsSQ = false /* default OFFalse */

/** This flag indicates, whether parsing errors during reading
 *  should be ignored, ie whether the parser should try to recover and
 *  parse the rest of the stream.
 *  This flag does not work for all parsing errors (at this time)
 *  making sense but was introduced afterwards.
 */
var DcmIgnoreParsingErrors = false /* default OFFalse */

/** This flag indicates, whether parsing should stop after a certain
 *  element in the stream was parsed. This is especially useful for
 *  datasets containing garbage at the end, usually after the Pixel
 *  Data attribute. To prevent the parser for "stumbling" over that
 *  garbage, it is possible to tell the parser to stop after a
 *  specific element. The flag is only sensitive to elements on
 *  dataset level, ie. inside sequence any occurence of the specified
 *  tag is ignored. Caution: Note that if Pixel Data is chosen
 *  as stop element, any attributes behind will not be parsed, e. g.
 *  any digital signature attributes coming after.
 *  Default is (0xffff,0xffff), which means that the feature is
 *  disabled.
 */
var DcmStopParsingAfterElement = true /* default OFTrue */

/** This flag influences behaviour when writing a dataset with items
 *  and sequences set to be encoded with explicit length. It is possible
 *  that the content of a sequence (or item) has an encoded length greater
 *  than the maximum 32-bit value that can be written to the sequence (item)
 *  length field. If this flag is enabled (OFTrue) then the encoding of the
 *  very sequence (item) is switched to undefined length encoding. Any
 *  contained items (sequences) will be encoded explicitly if possible.
 *  Default is OFTrue, i.e. encoding is switched to implicit if maximum
 *  size of length field is exceeded.
 */
var DcmWriteOversizedSeqsAndItemsUndefined = true /* default OFTrue */

/** This flag allows for ignoring the value of (0002,0000) File Meta Information
 *  Group Length which is useful in cases where this value is incorrect.  If the
 *  header length is ignored, the behavior is identical to the case when no value
 *  is available (i.e. all elements are read as long as the group number is 0x0002).
 */
var DcmIgnoreFileMetaInformationGroupLength = false /* default OFFalse */

type DcmObject struct {
	/// the DICOM attribute tag and VR for this object
	tag DcmTag
	/// the length of this attribute as read from stream, may be undefined length
	length uint32
	/// transfer state during read and write operations
	fTransferState E_TransferState
	/// number of bytes already read/written during transfer
	fTransferredBytes uint32
	/// error flag for this object.
	errorFlag ofstd.OFCondition
}

/** constructor.
 *  Create new object from given tag and length.
 *  @param tag DICOM tag for the new element
 *  @param len value length for the new element
 */
func NewDcmObject(t DcmTag, l uint32) *DcmObject {
	return &DcmObject{tag: t, length: l, fTransferState: ERW_init, fTransferredBytes: 0, errorFlag: ofstd.EC_Normal}
}

/** return the value representation assigned to this object.
 *  If object was read from a stream, this method returns the VR
 *  that was defined in the stream for this object. It is, therefore,
 *  possible that the VR does not match the one defined in the data
 *  dictionary for the tag assigned to this object.
 *  @return VR of this object
 */
func (o *DcmObject) GetVR() DcmEVR {
	return o.tag.GetEVR()
}

/** check if this element is a string type, based on the VR.
 *  Since the check is based on the VR and not on the class,
 *  the result of this method is not a guarantee that the object
 *  can be safely casted to one of the string-VR subclasses.
 *  @return true if this object is a string VR, false otherwise
 */
func (o *DcmObject) IsaString() bool {
	return o.tag.IsaString()
}

/** return the current transfer (read/write) state of this object.
 *  @return transfer state of this object
 */

func (o *DcmObject) TransferState() E_TransferState {
	return o.fTransferState
}

/** initialize the transfer state of this object. This method must be called
 *  before this object is written to a stream or read (parsed) from a stream.
 */
func (o *DcmObject) TransferInit() {
	o.fTransferState = ERW_init
	o.fTransferredBytes = 0

}

/** finalize the transfer state of this object. This method must be called
 *  when reading/writing this object from/to a stream has been completed.
 */
func (o *DcmObject) TransferEnd() {
	o.fTransferState = ERW_notInitialized
}

/** return the group number of the attribute tag for this object
 *  @return group number of the attribute tag for this object
 */
func (o *DcmObject) GetGTag() uint16 {
	return o.tag.GetGTag()
}

/** return the element number of the attribute tag for this object
 *  @return element number of the attribute tag for this object
 */
func (o *DcmObject) GetETag() uint16 {
	return o.tag.GetETag()
}

/** return const reference to the attribute tag for this object
 *  @return const reference to the attribute tag for this object
 */
func (o *DcmObject) GetTag() DcmTag {
	return o.tag
}

/** assign group tag (but not element tag) of the attribute tag for this object.
 *  This is sometimes useful when creating repeating group elements.
 *  @param gtag new attribute group tag
 */
func (o *DcmObject) SetGTag(g uint16) {
	o.tag.group = g
}

/** assign a new Value Representation (VR) to this object. This operation
 *  is only supported for very few subclasses derived from this class,
 *  in particular for classes handling pixel data which may either be
 *  of OB or OW value representation.
 *  @param vr value representation
 *  @return EC_Normal if successful, an error code otherwise
 */
func (o *DcmObject) SetVR(vr DcmEVR) ofstd.OFCondition {
	return EC_IllegalCall

}

/** returns true if the current object may be included in a digital signature
 *  @return true if signable, false otherwise
 */
func (o *DcmObject) IsSignable() bool {
	return o.tag.IsSignable()
}

/** returns true if the object contains an element with Unknown VR at any nesting level
 *  @return true if the object contains an element with Unknown VR, false otherwise
 */
func (o *DcmObject) ContainsUnknownVR() bool {
	return o.tag.IsUnknownVR()
}

/** check if this object contains non-ASCII characters
 *  @param checkAllStrings not used in this class
 *  @return always returns false, i.e. no extended characters used
 */
func (o *DcmObject) ContainsExtendedCharacters(checkAllStrings bool) bool {
	return false
}

/** check if this object is affected by SpecificCharacterSet
 *  @return always returns false, i.e. not affected by SpecificCharacterSet
 */
func (o *DcmObject) IsAffectedBySpecificCharacterSet() bool {
	return false
}

/** check if this object is empty
 *  @param normalize normalize value before checking (ignore non-significant characters)
 *  @return true if object is empty, i.e. has no value, false otherwise
 */
func (o *DcmObject) IsEmpty(normalize bool) bool {
	return (o.length == 0)
}

/** check if this element is a leaf node in a dataset tree.
 *  All subclasses of DcmElement except for DcmSequenceOfItems
 *  are leaf nodes, while DcmSequenceOfItems, DcmItem, DcmDataset etc.
 *  are not.
 *  @return true if leaf node, false otherwise.
 */
func (o *DcmObject) IsLeaf() bool {
	return false
}

/** return the current value of the Length field (which is different from the functionality
 *  of the public getLength method). Only needed for internal purposes and for checker tools
 *  that verify values against the length field.
 *  @return current value of length field
 */
func (o *DcmObject) GetLengthField() uint32 {
	return o.length
}

/** calculate the length of this DICOM element when encoded with the
 *  given transfer syntax and the given encoding type for sequences.
 *  For elements, the length includes the length of the tag, length field,
 *  VR field and the value itself, for items and sequences it returns
 *  the length of the complete item or sequence including delimitation tags
 *  if applicable. Never returns undefined length.
 *  @param xfer transfer syntax for length calculation
 *  @param enctype sequence encoding type for length calculation
 *  @return length of DICOM element
 */
func (o *DcmObject) CalcElementLength(xfer E_TransferSyntax, enctype E_EncodingType) uint32 {
	return 0
}

/** return identifier for this class. Every class derived from this class
 *  returns a unique value of type enum DcmEVR for this call. This is used
 *  as a "poor man's RTTI" to correctly identify instances derived from
 *  this class even on compilers not supporting RTTI.
 *  @return type identifier of this class
 */
func (o *DcmObject) Ident() DcmEVR {
	return o.GetVR()
}
