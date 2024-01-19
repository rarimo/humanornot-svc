/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type ConflictError struct {
	// Application-specific error code, expressed as a string value  - eth_address_taken - 1 - identity_verified - 2 - duplicate_provider_data - 3
	Code string `json:"code"`
	// Human-readable explanation specific to this occurrence of the problem
	Detail *string `json:"detail,omitempty"`
	// Object containing non-standard meta-information about the error
	Meta *map[string]interface{} `json:"meta,omitempty"`
	// HTTP status code applicable to this problem
	Status int32 `json:"status"`
	// Short, human-readable summary of the problem
	Title string `json:"title"`
}
