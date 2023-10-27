package data

import (
	"database/sql/driver"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	iden3core "github.com/iden3/go-iden3-core"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type UsersQ interface {
	Get() (*User, error)
	Select() ([]User, error)
	Insert(user *User) error
	Upsert(user *User) error
	Update(user *User) error

	Sort(sort pgdb.SortedOffsetPageParams) UsersQ
	WhereID(id uuid.UUID) UsersQ
	WhereStatus(status UserStatus) UsersQ
	WhereIdentityID(id *IdentityID) UsersQ
	WhereEthAddress(address common.Address) UsersQ
	WhereProviderHash(providerHash []byte) UsersQ
}

type User struct {
	ID        uuid.UUID  `db:"id"         structs:"id"`
	Status    UserStatus `db:"status"     structs:"status"`
	CreatedAt time.Time  `db:"created_at" structs:"created_at"`

	// IdentityID is a unique user's ident3 identity id
	IdentityID *IdentityID     `db:"identity_id" structs:"-"`
	EthAddress *common.Address `db:"eth_address" structs:"eth_address"`

	// ProviderData Store raw information that received from identity provider.
	// Its json structure depends on identity provider.
	ProviderData []byte `db:"provider_data" structs:"provider_data"`

	// ProviderHash is used for data uniqueness check.
	ProviderHash []byte    `db:"provider_hash" structs:"provider_hash"`
	ClaimID      uuid.UUID `db:"claim_id" structs:"claim_id"`
}

type UserStatus string

func (s UserStatus) String() string {
	return string(s)
}

const (
	UserStatusInitialized UserStatus = "initialized"
	UserStatusPending     UserStatus = "pending"
	UserStatusVerified    UserStatus = "verified"
	UserStatusUnverified  UserStatus = "unverified"
)

type IdentityID struct {
	iden3core.ID
}

func NewIdentityID(id iden3core.ID) *IdentityID {
	return &IdentityID{id}
}

func (id *IdentityID) Value() (driver.Value, error) {
	return id.Bytes(), nil
}

func (id *IdentityID) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion src.([]byte) failed")
	}

	identityID, err := iden3core.IDFromBytes(source)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal identityID from binary")
	}

	*id = IdentityID{
		ID: identityID,
	}

	return nil
}
