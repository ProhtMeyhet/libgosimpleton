package iotool

import(
	"errors"
)

// well, duh!
var IsDirectoryError = errors.New("That's a directory!")
var IsNotDirectoryError = errors.New("That's _not_ a directory!")
