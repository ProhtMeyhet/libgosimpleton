package simplee

import(
	"fmt"
	"testing"
)

func TestSimplee(t *testing.T) {
	p1, p2 := "Apple is %s", "crap"
	expected := fmt.Sprintf(p1, p2)
	e := New(p1)

	if message := doError(e).Error(); message != expected {
		t.Errorf("expected '%s', got '%s'", expected, message)
	}
}

func doError(e *ErrorF) error {
	return e.Format("crap")
}
