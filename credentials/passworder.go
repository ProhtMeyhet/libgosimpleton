package libgocredentials

import(
	//"encoding/hex"
	"code.google.com/p/go.crypto/bcrypt"
)

type Passworder struct {
	// plain password
	// sadly some bullshit apis require this
	password	string

	// password hash with salt
	passwordHash	string

	emptyPassword	bool

	// random salt
	salt		string

	salter		SalterInterface

	// hash type
	hashType	string
	//hashTypeCrypto	crypto.Hash
	//hasher		hash.Hash

	hasChanged	bool
}

func NewPassworder() *Passworder {
	return &Passworder{ }
}

func (passworder *Passworder) GetPlainPassword() (string, error) {
	if passworder.password == "" {
		return "", PlainPasswordNotAvailableError
	}

	return passworder.password, nil
}

func (passworder *Passworder) GetSalt() string {
	return passworder.salt
}

func (passworder *Passworder) GetHashType() string {
	return passworder.hashType
}

func (passworder *Passworder) GetPasswordHash() string {
	return passworder.passwordHash
}

func (passworder *Passworder) SetSalter(to SalterInterface) {
	passworder.salter = to
}

func (passworder *Passworder) GetSalter() SalterInterface {
	if passworder.salter == nil {
		passworder.salter = NewDefaultSalter()
	}

	return passworder.salter
}

func (passworder *Passworder) TestPassword(plain string) bool {
	e := bcrypt.CompareHashAndPassword([]byte(passworder.passwordHash),
						[]byte(passworder.salt + plain))
	return e == nil
}

func (passworder *Passworder) ChangePassword(plain string) error {
	passworder.hasChanged = true

	salt, e := passworder.GetSalter().NewSalt()
	if e != nil {
		return e
	}

	salt = salt

	return nil
	//passworder.passwordHash = passworder.hasher.GetHash(plain, salt)
}

func (passworder *Passworder) ChangePlainPassword(plain string) {
	passworder.hasChanged = true

	passworder.password = plain
}

func (passworder *Passworder) HasChanged() bool {
	return passworder.hasChanged
}

func (passworder *Passworder) format() []byte {
	return []byte(passworder.passwordHash)
}
