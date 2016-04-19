package core

import "fmt"

// DICOM File Meta Information
var (
	// FileMetaInformationGroupLength  Number of bytes following this File Meta Element (end of the Value field) up to and including the last File Meta Element of the Group 2 File Meta Information
	FileMetaInformationGroupLength = DcmTag{0x0002, 0x0000}
	// FileMetaInformationVersion This is a two byte field where each bit identifies a version of this File Meta Information header. In version 1 the first byte value is 00H and the second value byte value is 01H.
	FileMetaInformationVersion = DcmTag{0x0002, 0x0001}
	// MediaStorageSOPClassUID  Uniquely identifies the SOP Class associated with the Data Set.
	MediaStorageSOPClassUID = DcmTag{0x0002, 0x0002}
	// MediaStorageSOPInstanceUID Uniquely identifies the SOP Instance associated with the Data Set placed in the file and following the File Meta Information.
	MediaStorageSOPInstanceUID = DcmTag{0x0002, 0x0003}
	// TransferSyntaxUID Uniquely identifies the Transfer Syntax used to encode the following Data Set.
	TransferSyntaxUID = DcmTag{0x0002, 0x0010}
	// ImplementationClassUID Uniquely identifies the implementation that wrote this file and its content.
	ImplementationClassUID = DcmTag{0x0002, 0x0012}
	// ImplementationVersionName Identifies a version for an Implementation Class UID (0002,0012) using up to 16 characters of the repertoire
	ImplementationVersionName = DcmTag{0x0002, 0x0013}
	// SourceApplicationEntityTitle The DICOM Application Entity (AE) Title of the AE that wrote this file's content (or last updated it).
	SourceApplicationEntityTitle = DcmTag{0x0002, 0x0016}
	// SendingApplicationEntityTitle The DICOM Application Entity (AE) Title of the AE that sent this file's content over a network.
	SendingApplicationEntityTitle = DcmTag{0x0002, 0x0017}
	// ReceivingApplicationEntityTitle The DICOM Application Entity (AE) Title of the AE that received this file's content over a network.
	ReceivingApplicationEntityTitle = DcmTag{0x0002, 0x0018}
	// PrivateInformationCreatorUID The UID of the creator of the private information (0002,0102).
	PrivateInformationCreatorUID = DcmTag{0x0002, 0x0100}
	// PrivateInformation Contains Private Information placed in the File Meta Information.
	PrivateInformation = DcmTag{0x0002, 0x0102}
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
