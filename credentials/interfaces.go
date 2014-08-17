package libgocredentials

import(

)

type AuthenticationInterface interface {
	IsAuthenticated(user, password string) bool
}

type CredentialsInterface interface {
	Get(user string) (bool, UserInterface)
	New(user, password string) UserInterface
	Add(user UserInterface) error
	Modify(user UserInterface) error
	Remove(user UserInterface) error

	Next() (UserInterface, error)
	Reset()
}

type SalterInterface interface {
	// if to == 0, use DEFAULT_SALT_LENGTH
	SetSaltLength(to uint8)
	NewSalt() (string, error)
	GetSalt() string
	SetSalt(to string)
}

type UserInterface interface {
	GetName() string
	GetPassworder() PassworderInterface
	setPassworder(to PassworderInterface)
	ChangePassword(to string) error
	GetPasswordHash() string
	getIndex() uint64
	setIndex(to uint64)
	HasChanged() bool
}

type PassworderInterface interface {
	TestPassword(plain string) bool
	ChangePassword(to string) error
	HasChanged() bool
	GetPasswordHash() string
	GetSalt() string
	GetHashType() string
	format() []byte
}

type ParserInterface interface {
	parse(from string) error
}
