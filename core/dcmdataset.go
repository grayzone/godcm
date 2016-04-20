package core

// DcmDataSet is to contain the DICOM dataset from file.
type DcmDataSet struct {
	Elements []DcmElement
}

func (dataset *DcmDataSet) Read(stream *DcmFileStream, isExplicitVR bool) error {
	// for testing, implement the dcm tag with Implicit VR first.
	if isExplicitVR {
		return nil
	}
	for !stream.Eos() {
		//	for range [12]int{} {
		elem, err := stream.ReadDcmElement(isExplicitVR)
		if err != nil {
			return err
		}
		dataset.Elements = append(dataset.Elements, elem)
	}
	return nil
}
