package logging

import(

)

var (
	// the actual logger
	logger logInterface
	// chan for sending logs to
	logging chan *logEntry
)

// package init()
func init() {
	StartEarly()
}

// log a message
func doLog(level uint8, message ...interface{}) {
	LogFormat(level, "%v", message...)
}

// init
func initialise() {
	logging	= make(chan *logEntry, 10)
}

// DEPRECATED; rename to logFormat()
// log a formatted message
func LogFormat(level uint8, format string, message ...interface{}) {
	logging <-&logEntry{ level: level, message: message, format: format }
}
