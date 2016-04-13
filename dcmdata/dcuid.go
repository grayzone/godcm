package dcmdata

/// implementation version name for this version of the toolkit
const OFFIS_DTK_IMPLEMENTATION_VERSION_NAME = "OFFIS_DCMTK_GODCM"

/// implementation version name for this version of the toolkit, used for files received in "bit preserving" mode
const OFFIS_DTK_IMPLEMENTATION_VERSION_NAME2 = "OFFIS_DCMBP_GODCM"

/// UID root for DCMTK, registered for OFFIS with DIN in Germany
const OFFIS_UID_ROOT = "1.2.276.0.7230010.3"

/// UID root to be used when generating UIDs. By default uses the DCMTK root, but can be replaced at compile time.
const SITE_UID_ROOT = OFFIS_UID_ROOT /* default */

/*
** Useful UID prefixes. These can be whatever you want.
**
** These site UIDs are arbitary, non-standard, with no meaning
** and can be changed at any time.  Do _not_ rely on these values.
** Do _not_ assume any semantics when using these suffixes.
**
 */

/// UID root for study instance UIDs
const SITE_STUDY_UID_ROOT = SITE_UID_ROOT + ".1.2"

/// UID root for series instance UIDs
const SITE_SERIES_UID_ROOT = SITE_UID_ROOT + ".1.3"

/// UID root for SOP instance UIDs
const SITE_INSTANCE_UID_ROOT = SITE_UID_ROOT + ".1.4"

/** A private SOP Class UID which is used in a file meta-header when
 *  no real SOP Class is stored in the file. -- NON-STANDARD
 */
const UID_PrivateGenericFileSOPClass = SITE_UID_ROOT + ".1.0.1"

/// DICOM Defined Standard Application Context UID
const UID_StandardApplicationContext = "1.2.840.10008.3.1.1.1"

/*
** Defined Transfer Syntax UIDs
 */

/// Implicit VR Little Endian: Default Transfer Syntax for DICOM
const UID_LittleEndianImplicitTransferSyntax = "1.2.840.10008.1.2"

/// Explicit VR Little Endian
const UID_LittleEndianExplicitTransferSyntax = "1.2.840.10008.1.2.1"

/// Explicit VR Big Endian
const UID_BigEndianExplicitTransferSyntax = "1.2.840.10008.1.2.2"

/// Deflated Explicit VR Little Endian
const UID_DeflatedExplicitVRLittleEndianTransferSyntax = "1.2.840.10008.1.2.1.99"

/** JPEG Baseline (Process 1): Default Transfer Syntax
 * for Lossy JPEG 8 Bit Image Compression
 */
const UID_JPEGProcess1TransferSyntax = "1.2.840.10008.1.2.4.50"

/** JPEG Extended (Process 2 & 4): Default Transfer Syntax
 *  for Lossy JPEG 12 Bit Image Compression (Process 4 only)
 */
const UID_JPEGProcess2_4TransferSyntax = "1.2.840.10008.1.2.4.51"

/// JPEG Extended (Process 3 & 5) - RETIRED
const UID_JPEGProcess3_5TransferSyntax = "1.2.840.10008.1.2.4.52"

/// JPEG Spectral Selection, Non-Hierarchical (Process 6 & 8) - RETIRED
const UID_JPEGProcess6_8TransferSyntax = "1.2.840.10008.1.2.4.53"

/// JPEG Spectral Selection, Non-Hierarchical (Process 7 & 9) - RETIRED
const UID_JPEGProcess7_9TransferSyntax = "1.2.840.10008.1.2.4.54"

/// JPEG Full Progression, Non-Hierarchical (Process 10 & 12) - RETIRED
const UID_JPEGProcess10_12TransferSyntax = "1.2.840.10008.1.2.4.55"

/// JPEG Full Progression, Non-Hierarchical (Process 11 & 13) - RETIRED
const UID_JPEGProcess11_13TransferSyntax = "1.2.840.10008.1.2.4.56"

/// JPEG Lossless, Non-Hierarchical (Process 14)
const UID_JPEGProcess14TransferSyntax = "1.2.840.10008.1.2.4.57"

/// JPEG Lossless, Non-Hierarchical (Process 15) - RETIRED
const UID_JPEGProcess15TransferSyntax = "1.2.840.10008.1.2.4.58"

/// JPEG Extended, Hierarchical (Process 16 & 18) - RETIRED
const UID_JPEGProcess16_18TransferSyntax = "1.2.840.10008.1.2.4.59"

/// JPEG Extended, Hierarchical (Process 17 & 19) - RETIRED
const UID_JPEGProcess17_19TransferSyntax = "1.2.840.10008.1.2.4.60"

/// JPEG Spectral Selection, Hierarchical (Process 20 & 22) - RETIRED
const UID_JPEGProcess20_22TransferSyntax = "1.2.840.10008.1.2.4.61"

/// JPEG Spectral Selection, Hierarchical (Process 21 & 23) - RETIRED
const UID_JPEGProcess21_23TransferSyntax = "1.2.840.10008.1.2.4.62"

/// JPEG Full Progression, Hierarchical (Process 24 & 26) - RETIRED
const UID_JPEGProcess24_26TransferSyntax = "1.2.840.10008.1.2.4.63"

