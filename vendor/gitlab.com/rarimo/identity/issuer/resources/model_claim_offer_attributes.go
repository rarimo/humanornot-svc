/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type ClaimOfferAttributes struct {
	Body ClaimOfferBody `json:"body"`
	// The unique identity identifer who will get the claim
	From string `json:"from"`
	// The uniquer offer identifier
	Id   string `json:"id"`
	Thid string `json:"thid"`
	// The unique issuer identity identifier
	To string `json:"to"`
	// The Iden3 message media type
	Typ string `json:"typ"`
	// The Iden3 protocol message type
	Type string `json:"type"`
}
