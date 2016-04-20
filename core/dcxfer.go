package core

import (
	"errors"
	"fmt"
)

// ETransferSyntax : enumeration of all DICOM transfer syntaxes known to the toolkit
type ETransferSyntax int

const (
	//EXSUnknown : unknown transfer syntax or dataset created in-memory
	EXSUnknown = -1
	//EXSLittleEndianImplicit :  Implicit VR Little Endian
	EXSLittleEndianImplicit = 0
	/// Implicit VR Big Endian (pseudo transfer syntax that does not really exist)
	EXSBigEndianImplicit = 1
	/// Explicit VR Little Endian
	EXSLittleEndianExplicit ETransferSyntax = 2
	/// Explicit VR Big Endian
	EXSBigEndianExplicit = 3
	/// JPEG Baseline (lossy)
	EXSJPEGProcess1TransferSyntax = 4
	/// JPEG Extended Sequential (lossy, 8/12 bit)
	EXSJPEGProcess24TransferSyntax = 5
	/// JPEG Extended Sequential (lossy, 8/12 bit), arithmetic coding
	EXSJPEGProcess35TransferSyntax = 6
	/// JPEG Spectral Selection, Non-Hierarchical (lossy, 8/12 bit)
	EXSJPEGProcess68TransferSyntax = 7
	/// JPEG Spectral Selection, Non-Hierarchical (lossy, 8/12 bit), arithmetic coding
	EXSJPEGProcess79TransferSyntax = 8
	/// JPEG Full Progression, Non-Hierarchical (lossy, 8/12 bit)
	EXSJPEGProcess1012TransferSyntax = 9
	/// JPEG Full Progression, Non-Hierarchical (lossy, 8/12 bit), arithmetic coding
	EXSJPEGProcess1113TransferSyntax = 10
	/// JPEG Lossless with any selection value
	EXSJPEGProcess14TransferSyntax = 11
	/// JPEG Lossless with any selection value, arithmetic coding
	EXSJPEGProcess15TransferSyntax = 12
	/// JPEG Extended Sequential, Hierarchical (lossy, 8/12 bit)
	EXSJPEGProcess1618TransferSyntax = 13
	/// JPEG Extended Sequential, Hierarchical (lossy, 8/12 bit), arithmetic coding
	EXSJPEGProcess1719TransferSyntax = 14
	/// JPEG Spectral Selection, Hierarchical (lossy, 8/12 bit)
	EXSJPEGProcess2022TransferSyntax = 15
	/// JPEG Spectral Selection, Hierarchical (lossy, 8/12 bit), arithmetic coding
	EXSJPEGProcess2123TransferSyntax = 16
	/// JPEG Full Progression, Hierarchical (lossy, 8/12 bit)
	EXSJPEGProcess2426TransferSyntax = 17
	/// JPEG Full Progression, Hierarchical (lossy, 8/12 bit), arithmetic coding
	EXSJPEGProcess2527TransferSyntax = 18
	/// JPEG Lossless, Hierarchical
	EXSJPEGProcess28TransferSyntax = 19
	/// JPEG Lossless, Hierarchical, arithmetic coding
	EXSJPEGProcess29TransferSyntax = 20
	/// JPEG Lossless, Selection Value 1
	EXSJPEGProcess14SV1TransferSyntax = 21
	/// Run Length Encoding (lossless)
	EXSRLELossless = 22
	/// JPEG-LS (lossless)
	EXSJPEGLSLossless = 23
	/// JPEG-LS (lossless or near-lossless mode)
	EXSJPEGLSLossy = 24
	/// Deflated Explicit VR Little Endian
	EXSDeflatedLittleEndianExplicit = 25
	/// JPEG 2000 (lossless)
	EXSJPEG2000osslessOnly = 26
	/// JPEG 2000 (lossless or lossy)
	EXSJPEG2000 = 27
	/// MPEG2 Main Profile at Main Level
	EXSMPEG2MainProfileAtMainLevel = 28
	/// MPEG2 Main Profile at High Level
	EXSMPEG2MainProfileAtHighLevel = 29
	/// JPEG 2000 part 2 multi-component extensions (lossless)
	EXSJPEG2000MulticomponentLosslessOnly = 30
	/// JPEG 2000 part 2 multi-component extensions (lossless or lossy)
	EXSJPEG2000Multicomponent = 31
	/// JPIP Referenced
	EXSJPIPReferenced = 32
	/// JPIP Referenced Deflate
	EXSJPIPReferencedDeflate = 33
)

/** enumeration of byte orders
 */
type EByteOrder int

