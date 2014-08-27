package main

import(
	"fmt"
	"flag"
	"os"
)

type flagConfig struct {
	File, Type		string
	User			string
	Mode			string
	Login, Password		string
	SqlTable		string
	Timeout			int
	Verbose			bool
	Modify, Add, Remove, Authenticate	string
}

func newFlagConfig() *flagConfig {
	return &flagConfig{}
}

func (flags *flagConfig) parse() {
	flag.StringVar(&flags.File, "file", "shadow", "filename")
	flag.StringVar(&flags.Type, "type", TYPE_UNIX, "type of credentials-db")
	//flag.StringVar(&flags.User, "user", "", "name of user")
	//flag.StringVar(&flags.Password, "pass", "", "password")
	//flag.IntVar(&flags.Timeout, "timeout", 120, "timeout in seconds till abort (locks are released)")
	flag.BoolVar(&flags.Verbose, "v", false, "be verbose")

	flag.Parse()

	if flags.Verbose {
		fmt.Fprintf(os.Stderr, "number of args: %v\n", flag.NArg())
	}

	if flag.NArg() == 0 {
		flags.Mode = MODE_LIST
	} else if flag.NArg() > 1 {
		switch(flag.Arg(0)) {
		case MODE_MODIFY:
			flags.Mode = MODE_MODIFY
		case MODE_ADD:
			flags.Mode = MODE_ADD
		case MODE_REMOVE:
			flags.Mode = MODE_REMOVE
		case MODE_EXISTS:
			flags.Mode = MODE_EXISTS
		case MODE_LIST:
			flags.Mode = MODE_LIST
		default:
			flags.Mode = MODE_AUTHENTICATE
		}

		if flag.NArg() > 1 {
			if flags.Type == TYPE_SQL || flags.Type == TYPE_SQLITE {
				flags.SqlTable = flag.Arg(1)
				flags.File = flag.Arg(2)
				flags.User = flag.Arg(3)
			} else {
				flags.File = flag.Arg(1)
				flags.User = flag.Arg(2)
			}
		} else {
			flags.User = flag.Arg(1)
		}
	} else if flag.NArg() == 1 {
		flags.Mode = MODE_AUTHENTICATE
		flags.User = flag.Arg(0)
	} else {
		flags.Mode = MODE_LIST
	}
}

func (flags *flagConfig) usage() {
	flag.Usage()
}
