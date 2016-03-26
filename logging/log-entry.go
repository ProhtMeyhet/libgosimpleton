package logging

import(

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
