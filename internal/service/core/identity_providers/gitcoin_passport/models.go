package gcpsp

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"gitlab.com/rarimo/identity/kyc-service/internal/service/api/requests"
	"gitlab.com/rarimo/identity/kyc-service/resources"
)

const (
	defaultBaseURL         = "https://api.scorer.gitcoin.co/registry"
	submitPassportEndpoint = "/submit-passport"
	scoreEndpoint          = "/score"
	defaultRetryCount      = 3
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
	VerificationData resources.GitcoinPassportData

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

var (
	ErrUnexpectedStatus = errors.New("unexpected status")

	ErrUnexpectedStatusCode = errors.New("received unexpected status code")

	ErrScoreIsTooLow = errors.New("score is too low")
)

// Validate is a method that validates VerificationData
func (v VerificationData) Validate() error {
	return validation.Errors{
		"signature": validation.Validate(v.Signature, validation.Required),
		"address":   validation.Validate(v.Address, validation.Required, validation.By(requests.MustBeEthAddress)),
	}.Filter()
}
