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
		fmt.Fprintf(err.stderr, "%s %s: %s\n",
					err.NowClock(),
					LevelToString(level),
					SanitizeString(message))
	}
}

func (err *stderrLogger) Open() (e error) { return }
func (err *stderrLogger) Close() (e error) { return }
