package dcmdata

/// constant describing an unlimited VM
const DcmVariableVM = -1

/** attribute tag group/element range restrictions
 */
type DcmDictRangeRestriction int

const (
	/// integer range
	DcmDictRange_Unspecified = iota

	/// odd range
	DcmDictRange_Odd

	/// even range
	DcmDictRange_Even
)

type DcmDictEntry struct {
	/** upper limit of repeating group and element (lower limit is inherited
	 *   from DcmTagKey)
	 */
	DcmTagKey
	upperKey DcmTagKey
	/// value representation
	valueRepresentation DcmVR

	/// attribute name
	tagName string

	/// lower limit for VM
	valueMultiplicityMin int

	/// upper limit for VM
	valueMultiplicityMax int

	/// standard version name, may be NULL
	standardVersion string

	/// true if strings are copies (i.e. should be deleted upon destruction)
	stringsAreCopies bool

	/// restriction (even, odd, unrestricted) for group range
	groupRangeRestriction DcmDictRangeRestriction

	/// restriction (even, odd, unrestricted) for element range
	elementRangeRestriction DcmDictRangeRestriction

	/// private creator name, may be NULL
	privateCreator string
}

/** constructor
 *  @param g attribute tag group
 *  @param e attribute tag element
 *  @param vr value representation
 *  @param nam attribute name
 *  @param vmMin lower limit for value multiplicity
 *  @param vmMax upper limit for value multiplicity, DcmVariableVM for unlimited
 *  @param vers standard version name, may be NULL
 *  @param doCopyStrings true if strings should be copied, false if only referenced
 *  @param pcreator private creator name, may be NULL (for standard tags)
 */

func NewDcmDictEntry(g uint16, e uint16, vr DcmVR, nam string, vmMin int, vmMax int, vers string, doCopyStrings bool, pcreator string) *DcmDictEntry {
	entry := new(DcmDictEntry)
	entry.group = g
	entry.element = e
	entry.tagName = nam
	entry.valueRepresentation = vr
	entry.valueMultiplicityMin = vmMin
	entry.valueMultiplicityMax = vmMax
	entry.standardVersion = vers
	entry.privateCreator = pcreator
	entry.stringsAreCopies = doCopyStrings
	entry.upperKey = DcmTagKey{g, e}
	entry.groupRangeRestriction = DcmDictRange_Unspecified
	entry.elementRangeRestriction = DcmDictRange_Unspecified
	return entry
}

/** constructor for repeating tags
 *  @param g attribute tag group lower limit
 *  @param e attribute tag element lower limit
 *  @param ug attribute tag group upper limit
 *  @param ue attribute tag element upper limit
 *  @param vr value representation
 *  @param nam attribute name
 *  @param vmMin lower limit for value multiplicity
 *  @param vmMax upper limit for value multiplicity, DcmVariableVM for unlimited
 *  @param vers standard version name, may be NULL
 *  @param doCopyStrings true if strings should be copied, false if only referenced
 *  @param pcreator private creator name, may be NULL (for standard tags)
 */
func NewDcmDictEntryForRepeatingTags(g uint16, e uint16, ug uint16, ue uint16, vr DcmVR, nam string, vmMin int, vmMax int, vers string, doCopyStrings bool, pcreator string) *DcmDictEntry {
	entry := new(DcmDictEntry)
	entry.group = g
	entry.element = e
	entry.tagName = nam
	entry.valueRepresentation = vr
	entry.valueMultiplicityMin = vmMin
	entry.valueMultiplicityMax = vmMax
	entry.standardVersion = vers
	entry.privateCreator = pcreator
	entry.stringsAreCopies = doCopyStrings
	entry.upperKey = DcmTagKey{ug, ue}
	entry.groupRangeRestriction = DcmDictRange_Unspecified
	entry.elementRangeRestriction = DcmDictRange_Unspecified
	return entry
}

/// returns VR object by value
func (entry *DcmDictEntry) GetVR() DcmVR {
	return entry.valueRepresentation
}

/// returns VR code
func (entry *DcmDictEntry) GetEVR() DcmEVR {
	return entry.valueRepresentation.GetEVR()
}

/// returns standard version string, may be NULL
func (entry *DcmDictEntry) GetStandardVersion() string {
	return entry.standardVersion
}

/// returns tag name
func (entry *DcmDictEntry) GetTagName() string {
	return entry.tagName
}

/// returns private creator code, may be NULL
func (entry *DcmDictEntry) GetPrivateCreator() string {
	return entry.privateCreator
}

/// returns lower limit for VM (value multiplicity)
func (entry *DcmDictEntry) GetVMMin() int {
	return entry.valueMultiplicityMin
}

/// returns upper limit for VM (value multiplicity), DcmVariableVM for unlimited
func (entry *DcmDictEntry) GetVMMax() int {

	return entry.valueMultiplicityMax
}

/// returns true if element has a single valid VM value
func (entry *DcmDictEntry) IsFixedSingleVM() bool {
	return ((entry.valueMultiplicityMin != DcmVariableVM) && (entry.valueMultiplicityMin == entry.valueMultiplicityMax))
}

/// returns true if element has a fixed VM range
func (entry *DcmDictEntry) IsFixedRangeVM() bool {
	return ((entry.valueMultiplicityMin != DcmVariableVM) && (entry.valueMultiplicityMax != DcmVariableVM))
}

/// returns true if element has a variable VM range (no upper limit)
func (entry *DcmDictEntry) IsVariableRangeVM() bool {
	return ((entry.valueMultiplicityMin != DcmVariableVM) && (entry.valueMultiplicityMax == DcmVariableVM))
}

/** converts entry into repeating tag entry by defining an upper limit
 *  for group and element, taken from the given tag key.
 *  @param key tag key containing upper limit for group and element
 */
