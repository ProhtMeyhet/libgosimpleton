package logging

import(
	"testing"
	"bufio"
	"bytes"
	"os"
	"time"
)

func genericLoggerTest(t *testing.T, config LogConfigInterface, name string) {
	if e := Start(config); e != nil {
		t.Fatalf("%s: Open() failed! %s", name, e.Error())
	}

	if !logger.ShouldLog(ERROR) {
		t.Fatalf("log.Shouldlog failed for %s", LevelToString(ERROR))
	}

	Error("test")

	// wait for startup of goroutine
	time.Sleep(time.Second * 1)

	if e := Close(); e != nil {
		t.Errorf("%s: Close() failed! %s", name, e.Error())
	}
}

func TestFileLogger(t *testing.T) {
	config := NewDefaultFileConfig()
	config.Path = os.TempDir() + "/test" + time.Now().Format(DATE_FORMAT) +  "30385946.log"

	// remove old testfile if its there
	os.Remove(config.Path)
	defer os.Remove(config.Path)

	//file := NewFileLogger(config)

	// log something
	genericLoggerTest(t, config, "fileLogger")

	// now test what was logged
	expected := []byte("[ERROR]: test")
	if file, ok := logger.(*fileLogger); ok {
		testExpectedInFile(t, file, expected)
	} else {
		t.Fatal("logger isn't a *fileLogger!")
	}
}

func testExpectedInFile(t *testing.T, file *fileLogger, expected []byte) {
	var out []byte
	var needed []byte

	handle, e := os.Open(file.config.Path); if e != nil { t.Fatalf("couldn't open file: %s!", file.config.Path) }
	reader := bufio.NewReader(handle)
	// skip first line
	reader.ReadLine()
	out, _, _ = reader.ReadLine()

	lenPrefix := len(file.name + ": " + file.NowClock())

	if len(out) >= lenPrefix+14 {
		needed = out[lenPrefix+1:]
	} else {
		t.Errorf("read: '%s' - len %v, lennow: %v - , could't parse it",
			out, len(out), lenPrefix+15)
		t.Log("left logfile for analysis: " + file.config.Path)
		return
	}

	if !bytes.Equal(expected, needed) {
		t.Errorf("'%s' != '%s'", expected, needed)
		t.Log("left logfile for analysis: " + file.config.Path)
		return
	}
}
