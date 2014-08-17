package libgocredentials

import(
	"testing"
)

func TestPassworder(t *testing.T) {
	passworder := NewPassworder()
	salt := "bb3bf04b3853"
	hashType := "bcrypt"

	hash := "bb3bf04b3853$$2a$10$ADAvYz8SqTVh8iJ1nPOW9.l5ipZKNjXWv2Xhd5KNgFhu1rlPydizS"

	hashes := make([]string, 10)
	// missing first $
	hashes[0] = "bb3bf04b3853$2a$10$ADAvYz8SqTVh8iJ1nPOW9.l5ipZKNjXWv2Xhd5KNgFhu1rlPydizS"
	// missing salt
	hashes[1] = "$$2a$10$ADAvYz8SqTVh8iJ1nPOW9.l5ipZKNjXWv2Xhd5KNgFhu1rlPydizS"
	// missing salt and $ separator
	hashes[2] = "$2a$10$ADAvYz8SqTVh8iJ1nPOW9.l5ipZKNjXWv2Xhd5KNgFhu1rlPydizS"
	// cost not int
	hashes[3] = "$2a$af$ADAvYz8SqTVh8iJ1nPOW9.l5ipZKNjXWv2Xhd5KNgFhu1rlPydizS"

	if passworder.GetCost() != BCRYPT_DEFAULT_COST {
		t.Errorf("unexpected cost: %v", passworder.GetCost())
	}

	testPassworder(t, "passworder", passworder, hashType, hashes)

	if e := passworder.parse(hash); e != nil {
		t.Errorf("parse(hash) unexpected error: %s", e.Error())
	}

	if passworder.GetSalt() != salt {
		t.Errorf("parse(hash) salt: got %s, wanted %s", passworder.GetSalt(), salt)
	}
}
