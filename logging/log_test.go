package logging

import(
	"testing"

	"errors"
	"io/ioutil"
	"time"
)

func genericLoggerTest(t *testing.T, config LogConfigInterface, name string) {
	oldLevel := config.GetLevel(); defer config.SetLevel(oldLevel)
	config.SetLevel(EVERYTHING)
	oldEHandler := config.GetEHandler(); defer config.SetEHandler(oldEHandler)
	config.SetEHandler(func(name string, e error) {
		t.Errorf("error while logging: %v", e.Error())
	})

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

// a logger that just gives back errors for testing
const(
	WRONG Type = 222
)

type wrongLogger struct {
	nullLogger
}

func (wrong *wrongLogger) Open() (e error) {
	return errors.New("Now give three cheers!")
}
