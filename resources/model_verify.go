/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Verify struct {
	Key
	Attributes VerifyAttributes `json:"attributes"`
}
type VerifyRequest struct {
	Data     Verify   `json:"data"`
	Included Included `json:"included"`
}

type VerifyListRequest struct {
	Data     []Verify `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustVerify - returns Verify from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustVerify(key Key) *Verify {
	var verify Verify
	if c.tryFindEntry(key, &verify) {
		return &verify
	}
	return nil
}
