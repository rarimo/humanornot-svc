/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Nonce struct {
	Key
	Attributes NonceAttributes `json:"attributes"`
}
type NonceResponse struct {
	Data     Nonce    `json:"data"`
	Included Included `json:"included"`
}

type NonceListResponse struct {
	Data     []Nonce  `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustNonce - returns Nonce from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustNonce(key Key) *Nonce {
	var nonce Nonce
	if c.tryFindEntry(key, &nonce) {
		return &nonce
	}
	return nil
}