/// JPEG Full Progression, Hierarchical (Process 25 & 27) - RETIRED
const UID_JPEGProcess25_27TransferSyntax = "1.2.840.10008.1.2.4.64"

/// JPEG Lossless, Hierarchical (Process 28) - RETIRED
const UID_JPEGProcess28TransferSyntax = "1.2.840.10008.1.2.4.65"

/// JPEG Lossless, Hierarchical (Process 29) - RETIRED
const UID_JPEGProcess29TransferSyntax = "1.2.840.10008.1.2.4.66"

/** JPEG Lossless, Non-Hierarchical, First-Order Prediction (Process 14
 *  [Selection Value 1]): Default Transfer Syntax for Lossless JPEG Image Compression
 */
const UID_JPEGProcess14SV1TransferSyntax = "1.2.840.10008.1.2.4.70"

/// JPEG-LS Lossless Image Compression
const UID_JPEGLSLosslessTransferSyntax = "1.2.840.10008.1.2.4.80"

/// JPEG-LS Lossy (Near-Lossless) Image Compression
const UID_JPEGLSLossyTransferSyntax = "1.2.840.10008.1.2.4.81"

/// JPEG 2000 Image Compression (Lossless Only)
const UID_JPEG2000LosslessOnlyTransferSyntax = "1.2.840.10008.1.2.4.90"

/// JPEG 2000 Image Compression (Lossless or Lossy)
const UID_JPEG2000TransferSyntax = "1.2.840.10008.1.2.4.91"

/// JPEG 2000 Part 2 Multi-component Image Compression (Lossless Only)
const UID_JPEG2000Part2MulticomponentImageCompressionLosslessOnlyTransferSyntax = "1.2.840.10008.1.2.4.92"

/// JPEG 2000 Part 2 Multi-component Image Compression (Lossless or Lossy)
const UID_JPEG2000Part2MulticomponentImageCompressionTransferSyntax = "1.2.840.10008.1.2.4.93"

/// JPIP Referenced
const UID_JPIPReferencedTransferSyntax = "1.2.840.10008.1.2.4.94"

/// JPIP Referenced Deflate
const UID_JPIPReferencedDeflateTransferSyntax = "1.2.840.10008.1.2.4.95"

/// MPEG2 Main Profile @ Main Level
const UID_MPEG2MainProfileAtMainLevelTransferSyntax = "1.2.840.10008.1.2.4.100"

/// MPEG2 Main Profile @ High Level
const UID_MPEG2MainProfileAtHighLevelTransferSyntax = "1.2.840.10008.1.2.4.101"

/// RLE Lossless
const UID_RLELosslessTransferSyntax = "1.2.840.10008.1.2.5"

/** MIME encapsulation (Supplement 101) is only a pseudo transfer syntax used to
 *  refer to MIME encapsulated HL7 CDA documents from a DICOMDIR when stored
 *  on a DICOM storage medium. It is never used for network communication
 *  or encoding of DICOM objects.
 */
const UID_RFC2557MIMEEncapsulationTransferSyntax = "1.2.840.10008.1.2.6.1"

/** XML encoding (Supplement 114) is only a pseudo transfer syntax used to refer to
 *  encapsulated HL7 CDA documents from a DICOMDIR when stored on a DICOM storage
 *  medium. It is never used for network communication or encoding of DICOM objects.
 */
const UID_XMLEncodingTransferSyntax = "1.2.840.10008.1.2.6.2"

/*
** Defined SOP Class UIDs according to DICOM standard
 */

