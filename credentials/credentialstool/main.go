package main

import(
	"fmt"
	sqldb "database/sql"
	"os"
	"os/signal"
	"syscall"
	credentials "github.com/ProhtMeyhet/libgosimpleton/credentials"
	_ "github.com/mattn/go-sqlite3"
)

var flags = newFlagConfig()
var ignoreSignals = make(chan bool, 1)

func main() {
	go catchSignals()

	flags.parse()

	var database credentials.CredentialsInterface
	switch(flags.Type) {
	case TYPE_UNIX:
		if flags.Verbose {
			fmt.Println("type is unix")
		}

		if flags.File == "" {
			flags.usage()
			return
		}

		database = credentials.NewUnix(flags.File)
	case TYPE_SQLITE:
		if flags.Verbose {
			fmt.Printf("sqlite, file %s, table %s\n", flags.File, flags.SqlTable)
		}

		if flags.File == "" {
			fmt.Fprintf(os.Stderr, "SQL must given as table,database\n")
			return
		}
		config := &credentials.SqlConfig{
				Base: flags.File,
				Table: flags.SqlTable,
			    }

		var e error
		if config.Database, e = sqldb.Open("sqlite3", config.Base); e != nil {
			fmt.Fprintf(os.Stderr, "error opening sqlite database: %s\n", e.Error())
			os.Exit(-1)
		}

		database = credentials.NewSql(config)
	default:
		flags.usage()
		os.Exit(EXIT_UNKNOWN_TYPE)
	}

	DoModeSelect(database)
	database.Close()
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
	signals := make(chan os.Signal, 1)
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
