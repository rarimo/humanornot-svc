/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

// The data set required for the KYCFullNameCredentials
type KycFullNameCredentialsSchema struct {
	// The user's first name
	FirstName string `json:"first_name"`
	// The user's middle name
	MiddleName *string `json:"middle_name,omitempty"`
	// The user's name prefix
	NamePrefix *string `json:"name_prefix,omitempty"`
	// The user's name suffix
	NameSuffix *string `json:"name_suffix,omitempty"`
	// The user's second name
	SecondName string `json:"second_name"`
}
