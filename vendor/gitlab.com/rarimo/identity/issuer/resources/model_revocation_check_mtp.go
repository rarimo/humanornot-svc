/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type RevocationCheckMtp struct {
	// Is claim reverted or no
	Existence bool `json:"existence"`
	// List of merkle tree sibling as a way from claim to the merkle tree hash
	Siblings []string `json:"siblings"`
}