// Storage
const UID_RETIRED_StoredPrintStorage = "1.2.840.10008.5.1.1.27"
const UID_RETIRED_HardcopyGrayscaleImageStorage = "1.2.840.10008.5.1.1.29"
const UID_RETIRED_HardcopyColorImageStorage = "1.2.840.10008.5.1.1.30"
const UID_ComputedRadiographyImageStorage = "1.2.840.10008.5.1.4.1.1.1"
const UID_DigitalXRayImageStorageForPresentation = "1.2.840.10008.5.1.4.1.1.1.1"
const UID_DigitalXRayImageStorageForProcessing = "1.2.840.10008.5.1.4.1.1.1.1.1"
const UID_DigitalMammographyXRayImageStorageForPresentation = "1.2.840.10008.5.1.4.1.1.1.2"
const UID_DigitalMammographyXRayImageStorageForProcessing = "1.2.840.10008.5.1.4.1.1.1.2.1"
const UID_DigitalIntraOralXRayImageStorageForPresentation = "1.2.840.10008.5.1.4.1.1.1.3"
const UID_DigitalIntraOralXRayImageStorageForProcessing = "1.2.840.10008.5.1.4.1.1.1.3.1"
const UID_CTImageStorage = "1.2.840.10008.5.1.4.1.1.2"
const UID_EnhancedCTImageStorage = "1.2.840.10008.5.1.4.1.1.2.1"
const UID_RETIRED_UltrasoundMultiframeImageStorage = "1.2.840.10008.5.1.4.1.1.3"
const UID_UltrasoundMultiframeImageStorage = "1.2.840.10008.5.1.4.1.1.3.1"
const UID_MRImageStorage = "1.2.840.10008.5.1.4.1.1.4"
const UID_EnhancedMRImageStorage = "1.2.840.10008.5.1.4.1.1.4.1"
const UID_MRSpectroscopyStorage = "1.2.840.10008.5.1.4.1.1.4.2"
const UID_EnhancedMRColorImageStorage = "1.2.840.10008.5.1.4.1.1.4.3"
const UID_RETIRED_NuclearMedicineImageStorage = "1.2.840.10008.5.1.4.1.1.5"
const UID_RETIRED_UltrasoundImageStorage = "1.2.840.10008.5.1.4.1.1.6"
const UID_UltrasoundImageStorage = "1.2.840.10008.5.1.4.1.1.6.1"
const UID_EnhancedUSVolumeStorage = "1.2.840.10008.5.1.4.1.1.6.2"
const UID_SecondaryCaptureImageStorage = "1.2.840.10008.5.1.4.1.1.7"
const UID_MultiframeSingleBitSecondaryCaptureImageStorage = "1.2.840.10008.5.1.4.1.1.7.1"
const UID_MultiframeGrayscaleByteSecondaryCaptureImageStorage = "1.2.840.10008.5.1.4.1.1.7.2"
const UID_MultiframeGrayscaleWordSecondaryCaptureImageStorage = "1.2.840.10008.5.1.4.1.1.7.3"
const UID_MultiframeTrueColorSecondaryCaptureImageStorage = "1.2.840.10008.5.1.4.1.1.7.4"
const UID_RETIRED_StandaloneOverlayStorage = "1.2.840.10008.5.1.4.1.1.8"
const UID_RETIRED_StandaloneCurveStorage = "1.2.840.10008.5.1.4.1.1.9"
const UID_TwelveLeadECGWaveformStorage = "1.2.840.10008.5.1.4.1.1.9.1.1"
const UID_GeneralECGWaveformStorage = "1.2.840.10008.5.1.4.1.1.9.1.2"
const UID_AmbulatoryECGWaveformStorage = "1.2.840.10008.5.1.4.1.1.9.1.3"
const UID_HemodynamicWaveformStorage = "1.2.840.10008.5.1.4.1.1.9.2.1"
const UID_CardiacElectrophysiologyWaveformStorage = "1.2.840.10008.5.1.4.1.1.9.3.1"
const UID_BasicVoiceAudioWaveformStorage = "1.2.840.10008.5.1.4.1.1.9.4.1"
const UID_GeneralAudioWaveformStorage = "1.2.840.10008.5.1.4.1.1.9.4.2"
const UID_ArterialPulseWaveformStorage = "1.2.840.10008.5.1.4.1.1.9.5.1"
const UID_RespiratoryWaveformStorage = "1.2.840.10008.5.1.4.1.1.9.6.1"
const UID_RETIRED_StandaloneModalityLUTStorage = "1.2.840.10008.5.1.4.1.1.10"
const UID_RETIRED_StandaloneVOILUTStorage = "1.2.840.10008.5.1.4.1.1.11"
const UID_GrayscaleSoftcopyPresentationStateStorage = "1.2.840.10008.5.1.4.1.1.11.1"
const UID_ColorSoftcopyPresentationStateStorage = "1.2.840.10008.5.1.4.1.1.11.2"
const UID_PseudoColorSoftcopyPresentationStateStorage = "1.2.840.10008.5.1.4.1.1.11.3"
const UID_BlendingSoftcopyPresentationStateStorage = "1.2.840.10008.5.1.4.1.1.11.4"
const UID_XAXRFGrayscaleSoftcopyPresentationStateStorage = "1.2.840.10008.5.1.4.1.1.11.5"
const UID_XRayAngiographicImageStorage = "1.2.840.10008.5.1.4.1.1.12.1"
const UID_EnhancedXAImageStorage = "1.2.840.10008.5.1.4.1.1.12.1.1"
const UID_XRayRadiofluoroscopicImageStorage = "1.2.840.10008.5.1.4.1.1.12.2"
const UID_EnhancedXRFImageStorage = "1.2.840.10008.5.1.4.1.1.12.2.1"
const UID_XRay3DAngiographicImageStorage = "1.2.840.10008.5.1.4.1.1.13.1.1"
const UID_XRay3DCraniofacialImageStorage = "1.2.840.10008.5.1.4.1.1.13.1.2"
const UID_BreastTomosynthesisImageStorage = "1.2.840.10008.5.1.4.1.1.13.1.3"
const UID_RETIRED_XRayAngiographicBiPlaneImageStorage = "1.2.840.10008.5.1.4.1.1.12.3"
const UID_NuclearMedicineImageStorage = "1.2.840.10008.5.1.4.1.1.20"
const UID_RawDataStorage = "1.2.840.10008.5.1.4.1.1.66"
const UID_SpatialRegistrationStorage = "1.2.840.10008.5.1.4.1.1.66.1"
const UID_SpatialFiducialsStorage = "1.2.840.10008.5.1.4.1.1.66.2"
const UID_DeformableSpatialRegistrationStorage = "1.2.840.10008.5.1.4.1.1.66.3"
const UID_SegmentationStorage = "1.2.840.10008.5.1.4.1.1.66.4"
const UID_SurfaceSegmentationStorage = "1.2.840.10008.5.1.4.1.1.66.5"
const UID_RealWorldValueMappingStorage = "1.2.840.10008.5.1.4.1.1.67"
const UID_RETIRED_VLImageStorage = "1.2.840.10008.5.1.4.1.1.77.1"
const UID_VLEndoscopicImageStorage = "1.2.840.10008.5.1.4.1.1.77.1.1"
const UID_VideoEndoscopicImageStorage = "1.2.840.10008.5.1.4.1.1.77.1.1.1"
const UID_VLMicroscopicImageStorage = "1.2.840.10008.5.1.4.1.1.77.1.2"
const UID_VideoMicroscopicImageStorage = "1.2.840.10008.5.1.4.1.1.77.1.2.1"
const UID_VLSlideCoordinatesMicroscopicImageStorage = "1.2.840.10008.5.1.4.1.1.77.1.3"
const UID_VLPhotographicImageStorage = "1.2.840.10008.5.1.4.1.1.77.1.4"
const UID_VideoPhotographicImageStorage = "1.2.840.10008.5.1.4.1.1.77.1.4.1"
const UID_OphthalmicPhotography8BitImageStorage = "1.2.840.10008.5.1.4.1.1.77.1.5.1"
const UID_OphthalmicPhotography16BitImageStorage = "1.2.840.10008.5.1.4.1.1.77.1.5.2"
const UID_StereometricRelationshipStorage = "1.2.840.10008.5.1.4.1.1.77.1.5.3"
const UID_OphthalmicTomographyImageStorage = "1.2.840.10008.5.1.4.1.1.77.1.5.4"
const UID_VLWholeSlideMicroscopyImageStorage = "1.2.840.10008.5.1.4.1.1.77.1.6"
const UID_RETIRED_VLMultiFrameImageStorage = "1.2.840.10008.5.1.4.1.1.77.2"
const UID_LensometryMeasurementsStorage = "1.2.840.10008.5.1.4.1.1.78.1"
const UID_AutorefractionMeasurementsStorage = "1.2.840.10008.5.1.4.1.1.78.2"
const UID_KeratometryMeasurementsStorage = "1.2.840.10008.5.1.4.1.1.78.3"
const UID_SubjectiveRefractionMeasurementsStorage = "1.2.840.10008.5.1.4.1.1.78.4"
const UID_VisualAcuityMeasurementsStorage = "1.2.840.10008.5.1.4.1.1.78.5"
const UID_SpectaclePrescriptionReportStorage = "1.2.840.10008.5.1.4.1.1.78.6"
const UID_OphthalmicAxialMeasurementsStorage = "1.2.840.10008.5.1.4.1.1.78.7"
const UID_IntraocularLensCalculationsStorage = "1.2.840.10008.5.1.4.1.1.78.8"
const UID_MacularGridThicknessAndVolumeReportStorage = "1.2.840.10008.5.1.4.1.1.79.1"
const UID_OphthalmicVisualFieldStaticPerimetryMeasurementsStorage = "1.2.840.10008.5.1.4.1.1.80.1"
const UID_BasicTextSRStorage = "1.2.840.10008.5.1.4.1.1.88.11"
const UID_EnhancedSRStorage = "1.2.840.10008.5.1.4.1.1.88.22"
const UID_ComprehensiveSRStorage = "1.2.840.10008.5.1.4.1.1.88.33"
const UID_ProcedureLogStorage = "1.2.840.10008.5.1.4.1.1.88.40"
const UID_MammographyCADSRStorage = "1.2.840.10008.5.1.4.1.1.88.50"
const UID_KeyObjectSelectionDocumentStorage = "1.2.840.10008.5.1.4.1.1.88.59"
const UID_ChestCADSRStorage = "1.2.840.10008.5.1.4.1.1.88.65"
const UID_XRayRadiationDoseSRStorage = "1.2.840.10008.5.1.4.1.1.88.67"
const UID_ColonCADSRStorage = "1.2.840.10008.5.1.4.1.1.88.69"
const UID_ImplantationPlanSRDocumentStorage = "1.2.840.10008.5.1.4.1.1.88.70"
const UID_EncapsulatedPDFStorage = "1.2.840.10008.5.1.4.1.1.104.1"
const UID_EncapsulatedCDAStorage = "1.2.840.10008.5.1.4.1.1.104.2"
const UID_PositronEmissionTomographyImageStorage = "1.2.840.10008.5.1.4.1.1.128"
const UID_RETIRED_StandalonePETCurveStorage = "1.2.840.10008.5.1.4.1.1.129"
const UID_EnhancedPETImageStorage = "1.2.840.10008.5.1.4.1.1.130"
const UID_BasicStructuredDisplayStorage = "1.2.840.10008.5.1.4.1.1.131"
const UID_RTImageStorage = "1.2.840.10008.5.1.4.1.1.481.1"
const UID_RTDoseStorage = "1.2.840.10008.5.1.4.1.1.481.2"
const UID_RTStructureSetStorage = "1.2.840.10008.5.1.4.1.1.481.3"
const UID_RTBeamsTreatmentRecordStorage = "1.2.840.10008.5.1.4.1.1.481.4"
const UID_RTPlanStorage = "1.2.840.10008.5.1.4.1.1.481.5"
const UID_RTBrachyTreatmentRecordStorage = "1.2.840.10008.5.1.4.1.1.481.6"
const UID_RTTreatmentSummaryRecordStorage = "1.2.840.10008.5.1.4.1.1.481.7"
const UID_RTIonPlanStorage = "1.2.840.10008.5.1.4.1.1.481.8"
const UID_RTIonBeamsTreatmentRecordStorage = "1.2.840.10008.5.1.4.1.1.481.9"
const UID_GenericImplantTemplateStorage = "1.2.840.10008.5.1.4.43.1"
const UID_ImplantAssemblyTemplateStorage = "1.2.840.10008.5.1.4.44.1"
const UID_ImplantTemplateGroupStorage = "1.2.840.10008.5.1.4.45.1"

