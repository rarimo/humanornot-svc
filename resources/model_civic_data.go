/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

// Civic provider's data
type CivicData struct {
	// The user's address
	Address string `json:"address"`
	// One of the available chain's name
	ChainName string `json:"chain_name"`
	// The signature of the requested nonce to validate if the user owns the address
	Signature string `json:"signature"`
}
