package logging

import(
	"fmt"
	"io"
	"os"
)

/*
* logs to stderr
*/

type stderrLogger struct {
	abstractLogger

	stderr io.Writer
}

func NewStderrLogger(config LogConfigInterface) (stderr *stderrLogger) {
	stderr = &stderrLogger{}
	stderr.initialise(config)
	return
}

func (err *stderrLogger) initialise(config LogConfigInterface) {
	err.stderr = os.Stderr
	err.abstractLogger.initialise(config)
}

func (err *stderrLogger) Log(level uint8, message string) {
	if err.ShouldLog(level) {
		if err.interfaceConfig.IsPlain() {
			err.log(level, message, "%s %s: %s")
		} else {
			err.logColored(level, message)
		}
	}
}

func (err *stderrLogger) log(level uint8, message string, format string) {
	fmt.Fprintf(err.stderr, format,
				err.NowClock(),
				LevelToString(level),
				SanitizeString(message))
}

func (err *stderrLogger) logColored(level uint8, message string) {
	levelHighlight := ""
	switch level {
	case EMERGENCY, CRITICAL, ERROR:
		levelHighlight = "\033[31m" // red
        case WARNING:
		levelHighlight = "\033[35m" // magenta
        case INFO:
		levelHighlight = "\033[32m" // green
	case DEBUG:
		levelHighlight = "\033[33m" // yellow
	}

	err.log(level, message, "%s " + levelHighlight + "%s: %s\033[0m\n")
}

func (err *stderrLogger) Open() (e error) { return }
func (err *stderrLogger) Close() (e error) { return }
