package libgocredentials

import(
	//"fmt"
	"testing"
	"io"
	"io/ioutil"
	"os"
	//"time"
)

func TestUnix(t *testing.T) {
	tempFile, e := ioutil.TempFile(os.TempDir(), "testing")
	if e != nil {
		t.Errorf("couldn't create temp file! %s", e.Error())
	}
	tempFile.Close()
	defer os.Remove(tempFile.Name())

	unix := NewUnix(tempFile.Name())
	defer unix.Close()

	genericCredentialsTest(t,unix)
}

func genericCredentialsTest(t *testing.T, credentials CredentialsInterface) {
	var e error

	user1 := &Users {
		name: "brian johnson",
		password: "snowballed",
	}
	if user1.user, e = credentials.New(user1.name, user1.password); e != nil {
	    t.Errorf("Error creating user: %s", e.Error())
	}

	user2 := &Users {
		name: "cliff williams",
		password: "whole lotta rosie",
	}
	if user2.user, e = credentials.New(user2.name, user2.password); e != nil {
	    t.Errorf("Error creating user: %s", e.Error())
	}

	user3 := &Users {
		name: "phil rudd",
		password: "ballbreaker",
	}
	if user3.user, e = credentials.New(user3.name, user3.password); e != nil {
	    t.Errorf("Error creating user: %s", e.Error())
	}

	testAddUsers(t, credentials, user1, user2, user3)
	testNext(t, credentials, user1, user2, user3)

	e = credentials.Remove(user2.user)
	if e != nil {
		t.Errorf("error removing user! %s", e.Error())
	}

	if credentials.IsAuthenticated(user2.name, user2.password) {
		t.Errorf("removed user %s, but could authenticate!", user2.name)
	}

	// test if rest of file was written
	if !credentials.IsAuthenticated(user3.name, user3.password) {
		t.Errorf("couldn't authenticate user %v after remove!", user3.name)
	}

	testNext(t, credentials, user1, user3)
	testAddUsers(t, credentials, user2)
	testNext(t, credentials, user1, user3, user2)

	user3.newPassword = "stiff upper lip"
	if e = user3.user.ChangePassword(user3.newPassword); e != nil {
		t.Errorf("couldn't change password for user %s! %s",
				user3.name, e.Error())
	}

	if e = credentials.Modify(user3.user); e != nil {
		t.Errorf("couldn't update password for user %s! %s",
				user3.name, e.Error())
	}

	if credentials.IsAuthenticated(user3.name, user3.password) {
		t.Errorf("was able to authenticate with old password for user %s!",
				user3.name)
	}

	if !credentials.IsAuthenticated(user3.name, user3.newPassword) {
		t.Errorf("couldn't authenticate with new password for user %s!",
				user2.name)
	}

	// test if rest of file was written
	if !credentials.IsAuthenticated(user2.name, user2.password) {
		t.Errorf("couldn't authenticate user %v after modify!", user3.name)
	}

	// refresh user 2
	var found bool
	if found, user2.user = credentials.Get(user2.name); !found {
		t.Errorf("couldn't get user2")
	}

	if !credentials.IsAuthenticated(user2.name, user2.password) {
		t.Errorf("couldn't authenticate user %v after refresh!", user3.name)
	}
}

func testAddUsers(t *testing.T, credentials CredentialsInterface, users ...*Users) {
	_, indexTesting := credentials.(*Unix)

	for key, user := range users {
		if e := credentials.Add(user.user); e != nil {
			t.Errorf("couldn't add user %v! %s", user.name, e.Error())
		}

		// add again, must return error
		if e := credentials.Add(user.user); e == nil {
			t.Errorf("added the same user %v twice without error!", key)
		}

		if !credentials.IsAuthenticated(user.name, user.password) {
			t.Errorf("couldn't authenticate user %v!", user.name)
		}

		if indexTesting {
			if key != 0 && user.user.getIndex() == 0 {
				t.Errorf("index 0 on user %s!", user.name)
			}
		}
	}
}

func testNext(t *testing.T, credentials CredentialsInterface, users ...*Users) {
	credentials.Reset()

	_, orderTesting := credentials.(*Unix)

	i := 0
	for _, user := range users {
		if next, e := credentials.Next(); e != nil {
			if e == io.EOF {
				t.Errorf("Next() unexpected EOF!")
			} else {
				t.Errorf("Next() %s", e.Error())
			}
		} else if orderTesting &&  user.name != next.GetName() {
			t.Errorf("Next() %s, expected %s", next.GetName(), user.name)
		}

		i++
	}

	if len(users) != i {
		t.Errorf("Next() returned %v users, expected %v users", i, len(users))
	}

	credentials.Reset()
}

type Users struct {
	name, password, newPassword	string
	user				UserInterface
}
