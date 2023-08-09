package data

type MasterQ interface {
	New() MasterQ

	UsersQ() UsersQ
	NonceQ() NonceQ

	Transaction(func() error) error
}
