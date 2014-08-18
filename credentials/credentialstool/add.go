package main

import(
	"fmt"
	"os"
	credentials "github.com/ProhtMeyhet/libgosimpleton/credentials"
)

func DoAdd(database credentials.CredentialsInterface) {
	if flags.User == "" {
		fmt.Fprintf(os.Stderr, "user cannot be empty!\n")
		os.Exit(EXIT_USER_EMPTY)
	}

	message := "password for user " + flags.User + ": "
	retype := "Retype "
	password := getPassword(message, retype)

	newUser, e := database.New(flags.User, password)
	if e != nil {
		fmt.Fprintf(os.Stderr, "error creating user: %s\n", e.Error())
		os.Exit(EXIT_USER_CREATE)
	}

	ignoreSignals <- true
	e = database.Add(newUser)
	if e != nil {
		fmt.Fprintf(os.Stderr, "error adding user: %s\n", e.Error())
		if e == credentials.UserExistsError {
			os.Exit(EXIT_USER_EXISTS)
		} else {
			os.Exit(EXIT_ADD)
		}
	} else {
		fmt.Printf("user %s added!\n", newUser.GetName())
	}
	ignoreSignals <- false
}
