package libgocredentials

import(
	"crypto/rand"
	"errors"
)

var lenghtCannotBeZeroError := errors.New("Length can not be zero!")

type DefaultSalt struct {}

func (salt *DefaultSalt) GetSalt(length uint) ([]byte, error) {
	if length == 0 {
		return []byte, lengthCannotBeZeroError
	}

	random := make([]byte, length)
	if _, e : =rand.Read(random); e != nil {
		return []byte, e
	}

	return random, nil
}
