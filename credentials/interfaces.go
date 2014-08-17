package libgocredentials

import(
	//"crypto"
)

type AuthenticationInterface interface {
	IsAuthenticated(user, password string) bool
//	IsAuthenticatedPlain(user, password string) bool
}

type CredentialsInterface interface {
	Get(user string) (bool, UserInterface)
	New(user, password string) UserInterface
	Add(user UserInterface) error
	Modify(user UserInterface) error
	Remove(user UserInterface) error

	Next() (UserInterface, error)
	Reset()

	//TODO remove, was for debugging
	Print()
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

/*
credentialstool list messaged.cdb
credentialstool edit messaged.cdb [$user $user ....]
# user: passwordhash : hashtype : salt
# neo : asodfhasdf   : sha256 : 3498444
# edit username press u
# edit password press p
# edit password and hash press h

credentialstool changepw messaged.cdb $user [$user ...]
# enter new password:
# retype new password:

credentialstool changeuser messaged.cdb $user [$user ...]
# enter new username:

credentialstool changehash messaged.cdb $user [$user ...]
# enter new hash:
# enter new password:
# retype new password:

credentialstool userexists messaged.cdb $user [$user ...]

credentialstool testuser messaged.cdb $user [$user ...]
# enter password:
# hurray!
# nay!


credentialstool --type sqlite table@database@something.sqlite

credentialstool --type mysql table@database changeuser $user

*/
