package dcmdata

// Undefined Length Identifier
const DCM_UndefinedLength = 0xffffffff

/// object state during transfer (read/write) operations
type E_TransferState int

const (
	/// object prepared for transfer, no data transferred yet
	ERW_init E_TransferState = 0
	/// object transfer completed
	ERW_ready E_TransferState = 1
	/// object transfer in progress
	ERW_inWork E_TransferState = 2
	/// object not prepared for transfer operation
	ERW_notInitialized E_TransferState = 3
)

/// encoding type for sequences and sequence items
type E_EncodingType int

const (
	/// defined length
	EET_ExplicitLength = 0
	/// undefined length
	EET_UndefinedLength = 1
)

/// handling of group length elements when reading/writing a dataset
type E_GrpLenEncoding int

const (
	/// no change of group length values, WARNING: DO NOT USE THIS VALUE FOR WRITE OPERATIONS
	EGL_noChange = 0
	/// remove group length tags
	EGL_withoutGL = 1
	/// add group length tags for every group
	EGL_withGL = 2
	/// recalculate values for existing group length tags
	EGL_recalcGL = 3
)

/// handling of dataset trailing padding
type E_PaddingEncoding int

const (
	/// no change of padding tags
	EPD_noChange = 0
	/// remove all padding tags
	EPD_withoutPadding = 1
	/// add padding tags
	EPD_withPadding = 2
)
