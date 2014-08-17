package main

import(
	"flag"
)

type flagConfig struct {
	File, Type		string
	User			string
	Login, Password		string
	Mode			string
	Timeout			int
	List			bool
	Modify, Add, Remove, Authenticate	string
}

func newFlagConfig() *flagConfig {
	return &flagConfig{}
}

func (flags *flagConfig) parse() {
	flag.StringVar(&flags.File, "file", "shadow", "filename")
	flag.StringVar(&flags.Type, "type", TYPE_UNIX, "type of credentials-db")
	flag.StringVar(&flags.User, "user", "", "name of user")
	flag.StringVar(&flags.Password, "pass", "", "password")
	flag.IntVar(&flags.Timeout, "timeout", 120, "timeout in seconds till abort (locks are released)")

	flag.Parse()

	if flag.NArg() > 1 {
		switch(flag.Arg(0)) {
		case MODE_MODIFY:
			flags.Mode = MODE_MODIFY
		case MODE_ADD:
			flags.Mode = MODE_ADD
		case MODE_REMOVE:
			flags.Mode = MODE_REMOVE
		default:
			flags.Mode = MODE_LIST
		}

		if flag.NArg() > 2 {
			flags.File = flag.Arg(1)
			flags.User = flag.Arg(2)
		} else {
			flags.User = flag.Arg(1)
		}
	} else if flag.NArg() == 1{
		flags.Mode = MODE_MODIFY
		flags.User = flag.Arg(0)
	} else {
		flags.Mode = MODE_LIST
	}
}

func (flags *flagConfig) usage() {
	flag.Usage()
}
