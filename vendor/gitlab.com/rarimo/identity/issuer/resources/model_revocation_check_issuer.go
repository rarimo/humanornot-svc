/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type RevocationCheckIssuer struct {
	// Issuer's claims merkle tree root hash
	ClaimsTreeRoots string `json:"claims_tree_roots"`
	// Issuer's revocations merkle tree root hash
	RevocationTreeRoot string `json:"revocation_tree_root"`
	// Issuer's roots merkle tree root hash
	RootOfRoots string `json:"root_of_roots"`
	// Issuer state hash
	State string `json:"state"`
}
