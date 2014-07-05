package logging

import(

)

// the actual logger
var logger logInterface

// log a message
// it's your responsibility to call Start() first
func Log(level uint8, message string) {
	logger.Log(level, message)
}

// early logging does not require config
// it just stores messages in memory and can later
// flush them to another logger
func StartEarly() {
	logger = &bufferLogger{}
	logger.SetLevel(EVERYTHING)
}

// start logging
// give it a config
// flushes early log
func Start(config LogConfigInterface) (e error) {
	earlyLogging := logger

	logger, e = getLogger(config.GetType(), config)
	if e != nil {
		return
	}

	if earlyLogger, ok := earlyLogging.(earlyInterface); ok {
		earlyLogger.Flush(logger)
	}

	return logger.Open()
}

// either start logging or panic
func ForceStart(config LogConfigInterface) {
	e := Start(config)
	if e != nil {
		panic("Could not open Log! " + e.Error())
	}
}

// TODO
// add another logger
/*
func AddLogger(logType string) uint8 {
	if union, ok := logger.(unionInterface); !ok {
		union.Add(
	} else {
		union := NewUnionLogger()
		union.Add(logger)
		logger = union
	}
}*/

func Close() (e error) {
	return logger.Close()
}
