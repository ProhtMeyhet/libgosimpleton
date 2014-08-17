package libgocredentials

import(
	"testing"
)

// generic passworder test
func testPassworder(t *testing.T, name string, passworder PassworderInterface, hashType string,
			faultyHashes []string) {
	password := ""
	if e := passworder.ChangePassword(password); e == nil {
		t.Errorf("%s: ChangePassword() accepts empty password!", name)
	}

	password = "high voltage"
	if e := passworder.ChangePassword(password); e != nil {
		t.Errorf("%s: error changing password: %s", name, e.Error())
	}

	if !passworder.HasChanged() {
		t.Errorf("%s: HasChanged() returned false!", name)
	}

	if passworder.GetHashType() != hashType {
		t.Errorf("%s: got hashType %s, wanted %s", name, passworder.GetHashType(), hashType)
	}

	if passworder.GetSalt() == "" {
		t.Errorf("%s: changed password, but no salt found!", name)
	}

	if len(passworder.GetSalt()) != DEFAULT_SALT_LENGTH {
		t.Errorf("%s: len(GetSalt()) = %v, want %v", name, passworder.GetSalt(),
				DEFAULT_SALT_LENGTH)
	}

	if !passworder.TestPassword(password) {
		t.Errorf("TestPassword failed!")
	}

	if parser, ok := passworder.(ParserInterface); ok {
		for key, hash := range faultyHashes {
			if e := parser.parse(hash); e == nil {
				t.Errorf("%s: parse() accepted wrong format %v", name, key)
			}
		}
	}
}
