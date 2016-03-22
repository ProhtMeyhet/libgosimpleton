package logging

import(
	"testing"

	"time"
)

func TestLogApi(t *testing.T) {
	errorMessage := "I always voted at my party's call,\n" +
			"And I never thought of thinking for myself at all.\n" +
			"I thought so little, they rewarded me\n" +
			"By making me the Ruler of the Queen's Navee!"

	config := NewDefaultConfig()
	config.LogType = BUFFER
	config.Level = EVERYTHING

	if e := Start(config); e != nil {
		t.Fatalf("unable to open logger! %v", e.Error())
	}

	if !logger.ShouldLog(DEBUG) {
		t.Fatalf("log.Shouldlog failed for %s", LevelToString(DEBUG))
	}

	Debug(errorMessage)

	// wait for startup of goroutine
	time.Sleep(time.Second * 2)

	if actualLogger, ok := logger.(*bufferLogger); ok {
		if actualLogger.levels[0] != DEBUG {
			t.Errorf("expected levels[0] to be DEBUG! got %v", LevelToString(actualLogger.levels[0]))
		}

		if actualLogger.messages[0] != errorMessage {
			t.Errorf("expected messages[0] to be \n'%v'\n\ngot\n\n'%v'", errorMessage, actualLogger.messages[0])
		}
	} else {
		t.Fatalf("unable to cast logger to bufferLogger!")
	}

	if e := Close(); e != nil {
		t.Errorf("Close() failed: %v", e.Error())
	}
}