func (entry *DcmDictEntry) SetUpper(key DcmTagKey) {
	entry.upperKey = key
}

/** converts entry into repeating tag entry by defining an upper limit
 *  for tag group
 *  @param ug upper limit for tag group
 */
func (entry *DcmDictEntry) SetUpperGroup(ug uint16) {
	entry.upperKey.group = ug
}

/** converts entry into repeating tag entry by defining an upper limit
 *  for tag element
 *  @param ue upper limit for tag element
 */
func (entry *DcmDictEntry) SetUpperelement(ue uint16) {
	entry.upperKey.element = ue
}

/// returns upper limit for tag group
func (entry *DcmDictEntry) GetUpperGroup() uint16 {
	return entry.upperKey.group
}

/// returns upper limit for tag element
func (entry *DcmDictEntry) GetUpperElement() uint16 {
	return entry.upperKey.element
}

/// returns attribute tag as DcmTagKey object by value
func (entry *DcmDictEntry) GetKey() DcmTagKey {
	return entry.DcmTagKey
}

/// returns upper limits for attribute tag as DcmTagKey object by value
func (entry *DcmDictEntry) GetUpperKey() DcmTagKey {
	return entry.upperKey
}

/// returns true if entry is has a repeating group
func (entry *DcmDictEntry) IsRepeatingGroup() bool {
	return (entry.group != entry.GetUpperGroup())
}

/// returns true if entry is has a repeating element
func (entry *DcmDictEntry) IsRepeatingElement() bool {
	return (entry.element != entry.GetUpperElement())
}

/// returns true if entry is repeating (group or element)
func (entry *DcmDictEntry) IsRepeating() bool {
	return (entry.IsRepeatingGroup() || entry.IsRepeatingElement())
}

/// returns group range restriction
func (entry *DcmDictEntry) GetGroupRangeRestriction() DcmDictRangeRestriction {
	return entry.groupRangeRestriction
}

/// sets group range restriction
func (entry *DcmDictEntry) SetGroupRangeRestriction(rr DcmDictRangeRestriction) {

	entry.groupRangeRestriction = rr
}

/// returns element range restriction
func (entry *DcmDictEntry) GetElementRangeRestriction() DcmDictRangeRestriction {

	return entry.elementRangeRestriction
}

/// sets element range restriction
func (entry *DcmDictEntry) SetElementRangeRestriction(rr DcmDictRangeRestriction) {

	entry.elementRangeRestriction = rr
}

func dcm_inrange(x, a, b uint16) bool {
	return (x >= a) && (x <= b)
}

func dcm_is_odd(x uint16) bool {
	return ((x % 2) == 1)
}

func dcm_is_even(x uint16) bool {
	return ((x % 2) == 0)
}

/** checks if the private creator code equals the given string
 *  @param c string to compare with, may be NULL
 *  @return true if equal, false otherwise
 */
func (entry *DcmDictEntry) PrivateCreatorMatch(s string) bool {
	return entry.privateCreator == s

}

/** checks if the given tag key and private creator code are covered
 *  by this object.
 *  @param key tag key
 *  @param privCreator private creator, may be NULL
 *  @return true if this entry contains the given tag for the given private creator
 */
func (entry *DcmDictEntry) Contains(key DcmTagKey, privCreator string) bool {
	if (entry.GetGroupRangeRestriction() == DcmDictRange_Even) && dcm_is_odd(key.group) {
		return false
	} else if (entry.GetGroupRangeRestriction() == DcmDictRange_Odd) && dcm_is_even(key.group) {
		return false
	} else if (entry.GetElementRangeRestriction() == DcmDictRange_Even) && dcm_is_odd(key.element) {
		return false
	} else if (entry.GetElementRangeRestriction() == DcmDictRange_Odd) && dcm_is_even(key.element) {
		return false
	} else if !entry.PrivateCreatorMatch(privCreator) {
		return false
	} else {
		groupMatches := dcm_inrange(key.group, entry.group, entry.GetUpperGroup())
		found := groupMatches && dcm_inrange(key.element, entry.element, entry.GetUpperElement())
		if !found && groupMatches {
			found = dcm_inrange(key.element&0xFF, entry.element, entry.GetUpperElement())
		}
		return found
	}
}

/** checks if this entry contains the given name
 *  @param name attribute name, must not be NULL
 *  @return true if tagName matches the given string
 */
func (entry *DcmDictEntry) ContainsTagName(name string) bool { /* this contains named key */
	return entry.tagName == name
}

/** checks if this entry describes a true subset of tag range
 *  described by the given entry.
 *  @param e entry to compare with
 *  @return true if this object is subset of e
 */
func (entry *DcmDictEntry) Subset(e DcmDictEntry) bool { /* this is a subset of key */
	return ((entry.group >= e.group) &&
		(entry.GetUpperGroup() <= e.GetUpperGroup()) &&
		(entry.element >= e.element) &&
		(entry.GetUpperElement() <= e.GetUpperElement()) &&
		entry.PrivateCreatorMatch(e.privateCreator))
}

/** checks if this entry describes the same tag range as the given entry.
 *  @param e entry to compare with
 *  @return true if objects describe the same tag range
 */
func (entry *DcmDictEntry) SetEQ(e DcmDictEntry) bool { /* this is set equal to key */
	return ((entry.group == e.group) &&
		(entry.GetUpperGroup() == e.GetUpperGroup()) &&
		(entry.element == e.element) &&
		(entry.GetUpperElement() == e.GetUpperElement()) &&
		(entry.GetGroupRangeRestriction() == e.GetGroupRangeRestriction()) &&
		(entry.GetElementRangeRestriction() == e.GetElementRangeRestriction()) &&
		entry.PrivateCreatorMatch(e.privateCreator))
}
