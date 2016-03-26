package logging

import(
	"testing"

	"errors"
	"io/ioutil"
	"strings"
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

func TestNullLogger(t *testing.T) {
	config := NewDefaultConfig()
	config.LogType = NULL
	genericLoggerTest(t, config, "null")
}

func TestStderrLogger(t *testing.T) {
	config := NewDefaultConfig()
	config.LogType = STDERR
	genericLoggerTest(t, config, "stderr")
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

func genericLoggerTest(t *testing.T, config LogConfigInterface, name string) {
	oldLevel := config.GetLevel(); defer config.SetLevel(oldLevel)
	config.SetLevel(EVERYTHING)

	if e := Start(config); e != nil {
		t.Fatalf("%s: Open() failed! %s", name, e.Error())
	}

	if !logger.ShouldLog(ERROR) {
		t.Fatalf("log.Shouldlog failed for %s", LevelToString(ERROR))
	}

	// change stderr to write to /dev/null
	if actualLogger, ok := logger.(*stderrLogger); ok {
		actualLogger.stderr = ioutil.Discard
	}

	// TODO leave "test" for now until file_test is rewritten
	Error("test")
	Emergency("A British tar is a soaring soul,")
	Critical("As free as a mountain bird,")
	Warning("His energetic fist should be ready to resist")
	Notice("A dictatorial word.")
	Info("His nose should pant")
	Debug("and his lip should curl,")
	EmergencyFormat("%v", "His cheeks should flame")
	CriticalFormat("%v", "and his brow should furl,")
	ErrorFormat("%v", "His bosom should heave")
	WarningFormat("%v", "and his heart should glow,")
	NoticeFormat("%v", "And his fist be ever ready")
	InfoFormat("%v", "for a knock-down blow.")
	DebugFormat("%v", "Gilbert & Sullivan")
	// DEPRECATED TODO
	Log(DEBUG, "Exit Ralph")

	// wait for startup of goroutine
	time.Sleep(time.Second * 1)

	if e := Close(); e != nil {
		t.Errorf("%s: Close() failed! %s", name, e.Error())
	}

	t.Logf("%v of genericLogTest finished!", name)
}

func TestLimitString(t *testing.T) {
	s := strings.Repeat("Corcoran", 15)
	ss := LimitString(s)

	if len(ss) != 55 {
		t.Errorf("expected ss to be 55 chars long, got %v\n", len(ss))
	}

	if ss[50:55] != "....." {
		t.Errorf("ss ends with %v, wanted .....", ss[50:55])
	}

	s = "Cousin Hebe"
	ss = LimitString(s)
	if s != ss {
		t.Errorf("expected ss to be unchanged from s '%v', but got '%v'\n", ss)
	}
}

func TestSanitizeString(t *testing.T) {
	s := "abc\n d  d\t"
	ss := SanitizeString(s)
	expected := "abc d d"
	if ss != expected {
		t.Errorf("expected SanitizedString %v, got %v\n", expected, ss)
	}
}




const(
	WRONG Type = 222
)

type wrongLogger struct {
	nullLogger
}

func (wrong *wrongLogger) Open() (e error) {
	return errors.New("Now give three cheers!")
}
