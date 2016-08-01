package caching

// +build linux

import(
	"testing"
	"os"
	"time"

	"github.com/ProhtMeyhet/libgosimpleton/logging"
	"github.com/ProhtMeyhet/libgosimpleton"
)


func TestObservingFileCache(t *testing.T) {
	filename := os.TempDir() + "/testingoleeeeoleeeoleeee65496498"
	text1 := "I am the captain of the Pinafore\n"
	text2 := "And a right good captain, too!"
	textWhole := text1 + text2
	textCache := ""

	// write a temp file
	temp, e := os.Create(filename); //defer os.Remove(filename)
	if e != nil { t.Fatalf("error opening temp file: %v", e.Error()) }
	_, e = temp.WriteString(text1)
	if e != nil { t.Fatalf("error writing to temp file: %v", e.Error()) }

	// get temp file contents from cache
	cache := NewObservingFileCache()
	_, e = cache.Get(filename)
	if e != nil { t.Fatalf("cache error: %v", e.Error()) }

	// change file, should trigger inotify
	_, e = temp.WriteString(text2)
	if e != nil { t.Fatalf("error writing to temp file: %v", e.Error()) }
	if libgosimpleton.DEBUG { logging.Debug("changed file, expecting inotify") }

	// shouldn't take that long, just to be sure
	time.Sleep(time.Second * 2)

	// get the file again, should have been altered
	if textCache, e = cache.GetString(filename); e != nil { t.Errorf("error cache: %v", e.Error()) }

	if textCache != textWhole {
		t.Errorf("expected:\n\n %v \n\n got:\n\n %v", textWhole, textCache)
	}

	cache.Close()
	if len(cache.cache) > 0 {
		t.Errorf("expected cache.cache to be empty! still %v items!", len(cache.cache))
	}
}
