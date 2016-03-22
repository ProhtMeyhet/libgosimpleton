package caching

import(
	"testing"
	"os"
	"strings"
	"time"

	"github.com/ProhtMeyhet/libgosimpleton/logging"
)

/* 	TODO
*	add tests for remove
*/

func TestMain(m *testing.M) {
	logConfig := logging.NewDefaultConfig()
	logConfig.Name = "simpleton"
	logConfig.LogType = "stderr"
	logConfig.Level = logging.EVERYTHING
	logging.Start(logConfig)

	os.Exit(m.Run())
}

func TestMaxSize(t *testing.T) {
	filename1 := os.TempDir() + "/testingpinafore56464968468413"
	filename2 := filename1 + "2"
	filename3 := filename1 + "3"
	filename4 := filename1 + "4"
	text1 := "What never?\n"
	text2 := "Weeell, hardly ever!\n"
	text3 := strings.Repeat("ha", 256) // 512 ;-)

	// create files
	temp1, e := os.Create(filename1); defer os.Remove(filename1)
	if e != nil { t.Fatalf("error opening temp file: %v", e.Error()) }
	temp2, e := os.Create(filename2); defer os.Remove(filename2)
	if e != nil { t.Fatalf("error opening temp file: %v", e.Error()) }
	temp3, e := os.Create(filename3); defer os.Remove(filename3)
	if e != nil { t.Fatalf("error opening temp file: %v", e.Error()) }
	temp4, e := os.Create(filename4); defer os.Remove(filename4)
	if e != nil { t.Fatalf("error opening temp file: %v", e.Error()) }

	// write to files, TODO short write test
	_, e = temp1.WriteString(text1)
	if e != nil { t.Fatalf("error writing to temp file: %v", e.Error()) }
	_, e = temp2.WriteString(text2)
	if e != nil { t.Fatalf("error writing to temp file: %v", e.Error()) }
	_, e = temp3.WriteString(text3)
	if e != nil { t.Fatalf("error writing to temp file: %v", e.Error()) }
	_, e = temp4.WriteString(text3)
	if e != nil { t.Fatalf("error writing to temp file: %v", e.Error()) }

	// get file contents via cache
	cache := NewStupidFileCacheMaxSize(1024) // 1kb

	value, e := cache.Get(filename1)
	if e != nil { t.Fatalf("cache error: %v", e.Error()) }
	if value != text1 { t.Errorf("expected %v, got %v", text1, value) }

	value, e = cache.Get(filename2)
	if e != nil { t.Fatalf("cache error: %v", e.Error()) }
	if value != text2 { t.Errorf("expected %v, got %v", text2, value) }

	value, e = cache.Get(filename3)
	if e != nil { t.Fatalf("cache error: %v", e.Error()) }
	if value != text3 { t.Errorf("expected %v, got %v", text3, value) }

	if cache.Free() <= 0 {
		t.Errorf("unexpected: cache.Free() is 0!")
	}

	// exhaust cache, free should happen
	_, e = cache.Get(filename4)
	if e != nil { t.Fatalf("cache error: %v", e.Error()) }

	// cache free shouldn't take that long, just to be sure
	time.Sleep(time.Second * 2)

	if _, ok := cache.cache[filename3]; ok {
		t.Errorf("took longer then 2s to remove filename3 from cache!")
	}
}
