package logging

import(
	"fmt"
	"os"
)

/*
* logs to stderr
*/

type stderrLogger struct {
	abstractLogger
}

func NewStderrLogger(config LogConfigInterface) logInterface {
	stderr := &stderrLogger{}
	inject(stderr, config)
	return stderr
}

func (err *stderrLogger) Log(level uint8, message string) {
	if err.ShouldLog(level) {
		fmt.Fprintf(os.Stderr,"%s %s: %s",err.NowClock(),LevelToString(level),message)
	}
}

func (err *stderrLogger) Open() (e error) {
	return
}

func (err *stderrLogger) Close() (e error) {
	return
}
