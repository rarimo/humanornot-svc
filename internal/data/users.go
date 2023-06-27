package data

import (
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
	ID     uuid.UUID  `json:"id"     db:"id"`
	Status UserStatus `json:"status" db:"status"`

	// IdentityID is a unique user's ident3 identity id
	IdentityID iden3core.ID   `json:"identity_id" db:"identity_id"`
	EthAddress common.Address `json:"eth_address" db:"eth_address"`

	// ProviderData Store raw information that received from identity provider.
	// Its json structure depends on identity provider.
	ProviderData []byte `json:"provider_data" db:"provider_data"`
}

type UserStatus string

const (
	UserStatusInitialized UserStatus = "initialized"
	UserStatusPending     UserStatus = "pending"
	UserStatusVerified    UserStatus = "verified"
)
