/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type ClaimOffer struct {
	Key
	Attributes ClaimOfferAttributes `json:"attributes"`
}
type ClaimOfferResponse struct {
	Data     ClaimOffer `json:"data"`
	Included Included   `json:"included"`
}

type ClaimOfferListResponse struct {
	Data     []ClaimOffer `json:"data"`
	Included Included     `json:"included"`
	Links    *Links       `json:"links"`
}

// MustClaimOffer - returns ClaimOffer from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustClaimOffer(key Key) *ClaimOffer {
	var claimOffer ClaimOffer
	if c.tryFindEntry(key, &claimOffer) {
		return &claimOffer
	}
	return nil
}
