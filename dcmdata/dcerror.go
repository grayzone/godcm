package dcmdata

import "github.com/grayzone/godcm/ofstd"

/*
 *  DCMTK module numbers for modules which create their own error codes.
 *  Module numbers > 1023 are reserved for user code.
 */
const (
	OFM_dcmdata  = 1
	OFM_ctndisp  = 2 /* retired */
	OFM_dcmimgle = 3
	OFM_dcmimage = 4
	OFM_dcmjpeg  = 5
	OFM_dcmnet   = 6
	OFM_dcmprint = 7
	OFM_dcmpstat = 8
	OFM_dcmsign  = 9
	OFM_dcmsr    = 10
	OFM_dcmtls   = 11
	OFM_imagectn = 12
	OFM_wlistctn = 13 /* retired */
	OFM_dcmwlm   = 14
	OFM_dcmpps   = 15
	OFM_dcmdbsup = 16
	OFM_dcmppswm = 17
	OFM_dcmjp2k  = 18
	OFM_dcmjpls  = 19
	OFM_dcmwave  = 20
	OFM_dcmrt    = 21
	OFM_dcmloco  = 22
	OFM_dcmstcom = 23
	OFM_dcmppscu = 24
)

// condition constants

/// invalid tag
var EC_InvalidTag = *ofstd.NewOFCondition(OFM_dcmdata, 1, ofstd.OF_error, "Invalid tag")

/// tag not found
var EC_TagNotFound = *ofstd.NewOFCondition(OFM_dcmdata, 2, ofstd.OF_error, "Tag not found")

/// invalid VR
var EC_InvalidVR = *ofstd.NewOFCondition(OFM_dcmdata, 3, ofstd.OF_error, "Invalid VR")

/// invalid stream
var EC_InvalidStream = *ofstd.NewOFCondition(OFM_dcmdata, 4, ofstd.OF_error, "Invalid stream")

/// end of stream
var EC_EndOfStream = *ofstd.NewOFCondition(OFM_dcmdata, 5, ofstd.OF_error, "End of stream")

/// corrupted data
var EC_CorruptedData = *ofstd.NewOFCondition(OFM_dcmdata, 6, ofstd.OF_error, "Corrupted data")

/// illegal call, perhaps wrong parameters
var EC_IllegalCall = *ofstd.NewOFCondition(OFM_dcmdata, 7, ofstd.OF_error, "Illegal call, perhaps wrong parameters")

/// sequence end
var EC_SequEnd = *ofstd.NewOFCondition(OFM_dcmdata, 8, ofstd.OF_error, "Sequence end")

/// doubled tag
var EC_DoubledTag = *ofstd.NewOFCondition(OFM_dcmdata, 9, ofstd.OF_error, "Doubled tag")

/// I/O suspension or premature end of stream
var EC_StreamNotifyClient = *ofstd.NewOFCondition(OFM_dcmdata, 10, ofstd.OF_error, "I/O suspension or premature end of stream")

/// stream mode (R/W, random/sequence) is wrong
var EC_WrongStreamMode = *ofstd.NewOFCondition(OFM_dcmdata, 11, ofstd.OF_error, "Mode (R/W, random/sequence) is wrong")

/// item end
var EC_ItemEnd = *ofstd.NewOFCondition(OFM_dcmdata, 12, ofstd.OF_error, "Item end")

/// compressed/uncompressed pixel representation not found
var EC_RepresentationNotFound = *ofstd.NewOFCondition(OFM_dcmdata, 13, ofstd.OF_error, "Pixel representation not found")

/// Pixel representation cannot be changed to requested transfer syntax
var EC_CannotChangeRepresentation = *ofstd.NewOFCondition(OFM_dcmdata, 14, ofstd.OF_error, "Pixel representation cannot be changed")

/// Unsupported compression or encryption
var EC_UnsupportedEncoding = *ofstd.NewOFCondition(OFM_dcmdata, 15, ofstd.OF_error, "Unsupported compression or encryption")

// error code 16 is reserved for zlib-related error messages

/// Parser failure: Putback operation failed
var EC_PutbackFailed = *ofstd.NewOFCondition(OFM_dcmdata, 17, ofstd.OF_error, "Parser failure: Putback operation failed")

// error code 18 is reserved for file read error messages
// error code 19 is reserved for file write error messages

/// Too many compression filters
var EC_DoubleCompressionFilters = *ofstd.NewOFCondition(OFM_dcmdata, 20, ofstd.OF_error, "Too many compression filters")

/// Storage media application profile violated
var EC_ApplicationProfileViolated = *ofstd.NewOFCondition(OFM_dcmdata, 21, ofstd.OF_error, "Storage media application profile violated")

// error code 22 is reserved for dcmodify error messages

/// Invalid offset
var EC_InvalidOffset = *ofstd.NewOFCondition(OFM_dcmdata, 23, ofstd.OF_error, "Invalid offset")

/// Too many bytes requested
var EC_TooManyBytesRequested = *ofstd.NewOFCondition(OFM_dcmdata, 24, ofstd.OF_error, "Too many bytes requested")

// error code 25 is reserved for tag path parsing error messages

// Invalid basic offset table
var EC_InvalidBasicOffsetTable = *ofstd.NewOFCondition(OFM_dcmdata, 26, ofstd.OF_error, "Invalid basic offset table")

/// Element length is larger than (explicit) length of surrounding item
var EC_ElemLengthLargerThanItem = *ofstd.NewOFCondition(OFM_dcmdata, 27, ofstd.OF_error, "Length of element larger than explicit length of surrounding item")

/// File meta information header missing
var EC_FileMetaInfoHeaderMissing = *ofstd.NewOFCondition(OFM_dcmdata, 28, ofstd.OF_error, "File meta information header missing")

/// Item or sequence content larger than explicit 32-bit length field permits
var EC_SeqOrItemContentOverflow = *ofstd.NewOFCondition(OFM_dcmdata, 29, ofstd.OF_error, "Item or sequence content exceeds maximum of 32-bit length field")

/// Value Representation (VR) violated
var EC_ValueRepresentationViolated = *ofstd.NewOFCondition(OFM_dcmdata, 30, ofstd.OF_error, "Value Representation violated")

/// Value Multiplicity (VM) violated
var EC_ValueMultiplicityViolated = *ofstd.NewOFCondition(OFM_dcmdata, 31, ofstd.OF_error, "Value Multiplicity violated")

/// Maximum VR length violated
var EC_MaximumLengthViolated = *ofstd.NewOFCondition(OFM_dcmdata, 32, ofstd.OF_error, "Maximum VR length violated")

/// Element length is larger than 16-bit length field permits
var EC_ElemLengthExceeds16BitField = *ofstd.NewOFCondition(OFM_dcmdata, 33, ofstd.OF_error, "Length of element value exceeds maximum of 16-bit length field")
