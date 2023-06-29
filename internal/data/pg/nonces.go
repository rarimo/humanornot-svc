package pg

import (
	"database/sql"
	"gitlab.com/rarimo/identity/kyc-service/internal/data"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const nonceTableName = "nonce"

func newNonceQ(db *pgdb.DB) data.NonceQ {
	return &nonceQ{
		db:  db,
		sql: sq.StatementBuilder,
	}
}

type nonceQ struct {
	db  *pgdb.DB
	sql sq.StatementBuilderType
}

func (q *nonceQ) Get() (*data.Nonce, error) {
	var result data.Nonce
	err := q.db.Get(&result, q.sql.Select("*").From(nonceTableName))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to get nonce from db")
	}
	return &result, nil
}

func (q *nonceQ) Select() ([]data.Nonce, error) {
	var result []data.Nonce
	err := q.db.Select(&result, q.sql.Select("*").From(nonceTableName))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to select nonces from db")
	}
	return result, nil
}

func (q *nonceQ) Insert(value data.Nonce) (*data.Nonce, error) {
	clauses := structs.Map(value)

	var result data.Nonce
	stmt := sq.Insert(nonceTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert nonce to db")
	}
	return &result, nil
}

func (q *nonceQ) Update(value data.Nonce) (*data.Nonce, error) {
	clauses := structs.Map(value)

	var result data.Nonce
	stmt := q.sql.Update(nonceTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update nonce in db")
	}
	return &result, nil
}

func (q *nonceQ) Delete() error {
	err := q.db.Exec(q.sql.Delete(nonceTableName))
	if err != nil {
		return errors.Wrap(err, "failed to delete nonces from db")
	}
	return nil
}

func (q *nonceQ) FilterByAddress(addresses ...string) data.NonceQ {
	pred := sq.Eq{"address": addresses}
	q.sql = q.sql.Where(pred)
	return q
}

func (q *nonceQ) FilterExpired() data.NonceQ {
	q.sql = sq.StatementBuilder.Where("expires_at < ?", time.Now().Unix())
	return q
}
