package logging

import(
	"log/syslog"
)

type syslogLogger struct {
	abstractLogger
	logger *syslog.Writer
}

func NewSyslogLogger(config LogConfigInterface) *syslogLogger {
	sys := &syslogLogger{}
	inject(sys, config)
	return sys
}

func (sys *syslogLogger) Open() (e error) {
	sys.logger, e = syslog.New(DEBUG_SYSLOG, sys.name)
	return
}

func (sys *syslogLogger) Log(level uint8, message string) {
	if sys.ShouldLog(level) {
		switch(level) {
			case EMERGENCY:
				sys.lastE = sys.logger.Emerg(message)
			case CRITICAL:
				sys.lastE = sys.logger.Crit(message)
			//case ALERT:
			//	sys.lastE = syslog.logger.Alert(message)
			case ERROR:
				sys.lastE = sys.logger.Err(message)
			case WARNING:
				sys.lastE = sys.logger.Warning(message)
			case NOTICE:
				sys.lastE = sys.logger.Notice(message)
			case INFO:
				sys.lastE = sys.logger.Info(message)
			case DEBUG:
				sys.lastE = sys.logger.Debug(message)
		}
	}
}

func (sys *syslogLogger) Close() (e error) {
	return sys.logger.Close()
}
