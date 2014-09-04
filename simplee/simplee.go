package simplee

import(
	"fmt"
)

// E represents an error
type E struct {
	format, formatted string
}

// create a new error with string format
func New(to string) *E {
	if to == "" {
		panic("cannot initialise empty error!")
	}

	return &E{ format: to }
}

// returns a copy of this *E
func (e *E) Copy() *E {
	eNew := New(e.format)
	eNew.formatted = e.formatted
	return eNew
}

// returns format
func (e *E) GetFormat() string {
	return e.format
}

// formats string
func (e *E) GetFormatted(what ...interface{}) string {
	return fmt.Sprintf(e.format, what...)
}

// format error and returns a copy of that error
func (e *E) Format(what ...interface{}) *E {
	eCopy := e.Copy()
	eCopy.formatted = eCopy.GetFormatted(what...)
	return eCopy
}

// interface function, return formatted if possible
// otherwise returns format
func (e *E) Error() string {
	if e.formatted != "" {
		return e.formatted
	}

	return e.format
}

// tests if format is equal
// if given error is no *E, checks if format is equal to
// interface Error() output
func (e *E) IsEqual(to error) bool {
	if to == nil {
		return false
	}

	if ef, ok := to.(*E); ok {
		return e.format == ef.format
	}

	return e.format == to.Error()
}

// tests if format is equal and formatted is equal
// if not formatted, they are considered equal now
// if non *E given, tests if formatted equal to interface
// Error() output
func (e *E) IsSame(to error) bool {
	if to == nil {
		return false
	}

	// if both havn't been formatted, they are the same
	if ef, ok := to.(*E); ok {
		if e.format == ef.format {
			return e.formatted == ef.formatted
		} else {
			return false
		}
	}

	return e.formatted == to.Error()
}
