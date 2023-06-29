package data

type NonceQ interface {
	Get() (*Nonce, error)
	Select() ([]Nonce, error)
	Insert(value Nonce) (*Nonce, error)
	Update(value Nonce) (*Nonce, error)
	Delete() error

	FilterByAddress(addresses ...string) NonceQ
	FilterExpired() NonceQ
}

type Nonce struct {
	ID      int64  `db:"id" structs:"-"`
	Message string `db:"message" structs:"message"`
	Expires int64  `db:"expires_at" structs:"expires_at"`
	Address string `db:"address" structs:"address"`
}
