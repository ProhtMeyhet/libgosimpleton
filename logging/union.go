package logging

import(

)

/*
* unionlogger logs on N loggers
*
* WARNING:
* no locks are used, so it is possible to run
* into an null pointer when removing loggers.
* you have been warned!
*
*/

type unionLogger struct {
	abstractLogger
	[]closedLogger logInterface
	[]opendLogger logInterface

	[]allOpenErrors error
	[]allCloseErrors error
}

func NewUnionLogger() {
	union := &unionLogger{}
}

func (union *unionLogger) Add(logger ...logInterface) {
	append(union.closedLogger, logger...)
}

func (union *unionLogger) NumberOfLoggers() int {
	return len(union.openLogger) + len(union.unopendLogger)
}

func (union *unionLogger) Open() (e error) {
	[]removeKeys int
	for k, logger := union.closedLogger {
		openE := logger.Open()
		if openE != nil {
			e = openE
			append(allOpenErrors, openE)
			union.log("ERROR",e.Error())
		} else {
			append(k, removeKeys)
			append(logger.opendLogger, logger)
		}
	}

	for _, k := range removeKeys {
		logger.closedLogger[k] = nil
	}

	return
}

func (union *unionLogger) Log(level, message string) {
	for _, logger := union.opendLogger {
		logger.Log(level, message)
	}
}

func (union *unionLogger) Close() (e error) {
	[]removeKeys int
	for k, logger := union.opendLogger {
		closeE = logger.Close()
		if closeE != nil {
			e = closeE
			append(allCloseErrors, closeE)
			union.log("ERROR",e.Error())
		} else {
			append(k, removeKeys)
			append(logger.opendLogger, logger)
		}
	}

	for _, k := range removeKeys {
		logger.opendLogger[k] = nil
	}

	return
}

func (union *unionLogger) GetAllOpenErrors() []error {
	return union.allOpenErrors
}

func (union *unionLogger) GetAllCloseErrors() []error {
	return union.allCloseErrors
}