const (
	/// unknown
	EBOunknown = 0

	/// little endian
	EBOLittleEndian = 1

	/// big endian
	EBOBigEndian = 2
)

func (bo EByteOrder) String() string {
	var result string
	switch bo {
	case EBOunknown:
		result = "Unknown"
	case EBOLittleEndian:
		result = "LittleEndian"
	case EBOBigEndian:
		result = "BigEndian"
	}
	return result
}

/** enumeration of VR encoding options
 */
type EVRType int

const (
	/// implicit VR encoding
	EVTImplicit = 0

	/// explicit VR encoding
	EVTExplicit = 1
)

func (t EVRType) String() string {
	var result string
	switch t {
	case EVTImplicit:
		result = "Implicit"
	case EVTExplicit:
		result = "Explicit"
	}
	return result
}

/** enumeration of pixel data encapsulation options
 */
type EJPEGEncapsulated int

const (
	/// pixel data not encapsulated
	EJENotEncapsulated = 0

	/// pixel data encapsulated
	EJEEncapsulated = 1
)

/** enumeration of stream compression techniques
 */
type EStreamCompression int

const (

	/// no stream compression
	ESCnone = 0
	/// unsupported stream compression
	ESCunsupported = 1

	/// zlib stream compression
//	ESCzlib = 2 // not supported
)

// DcmXfer allows for a lookup of Transfer Syntax properties and readable descriptions
type DcmXfer struct {
	// transfer syntax UID
	XferID string

	/// transfer syntax name
	XferName string

	/// transfer syntax enum
	XferSyn ETransferSyntax

	/// transfer syntax byte order
	ByteOrder EByteOrder

	/// transfer syntax VR encoding (implicit/explicit)
	VRType EVRType

	/// transfer syntax encapsulated or native
	Encapsulated EJPEGEncapsulated

	/// 8-bit lossy JPEG process ID for this transfer syntax, 0 if not applicable
	JPEGProcess8 uint32

	/// 12-bit lossy JPEG process ID for this transfer syntax, 0 if not applicable
	JPEGProcess12 uint32

	/// flag indicating whether this transfer syntax has been retired from DICOM
	Retired bool

	/// transfer syntax stream compression type
	StreamCompression EStreamCompression
}

const ERRORXferName = "Unknown Transfer Syntax"

