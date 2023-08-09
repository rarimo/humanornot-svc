/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type OfferCallback struct {
	Body OfferCallbackBody `json:"body"`
	// The message receiver
	From string `json:"from"`
	// The callback id
	Id string `json:"id"`
	// The message receiver
	To string `json:"to"`
	// The message format
	Typ *string `json:"typ,omitempty"`
	// The message type
	Type string `json:"type"`
}
