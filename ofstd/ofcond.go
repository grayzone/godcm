package ofstd

type OFStatus int

const (
	/// no error, operation has completed successfully
	OF_ok = iota

	/// operation has not completed successfully
	OF_error

	/// application failure
	OF_failure
)

func (s OFStatus) String() string {
	var result string
	switch s {
	case OF_ok:
		result = "OF_ok"
	case OF_error:
		result = "OF_error"
	case OF_failure:
		result = "OF_failure"
	}
	return result
}

type OFCondition struct {
	module uint16
	code   uint16
	status OFStatus
	text   string
}

/*  global condition constants.
 *  All constants defined here use module number 0 which is reserved for
 *  global definitions. Other constants are defined elsewhere.
 */

/// condition constant: successful completion
var EC_Normal = OFCondition{0, 0, OF_ok, "Normal"}

/// condition constant: error, function called with illegal parameters
var EC_IllegalParameter = OFCondition{0, 1, OF_error, "Illegal parameter"}

/// condition constant: failure, memory exhausted
var EC_MemoryExhausted = OFCondition{0, 2, OF_failure, "Virtual Memory exhausted"}

func NewOFCondition(m uint16, c uint16, s OFStatus, t string) *OFCondition {
	return &OFCondition{m, c, s, t}
}

func MakeOFCondition(m uint16, c uint16, s OFStatus, t string) OFCondition {
	return OFCondition{m, c, s, t}
}

/// returns the module identifier for this object.
func (c *OFCondition) Module() uint16 {
	return c.module
}

/// returns the status code identifier for this object.
func (c *OFCondition) Code() uint16 {
	return c.code
}

/// returns the status for this object.
func (c *OFCondition) Status() OFStatus {
	return c.status
}

/// returns the error message text for this object.
func (c *OFCondition) Text() string {
	return c.text
}

/// returns true if status is OK
func (c *OFCondition) Good() bool {
	return c.status == OF_ok
}
