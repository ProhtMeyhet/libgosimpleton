package libgocredentials

import(
	"errors"
)

var EmptyError = errors.New("empty argument!")

var UserExistsError = errors.New("User exists!")
var UserDoesntExistError = errors.New("User does not exist!")

var PlainPasswordNotAvailableError = errors.New("Plain Password wasn't saved!")

var FileNotExistsError = errors.New("File doesn't exists!")
