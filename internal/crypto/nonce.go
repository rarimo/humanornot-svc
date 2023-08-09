package crypto

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// NewNonce generates a new nonce.
func NewNonce() (string, error) {
	hash := md5.New()

	if _, err := io.WriteString(hash, strconv.FormatInt(time.Now().Unix(), 10)); err != nil {
		return "", errors.Wrap(err, "failed to write string to hash")
	}

	nonce := make([]byte, 16)
	if _, err := rand.Read(nonce); err != nil {
		return "", errors.Wrap(err, "failed to generate nonce")
	}

	if _, err := io.WriteString(hash, fmt.Sprintf("%x", nonce)); err != nil {
		return "", errors.Wrap(err, "failed to write string to hash")
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// NonceToSignMessage returns a message with nonce to sign.
func NonceToSignMessage(nonce string) string {
	return fmt.Sprintf(
		"User Authentication\n\n"+
			"Below is your one-time nonce to make your authentication secure\n%s",
		nonce,
	)
}
