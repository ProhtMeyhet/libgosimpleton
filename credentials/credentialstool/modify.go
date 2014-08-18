package main

import(
	"fmt"
	"os"
	credentials "github.com/ProhtMeyhet/libgosimpleton/credentials"
)

func DoModify(database credentials.CredentialsInterface) {
	if flags.User == "" {
		fmt.Fprintf(os.Stderr, "user cannot be empty!\n")
		os.Exit(EXIT_USER_EMPTY)
	}

	found, user := database.Get(flags.User)
	if !found {
		fmt.Fprintf(os.Stderr, "user not found!\n")
		os.Exit(EXIT_USER_UNKNOWN)
	}

	message := "NEW password for user " + flags.User + ": "
	retype := "Retype "
	password := getPassword(message, retype)

	e := user.ChangePassword(password)
	if e != nil {
		fmt.Fprintf(os.Stderr, "error changing password: %s\n", e.Error())
		os.Exit(EXIT_CHANGE_PASSWORD_FAILED)
	}

	ignoreSignals <- true
	e = database.Modify(user)
	if e != nil {
		fmt.Fprintf(os.Stderr, "error modyfing user: %s\n", e.Error())
		os.Exit(EXIT_MODIFY)
	} else {
		fmt.Printf("modified user %s!\n", user.GetName())
	}
	ignoreSignals <- false
}
