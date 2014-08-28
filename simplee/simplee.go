package simplee

import(
	"fmt"
)

type ErrorF struct {
	format, formatted string
}

func New(to string) *ErrorF {
	return &ErrorF{ format: to }
}

func (errorF *ErrorF) Format(what ...interface{}) *ErrorF {
	if len(what) > 0 {
		errorF.formatted = fmt.Sprintf(errorF.format, what...)
	}

	return errorF
}

func (errorF *ErrorF) Error() string {
	if errorF.formatted != "" {
		return errorF.formatted
	}

	return errorF.format
}
