package libgocredentials

import(
	"testing"
)

func TestUnixPassworder(t *testing.T) {
	passworder := NewUnixPassworder()
	salt := "accadacca"
	hashType := "sha512"

	hash := "$6$accadacca$S52Nmn0XM83gWlM6j6LD.xnd.RZzqNKn7qqiGy/k5CCe.ZrHtYjnGsjwXfVfvMzHW9M4YZZdbBRnienrX66Te."

	hashes := make([]string, 10)
	// missing first $
	hashes[1] = "6$accadacca$S52Nmn0XM83gWlM6j6LD.xnd.RZzqNKn7qqiGy/k5CCe.ZrHtYjnGsjwXfVfvMzHW9M4YZZdbBRnienrX66Te."
	// missing hashType $6$
	hashes[2] = "accadacca$S52Nmn0XM83gWlM6j6LD.xnd.RZzqNKn7qqiGy/k5CCe.ZrHtYjnGsjwXfVfvMzHW9M4YZZdbBRnienrX66Te."
	// missing salt
	hashes[3] = "$6$S52Nmn0XM83gWlM6j6LD.xnd.RZzqNKn7qqiGy/k5CCe.ZrHtYjnGsjwXfVfvMzHW9M4YZZdbBRnienrX66Te."
	// missing hashType and salt
	hashes[4] = "$S52Nmn0XM83gWlM6j6LD.xnd.RZzqNKn7qqiGy/k5CCe.ZrHtYjnGsjwXfVfvMzHW9M4YZZdbBRnienrX66Te."

	if e := passworder.parse(hash); e != nil {
		t.Errorf("parse(hash) unexpected error: %s", e.Error())
	}

	if passworder.GetSalt() != salt {
		t.Errorf("parse(hash) salt: got %s, wanted %s", passworder.GetSalt(), salt)
	}

	testPassworder(t, "unix", passworder, hashType, hashes)
}
