package logging

import(
	"testing"
)

func TestSyslogLogger(t *testing.T) {
	config := NewDefaultConfig()
	syslogger, e := getLogger("syslog", config)

	if e != nil {
		t.Errorf("couldn't get syslogLogger! %s", e.Error())
	}

	genericLoggerTest(t, syslogger, "syslogger")

	if hasError, e := syslogger.HasError(); hasError {
		t.Errorf("syslogger has error: %s", e.Error())
	}
}
