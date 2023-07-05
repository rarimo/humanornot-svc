package crypto

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// NewNonce generates a new nonce.
func NewNonce() (string, error) {
	var (
		hash = md5.New()
		now  = time.Now().Unix()
	)

	if _, err := io.WriteString(hash, strconv.FormatInt(now, 10)); err != nil {
		return "", errors.Wrap(err, "failed to write string to hash")
	}
	if _, err := io.WriteString(hash, strconv.FormatInt(rand.Int63(), 10)); err != nil {
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
