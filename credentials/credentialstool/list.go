package main

import(
	"fmt"
	"os"
	"strconv"
	credentials "github.com/ProhtMeyhet/libgosimpleton/credentials"
	clitable "github.com/crackcomm/go-clitable"
)

func DoList(database credentials.CredentialsInterface) {
	database.Reset()

	table := clitable.New([]string{ "ID", "Name", "Salt", "HashType", "Hash" })
	for i := 0;; i++{
		row := make(map[string]interface{})

		user, e := database.Next()
		if e != nil {
			break
		}

		row["ID"] = strconv.Itoa(i)
		row["Name"] = user.GetName()
		row["Salt"] = user.GetPassworder().GetSalt()
		row["HashType"] = user.GetPassworder().GetHashType()
		row["Hash"] = user.GetPassworder().GetPasswordHash()
		table.AddRow(row)
	}

	database.Reset()

	if len(table.Rows) > 0 {
		table.Print()
	} else {
		fmt.Fprintf(os.Stderr, "empty or non-existing file given!\n")
		os.Exit(EXIT_NOT_EXISTS)
	}
}

