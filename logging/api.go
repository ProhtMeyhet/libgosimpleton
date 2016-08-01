package logging

import(
	"fmt"
)

// log to memory first, as this requires no real configuration.
// Start() will flush this logger to the actual logger
func StartEarly() {
	if _, ok := logger.(earlyInterface); ok {
		return // already started
	}

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
		earlyLogger.Close()
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
	PanicFormat("%v", message...)
}

// log with EMERGENCY and then panic with the message with format
func PanicFormat(format string, message ...interface{}) {
	stringMessage := fmt.Sprintf(format, message...)
	Emergency(stringMessage)
	panic(stringMessage)
}

// cannot continue, aborting immidiatly
func Emergency(message ...interface{}) {
	doLog(EMERGENCY, message...)
}

// require immidiate attention
func Critical(message ...interface{}) {
	doLog(CRITICAL, message...)
}

// not used
/*
func Alert(message ..interface{}) {
	doLolg(ALERT, message...)
}
*/

// an error that was recovered
func Error(message ...interface{}) {
	doLog(ERROR, message...)
}

// warning: space nearly full etc.
func Warning(message ...interface{}) {
	doLog(WARNING, message...)
}

// notice: behaviour will change, found this driver etc.
func Notice(message ...interface{}) {
	doLog(NOTICE, message...)
}

// info: version info, something succeded etc.
func Info(message ...interface{}) {
	doLog(INFO, message...)
}

// debug. favourite time of programming
func Debug(message ...interface{}) {
	doLog(DEBUG, message...)
}

// cannot continue, aborting immidiatly
func EmergencyFormat(format string, message ...interface{}) {
	LogFormat(EMERGENCY, format, message...)
}

// require immidiate attention
func CriticalFormat(format string, message ...interface{}) {
	LogFormat(CRITICAL, format, message...)
}

// not used
/*
func AlertFormat(format string, message ...interface{}) {
	LogFormat(ALERT, format, message...)
}
*/

// an error that was recovered
func ErrorFormat(format string, message ...interface{}) {
	LogFormat(ERROR, format, message...)
}

// warning: space nearly full etc.
func WarningFormat(format string, message ...interface{}) {
	LogFormat(WARNING, format, message...)
}

// notice: behaviour will change, found this driver etc.
func NoticeFormat(format string, message ...interface{}) {
	LogFormat(NOTICE, format, message...)
}

// info: version info, something succeded etc.
func InfoFormat(format string, message ...interface{}) {
	LogFormat(INFO, format, message...)
}

// debug, favourite time of programming
func DebugFormat(format string, message ...interface{}) {
	LogFormat(DEBUG, format, message...)
}

// DEPRECATED
// log a message; call Start() first.
func Log(level uint8, message ...interface{}) {
	doLog(level, message...)
}
