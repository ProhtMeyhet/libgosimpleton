package logging

import(
	"testing"
)

func TestNullLogger(t *testing.T) {
	config := NewDefaultConfig()
	config.LogType = NULL
	genericLoggerTest(t, config, "null")
}
