package libgocredentials

import(

)

type AuthenticationInterface interface {
	// authenticate against plain password
	IsAuthenticated(user, password string) bool
}

type CredentialsInterface interface {
	AuthenticationInterface

	// Get a user by string
	Get(user string) (bool, UserInterface)

	// create a new user, but do not commit that user yet
	New(user, password string) (UserInterface, error)

	// add a new user
	// see New()
	Add(user UserInterface) error

	// modify existing user
	// get the user via Get() and change that user
	Modify(user UserInterface) error

	// remove a user
	// take the user from Get()
	Remove(user UserInterface) error

	// read next user
	// always call Reset() before
	// always call Reset() after finished
	Next() (UserInterface, error)

	// must be called when finished with Next()
	Reset()
}

type SalterInterface interface {
	// if to == 0, use DEFAULT_SALT_LENGTH
	SetSaltLength(to uint8)

	// generate a new salt
	NewSalt() (string, error)

	// get current salt
	GetSalt() string

	// set a salt
	// used for password checking
	// not for you to fiddle out your own salt!
	SetSalt(to string)
}

type UserInterface interface {
	// get user name
	GetName() string

	// get passworder
	GetPassworder() PassworderInterface

	// set passworder
	// internal only
	setPassworder(to PassworderInterface)

	// change password
	// set hasChanged
	ChangePassword(to string) error

	// get current password hash
	// shortcut for user.GetPassworder().GetPasswordHash()
	GetPasswordHash() string

	// get an index
	// used internally
	getIndex() uint64
	setIndex(to uint64)

	// has this user been changed
	HasChanged() bool
}

type TransactionInterface interface {
	Begin()
	SetTransactionIsolationLevel(to string)
	Commit()
	Rollback()
}

type PassworderInterface interface {
	// test current password hash against plain string
	TestPassword(plain string) bool

	// change password
	// remeber to use salt!
	ChangePassword(to string) error

	// has this passworder changed
	HasChanged() bool

	// get current password hash
	GetPasswordHash() string

	// get current salt
	GetSalt() string

	// get hash type
	GetHashType() string

	// format password
	format() []byte
}

type ParserInterface interface {
	// well, parse it
	parse(from string) error
}
