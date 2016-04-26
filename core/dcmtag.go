package core

import "fmt"

// DICOM File Meta Information
var (
	// TagFileMetaInformationGroupLength  Number of bytes following this File Meta Element (end of the Value field) up to and including the last File Meta Element of the Group 2 File Meta Information
	TagFileMetaInformationGroupLength = DcmTag{0x0002, 0x0000}
	// TagFileMetaInformationVersion This is a two byte field where each bit identifies a version of this File Meta Information header. In version 1 the first byte value is 00H and the second value byte value is 01H.
	TagFileMetaInformationVersion = DcmTag{0x0002, 0x0001}
	// TagMediaStorageSOPClassUID  Uniquely identifies the SOP Class associated with the Data Set.
	TagMediaStorageSOPClassUID = DcmTag{0x0002, 0x0002}
	// TagMediaStorageSOPInstanceUID Uniquely identifies the SOP Instance associated with the Data Set placed in the file and following the File Meta Information.
	TagMediaStorageSOPInstanceUID = DcmTag{0x0002, 0x0003}
	// TagTransferSyntaxUID Uniquely identifies the Transfer Syntax used to encode the following Data Set.
	TagTransferSyntaxUID = DcmTag{0x0002, 0x0010}
	// TagImplementationClassUID Uniquely identifies the implementation that wrote this file and its content.
	TagImplementationClassUID = DcmTag{0x0002, 0x0012}
	// TagImplementationVersionName Identifies a version for an Implementation Class UID (0002,0012) using up to 16 characters of the repertoire
	TagImplementationVersionName = DcmTag{0x0002, 0x0013}
	// TagSourceApplicationEntityTitle The DICOM Application Entity (AE) Title of the AE that wrote this file's content (or last updated it).
	TagSourceApplicationEntityTitle = DcmTag{0x0002, 0x0016}
	// TagSendingApplicationEntityTitle The DICOM Application Entity (AE) Title of the AE that sent this file's content over a network.
	TagSendingApplicationEntityTitle = DcmTag{0x0002, 0x0017}
	// TagReceivingApplicationEntityTitle The DICOM Application Entity (AE) Title of the AE that received this file's content over a network.
	TagReceivingApplicationEntityTitle = DcmTag{0x0002, 0x0018}
	// TagPrivateInformationCreatorUID The UID of the creator of the private information (0002,0102).
	TagPrivateInformationCreatorUID = DcmTag{0x0002, 0x0100}
	// TagPrivateInformation Contains Private Information placed in the File Meta Information.
	TagPrivateInformation = DcmTag{0x0002, 0x0102}
)

// DICOM Sequence Items
var (
	DCMItem                     = DcmTag{0xfffe, 0xe000}
	DCMItemDelimitationItem     = DcmTag{0xfffe, 0xe00d}
	DCMSequenceDelimitationItem = DcmTag{0xfffe, 0xe0dd}
)

// DcmTag contains the group and element of the dicom element.
type DcmTag struct {
	Group   uint16
	Element uint16
}

// String is to convert tag to string
func (t DcmTag) String() string {
	return fmt.Sprintf("0x%04x%04x", t.Group, t.Element)
}