// DICOMDIR; was UID_BasicDirectoryStorageSOPClass in DCMTK versions prior to 3.5.3
const UID_MediaStorageDirectoryStorage = "1.2.840.10008.1.3.10"

/* Hanging Protocols Storage is a special case because hanging protocols use a different
   information model, i.e. there is no patient, study or series in a hanging protocol IOD. */
const UID_HangingProtocolStorage = "1.2.840.10008.5.1.4.38.1"

// Query/Retrieve
const UID_FINDPatientRootQueryRetrieveInformationModel = "1.2.840.10008.5.1.4.1.2.1.1"
const UID_MOVEPatientRootQueryRetrieveInformationModel = "1.2.840.10008.5.1.4.1.2.1.2"
const UID_GETPatientRootQueryRetrieveInformationModel = "1.2.840.10008.5.1.4.1.2.1.3"
const UID_FINDStudyRootQueryRetrieveInformationModel = "1.2.840.10008.5.1.4.1.2.2.1"
const UID_MOVEStudyRootQueryRetrieveInformationModel = "1.2.840.10008.5.1.4.1.2.2.2"
const UID_GETStudyRootQueryRetrieveInformationModel = "1.2.840.10008.5.1.4.1.2.2.3"
const UID_RETIRED_FINDPatientStudyOnlyQueryRetrieveInformationModel = "1.2.840.10008.5.1.4.1.2.3.1"
const UID_RETIRED_MOVEPatientStudyOnlyQueryRetrieveInformationModel = "1.2.840.10008.5.1.4.1.2.3.2"
const UID_RETIRED_GETPatientStudyOnlyQueryRetrieveInformationModel = "1.2.840.10008.5.1.4.1.2.3.3"
const UID_RETIRED_MOVECompositeInstanceRootRetrieve = "1.2.840.10008.5.1.4.1.2.4.2"
const UID_RETIRED_GETCompositeInstanceRootRetrieve = "1.2.840.10008.5.1.4.1.2.4.3"
const UID_RETIRED_GETCompositeInstanceRetrieveWithoutBulkData = "1.2.840.10008.5.1.4.1.2.5.3"