var xferList = []DcmXfer{
	{UIDLittleEndianImplicitTransferSyntax,
		"Little Endian Implicit",
		EXSLittleEndianImplicit,
		EBOLittleEndian,
		EVTImplicit,
		EJENotEncapsulated,
		0, 0,
		false,
		ESCnone},
	{"", // illegal type
		"Virtual Big Endian Implicit",
		EXSBigEndianImplicit,
		EBOBigEndian,
		EVTImplicit,
		EJENotEncapsulated,
		0, 0,
		false,
		ESCnone},
	{UIDLittleEndianExplicitTransferSyntax,
		"Little Endian Explicit",
		EXSLittleEndianExplicit,
		EBOLittleEndian,
		EVTExplicit,
		EJENotEncapsulated,
		0, 0,
		false,
		ESCnone},
	{UIDBigEndianExplicitTransferSyntax, // defined in dctypes.h
		"Big Endian Explicit",
		EXSBigEndianExplicit,
		EBOBigEndian,
		EVTExplicit,
		EJENotEncapsulated,
		0, 0,
		true,
		ESCnone},
	{UIDJPEGProcess1TransferSyntax,
		"JPEG Baseline",
		EXSJPEGProcess1TransferSyntax,
		EBOLittleEndian,
		EVTExplicit,
		EJEEncapsulated,
		1, 1,
		false,
		ESCnone},
	{UIDJPEGProcess24TransferSyntax,
		"JPEG Extended, Process 2+4",
		EXSJPEGProcess24TransferSyntax,
		EBOLittleEndian,
		EVTExplicit,
		EJEEncapsulated,
		2, 4,
		false,
		ESCnone},
	{UIDJPEGProcess35TransferSyntax,
		"JPEG Extended, Process 3+5",
		EXSJPEGProcess35TransferSyntax,
		EBOLittleEndian,
		EVTExplicit,
		EJEEncapsulated,
		3, 5,
		true,
		ESCnone},
	{UIDJPEGProcess68TransferSyntax,
		"JPEG Spectral Selection, Non-hierarchical, Process 6+8",
		EXSJPEGProcess68TransferSyntax,
		EBOLittleEndian,
		EVTExplicit,
		EJEEncapsulated,
		6, 8,
		true,
		ESCnone},
	{UIDJPEGProcess79TransferSyntax,
		"JPEG Spectral Selection, Non-hierarchical, Process 7+9",
		EXSJPEGProcess79TransferSyntax,
		EBOLittleEndian,
		EVTExplicit,
		EJEEncapsulated,
		7, 9,
		true,
		ESCnone},
	{UIDJPEGProcess1012TransferSyntax,
		"JPEG Full Progression, Non-hierarchical, Process 10+12",
		EXSJPEGProcess1012TransferSyntax,
		EBOLittleEndian,
		EVTExplicit,
		EJEEncapsulated,
		10, 12,
		true,
		ESCnone},
	{UIDJPEGProcess1113TransferSyntax,
		"JPEG Full Progression, Non-hierarchical, Process 11+13",
		EXSJPEGProcess1113TransferSyntax,
		EBOLittleEndian,
		EVTExplicit,
		EJEEncapsulated,
		11, 13,
		true,
		ESCnone},
	{UIDJPEGProcess14TransferSyntax,
		"JPEG Lossless, Non-hierarchical, Process 14",
		EXSJPEGProcess14TransferSyntax,
		EBOLittleEndian,
		EVTExplicit,
		EJEEncapsulated,
		14, 14,
		false,
		ESCnone},
	{UIDJPEGProcess15TransferSyntax,
		"JPEG Lossless, Non-hierarchical, Process 15",
		EXSJPEGProcess15TransferSyntax,
		EBOLittleEndian,
		EVTExplicit,
		EJEEncapsulated,
		15, 15,
		true,
		ESCnone},
	{UIDJPEGProcess1618TransferSyntax,
		"JPEG Extended, Hierarchical, Process 16+18",
		EXSJPEGProcess1618TransferSyntax,
		EBOLittleEndian,
		EVTExplicit,
		EJEEncapsulated,
		16, 18,
		true,
		ESCnone},
	{UIDJPEGProcess1719TransferSyntax,
		"JPEG Extended, Hierarchical, Process 17+19",
		EXSJPEGProcess1719TransferSyntax,
		EBOLittleEndian,
		EVTExplicit,
		EJEEncapsulated,
		17, 19,
		true,
		ESCnone},
	{UIDJPEGProcess2022TransferSyntax,
		"JPEG Spectral Selection, Hierarchical, Process 20+22",
		EXSJPEGProcess2022TransferSyntax,
		EBOLittleEndian,
		EVTExplicit,
		EJEEncapsulated,
		20, 22,
		true,
		ESCnone},
	{UIDJPEGProcess2123TransferSyntax,
		"JPEG Spectral Selection, Hierarchical, Process 21+23",
		EXSJPEGProcess2123TransferSyntax,
		EBOLittleEndian,
		EVTExplicit,
		EJEEncapsulated,
		21, 23,
		true,
		ESCnone},
	{UIDJPEGProcess2426TransferSyntax,
		"JPEG Full Progression, Hierarchical, Process 24+26",
		EXSJPEGProcess2426TransferSyntax,
		EBOLittleEndian,
		EVTExplicit,
		EJEEncapsulated,
		24, 26,
		true,
		ESCnone},
	{UIDJPEGProcess2527TransferSyntax,
		"JPEG Full Progression, Hierarchical, Process 25+27",
		EXSJPEGProcess2527TransferSyntax,
		EBOLittleEndian,
		EVTExplicit,
		EJEEncapsulated,
		25, 27,
		true,
		ESCnone},
	{UIDJPEGProcess28TransferSyntax,
		"JPEG Lossless, Hierarchical, Process 28",
		EXSJPEGProcess28TransferSyntax,
		EBOLittleEndian,
		EVTExplicit,
		EJEEncapsulated,
		28, 28,
		true,
		ESCnone},
	{UIDJPEGProcess29TransferSyntax,
		"JPEG Lossless, Hierarchical, Process 29",
		EXSJPEGProcess29TransferSyntax,
		EBOLittleEndian,
		EVTExplicit,
		EJEEncapsulated,
		29, 29,
		true,
		ESCnone},
	{UIDJPEGProcess14SV1TransferSyntax,
		"JPEG Lossless, Non-hierarchical, 1st Order Prediction",
		EXSJPEGProcess14SV1TransferSyntax,
		EBOLittleEndian,
		EVTExplicit,
		EJEEncapsulated,
		14, 14,
		false,
		ESCnone},
	{UIDRLELosslessTransferSyntax,
		"RLE Lossless",
		EXSRLELossless,
		EBOLittleEndian,
		EVTExplicit,
		EJEEncapsulated,
		0, 0,
		false,
		ESCnone},
	{UIDJPEGLSLosslessTransferSyntax,
		"JPEG-LS Lossless",
		EXSJPEGLSLossless,
		EBOLittleEndian,
		EVTExplicit,
		EJEEncapsulated,
		0, 0,
		false,
		ESCnone},
	{UIDJPEGLSLossyTransferSyntax,
		"JPEG-LS Lossy (Near-lossless)",
		EXSJPEGLSLossy,
		EBOLittleEndian,
		EVTExplicit,
		EJEEncapsulated,
		0, 0,
		false,
		ESCnone},
	{UIDDeflatedExplicitVRLittleEndianTransferSyntax,
		"Deflated Explicit VR Little Endian",
		EXSDeflatedLittleEndianExplicit,
		EBOLittleEndian,
		EVTExplicit,
		EJENotEncapsulated,
		0, 0,
		false,
		ESCunsupported},
	{UIDJPEG2000LosslessOnlyTransferSyntax,
		"JPEG 2000 (Lossless only)",
		EXSJPEG2000osslessOnly,
		EBOLittleEndian,
		EVTExplicit,
		EJEEncapsulated,
		0, 0,
		false,
		ESCnone},
	{UIDJPEG2000TransferSyntax,
		"JPEG 2000 (Lossless or Lossy)",
		EXSJPEG2000,
		EBOLittleEndian,
		EVTExplicit,
		EJEEncapsulated,
		0, 0,
		false,
		ESCnone},
	{UIDMPEG2MainProfileAtMainLevelTransferSyntax,
		"MPEG2 Main Profile @ Main Level",
		EXSMPEG2MainProfileAtMainLevel,
		EBOLittleEndian,
		EVTExplicit,
		EJEEncapsulated,
		0, 0,
		false,
		ESCnone},
	{UIDMPEG2MainProfileAtHighLevelTransferSyntax,
		"MPEG2 Main Profile @ High Level",
		EXSMPEG2MainProfileAtHighLevel,
		EBOLittleEndian,
		EVTExplicit,
		EJEEncapsulated,
		0, 0,
		false,
		ESCnone},
	{UIDJPEG2000Part2MulticomponentImageCompressionLosslessOnlyTransferSyntax,
		"JPEG 2000 Part 2 Multicomponent Image Compression (Lossless only)",
		EXSJPEG2000MulticomponentLosslessOnly,
		EBOLittleEndian,
		EVTExplicit,
		EJEEncapsulated,
		0, 0,
		false,
		ESCnone},
	{UIDJPEG2000Part2MulticomponentImageCompressionTransferSyntax,
		"JPEG 2000 Part 2 Multicomponent Image Compression (Lossless or Lossy)",
		EXSJPEG2000Multicomponent,
		EBOLittleEndian,
		EVTExplicit,
		EJEEncapsulated,
		0, 0,
		false,
		ESCnone},
	{UIDJPIPReferencedTransferSyntax,
		"JPIP Referenced",
		EXSJPIPReferenced,
		EBOLittleEndian,
		EVTExplicit,
		EJENotEncapsulated, // in fact, pixel data shall be referenced via (0028,7FE0) Pixel Data Provider URL
		0, 0,
		false,
		ESCnone},
	{UIDJPIPReferencedDeflateTransferSyntax,
		"JPIP Referenced Deflate",
		EXSJPIPReferencedDeflate,
		EBOLittleEndian,
		EVTExplicit,
		EJENotEncapsulated, // in fact, pixel data shall be referenced via (0028,7FE0) Pixel Data Provider URL
		0, 0,
		false,
		ESCunsupported},
}

// NewDcmXfer returns an new instance of DcmXfer
func NewDcmXfer(xfer ETransferSyntax) *DcmXfer {
	var f DcmXfer

	f.XferName = ERRORXferName
	f.XferSyn = EXSUnknown
	f.ByteOrder = EBOunknown
	f.VRType = EVTImplicit
	f.Encapsulated = EJENotEncapsulated
	f.JPEGProcess8 = 0
	f.JPEGProcess12 = 0
	f.Retired = false
	f.StreamCompression = ESCnone

	for _, v := range xferList {
		if v.XferSyn == xfer {
			f = v
			break
		}
	}
	return &f
}

// GetDcmXferByID get the transfer syntax by id.
func (xfer *DcmXfer) GetDcmXferByID() error {
	for _, v := range xferList {
		if v.XferID == xfer.XferID {
			*xfer = v
			return nil
		}
	}
	err := fmt.Sprintf("%s : unknown  transfer  syntax", xfer.XferID)
	return errors.New(err)
}

// IsExplicitVR returns true if transfer syntax is explicit VR, false otherwise
func (xfer *DcmXfer) IsExplicitVR() bool {
	return xfer.VRType == EVTExplicit
}
