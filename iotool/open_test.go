package iotool

import(
	"testing"

	"io/ioutil"
	"path/filepath"
	"time"
)

func TestOpen(t *testing.T) {
	Close(testOpen(t, "open.go"))
}

func TestCreate(t *testing.T) {
	prefix := "A dish fit for the gods"

	// Temporary() uses Create()
	handler, e := Temporary(ReadAndWrite(), prefix)
	if e != nil {
		t.Fatalf("unexpected error creating temporary file: '%v'", e)
	}
	defer Close(handler); defer Remove(handler); t.Logf("created file: '%v'", handler.Name())

	testWriteFile(t, handler)
	testReadFile(t, handler)
}

func TestTemporary(t *testing.T) {
	prefix := "Et tu, Brute"

	// see next debug if
	if DEBUG { random.Seed(1) }

	handler, e := Temporary(ReadAndWrite(), prefix)
	if e != nil {
		t.Fatalf("unexpected error creating temporary file: '%v'", e)
	}
	defer Close(handler); defer Remove(handler); t.Logf("created file: '%v'", handler.Name())

	if DEBUG {
		random.Seed(1)
		handler2, e := Temporary(ReadAndWrite(), prefix)
		if e != nil {
			t.Errorf("unexpected error creating temporary file: '%v'", e)
		}
		defer Close(handler2); defer Remove(handler2); t.Logf("created file: '%v'", handler2.Name())

		if handler.Name() == handler2.Name() {
			t.Errorf("recreated and opened the same file! EXCL not set? \n\t%v\n\t%v", handler.Name(), handler2.Name())
		}
	} else {
		t.Log("DEBUG is off, cannot test for recreating of random named temporary file!")
	}

	testWriteFile(t, handler)
	testReadFile(t, handler)

	random.Seed(1)
	name := temporaryName(prefix); t.Logf("temporary name1: '%v'", name)
	random.Seed(1)
	name2 := temporaryName(prefix); t.Logf("temporary name2: '%v'", name2)
	name3 := temporaryName(prefix); t.Logf("temporary name3: '%v'", name3)

	if name != name2 {
		t.Errorf("expected names to be equal, got '%v' and '%v'", name, name2)
	}

	if name == name3 || name2 == name3 {
		t.Errorf("expected names to be different, got '%v', '%v' and '%v'", name, name2, name3)
	}
}

func TestTroll(t *testing.T) {
	helper := ReadOnly().ToggleFileAdviceReadSequential()
	name := "Beware the ides of March. - (Act I, Scene II)."
	_, e := Open(helper, name); if e == nil {
		t.Fatalf("huh?! could open '%v'", name)
	}

	name = "./"
	_, e = Open(helper, name); if e != IsDirectoryError {
		t.Fatalf("error opening '%v': got '%v' or nil, expected %v", name, e, IsDirectoryError)
	}

	fileList := []string{"Cowards die many times before their deaths; The valiant never taste of death but once.",
			"Of all the wonders that I yet have heard, it seems to me most strange that men should fear;",
			"Seeing that death, a necessary end, will come when it will come. - (Act II, Scene II).",
			}

	errorCount := 0
	helper.SetE(func(path string, e error) {
		errorCount++
	})

	OpenFiles(helper, fileList...)

	time.Sleep(2 * time.Second)

	if errorCount != len(fileList) {
		t.Errorf("expected %v open errors, got %v", len(fileList), errorCount)
	}
}

func TestOpenFiles(t *testing.T) {
	fileList, e := filepath.Glob("*test.go"); if e != nil {
		t.Fatalf("unexpected glob error: %v", e)
	}
	t.Logf("using glob list: %v", fileList)

	helper := ReadOnly().ToggleFileAdviceReadSequential()
	files := testOpenFiles(t, helper, fileList...)
	Close(files...)

	files = testOpenFiles(t, helper, fileList...)

	// parallel close in separate goroutine
	closing := CloseChannel(helper)
	for _, file := range files {
		closing <-file
	}
}

func testOpen(t *testing.T, filename string) (file FileInterface) {
	helper := ReadOnly().ToggleFileAdviceReadSequential()
	file, e := Open(helper, filename); if e != nil {
		t.Fatalf("error opening file: '%v'", e)
	}; defer Close(file)

	testReadFile(t, file)

	return
}

func testOpenFiles(t *testing.T, helper *FileHelper, fileList ...string) (files []FileInterface) {
	helper.SetE(func(path string, e error) {
		t.Errorf("unexpected open error on '%v': %v", path, e)
	})
	for file := range OpenFiles(helper, fileList...) {
		testReadFile(t, file)
		files = append(files, file)
	}
	return
}

func testReadFile(t *testing.T, file FileInterface) {
	file.Seek(0, 0)

	info, e := file.Stat(); if e != nil {
		t.Fatalf("error getting file info: '%v'", e)
	}

	content, e := ioutil.ReadAll(file); if e != nil {
		t.Fatalf("error reading file: '%v'", e)
	}

	if int64(len(content)) != info.Size() {
		t.Errorf("expected to read %v, but read %v bytes!", info.Size(), len(content))
	}
}

func testWriteFile(t *testing.T, file FileInterface) {
	data := "Friends, Romans, countrymen, lend me your ears; I come to bury Caesar, not to praise him. - (Act III, Scene II)."
	written, e := file.WriteString(data)
	if e != nil {
		t.Errorf("unexpected error writing to file: '%v'", e)
	}

	if written != len(data) {
		t.Errorf("short write while testing: expected %v, got %v", len(data), written)
	}
}
