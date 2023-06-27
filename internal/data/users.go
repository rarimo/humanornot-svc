package data

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	iden3core "github.com/iden3/go-iden3-core"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type UsersQ interface {
	Get(id uuid.UUID) (*User, error)
	Select() ([]User, error)
	Insert(user *User) error
	Update(user *User) error

	Sort(sort pgdb.SortedOffsetPageParams) UsersQ
	WhereID(id uuid.UUID) UsersQ
	WhereStatus(status UserStatus) UsersQ
	WhereEthAddress(address common.Address) UsersQ
}

type User struct {
	ID        uuid.UUID  `db:"id"         structs:"id"`
	Status    UserStatus `db:"status"     structs:"status"`
	CreatedAt time.Time  `db:"created_at" structs:"created_at"`

	// IdentityID is a unique user's ident3 identity id
	IdentityID iden3core.ID   `db:"identity_id" structs:"identity_id"`
	EthAddress common.Address `db:"eth_address" structs:"eth_address"`

	// ProviderData Store raw information that received from identity provider.
	// Its json structure depends on identity provider.
	ProviderData []byte `db:"provider_data" structs:"provider_data"`
}

type UserStatus string

const (
	UserStatusInitialized UserStatus = "initialized"
	UserStatusPending     UserStatus = "pending"
	UserStatusVerified    UserStatus = "verified"
)
