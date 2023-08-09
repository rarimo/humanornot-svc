package civic

import (
	"github.com/ethereum/go-ethereum/common"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
)

var (
	// Interanl errors
	ErrVerifierNotFound = errors.New("verifier not found")

	// Unauthorized errors
	ErrInvalidGatewayToken = errors.New("invalid gateway token")
)

type ChainName string

func (c ChainName) String() string {
	return string(c)
}

const (
	EthereumChainName ChainName = "ethereum"
	PolygonChainName  ChainName = "polygon"
	ArbitrumChainName ChainName = "arbitrum"
	XDCChainName      ChainName = "xdc"
)

var chainNameFromString = map[string]ChainName{
	"ethereum": EthereumChainName,
	"polygon":  PolygonChainName,
	"arbitrum": ArbitrumChainName,
	"xdc":      XDCChainName,
}

type ProviderData struct {
	Address common.Address `json:"address"`
}

type VerificationData struct {
	ChainName string         `json:"chain_name"`
	Signature string         `json:"signature"`
	Address   common.Address `json:"address"`
}

func (v VerificationData) Validate() error {
	return validation.Errors{
		"chain_name": validation.Validate(
			v.ChainName, validation.Required, validation.By(MustBeChainName),
		),
		"signature": validation.Validate(
			v.Signature, validation.Required,
		),
		"address": validation.Validate(
			v.Address, validation.Required,
		),
	}.Filter()
}

func MustBeChainName(value interface{}) error {
	raw, ok := value.(string)
	if !ok {
		return validation.NewError("chain_name", "invalid data type")
	}

	if _, ok = chainNameFromString[raw]; !ok {
		return validation.NewError("chain_name", "invalid chain name")
	}

	return nil
}
