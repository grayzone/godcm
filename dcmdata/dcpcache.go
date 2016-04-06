package dcmdata

/** class handling one entry of the Private Tag Cache List
 */
type DcmPrivateTagCacheEntry struct {
	tagKey         DcmTagKey
	privateCreator string
}

/** constructor
 *  @param tk tag key for private creator element
 *  @param pc private creator name, must not be NULL or empty string
 */
func NewDcmPrivateTagCacheEntry(tk DcmTagKey, pc string) *DcmPrivateTagCacheEntry {
	return &DcmPrivateTagCacheEntry{tk, pc}
}

/** returns the private creator name
 */
func (entry *DcmPrivateTagCacheEntry) GetPrivateCreator() string {
	return entry.privateCreator
}

/** checks if this element is the private creator for the element
 *  with the given tag key
 *  @param tk tag key to check
 *  @return OFTrue if this element contains the matching private creator,
 *    OFFalse otherwise.
 */
func (entry *DcmPrivateTagCacheEntry) IsPrivateCreatorFor(tk DcmTagKey) bool {
	return (entry.tagKey.group == tk.group) && ((entry.tagKey.element << 8) == (tk.element & 0xFF00))

}

/** this class implements a cache of Private Creator elements
 *  and corresponding reserved tag numbers.
 */
type DcmPrivateTagCache struct {
	list []DcmPrivateTagCacheEntry
}

/** updates the private creator cache with the given object.
 *  If the object points to a private creator element,
 *  the tag key and creator code are added to the cache.
 *  Otherwise, the cache remains unmodified.
 */
func (cache *DcmPrivateTagCache) UpdateCache(obj DcmObject) {
	tag := obj.tag
	if obj.IsLeaf() && ((tag.GetGTag() & 1) == 1) && (tag.GetETag() <= 0xFF) && (tag.GetETag() >= 0x10) {
		var elem DcmElement
		elem.DcmObject = obj
		var c string
		status := elem.GetString(&c)
		if status.Good() {
			entry := DcmPrivateTagCacheEntry{tag.GetXTag(), c}
			cache.list = append(cache.list, entry)
		}

	}

}

/// resets the cache to default-constructed state
func (cache *DcmPrivateTagCache) Clear() {
	cache.list = nil
}

/** looks up the private creator name for the given private tag
 *  @param tk private tag to check
 *  @return private creator name if found, NULL otherwise.
 */
func (cache *DcmPrivateTagCache) FindPrivateCreator(tk DcmTagKey) string {
	for _, v := range cache.list {
		if v.IsPrivateCreatorFor(tk) {
			return v.GetPrivateCreator()
		}
	}
	return ""
}
