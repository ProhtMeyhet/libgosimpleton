package main

import(
	"fmt"
	"os"
	credentials "github.com/ProhtMeyhet/libgosimpleton/credentials"
)

func doExists(database credentials.CredentialsInterface) {
	if flags.User == "" {
		fmt.Fprintf(os.Stderr, "user cannot be empty!\n")
		os.Exit(EXIT_USER_EMPTY)
	}

	if flags.Verbose {
		if flags.File != "" {
			fmt.Printf("using file %s\n", flags.File)
		}
	}

	if flags.User == "" {
		fmt.Fprintf(os.Stderr, "user empty!")
		os.Exit(EXIT_USER_EMPTY)
	}

	if found, _ := database.Get(flags.User); !found {
		fmt.Fprintf(os.Stderr, "user '%s' not found!\n", flags.User)
		os.Exit(EXIT_USER_UNKNOWN)
	} else if flags.Verbose {
		fmt.Printf("user %s exists!\n", flags.User)
	}
}
