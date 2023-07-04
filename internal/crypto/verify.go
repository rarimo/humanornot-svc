package crypto

import (
	"fmt"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

const (
	EthereumMessageSalt = "\x19Ethereum Signed Message:\n%d%s"
)

var (
	AddressRegexp = regexp.MustCompile("^(0x)?[0-9a-fA-F]{40}$")
)

// VerifyEIP191Signature verifies the signature of a message using the EIP-191 standard.
// Returns true if the signature is valid, false otherwise.
func VerifyEIP191Signature(signature string, rawMessage string, address common.Address) (bool, error) {
	decodedSignature, err := hexutil.Decode(signature)
	if err != nil {
		return false, errors.Wrap(err, "failed to decode signature")
	}

	decodedSignature[64] -= 27 // Transform yellow paper V from 27/28 to 0/1

	ecdsaPublicKey, err := crypto.SigToPub(
		crypto.Keccak256(EIP191Prefix([]byte(rawMessage))),
		decodedSignature,
	)
	if err != nil || ecdsaPublicKey == nil {
		return false, errors.Wrap(err, "failed to recover public key from signature")
	}

	return address == crypto.PubkeyToAddress(*ecdsaPublicKey), nil
}

// EIP191Prefix returns the EIP-191 prefix for a message.
func EIP191Prefix(data []byte) []byte {
	return []byte(fmt.Sprintf(EthereumMessageSalt, len(data), data))
}