// Worklist
const UID_FINDModalityWorklistInformationModel = "1.2.840.10008.5.1.4.31"
const UID_FINDGeneralPurposeWorklistInformationModel = "1.2.840.10008.5.1.4.32.1"

// General Purpose Worklist
const UID_GeneralPurposeScheduledProcedureStepSOPClass = "1.2.840.10008.5.1.4.32.2"
const UID_GeneralPurposePerformedProcedureStepSOPClass = "1.2.840.10008.5.1.4.32.3"
const UID_GeneralPurposeWorklistManagementMetaSOPClass = "1.2.840.10008.5.1.4.32"

// MPPS
const UID_ModalityPerformedProcedureStepSOPClass = "1.2.840.10008.3.1.2.3.3"
const UID_ModalityPerformedProcedureStepRetrieveSOPClass = "1.2.840.10008.3.1.2.3.4"
const UID_ModalityPerformedProcedureStepNotificationSOPClass = "1.2.840.10008.3.1.2.3.5"

// Unified Worklist and Procedure Step
const UID_UnifiedWorklistAndProcedureStepServiceClass = "1.2.840.10008.5.1.4.34.4"
const UID_UnifiedProcedureStepPushSOPClass = "1.2.840.10008.5.1.4.34.4.1"
const UID_UnifiedProcedureStepWatchSOPClass = "1.2.840.10008.5.1.4.34.4.2"
const UID_UnifiedProcedureStepPullSOPClass = "1.2.840.10008.5.1.4.34.4.3"
const UID_UnifiedProcedureStepEventSOPClass = "1.2.840.10008.5.1.4.34.4.4"
const UID_UnifiedWorklistAndProcedureStepSOPInstance = "1.2.840.10008.5.1.4.34.5"

// Storage Commitment
const UID_StorageCommitmentPushModelSOPClass = "1.2.840.10008.1.20.1"
const UID_StorageCommitmentPushModelSOPInstance = "1.2.840.10008.1.20.1.1"
const UID_RETIRED_StorageCommitmentPullModelSOPClass = "1.2.840.10008.1.20.2"
const UID_RETIRED_StorageCommitmentPullModelSOPInstance = "1.2.840.10008.1.20.2.1"

