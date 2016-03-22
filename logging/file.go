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
	config *DefaultFileConfig
	logFile *os.File
}

func NewFileLogger(config LogConfigInterface) (file *fileLogger) {
	file = &fileLogger{}
	file.initialise(config)
	return
}

func (file *fileLogger) Log(level uint8, message string) {
	if file.ShouldLog(level) {
		file.logger.Printf(": %v [%v]: %v", file.NowClock(), LevelToString(level), message)
	}
}

func (file *fileLogger) Open() (e error) {
	if fileConfig, ok := file.interfaceConfig.(*DefaultFileConfig); !ok {
		e = errors.New("config is not DefaultFileConfig!")
		goto out
	} else {
		file.config = fileConfig
	}

	file.logFile, e = file.openLogFile(file.config.Path); if e != nil { goto out }
	file.logger = log.New(file.logFile, file.name, 0)

	// test if writing is possible
	e = file.logger.Output(10, ": " + file.NowClock() + " Logging started")

out:
	return
}

func (file *fileLogger) Close() error {
	if file.logFile != nil {
		return file.logFile.Close()
	}

	return nil
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
