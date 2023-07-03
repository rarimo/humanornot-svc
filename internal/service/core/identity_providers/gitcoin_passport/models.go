package gcpsp

import (
	"github.com/ethereum/go-ethereum/common"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"
)

const (
	submitPassportEndpoint = "/submit-passport"
	scoreEndpoint          = "/score"
)

const (
	statusDone       = "DONE"
	statusProcessing = "PROCESSING"
)

const (
	// retryInterval is a time interval between request retries
	retryInterval = time.Second * 3
)

type (
	// VerificationData is a data that is required by Gitcoin Passport to verify a user
	VerificationData struct {
		Signature string `json:"signature"`
		Address   string `json:"address"`
	}

	// SubmitPassportResponse is a data that is returned by Gitcoin Passport with status of request
	SubmitPassportResponse struct {
		ProviderData
		Status string `json:"status"`
	}

	// SubmitPassportRequest is a data that is sent to Gitcoin Passport
	SubmitPassportRequest struct {
		ScorerId int    `json:"scorer_id"`
		Address  string `json:"address"`
	}

	// ProviderData is a data that is returned by Gitcoin Passport
	ProviderData struct {
		Address string `json:"address"`
		Score   string `json:"score"`
	}
)

// Validate is a method that validates VerificationData
func (v VerificationData) Validate() error {
	return validation.Errors{
		"signature": validation.Validate(v.Signature, validation.Required),
		"address":   validation.Validate(v.Address, validation.Required, validation.By(validateAddress)),
	}.Filter()
}

// validateAddress is a validation.RuleFunc that validates address
func validateAddress(value interface{}) error {
	raw, ok := value.(string)
	if !ok {
		return validation.NewError("address", "invalid data type")
	}

	if !common.IsHexAddress(raw) {
		return validation.NewError("address", "invalid address")
	}

	return nil
}
