/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type NonceRequest struct {
	Key
	Attributes NonceRequestAttributes `json:"attributes"`
}
type NonceRequestResponse struct {
	Data     NonceRequest `json:"data"`
	Included Included     `json:"included"`
}

type NonceRequestListResponse struct {
	Data     []NonceRequest `json:"data"`
	Included Included       `json:"included"`
	Links    *Links         `json:"links"`
}

// MustNonceRequest - returns NonceRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustNonceRequest(key Key) *NonceRequest {
	var nonceRequest NonceRequest
	if c.tryFindEntry(key, &nonceRequest) {
		return &nonceRequest
	}
	return nil
}
