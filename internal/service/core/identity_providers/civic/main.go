package civic

import (
	"encoding/json"
	"math/big"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/logan/v3"

	"gitlab.com/rarimo/identity/kyc-service/internal/config"
	"gitlab.com/rarimo/identity/kyc-service/internal/data"
	"gitlab.com/rarimo/identity/kyc-service/internal/service/core/identity_providers/civic/contracts"
)

type Civic struct {
	log                 *logan.Entry
	GatekeeperNetworkId *big.Int
	ContractsVerifiers  map[ChainName]*contracts.GatewayVerifier
}

func New(log *logan.Entry, config *config.Civic) (*Civic, error) {
	contractsVerifiers, err := newContractsVerifiers(config)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create contracts verifiers")
	}

	return &Civic{
		log:                 log,
		GatekeeperNetworkId: config.GatekeeperNetworkId,
		ContractsVerifiers:  contractsVerifiers,
	}, nil
}

func (c *Civic) Verify(user *data.User, verifyDataRaw []byte) error {
	var verifyData VerificationData
	if err := json.Unmarshal(verifyDataRaw, &verifyData); err != nil {
		return errors.Wrap(err, "failed to unmarshal verification data")
	}

	return nil
}

func newContractsVerifiers(config *config.Civic) (map[ChainName]*contracts.GatewayVerifier, error) {
	ethereumVerifier, err := contracts.NewGatewayVerifier(config.GatewayTokenContract, config.EthereumRpc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Ethereum gateway verifier")
	}

	polygonVerifier, err := contracts.NewGatewayVerifier(config.GatewayTokenContract, config.PolygonRpc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Polygon gateway verifier")
	}

	arbitrumVerifier, err := contracts.NewGatewayVerifier(config.GatewayTokenContract, config.ArbitrumRpc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Arbitrum gateway verifier")
	}

	xdcVerifier, err := contracts.NewGatewayVerifier(config.GatewayTokenContract, config.XDCRpc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create XDC gateway verifier")
	}

	return map[ChainName]*contracts.GatewayVerifier{
		EthereumChainName: ethereumVerifier,
		PolygonChainName:  polygonVerifier,
		ArbitrumChainName: arbitrumVerifier,
		XDCChainName:      xdcVerifier,
	}, nil
}
