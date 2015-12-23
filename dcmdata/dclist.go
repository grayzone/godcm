package dcmdata

/// index indicating "end of list"
const DCM_EndOfListIndex = -1

/** helper class maintaining an entry in a DcmList double-linked list
 */

type DcmListNode struct {
	/// pointer to next node in double-linked list
	nextNode *DcmListNode
	/// pointer to previous node in double-linked list
	prevNode *DcmListNode
	/// pointer to DcmObject instance maintained by this list entry
	objNodeValue *DcmObject
}

func NewDcmListNode(obj DcmObject) *DcmListNode {
	return &DcmListNode{nil, nil, &obj}
}

/// return pointer to object maintained by this list node
func (n *DcmListNode) Value() *DcmObject {
	return n.objNodeValue
}

/// list position indicator
type E_ListPos int

const (

	/// at current position in list
	ELP_atpos = iota

	/// at list start
	ELP_first

	/// at list end
	ELP_last

	/// before current list position
	ELP_prev

	/// after current list position
	ELP_next
)

/** double-linked list class that maintains pointers to DcmObject instances.
 *  The remove operation does not delete the object pointed to, however,
 *  the destructor will delete all elements pointed to
 */
type DcmList struct {
	/// pointer to first node in list
	firstNode *DcmListNode

	/// pointer to last node in list
	lastNode *DcmListNode

	/// pointer to current node in list
	currentNode *DcmListNode

	/// number of elements in list
	cardinality uint32
}

/// return true if list is empty, false otherwise
func (l *DcmList) Empty() bool {
	return l.firstNode == nil
}

/// return true if current node exists, false otherwise
func (l *DcmList) Valid() bool {
	return l.currentNode != nil
}

/// return cardinality of list
func (l *DcmList) Card() uint32 {
	return l.cardinality
}

/** insert object at end of list
 *  @param obj pointer to object
 *  @return pointer to object
 */
func (dl *DcmList) Append(obj *DcmObject) *DcmObject {
	if obj != nil {
		if dl.Empty() { // list is empty !
			dl.currentNode = NewDcmListNode(*obj)
			dl.firstNode = dl.currentNode
			dl.lastNode = dl.currentNode
		} else {
			node := NewDcmListNode(*obj)
			node.nextNode = dl.firstNode
			dl.firstNode.prevNode = node
			dl.currentNode = node
			dl.lastNode = node
		}
		dl.cardinality = dl.cardinality + 1
	}
	return obj
}

/** insert object at start of list
 *  @param obj pointer to object
 *  @return pointer to object
 */
func (dl *DcmList) Prepend(obj *DcmObject) *DcmObject {
	if obj != nil {
		if dl.Empty() { // list is empty !
			dl.currentNode = NewDcmListNode(*obj)
			dl.firstNode = dl.currentNode
			dl.lastNode = dl.currentNode
		} else {
			node := NewDcmListNode(*obj)
			node.nextNode = dl.firstNode
			dl.firstNode.prevNode = node
			dl.currentNode = node
			dl.firstNode = node
		}
		dl.cardinality = dl.cardinality + 1
	}

	return obj
}

/** insert object relative to current position and indicator
 *  @param obj pointer to object
 *  @param pos position indicator
 *  @return pointer to object
 */
func (dl *DcmList) Insert(obj *DcmObject, pos E_ListPos) *DcmObject {
	if obj != nil {
		if dl.Empty() { // list is empty !
			dl.currentNode = NewDcmListNode(*obj)
			dl.firstNode = dl.currentNode
			dl.lastNode = dl.currentNode
			dl.cardinality = dl.cardinality + 1
		} else {
			if pos == ELP_last {
				dl.Append(obj) // cardinality++;
			} else if pos == ELP_first {
				dl.Prepend(obj) // cardinality++;
			} else if dl.Valid() != true {
				// set current node to the end if there is no predecessor or
				// there are successors to be determined
				dl.Append(obj) // cardinality++;
			} else if pos == ELP_prev { // insert before current node
				node := NewDcmListNode(*obj)
				if dl.currentNode.prevNode == nil {
					dl.firstNode = node // insert at the beginning
				} else {
					dl.currentNode.prevNode.nextNode = node
				}
				node.prevNode = dl.currentNode.prevNode
				node.nextNode = dl.currentNode
				dl.currentNode.prevNode = node
				dl.currentNode = node
				dl.cardinality = dl.cardinality + 1
			} else { //( pos==ELP_next || pos==ELP_atpos )
				// insert after current node
				node := NewDcmListNode(*obj)
				if dl.currentNode.nextNode == nil {
					dl.lastNode = node // append to the end
				} else {
					dl.currentNode.nextNode.prevNode = node
				}
				node.nextNode = dl.currentNode.nextNode
				node.prevNode = dl.currentNode
				dl.currentNode.nextNode = node
				dl.currentNode = node
				dl.cardinality = dl.cardinality + 1
			}
		}
	}
	return obj
}

/** remove current entry from list, return element
 *  @return pointer to removed element, which is not deleted
 */
func (l *DcmList) Remove() *DcmObject {
	if l.Empty() {
		return nil
	} else if l.Valid() != true {
		return nil
	} else {
		tmpnode := l.currentNode
		if l.currentNode.prevNode == nil {
			l.firstNode = l.currentNode.nextNode // delete first element
		} else {
			l.currentNode.prevNode.nextNode = l.currentNode.nextNode
		}
		if l.currentNode.nextNode == nil {
			l.lastNode = l.currentNode.prevNode // delete last element
		} else {
			l.currentNode.nextNode.prevNode = l.currentNode.prevNode
		}
		l.currentNode = l.currentNode.nextNode
		tmpobj := tmpnode.Value()
		l.cardinality = l.cardinality - 1
		return tmpobj
	}
}

/** get pointer to element in list at given position
 *  @param pos position indicator
 *  @return pointer to object
 */
func (l *DcmList) Get(pos E_ListPos) *DcmObject {
	return l.Seek(pos)
}

/** seek within element in list to given position
 *  (i.e. set current element to given position)
 *  @param pos position indicator
 *  @return pointer to new current object
 */
func (l *DcmList) Seek(pos E_ListPos) *DcmObject {
	switch pos {
	case ELP_first:
		l.currentNode = l.firstNode
	case ELP_last:
		l.currentNode = l.lastNode
	case ELP_prev:
		if l.Valid() {
			l.currentNode = l.currentNode.prevNode
		}
	case ELP_next:
		if l.Valid() {
			l.currentNode = l.currentNode.nextNode
		}

	}
	if l.Valid() {
		return l.currentNode.Value()
	} else {
		return nil
	}
}

/** seek within element in list to given element index
 *  (i.e. set current element to given index)
 *  @param absolute_position position index < card()
 *  @return pointer to new current object
 */
func (l *DcmList) Seek_to(absolute_position uint32) *DcmObject {
	var tmppos uint32
	if absolute_position < l.cardinality {
		tmppos = absolute_position
	} else {
		tmppos = l.cardinality
	}
	l.Seek(ELP_first)

	for i := uint32(0); i < tmppos; i++ {
		l.Seek(ELP_next)
	}
	return l.Get(ELP_atpos)
}
