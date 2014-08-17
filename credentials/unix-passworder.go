package libgocredentials

import(
	"strings"
)

type UnixPassworder struct {
	Passworder

	hashTypeCCrypt	string
	rawHash		string
}

func NewUnixPassworder() *UnixPassworder {
	return &UnixPassworder{ }
}

func NewUnixPassworderFromString(from string) *UnixPassworder {
	return &UnixPassworder{ rawHash: from }
}

func NewUnixPassworderParse(from string) (*UnixPassworder, error) {
	passworder := &UnixPassworder{ }
	e := passworder.parse(from)
	return passworder, e
}

func (passworder *UnixPassworder) TestPassword(plain string) bool {
	hash := CCrypt(plain, passworder.getFormatedSalt()) // passworder.getFormatedSalt())
	return hash == passworder.rawHash
}

func (passworder *UnixPassworder) ChangePassword(to string) (e error) {
	if to == "" {
		return PasswordEmptyError
	}

	salt, e := passworder.GetSalter().NewSalt()
	if e != nil {
		return
	}

	passworder.hasChanged = true
	passworder.salt = salt

	if passworder.hashTypeCCrypt == "" {
		passworder.hashTypeCCrypt = "6"
		passworder.hashType = "sha512"
	}

	passworder.rawHash = CCrypt(to, passworder.getFormatedSalt())

	return
}

// actually too much parsing, as we only need the crypt salt, but well...
func (passworder *UnixPassworder) parse(from string) (e error) {
	if from == "!" || from == "*" {
		passworder.emptyPassword = true
		return nil
	}

	passworder.rawHash = from

	splitted := strings.Split(from, "$")
	if len(splitted) != 4 {
		return InvalidUnixFormatError
	}

	passworder.hashType,
		passworder.hashTypeCCrypt,
		e = passworder.parseHashType(splitted[1])
	if e != nil {
		return e
	}

	passworder.salt = splitted[2]
	passworder.passwordHash = splitted[3]

	return nil
}

func (passworder *UnixPassworder) parseHashType(from string) (string, string, error) {
	switch(from) {
	case "1":
		return "md5", "1", nil
	case "5":
		return "sha256", "5", nil
	case "6":
		return "sha512", "6",  nil
	}

	return "", "", InvalidUnixFormatError
}

func (passworder *UnixPassworder) format() []byte {
	return []byte(passworder.rawHash)
}

func (passworder *UnixPassworder) getFormatedSalt() string {
	return "$" + passworder.hashTypeCCrypt + "$" + passworder.salt
}
