package data

type MasterQ interface {
	New() MasterQ

	UsersQ() UsersQ

	Transaction(func() error) error
}
