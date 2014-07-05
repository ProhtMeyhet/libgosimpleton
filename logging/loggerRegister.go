package logging

import(
	"errors"
)

func getLogger(from string, config LogConfigInterface) (newLogger logInterface,
								e error) {
	switch(from) {
	case "syslog":
		newLogger = NewSyslogLogger(config)
	// case "union"
	//	logger = NewUnionLogger()
	case "buffer":
		newLogger = NewBufferLogger(config)
	case "stderr":
		newLogger = NewStderrLogger(config)
	case "null":
		newLogger = NewNullLogger(config)
	case "file":
		newLogger = NewFileLogger(config)
	default:
		return newLogger, errors.New("Unknown logger: " + from)
	}

	return newLogger, nil
}
