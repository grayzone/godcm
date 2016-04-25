package core

// DcmDataSet is to contain the DICOM dataset from file.
type DcmDataSet struct {
	Elements []DcmElement
}

func (dataset *DcmDataSet) Read(stream *DcmFileStream, isExplicitVR bool, isReadValue bool, isReadPixel bool) error {
	for !stream.Eos() {
		//	for range [12]int{} {
		var elem DcmElement
		err := elem.ReadDcmElement(stream, isExplicitVR, isReadValue, isReadPixel)
		if err != nil {
			return err
		}
		dataset.Elements = append(dataset.Elements, elem)
	}
	return nil
}
