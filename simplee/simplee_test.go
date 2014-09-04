package simplee

import(
	"fmt"
	"errors"
	"testing"
)

var p1, p2, p3 = "Apple is %s", "crap", "still crap"

func TestSimplee(t *testing.T) {
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

	if e.IsEqual(nil) {
		t.Errorf("unexpected: e == nil")
	}

	if !e.IsEqual(doError2(e)) {
		t.Errorf("unexpected: e != doError2( e )!")
	}

	//IsSame()
	if e.IsSame(nil) {
		t.Errorf("unexpected: e.IsSame( nil )")
	}

	if !IsEqual(doError(e), doError2(e)) {
		t.Errorf("unexpected: doError( e ) != doError2( e )!")
	}

	if IsSame(doError(e), doError2(e)) {
		t.Errorf("unexpected: doError( e ) IsSame doError2( e ) )!")
	}

	if e.IsSame(doError2(e)) {
		t.Errorf("unexpected: e IsSame doError2( e )!")
	}
}

func doError(e *E) error {
	return e.Format(p2)
}

func doError2(e *E) error {
	return e.Format(p3)
}
