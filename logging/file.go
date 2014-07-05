package logging

import(
	"errors"
	"os"
	"log"
)

// take a wild guess, what this logger logs to

type fileLogger struct {
	abstractLogger
	logger *log.Logger
	path string
	logFile *os.File
}

func NewFileLogger(config LogConfigInterface) *fileLogger {
	file := &fileLogger{}
	inject(file, config)
	return file
}

func (file *fileLogger) Log(level uint8, message string) {
	if file.ShouldLog(level) {
		file.logger.Printf(": %v [%v]: %v", file.NowClock(),
			LevelToString(level), message)
	}
}

func (file *fileLogger) Open() (e error) {
	file.logFile, e = file.openLogFile(file.path)
	if e != nil {
		return
	}

	file.logger = log.New(file.logFile, file.name, 0)

	// test if we can write
	e = file.logger.Output(10,": " + file.NowClock() + " Logging started")
	return
}

func (file *fileLogger) Close() error {
	if file.logFile != nil {
		return file.logFile.Close()
	}
	return nil
}

func (file *fileLogger) SetPath(path string) {
	file.path = path
}

func (file *fileLogger) openLogFile(logFileName string) (logFile *os.File, e error) {
	logFile, e = os.OpenFile(logFileName, os.O_APPEND|os.O_WRONLY, 0777)
	if e != nil {
		if os.IsNotExist(e) {
			return file.createLogFile(logFileName)
		}
	}

	return
}

func (file *fileLogger) createLogFile(logFileName string) (logFile *os.File, e error) {
	logFile, e = os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if e != nil {
		if os.IsExist(e) {
			return file.openLogFile(logFileName)
		} else {
			e = errors.New("Could not create log file " + logFileName)
		}
	}

	return
}
