package logging

import(
	"errors"
)

func getLogger(from Type, config LogConfigInterface) (newLogger logInterface, e error) {
	switch(from) {
	case SYS:
		newLogger = NewSyslogLogger(config)
	// case UNION:
	//	logger = NewUnionLogger()
	case BUFFER:
		newLogger = newBufferLogger(config)
	case STDERR:
		newLogger = NewStderrLogger(config)
	case NULL:
		newLogger = NewNullLogger(config)
	case FILE:
		newLogger = NewFileLogger(config)
	default:
		e = errors.New("Unknown logger!")
	}

	return
}
