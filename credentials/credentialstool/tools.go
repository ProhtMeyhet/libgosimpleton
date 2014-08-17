package main

import(
	"fmt"
	"os"
	"code.google.com/p/gopass"
)

func getPassword(message, retype string) (password string) {
	previousPassword := ""
	for {
		password, e := gopass.GetPass(message)
		if e != nil {
			fmt.Fprintf(os.Stderr, "error GetPass(): %s\n", e.Error())
			os.Exit(EXIT_READ_PASSWORD)
		}

		if previousPassword == "" {
			previousPassword = password
			message = retype + message
		} else if previousPassword == password {
			return password
		} else {
			fmt.Println("passwords don't match!\n")
			previousPassword = ""
			message = message[len(retype):]
		}
	}

	return
}
