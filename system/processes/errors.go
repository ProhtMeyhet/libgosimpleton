package processes

import(
	"errors"
)

var(
	INVALID_PROCESS_ID_ERROR =		errors.New("invalid process")
	NO_SUCH_PROCESS_ERROR =			errors.New("no such process")
	UNEXPECTED_STAT_FORMAT_ERROR =		errors.New("unexepcted /proc/[pid]/stat format")
	UNEXPECTED_CMDLINE_ERROR =		errors.New("unexepcted error reading /proc/[pid]/cmdline")
)
