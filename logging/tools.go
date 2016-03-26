package logging

import(
	"strings"
)

func LevelToInt(level string) uint8 {
	switch strings.ToUpper(level) {
	case "EMERGENCY":
		return EMERGENCY
	case "CRITICAL":
		return CRITICAL
	case "ALERT":
		return ALERT
	case "ERROR":
		return ERROR
	case "WARNING":
		return WARNING
	case "NOTICE":
		return NOTICE
	case "INFO":
		return INFO
	case "DEBUG":
		return DEBUG
	case "FATAL":
		return FATAL
	case "WARNINGS":
		return WARNINGS
	case "VERBOSE":
		return VERBOSE
	case "ATTENTION":
		return ATTENTION
	case "DEBUGGING":
		return DEBUGGING
	case "EVERYTHING":
		return EVERYTHING
	}
	return 0
}

func LevelToString(level uint8) string {
	switch level {
	case EMERGENCY:
		return "EMERGENCY"
	case CRITICAL:
		return "CRITICAL"
	//case ALERT:
	//	return "ALERT"
	case ERROR:
		return "ERROR"
	case WARNING:
		return "WARNING"
	case NOTICE:
		return "NOTICE"
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	case FATAL:
		return "FATAL"
	case WARNINGS:
		return "WARNINGS"
	case VERBOSE:
		return "VERBOSE"
	case ATTENTION:
		return "ATTENTION"
	case DEBUGGING:
		return "DEBUGGING"
	case EVERYTHING:
		return "EVERYTHING"
	}
	return "UnknownLogLevel"
}

func LimitString(input string) string {
	// limit to 55 chars
	if len(input) > 50 {
		return input[:50] + "....."
	}

	return input
}

// generic string sanitizing
// specific sanitizing must still be done in
// specific loggers
func SanitizeString(input string) string {
	// remove line breaks
	input = strings.Replace(input, "\n", "", -1)

	// remove tab
	input = strings.Replace(input, "\t", "", -1)

	// remove repeated spaces
	input = strings.Replace(input, "  ", " ", -1)

	return input
}
