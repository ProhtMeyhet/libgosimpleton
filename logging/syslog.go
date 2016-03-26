package logging

import(
	"log/syslog"
)

type syslogLogger struct {
	abstractLogger
	logger *syslog.Writer
}

func NewSyslogLogger(config LogConfigInterface) (sys *syslogLogger) {
	sys = &syslogLogger{}
	sys.initialise(config)
	return
}

func (sys *syslogLogger) Open() (e error) {
	sys.logger, e = syslog.New(DEBUG_SYSLOG, sys.name)
	return
}

func (sys *syslogLogger) Log(level uint8, message string) {
	var e error
	if sys.ShouldLog(level) {
		switch(level) {
			case EMERGENCY:
				e = sys.logger.Emerg(message)
			case CRITICAL:
				e = sys.logger.Crit(message)
			// not used
			//case ALERT:
			//	e = syslog.logger.Alert(message)
			case ERROR:
				e = sys.logger.Err(message)
			case WARNING:
				e = sys.logger.Warning(message)
			case NOTICE:
				e = sys.logger.Notice(message)
			case INFO:
				e = sys.logger.Info(message)
			case DEBUG:
				e = sys.logger.Debug(message)
		}
	}

	if e != nil {
		sys.interfaceConfig.HandleE(sys.name, e)
	}
}

func (sys *syslogLogger) Close() (e error) {
	return sys.logger.Close()
}
