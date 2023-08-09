/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type VerifyAttributes struct {
	// Iden3 identity's id
	IdentityId   string          `json:"identity_id"`
	ProviderData json.RawMessage `json:"provider_data"`
}
