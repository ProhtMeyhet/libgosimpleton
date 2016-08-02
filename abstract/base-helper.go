package abstract

import(
	"io"
	"os"
)

// used to reduce function parameters
type BaseHelper struct {
	// try not to leave caches (especially file caches) behind
	noCache		bool

	// call this function on every error
	onE		func(*BaseHelper, string, error)
}

// shiny and fresh
func NewBaseHelper(aOnE func(*BaseHelper, string, error)) (helper *BaseHelper) {
	helper = &BaseHelper{}
	helper.Initialise(aOnE)
	return
}

// i n i t
func (helper *BaseHelper) Initialise(aOnE func(*BaseHelper, string, error)) {
	if aOnE == nil {
		helper.onE = func(*BaseHelper, string, error) {}
	} else {
		helper.onE = aOnE
	}
}

// cache you should
func (helper *BaseHelper) ShouldCache() bool {
	return !helper.noCache
}

// toggle cache you should
func (helper *BaseHelper) ToggleCache() *BaseHelper {
	helper.noCache = !helper.noCache
	return helper
}

// set error function
func (helper *BaseHelper) SetE(to func(*BaseHelper, string, error)) *BaseHelper {
	helper.onE = to
	return helper
}

// raise an error
func (helper *BaseHelper) RaiseError(name string, e error) {
	if helper.onE != nil {
		helper.onE(helper, name, e)
	}
}

func (helper *BaseHelper) Copy(from interface{}) *BaseHelper {
	if baseHelper, ok := from.(*BaseHelper); !ok {
		panic("could not cast to *BaseHelper !")
	} else {
		helper.noCache = baseHelper.noCache
		helper.onE = baseHelper.onE
	}

	return helper
}

// just write plainly to stderr
func WriteErrorsToStderr(helper *BaseHelper, name string, e error) {
	io.WriteString(os.Stderr, "error with: " + name + ": " + e.Error())
}

// ignore errors. doesn't lift a finger.
func IgnoreErrors(helper *BaseHelper, name string, e error) { }
