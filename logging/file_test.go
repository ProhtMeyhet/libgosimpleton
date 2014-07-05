package logging

import(
	"testing"
	"bufio"
	"bytes"
	"os"
	"os/user"
	"time"
)

func genericLoggerTest(t *testing.T, log logInterface, name string) {
	if e := log.Open(); e != nil {
		t.Errorf("%s: Open() failed! %s", name, e.Error())
	}

	if !log.ShouldLog(ERROR) {
		t.Errorf("log.Shouldlog failed for %s", LevelToString(ERROR))
	}

	log.Log(ERROR, "test")

	if e := log.Close(); e != nil {
		t.Errorf("%s: Close() failed! %s", name, e.Error())
	}
}

func TestFileLogger(t *testing.T) {
	// use users home, so we now we have write access
	thisUser, _ := user.Current()

	config := NewDefaultFileConfig()
	config.Path = thisUser.HomeDir + "/test" +
		time.Now().Format(DATE_FORMAT) +  "30385946.log"

	os.Remove(config.Path)

	file := NewFileLogger(config)

	// log something
	genericLoggerTest(t, file, "fileLogger")

	// now test what was logged
	expected := []byte("[ERROR]: test")
	var out []byte
	var needed []byte

	handle, _ := os.Open(config.Path)
	reader := bufio.NewReader(handle)
	// skip first line
	_, _, _ = reader.ReadLine()
	out, _, _ = reader.ReadLine()

	lenPrefix := len(file.name + ": " + file.NowClock())

	if len(out) >= lenPrefix+14 {
		needed = out[lenPrefix+1:]
	} else {
		t.Errorf("read: '%s' - len %v, lennow: %v - , could't parse it",
			out, len(out), lenPrefix+15)
		t.Log("left logfile for analysis: " + config.Path)
		return
	}

	if !bytes.Equal(expected, needed) {
		t.Errorf("'%s' != '%s'", expected, needed)
		t.Log("left logfile for analysis: " + config.Path)
		return
	}

	os.Remove(config.Path)
}
