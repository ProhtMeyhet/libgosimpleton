package abstract

import(
	"testing"

	"errors"
)

func TestBaseHelper(t *testing.T) {
	errorFunc := func(helper *BaseHelper, message string, e error) {}
	helper := NewBaseHelper(errorFunc)
	testBaseHelper(t, helper, false)

	secondHelper := NewBaseHelper(errorFunc); secondHelper.Copy(helper)
	testBaseHelper(t, secondHelper, true)
}

func testBaseHelper(t *testing.T, helper *BaseHelper, toggleUpsideDown bool) {
	errorCalled := false
	errorFunc := func(helper *BaseHelper, message string, e error) {
		errorCalled = true
	}
	helper.SetE(errorFunc)
	helper.RaiseError("Hamlet, Act III, Scene II", errors.New("I will speak daggers to her, but use none"))
	if !errorCalled {
		t.Error("errorCalled is false!")
	}

	if toggleUpsideDown {
		if helper.ShouldCache() {
			t.Error("helper.ShouldCache should be false, is true")
		}

		helper.ToggleCache()
		if !helper.ShouldCache() {
			t.Error("helper.ShouldCache should be true, is false")
		}
	} else {
		if !helper.ShouldCache() {
			t.Error("helper.ShouldCache should be true, is false")
		}

		helper.ToggleCache()
		if helper.ShouldCache() {
			t.Error("helper.ShouldCache should be false, is true")
		}
	}
}
