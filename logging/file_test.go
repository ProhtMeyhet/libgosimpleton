package logging

import(
	"testing"
	"bufio"
	"bytes"
	"os"
	"time"
)

func TestFileLogger(t *testing.T) {
	config := NewDefaultFileConfig()
	config.Path = os.TempDir() + "/test" + time.Now().Format(DATE_FORMAT) +  "30385946.log"
	if config.GetPath() != config.Path {
		t.Errorf("expected %v from GetPath(), got %v", config.Path, config.GetPath())
	}

	// remove old testfile if its there
	os.Remove(config.Path);	defer os.Remove(config.Path)


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

func TestTrollFile(t *testing.T) {
	config := NewDefaultFileConfig()
	config.Path = os.TempDir() + "/test" + time.Now().Format(DATE_FORMAT) +  "30385946.log"
	file := NewFileLogger(config); defer file.Close()
	if e := file.Close(); e != nil {
		t.Errorf("unexpected error on Close(): %v", e.Error())
	}

	if e := file.Open(); e != nil {
		t.Errorf("unexpected error on Open(): %v", e.Error())
	}

	if _, e := file.openLogFile(config.Path); e != nil {
		t.Errorf("unexpected error on openLogFile(): %v", e.Error())
	}

	file = NewFileLogger(NewDefaultConfig()); defer file.Close()
	if e := file.Open(); e == nil {
		t.Errorf("fileLogger accepts DefaultConfig!\n")
	}

	wrongConfig := NewDefaultFileConfig()
	wrongConfig.Path = "/mrs/crips/little/buttercup.log" + time.Now().Format(DATE_FORMAT)
	file = NewFileLogger(wrongConfig); defer file.Close()
	if e := file.Open(); e == nil {
		t.Errorf("fileLogger opens non existing file!\n")
	}
}