// Hanging Protocols
const UID_FINDHangingProtocolInformationModel = "1.2.840.10008.5.1.4.38.2"
const UID_MOVEHangingProtocolInformationModel = "1.2.840.10008.5.1.4.38.3"

// Relevant Patient Information Query
const UID_GeneralRelevantPatientInformationQuery = "1.2.840.10008.5.1.4.37.1"
const UID_BreastImagingRelevantPatientInformationQuery = "1.2.840.10008.5.1.4.37.2"
const UID_CardiacRelevantPatientInformationQuery = "1.2.840.10008.5.1.4.37.3"

// Color Palette Storage and Query/Retrieve
const UID_ColorPaletteStorage = "1.2.840.10008.5.1.4.39.1"
const UID_FINDColorPaletteInformationModel = "1.2.840.10008.5.1.4.39.2"
const UID_MOVEColorPaletteInformationModel = "1.2.840.10008.5.1.4.39.3"
const UID_GETColorPaletteInformationModel = "1.2.840.10008.5.1.4.39.4"

// Implant Template Query/Retrieve
const UID_FINDGenericImplantTemplateInformationModel = "1.2.840.10008.5.1.4.43.2"
const UID_MOVEGenericImplantTemplateInformationModel = "1.2.840.10008.5.1.4.43.3"
const UID_GETGenericImplantTemplateInformationModel = "1.2.840.10008.5.1.4.43.4"
const UID_FINDImplantAssemblyTemplateInformationModel = "1.2.840.10008.5.1.4.44.2"
const UID_MOVEImplantAssemblyTemplateInformationModel = "1.2.840.10008.5.1.4.44.3"
const UID_GETImplantAssemblyTemplateInformationModel = "1.2.840.10008.5.1.4.44.4"
const UID_FINDImplantTemplateGroupInformationModel = "1.2.840.10008.5.1.4.45.2"
const UID_MOVEImplantTemplateGroupInformationModel = "1.2.840.10008.5.1.4.45.3"
const UID_GETImplantTemplateGroupInformationModel = "1.2.840.10008.5.1.4.45.4"

// Print
const UID_BasicFilmSessionSOPClass = "1.2.840.10008.5.1.1.1"
const UID_BasicFilmBoxSOPClass = "1.2.840.10008.5.1.1.2"
const UID_BasicGrayscaleImageBoxSOPClass = "1.2.840.10008.5.1.1.4"
const UID_BasicColorImageBoxSOPClass = "1.2.840.10008.5.1.1.4.1"
const UID_RETIRED_ReferencedImageBoxSOPClass = "1.2.840.10008.5.1.1.4.2"
const UID_BasicGrayscalePrintManagementMetaSOPClass = "1.2.840.10008.5.1.1.9"
const UID_RETIRED_ReferencedGrayscalePrintManagementMetaSOPClass = "1.2.840.10008.5.1.1.9.1"
const UID_PrintJobSOPClass = "1.2.840.10008.5.1.1.14"
const UID_BasicAnnotationBoxSOPClass = "1.2.840.10008.5.1.1.15"
const UID_PrinterSOPClass = "1.2.840.10008.5.1.1.16"
const UID_PrinterConfigurationRetrievalSOPClass = "1.2.840.10008.5.1.1.16.376"
const UID_PrinterSOPInstance = "1.2.840.10008.5.1.1.17"
const UID_PrinterConfigurationRetrievalSOPInstance = "1.2.840.10008.5.1.1.17.376"
const UID_BasicColorPrintManagementMetaSOPClass = "1.2.840.10008.5.1.1.18"
const UID_RETIRED_ReferencedColorPrintManagementMetaSOPClass = "1.2.840.10008.5.1.1.18.1"
const UID_VOILUTBoxSOPClass = "1.2.840.10008.5.1.1.22"
const UID_PresentationLUTSOPClass = "1.2.840.10008.5.1.1.23"
const UID_RETIRED_ImageOverlayBoxSOPClass = "1.2.840.10008.5.1.1.24"
const UID_RETIRED_BasicPrintImageOverlayBoxSOPClass = "1.2.840.10008.5.1.1.24.1"
const UID_RETIRED_PrintQueueSOPInstance = "1.2.840.10008.5.1.1.25"
const UID_RETIRED_PrintQueueManagementSOPClass = "1.2.840.10008.5.1.1.26"
const UID_RETIRED_PullPrintRequestSOPClass = "1.2.840.10008.5.1.1.31"
const UID_RETIRED_PullStoredPrintManagementMetaSOPClass = "1.2.840.10008.5.1.1.32"

// Detached Management
const UID_RETIRED_DetachedPatientManagementSOPClass = "1.2.840.10008.3.1.2.1.1"
const UID_RETIRED_DetachedPatientManagementMetaSOPClass = "1.2.840.10008.3.1.2.1.4"
const UID_RETIRED_DetachedVisitManagementSOPClass = "1.2.840.10008.3.1.2.2.1"
const UID_RETIRED_DetachedStudyManagementSOPClass = "1.2.840.10008.3.1.2.3.1"
const UID_RETIRED_DetachedResultsManagementSOPClass = "1.2.840.10008.3.1.2.5.1"
const UID_RETIRED_DetachedResultsManagementMetaSOPClass = "1.2.840.10008.3.1.2.5.4"
const UID_RETIRED_DetachedStudyManagementMetaSOPClass = "1.2.840.10008.3.1.2.5.5"
const UID_RETIRED_DetachedInterpretationManagementSOPClass = "1.2.840.10008.3.1.2.6.1"

