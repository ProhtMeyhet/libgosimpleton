package libgocredentials

import(

)

type User struct {
	name		string
	passworder	PassworderInterface
	hasChanged	bool

	index		uint64
}

func NewUser() *User {
	return &User{ }
}

func CreateUser(to string) *User {
	return &User{ name: to }
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) GetPasswordHash() string {
	return user.GetPassworder().GetPasswordHash()
}

func (user *User) ChangePassword(to string) error {
	user.hasChanged = true
	return user.GetPassworder().ChangePassword(to)
}

func (user *User) GetPassworder() PassworderInterface {
	if user.passworder == nil {
		user.passworder = NewPassworder()
	}

	return user.passworder
}

func (user *User) setPassworder(to PassworderInterface) {
	user.passworder = to
}

func (user *User) getIndex() uint64 {
	return user.index
}

func (user *User) setIndex(to uint64) {
	user.index = to
}

func (user *User) HasChanged() bool {
	return user.hasChanged
}
