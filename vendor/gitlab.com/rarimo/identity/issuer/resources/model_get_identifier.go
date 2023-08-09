/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type GetIdentifier struct {
	Key
	Attributes *GetIdentifierAttributes `json:"attributes,omitempty"`
}
type GetIdentifierResponse struct {
	Data     GetIdentifier `json:"data"`
	Included Included      `json:"included"`
}

type GetIdentifierListResponse struct {
	Data     []GetIdentifier `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
}

// MustGetIdentifier - returns GetIdentifier from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustGetIdentifier(key Key) *GetIdentifier {
	var getIdentifier GetIdentifier
	if c.tryFindEntry(key, &getIdentifier) {
		return &getIdentifier
	}
	return nil
}
