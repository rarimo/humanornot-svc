package data

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type NonceQ interface {
	Get() (*Nonce, error)
	Select() ([]Nonce, error)
	Insert(value *Nonce) error
	Update(value *Nonce) error
	Delete() error

	WhereEthAddress(addresses ...common.Address) NonceQ
	WhereExpiresAtLt(expiresAt time.Time) NonceQ
	WhereExpiresAtGt(expiresAt time.Time) NonceQ
}

type Nonce struct {
	ID         int64          `db:"id"          structs:"-"`
	Nonce      string         `db:"nonce"       structs:"nonce"`
	ExpiresAt  time.Time      `db:"expires_at"  structs:"expires_at"`
	EthAddress common.Address `db:"eth_address" structs:"eth_address"`
}
