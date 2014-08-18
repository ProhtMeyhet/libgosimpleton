package libgocredentials

import(
	"strconv"
	"strings"
	"code.google.com/p/go.crypto/bcrypt"
)

type Passworder struct {
	emptyPassword	bool
	hasChanged	bool

	// password hash with salt
	passwordHash	string
	rawHash		string

	// random salt
	salt		string
	salter		SalterInterface

	// hash type
	hashType	string
	cost		int
}

func NewPassworder() *Passworder {
	return &Passworder{ hashType: "bcrypt" }
}

func NewPassworderParse(from string) (*Passworder, error) {
	passworder := NewPassworder()
	e := passworder.parse(from)
	return passworder, e
}

func NewPassworderParsed(hash, hashType, salt string) *Passworder {
	return &Passworder{ passwordHash: hash,
				salt: salt,
				hashType: hashType,
				}
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

func (passworder *Passworder) GetCost() int {
	if passworder.cost == 0 {
		passworder.cost = BCRYPT_DEFAULT_COST
	}

	return passworder.cost
}

func (passworder *Passworder) SetCost(to int) {
	passworder.cost = to
}

func (passworder *Passworder) TestPassword(plain string) bool {
	e := bcrypt.CompareHashAndPassword([]byte(passworder.passwordHash),
						[]byte(passworder.salt + plain))
	return e == nil
}

func (passworder *Passworder) ChangePassword(plain string) (e error) {
	if plain == "" {
		return PasswordEmptyError
	}

	salt := ""
	if salt, e = passworder.GetSalter().NewSalt(); e != nil {
		return
	}

	toHash := []byte(salt + plain)
	byteHash, e := bcrypt.GenerateFromPassword(toHash, passworder.GetCost())

	if e == nil {
		passworder.passwordHash = string(byteHash)
		passworder.hasChanged = true
		passworder.salt = salt
	}

	return
}

func (passworder *Passworder) HasChanged() bool {
	return passworder.hasChanged
}

func (passworder *Passworder) parse(from string) (e error) {
	splitted := strings.SplitN(from, "$", 2)
	if splitted[0] == "2a" || splitted[0] == "" {
		return InvalidHashFormatError
	}
	passworder.salt = splitted[0]
	passworder.rawHash = splitted[1]

	splitted = strings.Split(splitted[1], "$")
	if len(splitted) != 4 || splitted[0] != "" {
		return InvalidHashFormatError
	}

	if splitted[1] != "2a" {
		return UnknownHashError
	}

	if passworder.cost, e = strconv.Atoi(splitted[2]); e != nil {
		return e
	}

	passworder.passwordHash = splitted[2]

	return
}

func (passworder *Passworder) format() []byte {
	return []byte(passworder.salt + "$" + passworder.passwordHash)
}
