package libgocredentials

import(

)

type UnixUser struct {
	User

	daysPasswordChanged		string
	daysPasswordCanBeChanged	string
	daysPasswordMustBeChanged	string
	daysPasswordChangeWarning	string
	daysExpiration			string
	daysExpiered			string
	reserved			string
}

func NewUnixUser(username string) (user *UnixUser) {
	user = &UnixUser{ }
	user.name = username
	return
}

func CreateUnixUser(username, password string) (user *UnixUser, e error) {
	user = &UnixUser{ }
	user.name = username

	// set some default values
	// taken from openSuse
	user.daysPasswordChanged =		"16000"
	user.daysPasswordCanBeChanged =		"0"
	user.daysPasswordMustBeChanged =	"99999"
	user.daysPasswordChangeWarning =	"7"

	user.setPassworder(NewUnixPassworder())
	e = user.GetPassworder().ChangePassword(password)

	return
}
