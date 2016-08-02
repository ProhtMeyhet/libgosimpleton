package iotool

import(
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/ProhtMeyhet/libgosimpleton/simpleton"
)

var random *rand.Rand

// open a file with FileHelper. if its a dir, return nil.
// it's your responsibility to close the file...
func Open(fileHelper *FileHelper, path string) (handler FileInterface, e error) {
	// FIXME TODO
/*	if path == abstract.STDIN_TOKEN {
		return NewFakeStdFile(fileHelper.Stdin(), fileHelper.Stdout())
	}*/

	handler, e = os.OpenFile(path, fileHelper.OpenFlags(), fileHelper.Permissions())
	if e != nil { fileHelper.RaiseError(path, e); goto out }
	if !fileHelper.DoNotTestForDirectory() {
		// ie := darn shadowing...
		fileInfo, ie := handler.Stat(); if ie != nil { e = ie; fileHelper.RaiseError(path, ie); goto out }
		if fileInfo.IsDir() { e = IsDirectoryError; fileHelper.RaiseError(path, e); goto out }
	}
	fileHelper.ApplyFileAdvice(handler)

out:
	return
}

// open N files. please keep limits like ulimit file.max in mind.
// see also parallel.OpenFiles() for a parallel version and parallel.OpenFilesIntoChannel()
// consider using iotool.ToggleFileAdviceWillNeed(). for error handling, use FileHelper.SetE().
// it's your responsibility to close the files...
func OpenFiles(fileHelper *FileHelper, files ...string) <-chan FileInterface {
	return OpenFilesFromChannel(fileHelper, simpleton.StringListToChannel(files...))
}

// open N files from channel. please keep limits like ulimit file.max in mind.
// see also parallel.OpenFiles() for a parallel version and parallel.OpenFilesIntoChannel()
// consider using iotool.ToggleFileAdviceWillNeed(). for error handling, use FileHelper.SetE().
// it's your responsibility to close the files...
func OpenFilesFromChannel(fileHelper *FileHelper, fileList <-chan string) <-chan FileInterface {
	handlers := make(chan FileInterface, fileHelper.WorkerBuffers())
	go func() {
		for path := range fileList {
			// no need to call fileHelper.RaiseError, it's called in Open() already
			if handler, e := Open(fileHelper, path); e == nil {
				handlers <-handler
			}
		}; close(handlers)
	}()
	return handlers
}

// create a file. if you need it truncated, use FileHelper.ToggleTruncate().
func Create(fileHelper *FileHelper, path string) (handler FileInterface, e error) {
	if !fileHelper.HasCreate() {
		fileHelper.ToggleCreate()
	}

	return Open(fileHelper, path)
}

// create a temporary file. only makes sure the file is unique. file will not be removed.
// prefix should be your program name.
func Temporary(fileHelper *FileHelper, prefix string) (handler FileInterface, e error) {
	if !fileHelper.HasExclusive() {
		fileHelper.ToggleExclusive()
	}

again:
	handler, e = Create(fileHelper, TemporaryName(prefix))
	if e != nil && os.IsExist(e) { goto again }

	// pass through all other errors, there isn't anything that can be done here
	return
}

// return a filename with full path in os.TempDir(). prefix should be your program name.
func TemporaryName(prefix string) string {
	// if not testing, seed with current time.
	if !DEBUG {
		if random == nil {
			// gotta seed manually these days...
			random = rand.New(rand.NewSource(time.Now().UnixNano()))
		}
	}
	return temporaryName(prefix)
}

// for testing purposes this function was split
func temporaryName(prefix string) string {
	directory := os.TempDir()
	name := prefix + PREFIX_SEPARATOR + strconv.Itoa(int(uint(random.Int63())))
	// FIXME: why the bloody fuck is os.PathSeparator a rune ?!?
	return directory + string(os.PathSeparator) + name
}

// close a bunch of files from a list.
func Close(fileList ...FileInterface) {
	for _, file := range fileList {
		file.Close()
	}
}

// give back a close chan to push FileInterfaces to; when done, close the channel.
func CloseChannel(fileHelper *FileHelper) chan<- FileInterface {
	closing := make(chan FileInterface, 16)
	CloseFilesFromChannel(fileHelper, closing)
	return closing
}

// close a bunch of files via a channel
func CloseFilesFromChannel(fileHelper *FileHelper, fileList chan FileInterface) {
	go func() {
		for handler := range fileList {
			if handler != nil {
				e := handler.Close(); if e != nil {
					fileHelper.RaiseError(handler.Name(), e)
				}
			}
		}
	}()
}


/* convinience functions */
// TODO add more

func Remove(handler FileInterface) error {
	return os.Remove(handler.Name())
}

// use FileHelper.ToggleTruncate(), otherwise os.Truncate().
// func Truncate(handler FileInterface, size int64) error {}
