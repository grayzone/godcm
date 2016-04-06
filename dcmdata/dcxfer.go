package dcmdata

/*
typedef unsigned long Uint32;
typedef unsigned short Uint16;
int FindMachineTransferSyntax(){
	int localByteOrderFlag;
	union
    {
        Uint32 ul;
        char uc[4];
    } tl;
    union
    {
        Uint16 us;
        char uc[2];
    } ts;
    tl.ul = 1;
    ts.us = 1;
    if (tl.uc[0] == 1 && !(tl.uc[1] | tl.uc[2] | tl.uc[3]) && ts.uc[0] == 1 && !(ts.uc[1]))
        localByteOrderFlag = 1;
    else if (tl.uc[3] == 1 && !(tl.uc[0] | tl.uc[1] | tl.uc[2]) && ts.uc[1] == 1 && !(ts.uc[0]))
        localByteOrderFlag = 2;
    else
        localByteOrderFlag = 0;
    return localByteOrderFlag;
}
*/
import "C"

/** enumeration of all DICOM transfer syntaxes known to the toolkit
 */
type E_TransferSyntax int

const (
	/// unknown transfer syntax or dataset created in-memory
	EXS_Unknown = -1
	/// Implicit VR Little Endian
	EXS_LittleEndianImplicit = 0
	/// Implicit VR Big Endian (pseudo transfer syntax that does not really exist)
	EXS_BigEndianImplicit = 1
	/// Explicit VR Little Endian
	EXS_LittleEndianExplicit E_TransferSyntax = 2
	/// Explicit VR Big Endian
	EXS_BigEndianExplicit = 3
	/// JPEG Baseline (lossy)
	EXS_JPEGProcess1TransferSyntax = 4
	/// JPEG Extended Sequential (lossy, 8/12 bit)
	EXS_JPEGProcess2_4TransferSyntax = 5
	/// JPEG Extended Sequential (lossy, 8/12 bit), arithmetic coding
	EXS_JPEGProcess3_5TransferSyntax = 6
	/// JPEG Spectral Selection, Non-Hierarchical (lossy, 8/12 bit)
	EXS_JPEGProcess6_8TransferSyntax = 7
	/// JPEG Spectral Selection, Non-Hierarchical (lossy, 8/12 bit), arithmetic coding
	EXS_JPEGProcess7_9TransferSyntax = 8
	/// JPEG Full Progression, Non-Hierarchical (lossy, 8/12 bit)
	EXS_JPEGProcess10_12TransferSyntax = 9
	/// JPEG Full Progression, Non-Hierarchical (lossy, 8/12 bit), arithmetic coding
	EXS_JPEGProcess11_13TransferSyntax = 10
	/// JPEG Lossless with any selection value
	EXS_JPEGProcess14TransferSyntax = 11
	/// JPEG Lossless with any selection value, arithmetic coding
	EXS_JPEGProcess15TransferSyntax = 12
	/// JPEG Extended Sequential, Hierarchical (lossy, 8/12 bit)
	EXS_JPEGProcess16_18TransferSyntax = 13
	/// JPEG Extended Sequential, Hierarchical (lossy, 8/12 bit), arithmetic coding
	EXS_JPEGProcess17_19TransferSyntax = 14
	/// JPEG Spectral Selection, Hierarchical (lossy, 8/12 bit)
	EXS_JPEGProcess20_22TransferSyntax = 15
	/// JPEG Spectral Selection, Hierarchical (lossy, 8/12 bit), arithmetic coding
	EXS_JPEGProcess21_23TransferSyntax = 16
	/// JPEG Full Progression, Hierarchical (lossy, 8/12 bit)
	EXS_JPEGProcess24_26TransferSyntax = 17
	/// JPEG Full Progression, Hierarchical (lossy, 8/12 bit), arithmetic coding
	EXS_JPEGProcess25_27TransferSyntax = 18
	/// JPEG Lossless, Hierarchical
	EXS_JPEGProcess28TransferSyntax = 19
	/// JPEG Lossless, Hierarchical, arithmetic coding
	EXS_JPEGProcess29TransferSyntax = 20
	/// JPEG Lossless, Selection Value 1
	EXS_JPEGProcess14SV1TransferSyntax = 21
	/// Run Length Encoding (lossless)
	EXS_RLELossless = 22
	/// JPEG-LS (lossless)
	EXS_JPEGLSLossless = 23
	/// JPEG-LS (lossless or near-lossless mode)
	EXS_JPEGLSLossy = 24
	/// Deflated Explicit VR Little Endian
	EXS_DeflatedLittleEndianExplicit = 25
	/// JPEG 2000 (lossless)
	EXS_JPEG2000osslessOnly = 26
	/// JPEG 2000 (lossless or lossy)
	EXS_JPEG2000 = 27
	/// MPEG2 Main Profile at Main Level
	EXS_MPEG2MainProfileAtMainLevel = 28
	/// MPEG2 Main Profile at High Level
	EXS_MPEG2MainProfileAtHighLevel = 29
	/// JPEG 2000 part 2 multi-component extensions (lossless)
	EXS_JPEG2000MulticomponentLosslessOnly = 30
	/// JPEG 2000 part 2 multi-component extensions (lossless or lossy)
	EXS_JPEG2000Multicomponent = 31
	/// JPIP Referenced
	EXS_JPIPReferenced = 32
	/// JPIP Referenced Deflate
	EXS_JPIPReferencedDeflate = 33
)

