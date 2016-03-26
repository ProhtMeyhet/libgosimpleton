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
	errorMessage2 := "He thought so little, they rewarded he\n" +
			  "By making him the Ruler of the Queen's Navee!"

	config := NewDefaultConfig()
	config.LogType = BUFFER
	config.Level = EVERYTHING
	config.Name = "Sir Joseph Porter"
	logger.SetName(config.Name)
	if config.GetName() != config.Name {
		t.Errorf("expected name = %v, got %v", config.Name, config.GetName())
	}
	if logger.GetName() != config.Name {
		t.Errorf("expected name = %v, got %v", config.Name, logger.GetName())
	}

	logger.SetLevel(DEBUG)
	if logger.GetLevel() != DEBUG {
		t.Errorf("expected level = %v, got %v", DEBUG, logger.GetLevel())
	}
	logger.SetLevel(EVERYTHING)

	// twice; this is intenional!
	StartEarly()
	StartEarly()
	if _, ok := logger.(earlyInterface); !ok {
		t.Fatal("StartEarly(): logger is not an earlyInterface!")
	}

	Error(errorMessage)

	// wait for startup of goroutine
	time.Sleep(time.Second * 2)

	if actualLogger, ok := logger.(*bufferLogger); ok {
		if len(actualLogger.levels) != 1 {
			t.Fatalf("expected actualLogger.levels len() to be 1, got %v", len(actualLogger.levels))
		}
	}

	if e := Start(config); e != nil {
		t.Fatalf("unable to open logger! %v", e.Error())
	}

	if !logger.ShouldLog(DEBUG) {
		t.Fatalf("log.Shouldlog failed for %s", LevelToString(DEBUG))
	}

	if actualLogger, ok := logger.(*bufferLogger); ok {
		if len(actualLogger.levels) != 1 {
			t.Fatalf("%p expected actualLogger.levels len() after Start() to be 1, got %v", actualLogger, len(actualLogger.levels))
		}
	}

	Debug(errorMessage2)

	// wait for startup of goroutine
	time.Sleep(time.Second * 2)

	if actualLogger, ok := logger.(*bufferLogger); ok {
		if len(actualLogger.levels) != 2 {
			if len(actualLogger.levels) > 0 {
				t.Log(LevelToString(actualLogger.levels[0]))
			}
			t.Fatalf("expected actualLogger.levels len() to be 2, got %v", len(actualLogger.levels))
		}

		if actualLogger.levels[0] != ERROR {
			t.Errorf("expected levels[0] to be ERROR! got %v", LevelToString(actualLogger.levels[0]))
		}

		if actualLogger.messages[0] != errorMessage {
			t.Errorf("expected messages[0] to be \n'%v'\n\ngot\n\n'%v'", errorMessage, actualLogger.messages[0])
		}

		if actualLogger.levels[1] != DEBUG {
			t.Errorf("expected levels[0] to be DEBUG! got %v", LevelToString(actualLogger.levels[1]))
		}

		if actualLogger.messages[1] != errorMessage2 {
			t.Errorf("expected messages[0] to be \n'%v'\n\ngot\n\n'%v'", errorMessage, actualLogger.messages[1])
		}
	} else {
		t.Fatalf("unable to cast logger to bufferLogger!")
	}

	if e := Close(); e != nil {
		t.Errorf("Close() failed: %v", e.Error())
	}
}

func TestStart(t *testing.T) {
	defer func() {
		recover()
	}()

	config := NewDefaultConfig()
	config.LogType = WRONG

	e := Start(config); if e == nil {
		t.Errorf("Start() didn't return an error!\n")
	}

	ForceStart(config)

	// shouldn't be reached, ForceStart should panic()
	t.Errorf("ForceStart didn't panic!")
}

func TestPanic(t *testing.T) {
	defer func() {
		recover()
	}()

	config := NewDefaultConfig()
	config.LogType = NULL

	Panic("Kind Captain, I've important information")

	// shouldn't be reached, ForceStart should panic()
	t.Errorf("Panic didn't panic!")
}