// Procedure Log
const UID_ProceduralEventLoggingSOPClass = "1.2.840.10008.1.40"
const UID_ProceduralEventLoggingSOPInstance = "1.2.840.10008.1.40.1"

// Substance Administration
const UID_SubstanceAdministrationLoggingSOPClass = "1.2.840.10008.1.42"
const UID_SubstanceAdministrationLoggingSOPInstance = "1.2.840.10008.1.42.1"
const UID_ProductCharacteristicsQuerySOPClass = "1.2.840.10008.5.1.4.41"
const UID_SubstanceApprovalQuerySOPClass = "1.2.840.10008.5.1.4.42"

// Media Creation
const UID_MediaCreationManagementSOPClass = "1.2.840.10008.5.1.1.33"

// SOP Class Relationship Negotiation
const UID_StorageServiceClass = "1.2.840.10008.4.2"

// Instance Availability Notification
const UID_InstanceAvailabilityNotificationSOPClass = "1.2.840.10008.5.1.4.33"

// Application Hosting
const UID_NativeDICOMModel = "1.2.840.10008.7.1.1"
const UID_AbstractMultiDimensionalImageModel = "1.2.840.10008.7.1.2"

// Other
const UID_VerificationSOPClass = "1.2.840.10008.1.1"
const UID_RETIRED_BasicStudyContentNotificationSOPClass = "1.2.840.10008.1.9"
const UID_RETIRED_StudyComponentManagementSOPClass = "1.2.840.10008.3.1.2.3.2"

// Coding Schemes
const UID_DICOMControlledTerminologyCodingScheme = "1.2.840.10008.2.16.4"
const UID_DICOMUIDRegistryCodingScheme = "1.2.840.10008.2.6.1"

// Configuration Management LDAP UIDs
const UID_LDAP_dicomDeviceName = "1.2.840.10008.15.0.3.1"
const UID_LDAP_dicomDescription = "1.2.840.10008.15.0.3.2"
const UID_LDAP_dicomManufacturer = "1.2.840.10008.15.0.3.3"
const UID_LDAP_dicomManufacturerModelName = "1.2.840.10008.15.0.3.4"
const UID_LDAP_dicomSoftwareVersion = "1.2.840.10008.15.0.3.5"
const UID_LDAP_dicomVendorData = "1.2.840.10008.15.0.3.6"
const UID_LDAP_dicomAETitle = "1.2.840.10008.15.0.3.7"
const UID_LDAP_dicomNetworkConnectionReference = "1.2.840.10008.15.0.3.8"
const UID_LDAP_dicomApplicationCluster = "1.2.840.10008.15.0.3.9"
const UID_LDAP_dicomAssociationInitiator = "1.2.840.10008.15.0.3.10"
const UID_LDAP_dicomAssociationAcceptor = "1.2.840.10008.15.0.3.11"
const UID_LDAP_dicomHostname = "1.2.840.10008.15.0.3.12"
const UID_LDAP_dicomPort = "1.2.840.10008.15.0.3.13"
const UID_LDAP_dicomSOPClass = "1.2.840.10008.15.0.3.14"
const UID_LDAP_dicomTransferRole = "1.2.840.10008.15.0.3.15"
const UID_LDAP_dicomTransferSyntax = "1.2.840.10008.15.0.3.16"
const UID_LDAP_dicomPrimaryDeviceType = "1.2.840.10008.15.0.3.17"
const UID_LDAP_dicomRelatedDeviceReference = "1.2.840.10008.15.0.3.18"
const UID_LDAP_dicomPreferredCalledAETitle = "1.2.840.10008.15.0.3.19"
const UID_LDAP_dicomTLSCyphersuite = "1.2.840.10008.15.0.3.20"
const UID_LDAP_dicomAuthorizedNodeCertificateReference = "1.2.840.10008.15.0.3.21"
const UID_LDAP_dicomThisNodeCertificateReference = "1.2.840.10008.15.0.3.22"
const UID_LDAP_dicomInstalled = "1.2.840.10008.15.0.3.23"
const UID_LDAP_dicomStationName = "1.2.840.10008.15.0.3.24"
const UID_LDAP_dicomDeviceSerialNumber = "1.2.840.10008.15.0.3.25"
const UID_LDAP_dicomInstitutionName = "1.2.840.10008.15.0.3.26"
const UID_LDAP_dicomInstitutionAddress = "1.2.840.10008.15.0.3.27"
const UID_LDAP_dicomInstitutionDepartmentName = "1.2.840.10008.15.0.3.28"
const UID_LDAP_dicomIssuerOfPatientID = "1.2.840.10008.15.0.3.29"
const UID_LDAP_dicomPreferredCallingAETitle = "1.2.840.10008.15.0.3.30"
const UID_LDAP_dicomSupportedCharacterSet = "1.2.840.10008.15.0.3.31"
const UID_LDAP_dicomConfigurationRoot = "1.2.840.10008.15.0.4.1"
const UID_LDAP_dicomDevicesRoot = "1.2.840.10008.15.0.4.2"
const UID_LDAP_dicomUniqueAETitlesRegistryRoot = "1.2.840.10008.15.0.4.3"
const UID_LDAP_dicomDevice = "1.2.840.10008.15.0.4.4"
const UID_LDAP_dicomNetworkAE = "1.2.840.10008.15.0.4.5"
const UID_LDAP_dicomNetworkConnection = "1.2.840.10008.15.0.4.6"
const UID_LDAP_dicomUniqueAETitle = "1.2.840.10008.15.0.4.7"
const UID_LDAP_dicomTransferCapability = "1.2.840.10008.15.0.4.8"

