package logging

import(
	"fmt"
)

var (
	// the actual logger
	logger logInterface
	// chan for sending logs to
	logging chan *logEntry
)

// a log entry
type logEntry struct {
	// entry log level
	level uint8
	// entry message
	message []interface{}
	// format inside the goroutine rather then before
	format string
}

// package init()
func init() {
	StartEarly()
}

// log to memory first, as this requires no real configuration.
// Start() will flush this logger to the actual logger
func StartEarly() {
	initialise()
	config := NewDefaultConfig()
	config.Level = EVERYTHING
	logger = newBufferLogger(config)
	go logger.run()
}

// either start logging or panic
func ForceStart(config LogConfigInterface) {
	if e := Start(config); e != nil {
		panic("Could not open Log! " + e.Error())
	}
}

// start logging; flushes early log from StartEarly()
func Start(config LogConfigInterface) (e error) {
	earlyLogging := logger

	logger, e = getLogger(config.GetType(), config); if e != nil { goto out }
	e = logger.Open(); if e != nil { goto out }

	if earlyLogger, ok := earlyLogging.(earlyInterface); ok {
		earlyLogger.Flush(logger)
		Close()
	}

	initialise()
	go logger.run()

out:
	return
}

// the end is at hand
func Close() (e error) {
	close(logging)
	return logger.Close()
}

// log with EMERGENCY and then panic with the message
func Panic(message ...interface{}) {
	stringMessage := fmt.Sprintf("%v", message...)
	Emergency(stringMessage)
	panic(stringMessage)
}

func Emergency(message ...interface{}) {
	doLog(EMERGENCY, message...)
}

func Critical(message ...interface{}) {
	doLog(CRITICAL, message...)
}

func Error(message ...interface{}) {
	doLog(ERROR, message...)
}

func Warning(message ...interface{}) {
	doLog(WARNING, message...)
}

func Notice(message ...interface{}) {
	doLog(NOTICE, message...)
}

func Info(message ...interface{}) {
	doLog(INFO, message...)
}

func Debug(message ...interface{}) {
	doLog(DEBUG, message...)
}

func EmergencyFormat(format string, message ...interface{}) {
	LogFormat(EMERGENCY, format, message...)
}

func CriticalFormat(format string, message ...interface{}) {
	LogFormat(CRITICAL, format, message...)
}

func ErrorFormat(format string, message ...interface{}) {
	LogFormat(ERROR, format, message...)
}

func WarningFormat(format string, message ...interface{}) {
	LogFormat(WARNING, format, message...)
}

func NoticeFormat(format string, message ...interface{}) {
	LogFormat(NOTICE, format, message...)
}

func InfoFormat(format string, message ...interface{}) {
	LogFormat(INFO, format, message...)
}

func DebugFormat(format string, message ...interface{}) {
	LogFormat(DEBUG, format, message...)
}

// DEPRECATED
// log a message; call Start() first.
func Log(level uint8, message ...interface{}) {
	doLog(level, message...)
}

// DEPRECATED; rename to logFormat()
// log a formatted message
func LogFormat(level uint8, format string, message ...interface{}) {
	logging <-&logEntry{ level: level, message: message, format: format }
}

// log a message
func doLog(level uint8, message ...interface{}) {
	LogFormat(level, "%v", message...)
}

// init
func initialise() {
	logging	= make(chan *logEntry, 10)
}
