package core

/*
** Defined Transfer Syntax UIDs
 */
var (
	// UIDLittleEndianImplicitTransferSyntax : Implicit VR Little Endian: Default Transfer Syntax for DICOM
	UIDLittleEndianImplicitTransferSyntax = "1.2.840.10008.1.2"

	// UIDLittleEndianExplicitTransferSyntax : Explicit VR Little Endian
	UIDLittleEndianExplicitTransferSyntax = "1.2.840.10008.1.2.1"

	// UIDBigEndianExplicitTransferSyntax Explicit VR Big Endian
	UIDBigEndianExplicitTransferSyntax = "1.2.840.10008.1.2.2"

	// UIDDeflatedExplicitVRLittleEndianTransferSyntax :　Deflated Explicit VR Little Endian
	UIDDeflatedExplicitVRLittleEndianTransferSyntax = "1.2.840.10008.1.2.1.99"

	// UIDJPEGProcess1TransferSyntax ：JPEG Baseline (Process 1): Default Transfer Syntax
	// for Lossy JPEG 8 Bit Image Compression
	UIDJPEGProcess1TransferSyntax = "1.2.840.10008.1.2.4.50"

	// UIDJPEGProcess24TransferSyntax : JPEG Extended (Process 2 & 4): Default Transfer Syntax
	//  for Lossy JPEG 12 Bit Image Compression (Process 4 only)
	UIDJPEGProcess24TransferSyntax = "1.2.840.10008.1.2.4.51"

	// UIDJPEGProcess35TransferSyntax : JPEG Extended (Process 3 & 5) - RETIRED
	UIDJPEGProcess35TransferSyntax = "1.2.840.10008.1.2.4.52"

	// UIDJPEGProcess68TransferSyntax : JPEG Spectral Selection, Non-Hierarchical (Process 6 & 8) - RETIRED
	UIDJPEGProcess68TransferSyntax = "1.2.840.10008.1.2.4.53"

	// UIDJPEGProcess79TransferSyntax : JPEG Spectral Selection, Non-Hierarchical (Process 7 & 9) - RETIRED
	UIDJPEGProcess79TransferSyntax = "1.2.840.10008.1.2.4.54"

	// UIDJPEGProcess1012TransferSyntax : JPEG Full Progression, Non-Hierarchical (Process 10 & 12) - RETIRED
	UIDJPEGProcess1012TransferSyntax = "1.2.840.10008.1.2.4.55"

	// UIDJPEGProcess1113TransferSyntax : JPEG Full Progression, Non-Hierarchical (Process 11 & 13) - RETIRED
	UIDJPEGProcess1113TransferSyntax = "1.2.840.10008.1.2.4.56"

	// UIDJPEGProcess14TransferSyntax : JPEG Lossless, Non-Hierarchical (Process 14)
	UIDJPEGProcess14TransferSyntax = "1.2.840.10008.1.2.4.57"

	// UIDJPEGProcess15TransferSyntax : JPEG Lossless, Non-Hierarchical (Process 15) - RETIRED
	UIDJPEGProcess15TransferSyntax = "1.2.840.10008.1.2.4.58"

	// UIDJPEGProcess1618TransferSyntax : JPEG Extended, Hierarchical (Process 16 & 18) - RETIRED
	UIDJPEGProcess1618TransferSyntax = "1.2.840.10008.1.2.4.59"

	// UIDJPEGProcess1719TransferSyntax : JPEG Extended, Hierarchical (Process 17 & 19) - RETIRED
	UIDJPEGProcess1719TransferSyntax = "1.2.840.10008.1.2.4.60"

	// UIDJPEGProcess2022TransferSyntax : JPEG Spectral Selection, Hierarchical (Process 20 & 22) - RETIRED
	UIDJPEGProcess2022TransferSyntax = "1.2.840.10008.1.2.4.61"

	// UIDJPEGProcess2123TransferSyntax : JPEG Spectral Selection, Hierarchical (Process 21 & 23) - RETIRED
	UIDJPEGProcess2123TransferSyntax = "1.2.840.10008.1.2.4.62"

	// UIDJPEGProcess2426TransferSyntax : JPEG Full Progression, Hierarchical (Process 24 & 26) - RETIRED
	UIDJPEGProcess2426TransferSyntax = "1.2.840.10008.1.2.4.63"

	// UIDJPEGProcess2527TransferSyntax : JPEG Full Progression, Hierarchical (Process 25 & 27) - RETIRED
	UIDJPEGProcess2527TransferSyntax = "1.2.840.10008.1.2.4.64"

	// UIDJPEGProcess28TransferSyntax : JPEG Lossless, Hierarchical (Process 28) - RETIRED
	UIDJPEGProcess28TransferSyntax = "1.2.840.10008.1.2.4.65"

	// UIDJPEGProcess29TransferSyntax :  JPEG Lossless, Hierarchical (Process 29) - RETIRED
	UIDJPEGProcess29TransferSyntax = "1.2.840.10008.1.2.4.66"

	// UIDJPEGProcess14SV1TransferSyntax : JPEG Lossless, Non-Hierarchical, First-Order Prediction (Process 14
	//  [Selection Value 1]): Default Transfer Syntax for Lossless JPEG Image Compression
	UIDJPEGProcess14SV1TransferSyntax = "1.2.840.10008.1.2.4.70"

	// UIDJPEGLSLosslessTransferSyntax : JPEG-LS Lossless Image Compression
	UIDJPEGLSLosslessTransferSyntax = "1.2.840.10008.1.2.4.80"

	// UIDJPEGLSLossyTransferSyntax :  JPEG-LS Lossy (Near-Lossless) Image Compression
	UIDJPEGLSLossyTransferSyntax = "1.2.840.10008.1.2.4.81"

	// UIDJPEG2000LosslessOnlyTransferSyntax :  JPEG 2000 Image Compression (Lossless Only)
	UIDJPEG2000LosslessOnlyTransferSyntax = "1.2.840.10008.1.2.4.90"

	// UIDJPEG2000TransferSyntax :  JPEG 2000 Image Compression (Lossless or Lossy)
	UIDJPEG2000TransferSyntax = "1.2.840.10008.1.2.4.91"

	// UIDJPEG2000Part2MulticomponentImageCompressionLosslessOnlyTransferSyntax : JPEG 2000 Part 2 Multi-component Image Compression (Lossless Only)
	UIDJPEG2000Part2MulticomponentImageCompressionLosslessOnlyTransferSyntax = "1.2.840.10008.1.2.4.92"

	// UIDJPEG2000Part2MulticomponentImageCompressionTransferSyntax :  JPEG 2000 Part 2 Multi-component Image Compression (Lossless or Lossy)
	UIDJPEG2000Part2MulticomponentImageCompressionTransferSyntax = "1.2.840.10008.1.2.4.93"

	// UIDJPIPReferencedTransferSyntax :  JPIP Referenced
	UIDJPIPReferencedTransferSyntax = "1.2.840.10008.1.2.4.94"

	// UIDJPIPReferencedDeflateTransferSyntax :  JPIP Referenced Deflate
	UIDJPIPReferencedDeflateTransferSyntax = "1.2.840.10008.1.2.4.95"

	// UIDMPEG2MainProfileAtMainLevelTransferSyntax :  MPEG2 Main Profile @ Main Level
	UIDMPEG2MainProfileAtMainLevelTransferSyntax = "1.2.840.10008.1.2.4.100"

	// UIDMPEG2MainProfileAtHighLevelTransferSyntax :  MPEG2 Main Profile @ High Level
	UIDMPEG2MainProfileAtHighLevelTransferSyntax = "1.2.840.10008.1.2.4.101"

	// UIDRLELosslessTransferSyntax :  RLE Lossless
	UIDRLELosslessTransferSyntax = "1.2.840.10008.1.2.5"
)