// UTC Synchronization Frame of Reference (CP 432)
const UID_UniversalCoordinatedTimeSynchronizationFrameOfReference = "1.2.840.10008.15.1.1"

// Well-known Frame of References
const UID_TalairachBrainAtlasFrameOfReference = "1.2.840.10008.1.4.1.1"
const UID_SPM2T1FrameOfReference = "1.2.840.10008.1.4.1.2"
const UID_SPM2T2FrameOfReference = "1.2.840.10008.1.4.1.3"
const UID_SPM2PDFrameOfReference = "1.2.840.10008.1.4.1.4"
const UID_SPM2EPIFrameOfReference = "1.2.840.10008.1.4.1.5"
const UID_SPM2FILT1FrameOfReference = "1.2.840.10008.1.4.1.6"
const UID_SPM2PETFrameOfReference = "1.2.840.10008.1.4.1.7"
const UID_SPM2TRANSMFrameOfReference = "1.2.840.10008.1.4.1.8"
const UID_SPM2SPECTFrameOfReference = "1.2.840.10008.1.4.1.9"
const UID_SPM2GRAYFrameOfReference = "1.2.840.10008.1.4.1.10"
const UID_SPM2WHITEFrameOfReference = "1.2.840.10008.1.4.1.11"
const UID_SPM2CSFFrameOfReference = "1.2.840.10008.1.4.1.12"
const UID_SPM2BRAINMASKFrameOfReference = "1.2.840.10008.1.4.1.13"
const UID_SPM2AVG305T1FrameOfReference = "1.2.840.10008.1.4.1.14"
const UID_SPM2AVG152T1FrameOfReference = "1.2.840.10008.1.4.1.15"
const UID_SPM2AVG152T2FrameOfReference = "1.2.840.10008.1.4.1.16"
const UID_SPM2AVG152PDFrameOfReference = "1.2.840.10008.1.4.1.17"
const UID_SPM2SINGLESUBJT1FrameOfReference = "1.2.840.10008.1.4.1.18"
const UID_ICBM452T1FrameOfReference = "1.2.840.10008.1.4.2.1"
const UID_ICBMSingleSubjectMRIFrameOfReference = "1.2.840.10008.1.4.2.2"

// Well-known SOP Instances for Color Palettes
const UID_HotIronColorPaletteSOPInstance = "1.2.840.10008.1.5.1"
const UID_PETColorPaletteSOPInstance = "1.2.840.10008.1.5.2"
const UID_HotMetalBlueColorPaletteSOPInstance = "1.2.840.10008.1.5.3"
const UID_PET20StepColorPaletteSOPInstance = "1.2.840.10008.1.5.4"

// Private DCMTK UIDs

// Private SOP Class UID used to shutdown external network applications
const UID_PrivateShutdownSOPClass = "1.2.276.0.7230010.3.4.1915765545.18030.917282194.0"

/* DRAFT SUPPLEMENTS - EXPERIMENTAL USE ONLY */

/*
 * The following UIDs were defined in "frozen draft for trial implementation" versions
 * of various DICOM supplements and are or will be changed before final text.
 * Since it is likely that trial implementations exist, we leave the UIDs in the dictionary.
 */

/* Supplement 23 Frozen Draft (November 1997) */
const UID_DRAFT_SRTextStorage = "1.2.840.10008.5.1.4.1.1.88.1"
const UID_DRAFT_SRAudioStorage = "1.2.840.10008.5.1.4.1.1.88.2"
const UID_DRAFT_SRDetailStorage = "1.2.840.10008.5.1.4.1.1.88.3"
const UID_DRAFT_SRComprehensiveStorage = "1.2.840.10008.5.1.4.1.1.88.4"

/* Supplement 30 Draft 08 for Demonstration (October 1997) */
const UID_DRAFT_WaveformStorage = "1.2.840.10008.5.1.4.1.1.9.1"

/* Supplement 74 Frozen Draft (October 2007) */
const UID_DRAFT_RTBeamsDeliveryInstructionStorage = "1.2.840.10008.5.1.4.34.1"
const UID_DRAFT_RTConventionalMachineVerification = "1.2.840.10008.5.1.4.34.2"
const UID_DRAFT_RTIonMachineVerification = "1.2.840.10008.5.1.4.34.3"
