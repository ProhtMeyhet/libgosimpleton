package database

import(

)


type UserPasswordDatabaseFile struct {
	user string
}

func NewUserPasswordDatabaseFile(Namespace, Table, User string) *UserPasswordDatabaseFile {
	return &UserPasswordDatabaseFile{ user: User }
}

func (database *UserPasswordDatabaseFile) UseUser(user string) {

}

func (database *UserPasswordDatabaseFile) CreateUser(password, salt string) error {

}

func (database *UserPasswordDatabaseFile) ChangePassword(password, salt string) error {

}

func (database *UserPasswordDatabaseFile) GetPassword() (bool, string) {

}

func (database *UserPasswordDatabaseFile) HasUser() bool {

}

func (database *UserPasswordDatabaseFile) IsAuthenticated(password, salt string) bool {

}

func (database *UserPasswordDatabaseFile) DeleteUser() error {

}
