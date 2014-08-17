package libgocredentials

import(
	//"fmt"
	"testing"
	"io"
	"io/ioutil"
	"os"
	//"time"
)

func TestCredentialsUnix(t *testing.T) {
	tempFile, e := ioutil.TempFile(os.TempDir(), "testing")
	if e != nil {
		t.Errorf("couldn't create temp file! %s", e.Error())
	}
	tempFile.Close()
	defer os.Remove(tempFile.Name())

	unix := NewUnix(tempFile.Name())
	defer unix.Close()

	user1 := &Users {
		name: "brian johnson",
		password: "snowballed",
	}
	user1.user = CreateUnixUser(user1.name, user1.password)

	user2 := &Users {
		name: "cliff williams",
		password: "whole lotta rosie",
	}
	user2.user = CreateUnixUser(user2.name, user2.password)

	user3 := &Users {
		name: "phil rudd",
		password: "ballbreaker",
	}
	user3.user = CreateUnixUser(user3.name, user3.password)

	testAddUsers(t, unix, user1, user2, user3)
	testNext(t, unix, user1, user2, user3)

	e = unix.Remove(user2.user)
	if e != nil {
		t.Errorf("error removing user! %s", e.Error())
	}

	if unix.IsAuthenticated(user2.name, user2.password) {
		t.Errorf("removed user %s, but could authenticate!", user2.name)
	}

	// test if rest of file was written
	if !unix.IsAuthenticated(user3.name, user3.password) {
		t.Errorf("couldn't authenticate user %v after remove!", user3.name)
	}


	testNext(t, unix, user1, user3)
	testAddUsers(t, unix, user2)
	testNext(t, unix, user1, user3, user2)

	user3.newPassword = "stiff upper lip"
	if e = user3.user.ChangePassword(user3.newPassword); e != nil {
		t.Errorf("couldn't change password for user %s! %s",
				user3.name, e.Error())
	}

	if e = unix.Modify(user3.user); e != nil {
		t.Errorf("couldn't update password for user %s! %s",
				user3.name, e.Error())
	}

	if unix.IsAuthenticated(user3.name, user3.password) {
		t.Errorf("was able to authenticate with old password for user %s!",
				user3.name)
	}

	if !unix.IsAuthenticated(user3.name, user3.newPassword) {

		t.Errorf("couldn't authenticate with new password for user %s!",
				user2.name)
	}

	// test if rest of file was written
	if !unix.IsAuthenticated(user2.name, user2.password) {
		t.Errorf("couldn't authenticate user %v after modify!", user3.name)
	}

	// refresh user 2
	var found bool
	if found, user2.user = unix.Get(user2.name); !found {
		t.Errorf("couldn't get user2")
	}

	if !unix.IsAuthenticated(user2.name, user2.password) {
		t.Errorf("couldn't authenticate user %v after refresh!", user3.name)
	}
}

func testAddUsers(t *testing.T, unix *Unix, users ...*Users) {
	for key, user := range users {
		if e := unix.Add(user.user); e != nil {
			t.Errorf("couldn't add user %v! %s", user.name, e.Error())
		}

		// add again, must return error
		if e := unix.Add(user.user); e == nil {
			t.Errorf("added the same user %v twice without error!", key)
		}

		if !unix.IsAuthenticated(user.name, user.password) {
			t.Errorf("couldn't authenticate user %v!", key)
		}

		//index testing
		if key != 0 && user.user.getIndex() == 0 {
			t.Errorf("index 0 on user %s!", user.name)
		}
	}
}

func testNext(t *testing.T, unix *Unix, users ...*Users) {
	unix.Reset()
	for _, user := range users {
		if next, e := unix.Next(); e != nil {
			if e == io.EOF {
				t.Errorf("Next() unexpected EOF!")
			} else {
				t.Errorf("Next() %s", e.Error())
			}
		} else {
			if user.name != next.GetName() {
				t.Errorf("Next() %s, expected %s", next.GetName(),
						user.name)
			}
		}
	}
}

type Users struct {
	name, password, newPassword	string
	user				UserInterface
}
