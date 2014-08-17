package libgocredentials

import(
	"testing"
	//"fmt"
	//"os/exec"
	//"strings"
)

func TestCgoCrypt(t *testing.T) {
	password := "acdc"
	salt := "$6$accadacca$"
	hash := CCrypt(password, salt)
	expected := "$6$accadacca$S52Nmn0XM83gWlM6j6LD.xnd.RZzqNKn7qqiGy/k5CCe.ZrHtYjnGsjwXfVfvMzHW9M4YZZdbBRnienrX66Te."

	if hash != expected {
		t.Errorf("unexpected hash!")
	}

	//TODO
	/*
	// compare with perls output of crypt
	perlSalt := strings.Replace(salt, "$", "\\$", -1)
	perlArguments := fmt.Sprintf("-e 'print crypt(\"%s\",\"%s\")'", password,
					perlSalt)
	cmd := exec.Command("perl", perlArguments)
	//cmd := exec.Command("tr", "a-z", "A-Z")
	//cmd = exec.Command("sleep", "5")

	out, e := cmd.Output()
	if e != nil {
		t.Errorf("couldn't run command: %s", e.Error())
	}

	t.Errorf("%s\n", hash)
	t.Errorf("%s\n", out)
	t.Errorf("%s\n", perlArguments)
	*/
}
