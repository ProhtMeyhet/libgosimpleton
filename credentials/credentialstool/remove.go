package main

import(
	"fmt"
	"os"
	credentials "github.com/ProhtMeyhet/libgosimpleton/credentials"
)

func DoRemove(database credentials.CredentialsInterface) {
	if flags.User == "" {
		fmt.Fprintf(os.Stderr, "user cannot be empty!\n")
		os.Exit(EXIT_USER_EMPTY)
	}

	found, user := database.Get(flags.User)
	if !found {
		fmt.Fprintf(os.Stderr, "user doesn't exists!\n")
		os.Exit(EXIT_USER_UNKNOWN)
	}

	ignoreSignals <- true
	e := database.Remove(user)
	if e != nil {
		fmt.Fprintf(os.Stderr, "error removing user: %s\n", e.Error())
		os.Exit(EXIT_REMOVE)
	} else {
		fmt.Printf("removed user %s\n", user.GetName())
	}
	ignoreSignals <- false
}
