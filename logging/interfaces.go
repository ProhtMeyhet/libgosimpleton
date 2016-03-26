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

	// error handling
	GetEHandler() func(string, error)
	SetEHandler(to func(string, error))
	HandleE(string, error)

	// plain mode (no colors, no highlights etc)
	IsPlain() bool
	// set to plain mode (default: false)
	TogglePlain()
}

// union logger can log to more then one logger
type unionInterface interface {
	AddLogger(log logInterface) uint8
	RemoveLogger(id uint8)
}

// used for early logging
type earlyInterface interface {
	// flush buffered Logs to
	Flush(to logInterface) error
	Close() error
}

// interface for loggers, see abstractLogger & fileLogger
type logInterface interface {
	Open() error
	Close() error

	// log with this level, watch if level satisfies
	// current logLevel (given via SetLevel())
	Log(level uint8, message string)

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
