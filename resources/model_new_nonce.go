/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type NewNonce struct {
	Key
	Attributes NewNonceAttributes `json:"attributes"`
}
type NewNonceResponse struct {
	Data     NewNonce `json:"data"`
	Included Included `json:"included"`
}

type NewNonceListResponse struct {
	Data     []NewNonce `json:"data"`
	Included Included   `json:"included"`
	Links    *Links     `json:"links"`
}

// MustNewNonce - returns NewNonce from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustNewNonce(key Key) *NewNonce {
	var newNonce NewNonce
	if c.tryFindEntry(key, &newNonce) {
		return &newNonce
	}
	return nil
}
