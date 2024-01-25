package kleros

import (
	"encoding/json"
	"time"

	"github.com/ethereum/go-ethereum/common"
	cryptoPkg "github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/logan/v3"

	"github.com/rarimo/kyc-service/internal/config"
	"github.com/rarimo/kyc-service/internal/crypto"
	"github.com/rarimo/kyc-service/internal/data"
	providers "github.com/rarimo/kyc-service/internal/service/core/identity_providers"
	"github.com/rarimo/kyc-service/internal/service/core/identity_providers/kleros/contracts"
	"github.com/rarimo/kyc-service/internal/service/core/issuer"
)

type Kleros struct {
	logger                  *logan.Entry
	masterQ                 data.MasterQ
	proofOfHumanityContract *contracts.ProofOfHumanity
}

func NewIdentityProvider(
	logger *logan.Entry, masterQ data.MasterQ, config *config.Kleros,
) (*Kleros, error) {
	pohContract, err := contracts.NewProofOfHumanity(config.ProofOfHumanityContract, config.EthereumRpc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create proof of humanity contract")
	}

	return &Kleros{
		logger:                  logger,
		masterQ:                 masterQ,
		proofOfHumanityContract: pohContract,
	}, nil
}

func (k *Kleros) Verify(
	user *data.User, verifyDataRaw []byte,
) (*issuer.IdentityProvidersCredentialSubject, []byte, error) {
	var verifyData VerificationData
	if err := json.Unmarshal(verifyDataRaw, &verifyData); err != nil {
		return nil, nil, errors.Wrap(err, "failed to unmarshal verification data")
	}

	if err := verifyData.Validate(); err != nil {
		return nil, nil, providers.ErrInvalidVerificationData
	}

	userAddr := common.HexToAddress(verifyData.Address)

	if err := k.verifySignature(verifyData.Signature, userAddr); err != nil {
		return nil, nil, errors.Wrap(err, "failed to verify signature")
	}

	if err := k.checkIfIsRegistered(userAddr); err != nil {
		return nil, nil, errors.Wrap(err, "failed to check if user is registered")
	}

	providerDataRaw, err := json.Marshal(ProviderData{
		Address: userAddr,
	})
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to marshal provider data")
	}

	user.EthAddress = &userAddr
	user.Status = data.UserStatusVerified
	user.ProviderData = providerDataRaw

	credentialSubject := issuer.NewEmptyIdentityProvidersCredentialSubject()
	credentialSubject.Provider = issuer.KlerosProviderName
	credentialSubject.Address = userAddr.String()
	credentialSubject.ProviderMetadata = "none"

	return credentialSubject, cryptoPkg.Keccak256(
		userAddr.Bytes(),
		providers.KlerosIdentityProvider.Bytes(),
	), nil
}

// verifySignature verifies user's signature
func (k *Kleros) verifySignature(signature string, userAddr common.Address) error {
	nonce, err := k.masterQ.NonceQ().
		WhereEthAddress(userAddr).
		WhereExpiresAtGt(time.Now()).
		Get()
	if err != nil {
		return errors.Wrap(err, "failed to get nonce")
	}
	if nonce == nil {
		return providers.ErrNonceNotFound
	}

	valid, err := crypto.VerifyEIP191Signature(
		signature,
		crypto.NonceToSignMessage(nonce.Nonce),
		userAddr,
	)
	if err != nil || !valid {
		return providers.ErrInvalidUsersSignature
	}

	return errors.Wrap(k.masterQ.NonceQ().WhereEthAddress(userAddr).Delete(), "failed to delete nonce")
}

func (k *Kleros) checkIfIsRegistered(userAddress common.Address) error {
	isRegistered, err := k.proofOfHumanityContract.IsRegistered(nil, userAddress)
	if err != nil {
		return errors.Wrap(err, "failed to call isRegistered contract method")
	}
	if !isRegistered {
		return ErrIsNotRegistered
	}
	return nil
}
