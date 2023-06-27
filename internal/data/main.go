package data

type MasterQ interface {
	New() MasterQ

	VerificationUsersQ() UsersQ

	Transaction(func() error) error
}
