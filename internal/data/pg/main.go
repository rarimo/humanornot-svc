package pg

import (
	"gitlab.com/distributed_lab/kit/pgdb"

	"gitlab.com/rarimo/identity/kyc-service/internal/data"
)

const (
	usersTableName = "users"
)

const (
	idColumnName         = "id"
	statusColumnName     = "status"
	ethAddressColumnName = "eth_address"
	identityIDColumnName = "identity_id"
)

var sortColumns = map[string]string{
	statusColumnName: statusColumnName,
}

type masterQ struct {
	db *pgdb.DB
}

func NewMasterQ(db *pgdb.DB) data.MasterQ {
	return &masterQ{
		db: db.Clone(),
	}
}

func (q *masterQ) New() data.MasterQ {
	return NewMasterQ(q.db)
}

func (q *masterQ) UsersQ() data.UsersQ {
	return NewUsersQ(q.db)
}

func (q *masterQ) NonceQ() data.NonceQ {
	return newNonceQ(q.db)
}

func (q *masterQ) Transaction(fn func() error) error {
	return q.db.Transaction(func() error {
		return fn()
	})
}
