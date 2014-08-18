package main

import(
	"fmt"
	"os"
	credentials "github.com/ProhtMeyhet/libgosimpleton/credentials"
	"code.google.com/p/gopass"
)

func DoAuthenticate(database credentials.CredentialsInterface) {
	if flags.User == "" {
		fmt.Fprintf(os.Stderr, "user cannot be empty!\n")
		os.Exit(EXIT_USER_EMPTY)
	}

	if flags.Verbose {
		if flags.File != "" {
			fmt.Printf("using file %s\n", flags.File)
		}
	}

	password, e := gopass.GetPass("password for user " + flags.User + ": ")
	if e != nil {
		fmt.Fprintf(os.Stderr, "error getting password: %s\n", e.Error())
		os.Exit(EXIT_READ_PASSWORD)
	}

	if authenticator, ok := database.(credentials.AuthenticationInterface); ok {
		if !authenticator.IsAuthenticated(flags.User, password) {
			fmt.Fprintf(os.Stderr, "couldn't authenticate user %s\n",
					flags.User)
			os.Exit(EXIT_WRONG_AUTHENTICATION)
		}
	} else {
		fmt.Fprintf(os.Stderr, "library error: couldn't typecast!\n")
		os.Exit(EXIT_LIBRARY_ERROR)
	}
}