/** enumeration of byte orders
 */
type E_ByteOrder int

const (
	/// unknown
	EBO_unknown = 0

	/// little endian
	EBO_LittleEndian = 1

	/// big endian
	EBO_BigEndian = 2
)

func (bo E_ByteOrder) String() string {
	var result string
	switch bo {
	case EBO_unknown:
		result = "EBO_unknown"
	case EBO_LittleEndian:
		result = "EBO_LittleEndian"
	case EBO_BigEndian:
		result = "EBO_BigEndian"
	}
	return result
}

/** enumeration of VR encoding options
 */
type E_VRType int

const (
	/// implicit VR encoding
	EVT_Implicit = 0

	/// explicit VR encoding
	EVT_Explicit = 1
)

/** enumeration of pixel data encapsulation options
 */
type E_JPEGEncapsulated int

const (
	/// pixel data not encapsulated
	EJE_NotEncapsulated = 0

	/// pixel data encapsulated
	EJE_Encapsulated = 1
)

/** enumeration of stream compression techniques
 */
type E_StreamCompression int

const (

	/// no stream compression
	ESC_none = 0
	/// unsupported stream compression
	ESC_unsupported = 1

	/// zlib stream compression
//	ESC_zlib = 2 // not supported
)

/** a class that allows for a lookup of Transfer Syntax properties and readable descriptions
 */
type DcmXfer struct {
	/// transfer syntax UID
	xferID string

	/// transfer syntax name
	xferName string

	/// transfer syntax enum
	xferSyn E_TransferSyntax

	/// transfer syntax byte order
	byteOrder E_ByteOrder

	/// transfer syntax VR encoding (implicit/explicit)
	vrType E_VRType

	/// transfer syntax encapsulated or native
	encapsulated E_JPEGEncapsulated

	/// 8-bit lossy JPEG process ID for this transfer syntax, 0 if not applicable
	JPEGProcess8 uint32

	/// 12-bit lossy JPEG process ID for this transfer syntax, 0 if not applicable
	JPEGProcess12 uint32

	/// flag indicating whether this transfer syntax has been retired from DICOM
	retired bool

	/// transfer syntax stream compression type
	streamCompression E_StreamCompression
}

type s_XferNames struct {
	xferID            string
	xferName          string
	xfer              E_TransferSyntax
	byteOrder         E_ByteOrder
	vrType            E_VRType
	encapsulated      E_JPEGEncapsulated
	JPEGProcess8      uint32
	JPEGProcess12     uint32
	retired           bool
	streamCompression E_StreamCompression
}

const ERROR_XferName = "Unknown Transfer Syntax"

