package logging

import(

)

// configuration
type LogConfigInterface interface {
	// name of you program
	GetName() string
	// main log type (null, stderr, file, syslog ...)
	GetType() Type
	// additional loggers
	//GetSecondaryTypes() []string
	// verbosity, see tools.go:LevelToString
	GetLevel() uint8
	SetLevel(to uint8)
}

// union logger can log to more then one logger
type unionInterface interface {
	AddLogger(log logInterface) uint8
	RemoveLogger(id uint8)
}

// used for early logging
type earlyInterface interface {
	// flush buffered Logs to
	Flush(to logInterface)
}

// interface for loggers, see abstractLogger & fileLogger
type logInterface interface {
	Open() error
	Close() error

	// log with this level, watch if level satisfies
	// current logLevel (given via SetLevel())
	Log(level uint8, message string)

	// return if there has been an error and return last one
	HasError() (bool, error)

	// should message be logged for level
	ShouldLog(level uint8) bool

	// log only messages with this or lower level
	// lower is more important
	SetLevel(level uint8)
	GetLevel() uint8

	SetName(name string)
	GetName() string

	run()
}
