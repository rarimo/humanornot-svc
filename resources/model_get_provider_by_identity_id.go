/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type GetProviderByIdentityId struct {
	Key
	Attributes GetProviderByIdentityIdAttributes `json:"attributes"`
}
type GetProviderByIdentityIdResponse struct {
	Data     GetProviderByIdentityId `json:"data"`
	Included Included                `json:"included"`
}

type GetProviderByIdentityIdListResponse struct {
	Data     []GetProviderByIdentityId `json:"data"`
	Included Included                  `json:"included"`
	Links    *Links                    `json:"links"`
}

// MustGetProviderByIdentityId - returns GetProviderByIdentityId from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustGetProviderByIdentityId(key Key) *GetProviderByIdentityId {
	var getProviderByIdentityId GetProviderByIdentityId
	if c.tryFindEntry(key, &getProviderByIdentityId) {
		return &getProviderByIdentityId
	}
	return nil
}
