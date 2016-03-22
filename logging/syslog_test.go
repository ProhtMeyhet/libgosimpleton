package logging

import(
	"testing"
)

func TestSyslogLogger(t *testing.T) {
	config := NewDefaultConfig()
	config.LogType = SYS

	genericLoggerTest(t, config, "syslogger")
}
