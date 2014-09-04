package simplee

import(
	"errors"
)

// tests if format is equal
// if given error is no *E, checks if format is equal to
// interface Error() output
func IsEqual(e, e2 error) bool {
	if simplee, ok := e.(*E); ok {
		return simplee.IsEqual(e2)
	}

	if simplee, ok := e2.(*E); ok {
		return simplee.IsEqual(e)
	}

	return e.Error() == e2.Error()
}

// tests if format is equal and formatted is equal
// if not formatted, they are considered equal now
// if non *E given, tests if formatted equal to interface
// Error() output
func IsSame(e, e2 error) bool {
	if simplee, ok := e.(*E); ok {
		return simplee.IsSame(e2)
	}

	if simplee, ok := e2.(*E); ok {
		return simplee.IsSame(e)
	}

	return e.Error() == e2.Error()
}

// returns a copy of *E
// or, if error given, an errors.New() with interface
// Error() as input
func Copy(e error) error {
	if e == nil {
		return e
	}

	if simplee, ok := e.(*E); ok {
		return simplee.Copy()
	}

	return errors.New(e.Error())
}
