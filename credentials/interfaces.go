package libgocredentials

import(

)

type AuthenticationInterface interface {
	IsAuthenticated(user, password) bool
}

type CredentialsInterface interface {
	GetHashType() uint
	SetHashType(to uint) error
	GetSaltType() SaltInterface
	SetSaltType(to SaltInterface)

	GetHash(from string) string

	Add(user, password string) error
	ChangePassword(user, password string) error
}


type SaltInterface interface {
	GetSalt(lenght int) string
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
