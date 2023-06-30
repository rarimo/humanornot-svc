package crypto

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"regexp"
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
	decodedSignature, err := decodeSignature(signature)
	if err != nil {
		return false, errors.Wrap(err, "failed to decode signature")
	}

	ecdsaPublicKey, err := crypto.SigToPub(
		crypto.Keccak256(EIP191Prefix([]byte(rawMessage))),
		decodedSignature,
	)
	if err != nil || ecdsaPublicKey == nil {
		return false, errors.Wrap(err, "failed to recover public key from signature")
	}

	return address == crypto.PubkeyToAddress(*ecdsaPublicKey), nil
}

// DecodeSignature decodes a signature from a hex string.
func decodeSignature(signature string) ([]byte, error) {
	signatureBytes, err := hexutil.Decode(signature)
	if err != nil {
		return nil, err
	}

	if len(signatureBytes) != 65 {
		return nil, errors.New("bad signature length")
	}

	if signatureBytes[64] != 27 && signatureBytes[64] != 28 {
		return nil, errors.New("bad recovery byte")
	}

	signatureBytes[64] -= 27

	return signatureBytes, nil
}

// EIP191Prefix returns the EIP-191 prefix for a message.
func EIP191Prefix(data []byte) []byte {
	return []byte(fmt.Sprintf(EthereumMessageSalt, len(data), data))
}
