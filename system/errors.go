package system

import(
	"errors"
)

var (
	UNEXPECTED_LOADAVG_FORMAT_ERROR =	errors.New("unexepcted /proc/loadavg format")
)
