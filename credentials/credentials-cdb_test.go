package libgocredentials

import(
	"testing"
)

func TestCdb(t *testing.T) {
	cdb := NewCdb("/tmp/test-cdb.cdb")

	genericCredentialsTest(cdb, t)
}

func genericCredentialsTest(credentials CredentialsInterface, t *testing.T) {
	if e := credentials.Add("testuser", "testpw"); e != nil {
		t.Errorf("couldn't add user! %s", e.Error())
	}

	authenticated, e := credentials.IsAuthententicated("testuser",
				credentials.GetHash("testpw"))
	if !authenticated || e != nil {
		t.Errorf("couldn't authenticate! %s", e.Error())
	}

}
