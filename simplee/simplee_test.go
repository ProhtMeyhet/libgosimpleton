package simplee

import(
	"fmt"
	"errors"
	"testing"
)

func TestSimplee(t *testing.T) {
	p1, p2 := "Apple is %s", "crap"
	expected := fmt.Sprintf(p1, p2)
	e := New(p1)

	if message := doError(e).Error(); message != expected {
		t.Errorf("expected '%s', got '%s'", expected, message)
	}

	if e2 := New(p1); !e.IsEqual(e2) {
		t.Errorf("unexpected: e != e2")
	}

	if e3 := errors.New(p1); !e.IsEqual(e3) {
		t.Errorf("unexpected: e != e3")
	}

	if e4 := New(p2); e.IsEqual(e4) {
		t.Errorf("unexpected: e == e4")
	}

	if e5 := errors.New(p2); e.IsEqual(e5) {
		t.Errorf("unexpected: e == e5")
	}
}

func doError(e *ErrorF) error {
	return e.Format("crap")
}
