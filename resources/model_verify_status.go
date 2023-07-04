/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type VerifyStatus struct {
	Key
	Attributes VerifyStatusAttributes `json:"attributes"`
}
type VerifyStatusResponse struct {
	Data     VerifyStatus `json:"data"`
	Included Included     `json:"included"`
}

type VerifyStatusListResponse struct {
	Data     []VerifyStatus `json:"data"`
	Included Included       `json:"included"`
	Links    *Links         `json:"links"`
}

// MustVerifyStatus - returns VerifyStatus from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustVerifyStatus(key Key) *VerifyStatus {
	var verifyStatus VerifyStatus
	if c.tryFindEntry(key, &verifyStatus) {
		return &verifyStatus
	}
	return nil
}