var xferNames = []s_XferNames{
	{UID_LittleEndianImplicitTransferSyntax,
		"Little Endian Implicit",
		EXS_LittleEndianImplicit,
		EBO_LittleEndian,
		EVT_Implicit,
		EJE_NotEncapsulated,
		0, 0,
		false,
		ESC_none},
	{"", // illegal type
		"Virtual Big Endian Implicit",
		EXS_BigEndianImplicit,
		EBO_BigEndian,
		EVT_Implicit,
		EJE_NotEncapsulated,
		0, 0,
		false,
		ESC_none},
	{UID_LittleEndianExplicitTransferSyntax,
		"Little Endian Explicit",
		EXS_LittleEndianExplicit,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_NotEncapsulated,
		0, 0,
		false,
		ESC_none},
	{UID_BigEndianExplicitTransferSyntax, // defined in dctypes.h
		"Big Endian Explicit",
		EXS_BigEndianExplicit,
		EBO_BigEndian,
		EVT_Explicit,
		EJE_NotEncapsulated,
		0, 0,
		true,
		ESC_none},
	{UID_JPEGProcess1TransferSyntax,
		"JPEG Baseline",
		EXS_JPEGProcess1TransferSyntax,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_Encapsulated,
		1, 1,
		false,
		ESC_none},
	{UID_JPEGProcess2_4TransferSyntax,
		"JPEG Extended, Process 2+4",
		EXS_JPEGProcess2_4TransferSyntax,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_Encapsulated,
		2, 4,
		false,
		ESC_none},
	{UID_JPEGProcess3_5TransferSyntax,
		"JPEG Extended, Process 3+5",
		EXS_JPEGProcess3_5TransferSyntax,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_Encapsulated,
		3, 5,
		true,
		ESC_none},
	{UID_JPEGProcess6_8TransferSyntax,
		"JPEG Spectral Selection, Non-hierarchical, Process 6+8",
		EXS_JPEGProcess6_8TransferSyntax,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_Encapsulated,
		6, 8,
		true,
		ESC_none},
	{UID_JPEGProcess7_9TransferSyntax,
		"JPEG Spectral Selection, Non-hierarchical, Process 7+9",
		EXS_JPEGProcess7_9TransferSyntax,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_Encapsulated,
		7, 9,
		true,
		ESC_none},
	{UID_JPEGProcess10_12TransferSyntax,
		"JPEG Full Progression, Non-hierarchical, Process 10+12",
		EXS_JPEGProcess10_12TransferSyntax,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_Encapsulated,
		10, 12,
		true,
		ESC_none},
	{UID_JPEGProcess11_13TransferSyntax,
		"JPEG Full Progression, Non-hierarchical, Process 11+13",
		EXS_JPEGProcess11_13TransferSyntax,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_Encapsulated,
		11, 13,
		true,
		ESC_none},
	{UID_JPEGProcess14TransferSyntax,
		"JPEG Lossless, Non-hierarchical, Process 14",
		EXS_JPEGProcess14TransferSyntax,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_Encapsulated,
		14, 14,
		false,
		ESC_none},
	{UID_JPEGProcess15TransferSyntax,
		"JPEG Lossless, Non-hierarchical, Process 15",
		EXS_JPEGProcess15TransferSyntax,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_Encapsulated,
		15, 15,
		true,
		ESC_none},
	{UID_JPEGProcess16_18TransferSyntax,
		"JPEG Extended, Hierarchical, Process 16+18",
		EXS_JPEGProcess16_18TransferSyntax,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_Encapsulated,
		16, 18,
		true,
		ESC_none},
	{UID_JPEGProcess17_19TransferSyntax,
		"JPEG Extended, Hierarchical, Process 17+19",
		EXS_JPEGProcess17_19TransferSyntax,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_Encapsulated,
		17, 19,
		true,
		ESC_none},
	{UID_JPEGProcess20_22TransferSyntax,
		"JPEG Spectral Selection, Hierarchical, Process 20+22",
		EXS_JPEGProcess20_22TransferSyntax,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_Encapsulated,
		20, 22,
		true,
		ESC_none},
	{UID_JPEGProcess21_23TransferSyntax,
		"JPEG Spectral Selection, Hierarchical, Process 21+23",
		EXS_JPEGProcess21_23TransferSyntax,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_Encapsulated,
		21, 23,
		true,
		ESC_none},
	{UID_JPEGProcess24_26TransferSyntax,
		"JPEG Full Progression, Hierarchical, Process 24+26",
		EXS_JPEGProcess24_26TransferSyntax,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_Encapsulated,
		24, 26,
		true,
		ESC_none},
	{UID_JPEGProcess25_27TransferSyntax,
		"JPEG Full Progression, Hierarchical, Process 25+27",
		EXS_JPEGProcess25_27TransferSyntax,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_Encapsulated,
		25, 27,
		true,
		ESC_none},
	{UID_JPEGProcess28TransferSyntax,
		"JPEG Lossless, Hierarchical, Process 28",
		EXS_JPEGProcess28TransferSyntax,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_Encapsulated,
		28, 28,
		true,
		ESC_none},
	{UID_JPEGProcess29TransferSyntax,
		"JPEG Lossless, Hierarchical, Process 29",
		EXS_JPEGProcess29TransferSyntax,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_Encapsulated,
		29, 29,
		true,
		ESC_none},
	{UID_JPEGProcess14SV1TransferSyntax,
		"JPEG Lossless, Non-hierarchical, 1st Order Prediction",
		EXS_JPEGProcess14SV1TransferSyntax,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_Encapsulated,
		14, 14,
		false,
		ESC_none},
	{UID_RLELosslessTransferSyntax,
		"RLE Lossless",
		EXS_RLELossless,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_Encapsulated,
		0, 0,
		false,
		ESC_none},
	{UID_JPEGLSLosslessTransferSyntax,
		"JPEG-LS Lossless",
		EXS_JPEGLSLossless,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_Encapsulated,
		0, 0,
		false,
		ESC_none},
	{UID_JPEGLSLossyTransferSyntax,
		"JPEG-LS Lossy (Near-lossless)",
		EXS_JPEGLSLossy,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_Encapsulated,
		0, 0,
		false,
		ESC_none},
	{UID_DeflatedExplicitVRLittleEndianTransferSyntax,
		"Deflated Explicit VR Little Endian",
		EXS_DeflatedLittleEndianExplicit,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_NotEncapsulated,
		0, 0,
		false,
		ESC_unsupported},
	{UID_JPEG2000LosslessOnlyTransferSyntax,
		"JPEG 2000 (Lossless only)",
		EXS_JPEG2000osslessOnly,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_Encapsulated,
		0, 0,
		false,
		ESC_none},
	{UID_JPEG2000TransferSyntax,
		"JPEG 2000 (Lossless or Lossy)",
		EXS_JPEG2000,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_Encapsulated,
		0, 0,
		false,
		ESC_none},
	{UID_MPEG2MainProfileAtMainLevelTransferSyntax,
		"MPEG2 Main Profile @ Main Level",
		EXS_MPEG2MainProfileAtMainLevel,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_Encapsulated,
		0, 0,
		false,
		ESC_none},
	{UID_MPEG2MainProfileAtHighLevelTransferSyntax,
		"MPEG2 Main Profile @ High Level",
		EXS_MPEG2MainProfileAtHighLevel,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_Encapsulated,
		0, 0,
		false,
		ESC_none},
	{UID_JPEG2000Part2MulticomponentImageCompressionLosslessOnlyTransferSyntax,
		"JPEG 2000 Part 2 Multicomponent Image Compression (Lossless only)",
		EXS_JPEG2000MulticomponentLosslessOnly,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_Encapsulated,
		0, 0,
		false,
		ESC_none},
	{UID_JPEG2000Part2MulticomponentImageCompressionTransferSyntax,
		"JPEG 2000 Part 2 Multicomponent Image Compression (Lossless or Lossy)",
		EXS_JPEG2000Multicomponent,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_Encapsulated,
		0, 0,
		false,
		ESC_none},
	{UID_JPIPReferencedTransferSyntax,
		"JPIP Referenced",
		EXS_JPIPReferenced,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_NotEncapsulated, // in fact, pixel data shall be referenced via (0028,7FE0) Pixel Data Provider URL
		0, 0,
		false,
		ESC_none},
	{UID_JPIPReferencedDeflateTransferSyntax,
		"JPIP Referenced Deflate",
		EXS_JPIPReferencedDeflate,
		EBO_LittleEndian,
		EVT_Explicit,
		EJE_NotEncapsulated, // in fact, pixel data shall be referenced via (0028,7FE0) Pixel Data Provider URL
		0, 0,
		false,
		ESC_unsupported},
}

