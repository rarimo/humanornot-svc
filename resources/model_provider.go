/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Provider struct {
	Key
	Attributes ProviderAttributes `json:"attributes"`
}
type ProviderResponse struct {
	Data     Provider `json:"data"`
	Included Included `json:"included"`
}

type ProviderListResponse struct {
	Data     []Provider `json:"data"`
	Included Included   `json:"included"`
	Links    *Links     `json:"links"`
}

// MustProvider - returns Provider from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustProvider(key Key) *Provider {
	var provider Provider
	if c.tryFindEntry(key, &provider) {
		return &provider
	}
	return nil
}
