package kleros

import (
	"github.com/ethereum/go-ethereum/common"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/kyc-service/internal/service/api/requests"
	"github.com/rarimo/kyc-service/resources"
)

type (
	// VerificationData is a data that is required by Kleros to verify a user
	VerificationData resources.KlerosData
)

type ProviderData struct {
	Address common.Address `json:"address"`
}

// Validate is a method that validates VerificationData
func (v VerificationData) Validate() error {
	return validation.Errors{
		"signature": validation.Validate(v.Signature, validation.Required),
		"address":   validation.Validate(v.Address, validation.Required, validation.By(requests.MustBeEthAddress)),
	}.Filter()
}
