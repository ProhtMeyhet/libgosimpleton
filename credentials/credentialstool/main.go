package main

import(
	"fmt"
	"os"
	"os/signal"
	"syscall"
	credentials "github.com/ProhtMeyhet/libgosimpleton/credentials"
)

var flags = newFlagConfig()
var ignoreSignals = make(chan bool, 1)

func main() {
	go catchSignals()

	flags.parse()

	switch(flags.Type) {
	case TYPE_UNIX:
		if flags.File == "" {
			flags.usage()
			return
		}

		database := credentials.NewUnix(flags.File)

		DoModeSelect(database)
	//case TYPE_SQL:

	default:
		flags.usage()
		os.Exit(EXIT_UNKNOWN_TYPE)
	}
}

func DoModeSelect(database credentials.CredentialsInterface) {
	switch(flags.Mode) {
	case MODE_AUTHENTICATE:
		DoAuthenticate(database)
	case MODE_LIST:
		DoList(database)
	case MODE_MODIFY:
		DoModify(database)
	case MODE_ADD:
		DoAdd(database)
	case MODE_REMOVE:
		DoRemove(database)
	case MODE_EXISTS:
		DoExists(database)
	default:
		flags.usage()
		os.Exit(EXIT_UNKNOWN_MODE)
	}
}

func catchSignals() {
	signals := make(chan os.Signal, 10)
	signal.Notify(signals, syscall.SIGHUP, syscall.SIGINT, syscall.SIGKILL,
			syscall.SIGQUIT, syscall.SIGTERM)

infinite:
	for {
		select {
		case <-signals:
			// exit pretty
			fmt.Println("")
			os.Exit(-1)
			break infinite
		case <-ignoreSignals:
			// block until released
			<-ignoreSignals
		}
	}
}
