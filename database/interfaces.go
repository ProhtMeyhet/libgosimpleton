package database

import(

)

type UserPasswordDatabaseInterface interface {
	UseUser(user string)
	CreateUser(password, salt string) error
	ChangePassword(password, salt string) error
	GetPassword() bool, string
	HasUser() bool
	IsAuthenticated(password, salt string) bool
	DeleteUser() error
	SetHashFunction(hashFunction string) error
}
