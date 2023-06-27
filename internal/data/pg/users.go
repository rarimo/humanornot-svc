package pg

import (
	"database/sql"
	"fmt"

	"github.com/fatih/structs"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/kit/pgdb"

	sq "github.com/Masterminds/squirrel"
	"github.com/ethereum/go-ethereum/common"

	"gitlab.com/rarimo/identity/kyc-service/internal/data"
)

type usersQ struct {
	db  *pgdb.DB
	sel sq.SelectBuilder
}

func NewUsersQ(db *pgdb.DB) data.UsersQ {
	return &usersQ{
		db:  db.Clone(),
		sel: sq.Select("*").From(usersTableName),
	}
}

func (q *usersQ) Select() ([]data.User, error) {
	var result []data.User

	err := q.db.Select(&result, q.sel)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to select rows")
	}

	return result, nil
}

func (q *usersQ) Get(id uuid.UUID) (*data.User, error) {
	var result data.User

	err := q.db.Get(&result,
		sq.Select("*").
			From(usersTableName).
			Where(sq.Eq{idColumnName: id}),
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to select rows")
	}

	return &result, nil
}

func (q *usersQ) Insert(user *data.User) error {
	err := q.db.Get(&user.ID,
		sq.Insert(usersTableName).
			SetMap(structs.Map(user)).
			Suffix(fmt.Sprintf("returning %s", idColumnName)),
	)
	if err != nil {
		return errors.Wrap(err, "failed to insert rows")
	}

	return nil
}

func (q *usersQ) Update(user *data.User) error {
	err := q.db.Exec(
		sq.Update(usersTableName).
			SetMap(structs.Map(user)).
			Where(sq.Eq{idColumnName: user.ID}),
	)
	if err != nil {
		return errors.Wrap(err, "failed to update rows")
	}

	return nil
}

func (q *usersQ) Sort(sort pgdb.SortedOffsetPageParams) data.UsersQ {
	q.sel = sort.ApplyTo(q.sel, sortColumns)
	return q
}

func (q *usersQ) WhereID(id uuid.UUID) data.UsersQ {
	q.sel = q.sel.Where(sq.Eq{idColumnName: id})
	return q
}

func (q *usersQ) WhereStatus(status data.UserStatus) data.UsersQ {
	q.sel = q.sel.Where(sq.Eq{statusColumnName: status})
	return q
}

func (q *usersQ) WhereEthAddress(address common.Address) data.UsersQ {
	q.sel = q.sel.Where(sq.Eq{ethAddressColumnName: address})
	return q
}
