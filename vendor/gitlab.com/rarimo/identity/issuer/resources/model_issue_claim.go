/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type IssueClaim struct {
	Key
	Attributes IssueClaimAttributes `json:"attributes"`
}
type IssueClaimRequest struct {
	Data     IssueClaim `json:"data"`
	Included Included   `json:"included"`
}

type IssueClaimListRequest struct {
	Data     []IssueClaim `json:"data"`
	Included Included     `json:"included"`
	Links    *Links       `json:"links"`
}

// MustIssueClaim - returns IssueClaim from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustIssueClaim(key Key) *IssueClaim {
	var issueClaim IssueClaim
	if c.tryFindEntry(key, &issueClaim) {
		return &issueClaim
	}
	return nil
}
