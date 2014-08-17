package libgocredentials

import(
	"crypto/rand"
	"encoding/hex"
	//"errors"
)

type DefaultSalter struct {
	salt string
	length uint8
}

func NewDefaultSalter() *DefaultSalter {
	return &DefaultSalter{ }
}

func (salt *DefaultSalter) SetSalt(to string) {
	salt.salt = to
}

func (salt *DefaultSalter) GetSalt() string {
	return salt.salt
}

func (salt *DefaultSalter) NewSalt() (random string, e error) {
	if salt.length == 0 {
		salt.length = DEFAULT_SALT_LENGTH
	}

	randomBytes := make([]byte, salt.length / 2)
	_, e = rand.Read(randomBytes)
	random = hex.EncodeToString(randomBytes)

	return
}

func (salt *DefaultSalter) SetSaltLength(to uint8) {
	salt.length = to
}