var DIM_OF_XferNames = len(xferNames)

/** global constant describing the byte order on the machine the application
 *  is currently executing on. This is runtime and not compile time information
 *  because of "fat" binaries that can be executed on multiple CPU types (e.g. NeXTStep)
 */
// global constant: local byte order (little or big endian)
var GLocalByteOrder E_ByteOrder = FindMachineTransferSyntax()

func FindMachineTransferSyntax() E_ByteOrder {

	return (E_ByteOrder)(C.FindMachineTransferSyntax())

}

func NewDcmXfer(xfer E_TransferSyntax) *DcmXfer {
	var f DcmXfer

	f.xferName = ERROR_XferName
	f.xferSyn = EXS_Unknown
	f.byteOrder = EBO_unknown
	f.vrType = EVT_Implicit
	f.encapsulated = EJE_NotEncapsulated
	f.JPEGProcess8 = 0
	f.JPEGProcess12 = 0
	f.retired = false
	f.streamCompression = ESC_none

	for _, v := range xferNames {
		if v.xfer == xfer {
			f.xferSyn = v.xfer
			f.xferID = v.xferID
			f.xferName = v.xferName
			f.byteOrder = v.byteOrder
			f.vrType = v.vrType
			f.encapsulated = v.encapsulated
			f.JPEGProcess8 = v.JPEGProcess8
			f.JPEGProcess12 = v.JPEGProcess12
			f.retired = v.retired
			f.streamCompression = v.streamCompression
			break
		}
	}
	return &f
}

/// return true if transfer syntax is explicit VR, false otherwise
func (xfer *DcmXfer) IsExplicitVR() bool {
	return xfer.vrType == EVT_Explicit
}

func (xfer *DcmXfer) SizeofTagHeader(evr DcmEVR) uint32 {
	// all implicit VRs have the same format
	if !xfer.IsExplicitVR() {
		return 8 // for Tag und Length
	}
	// some VRs have an extended format
	var vr DcmVR
	vr.SetVR(evr)
	if vr.UsesExtendedLengthEncoding() {
		return 12 // for Tag, Length, VR und reserved
	}
	return 8 // for Tag, Length und VR
}
