package logging

import(
	"testing"

	"strings"
)

func TestLimitString(t *testing.T) {
	s := strings.Repeat("Corcoran", 15)
	ss := LimitString(s)

	if len(ss) != 55 {
		t.Errorf("expected ss to be 55 chars long, got %v\n", len(ss))
	}

	if ss[50:55] != "....." {
		t.Errorf("ss ends with %v, wanted .....", ss[50:55])
	}

	s = "Cousin Hebe"
	ss = LimitString(s)
	if s != ss {
		t.Errorf("expected ss to be unchanged from s '%v', but got '%v'\n", ss)
	}
}

func TestSanitizeString(t *testing.T) {
	s := "abc\n d  d\t"
	ss := SanitizeString(s)
	expected := "abc d d"
	if ss != expected {
		t.Errorf("expected SanitizedString %v, got %v\n", expected, ss)
	}
}
