package libgocredentials

import(
	"errors"
	"io"
)

var EmptyError = errors.New("empty argument!")

var UserExistsError = errors.New("User exists!")
var UserDoesntExistError = errors.New("User does not exist!")

var PlainPasswordNotAvailableError = errors.New("Plain Password wasn't saved!")
var PasswordEmptyError = errors.New("Password is empty!")

var FileNotExistsError = errors.New("File doesn't exists!")
var UnknownHashError = errors.New("Hash unknown!")
var InvalidUnixFormatError = errors.New("Invalid Unix password Format!")
var InvalidHashFormatError = errors.New("Invalid Hash Format!")
var EOFError = io.EOF
