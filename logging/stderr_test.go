package logging

import(
	"testing"
)

func TestStderrLogger(t *testing.T) {
	config := NewDefaultConfig()
	config.LogType = STDERR
	genericLoggerTest(t, config, "stderr")
}
