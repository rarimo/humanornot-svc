package unstoppabledomains

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

const (
	EthereumMessageSalt = "\x19Ethereum Signed Message:\n%d%s"
)

func verifyEIP191Signature(signature []byte, message []byte, address common.Address) (bool, error) {
	ecdsaPublicKey, err := crypto.SigToPub(
		crypto.Keccak256(EIP191Prefix(message)),
		signature,
	)
	if err != nil {
		return false, errors.Wrap(err, "failed to recover public key from signature")
	}

	return address == crypto.PubkeyToAddress(*ecdsaPublicKey), nil
}

func EIP191Prefix(data []byte) []byte {
	return []byte(fmt.Sprintf(EthereumMessageSalt, len(data), data))
}
