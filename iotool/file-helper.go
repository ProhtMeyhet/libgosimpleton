package iotool

// +build linux

import(
	"io"
	"os"
	"sync"

	"golang.org/x/sys/unix"

	"github.com/ProhtMeyhet/libgosimpleton"
	"github.com/ProhtMeyhet/libgosimpleton/abstract"
)

// this struct is used to reduce function parameters.
type FileHelper struct {
	sync.Mutex

	// the more help, the better. lots and lots and lots of help needed!
	abstract.BaseHelper
	abstract.WorkerHelper

	// os.O_RDONLY and such
	openFlags			int

	// file permissions
	permissions			os.FileMode

	// turn off checking if it's a directory. one syscall less, but the caller has to do the work.
	doNotTestForDirectory		bool

	// the read size for a buffer for io.Reader
	// FIXME -> uint
	readSize			int

	// toggled by setting stdout and stdin
	supportCli			bool
	// must not be empty!
	stdinToken			string
	stdout				io.Writer
	stdin				io.Reader
	// as a special case also support stderr
	stderr				io.Writer

	// reset all file advices
	fileAdviceNormal		bool

	// previous cached stat calls
					// map[path]os.FileInfo
	fileInfo			map[string]FileInfoInterface
//	stats				map[string]os.FileInfo
//	lstats				map[string]os.FileInfo

	// Access data only once.
	// FIXME: this is not implemented in the Linux kernel! (maybe doable via madvice?)
	fileAdviceNoReuse		bool

	// Do not expect access in the near future. Subsequent access of pages
	// in this range will succeed, but will result either in reloading of
	// the memory contents from the underlying mapped file or zero-fill-in-demand
	// pages for mappings without an underlying file.
	// Linux: Drop the file from cache. Note this is automatically done when files are unlinked.
	fileAdviceDontNeed		bool

	// Expect page references in random order
	// Linux: Sets FMODE_RANDOM on the file descriptor (fd)
	fileAdviceReadRandom		bool

	// Expect page references in sequential order 
	// Linux: Doubles the size of read ahead done for file
	fileAdviceReadSequential	bool

	// Expect access in the near future
	// Linux: _synchronously_ prepopulate the buffer cache with the file
	fileAdviceWillNeed		bool
}

// fresh and shiny
func newFileHelper() *FileHelper {
	helper := &FileHelper{}
	// FIXME onE
	helper.BaseHelper.Initialise(abstract.IgnoreErrors)
	helper.WorkerHelper.Initialise()
	helper.readSize = READ_BUFFER_SIZE
	helper.fileInfo = make(map[string]FileInfoInterface)
	return helper
}

// open for read only
func ReadOnly() *FileHelper {
	helper := newFileHelper()
	return helper.ReadOnly()
}

// open for write only
func WriteOnly() *FileHelper {
	helper := newFileHelper()
	return helper.WriteOnly()
}

// open read and write
func ReadAndWrite() *FileHelper {
	helper := newFileHelper()
	return helper.ReadAndWrite()
}

// set the file advices on file descriptors (fd). ignore all errors (just like coreutils does)
func (helper *FileHelper) ApplyFileAdvice(to FileInterface) {
	if to == nil { return }

	fd := int(to.Fd())

	if helper.fileAdviceNormal {
		unix.Fadvise(fd, 0, 0, 0) // 0 == POSIX_FADV_NORMAL
		helper.fileAdviceNormal = false
		return
	}

	if helper.ShouldFileAdviceReadRandom() {
		unix.Fadvise(fd, 0, 0, 1) // 1 == POSIX_FADV_RANDOM
	}

	if helper.ShouldFileAdviceReadSequential() {
		unix.Fadvise(fd, 0, 0, 2) // 2 == POSIX_FADV_SEQUENTIAL
	}

	// go, as the linux kernel is immidiatly calling force_page_cache_readahead
	if helper.ShouldFileAdviceWillNeed() {
		go unix.Fadvise(fd, 0, 0, 3) // 3 == POSIX_FADV_WILLNEED
	}

	// go, as the linux kernel is immidiatly doing some work
	if helper.ShouldFileAdviceDontNeed() {
		go unix.Fadvise(fd, 0, 0, 4) // 4 == POSIX_FADV_DONTNEED
	}

	if helper.ShouldFileAdviceNoReuse() {
		unix.Fadvise(fd, 0, 0, 5) // 5 == POSIX_FADV_NOREUSE
	}
}

// resets all file advices on $to and in this helper
func (helper *FileHelper) ResetFileAdvice(to FileInterface) {
	helper.fileAdviceNormal = true
	helper.fileAdviceReadRandom = false
	helper.fileAdviceReadSequential = false
	helper.fileAdviceWillNeed = false
	helper.fileAdviceDontNeed = false
	helper.fileAdviceNoReuse = false
	helper.ApplyFileAdvice(to)
}

// copy several values from a helper
func (helper *FileHelper) Copy(from interface{}) *FileHelper {
	// FIXME
//	helper.BaseHelper.Copy(from)
//	helper.WorkerHelper.Copy(from)

	if fileHelper, ok := from.(*FileHelper); ok {
		helper.fileAdviceNormal = fileHelper.fileAdviceNormal
		helper.fileAdviceReadRandom = fileHelper.fileAdviceReadRandom
		helper.fileAdviceReadSequential = fileHelper.fileAdviceReadSequential
		helper.fileAdviceWillNeed = fileHelper.fileAdviceWillNeed
		helper.fileAdviceDontNeed = fileHelper.fileAdviceDontNeed
		helper.fileAdviceNoReuse = fileHelper.fileAdviceNoReuse

		if fileHelper.HasAppend() { helper.ToggleAppend() }
		if fileHelper.HasCreate() { helper.ToggleCreate() }
		if fileHelper.HasExclusive() { helper.ToggleExclusive() }
		if fileHelper.HasSynchronous() { helper.ToggleSynchronous() }
		if fileHelper.HasTruncate() { helper.ToggleTruncate() }

		for key := range fileHelper.fileInfo {
			if _, ok := helper.fileInfo[key]; !ok {
				helper.fileInfo[key] = fileHelper.fileInfo[key]
			}
		}
	}

	return helper
}

// get read size
func (helper *FileHelper) ReadSize() int {
	return helper.readSize
}

func (helper *FileHelper) SetReadSize(to int) *FileHelper {
	// it wasn't a very wise decision to use signed ints for lengths...
	if to > 0 {
		helper.readSize = to
	}
	return helper
}

// FIXME: correct lstat cache
func (helper *FileHelper) FileInfo(path string, lstat bool) (FileInfoInterface, error) {
	var e error
	if _, ok := helper.fileInfo[path]; !ok {
		var info os.FileInfo
		if lstat {
			info, e = os.Lstat(path); if e != nil { return nil, e }
		} else {
			info, e = os.Stat(path); if e != nil { return nil, e }
		}
		helper.Lock(); helper.fileInfo[path] = NewFileInfo(path, info); helper.Unlock()
	}

	return helper.fileInfo[path], e
}

// return if cli output is supported
func (helper *FileHelper) SupportCli() bool {
	return helper.supportCli
}

// toggle supportCli by setting STDIN and STDOUT in SetStdinStdout()

// return if string is equal to STDIN token (usualley "-")
func (helper *FileHelper) IsStdinToken(what string) bool {
	return helper.stdinToken == what
}

// get STDIN to use
func (helper *FileHelper) Stdin() io.Reader {
	return helper.stdin
}

// get STDOUT to use
func (helper *FileHelper) Stdout() io.Writer {
	return helper.stdout
}

// get STDERR to use
func (helper *FileHelper) Stderr() io.Writer {
	return helper.stderr
}

// set STDIN and STDOUT. only support supplying both
func (helper *FileHelper) SetStdinStdout(astdinToken string, astdin io.Reader, astdout io.Writer) {
	helper.supportCli = true
	// disallow empty value
	if astdinToken != "" {
		helper.stdinToken = astdinToken
	} else {
		helper.stdinToken = STDIN_TOKEN
	}
	helper.stdin = astdin
	helper.stdout = astdout
}

// set STDERR to a io.Writer
func (helper *FileHelper) SetStderr(astderr io.Writer) {
	helper.stderr = astderr
}

// sets to read only and discards all other flags
func (helper *FileHelper) ReadOnly() *FileHelper {
	helper.openFlags = os.O_RDONLY
	return helper
}

// sets to write only and discards all other flags
func (helper *FileHelper) WriteOnly() *FileHelper {
	helper.openFlags = os.O_WRONLY
	return helper
}

// sets to read and write and discards all other flags
func (helper *FileHelper) ReadAndWrite() *FileHelper {
	helper.openFlags = os.O_RDWR
	return helper
}

// is append already active
func (helper *FileHelper) HasAppend() bool {
	return helper.openFlags ^ os.O_APPEND < helper.openFlags
}

// add append to flags
func (helper *FileHelper) ToggleAppend() *FileHelper {
	helper.openFlags ^= os.O_APPEND
	return helper
}

// is append already active
func (helper *FileHelper) HasCreate() bool {
	return helper.openFlags ^ os.O_CREATE < helper.openFlags
}

// add create to flags
func (helper *FileHelper) ToggleCreate() *FileHelper {
	helper.openFlags ^= os.O_CREATE
	helper.permissions = 0666
	return helper
}

// is exlusive already active
func (helper *FileHelper) HasExclusive() bool {
	return helper.openFlags ^ os.O_EXCL < helper.openFlags
}

// add exclusive to flags
func (helper *FileHelper) ToggleExclusive() *FileHelper {
	helper.openFlags ^= os.O_EXCL
	return helper
}

// is synchronized already active
func (helper *FileHelper) HasSynchronous() bool {
	return helper.openFlags ^ os.O_SYNC < helper.openFlags
}
// add sync to flags
func (helper *FileHelper) ToggleSynchronous() *FileHelper {
	helper.openFlags ^= os.O_SYNC
	return helper
}

// is truncate already active
func (helper *FileHelper) HasTruncate() bool {
	return helper.openFlags ^ os.O_TRUNC < helper.openFlags
}
// add trunc to flags
func (helper *FileHelper) ToggleTruncate() *FileHelper {
	helper.openFlags ^= os.O_TRUNC
	return helper
}

// return open flags
func (helper *FileHelper) OpenFlags() int {
	return helper.openFlags
}

// return permissions
func (helper *FileHelper) Permissions() os.FileMode {
	return helper.permissions
}

// set permissions
func (helper *FileHelper) SetPermissions(to os.FileMode) *FileHelper {
	helper.permissions = to
	return helper
}

// get if test for directory
func (helper *FileHelper) DoNotTestForDirectory() bool {
	return helper.doNotTestForDirectory
}

// please do file advice DONT_NEED
func (helper *FileHelper) ToggleFileAdviceDontNeed() *FileHelper {
	if !libgosimpleton.SET_FILE_ADVICE_DONTNEED { goto out }
	helper.fileAdviceDontNeed = !helper.fileAdviceDontNeed

out:
	return helper
}

// thats a question to be answered
func (helper *FileHelper) ShouldFileAdviceDontNeed() bool {
	return helper.fileAdviceDontNeed
}

// please do file advice SEQUENTIAL_READ
func (helper *FileHelper) ToggleFileAdviceReadSequential() *FileHelper {
	helper.fileAdviceReadSequential = !helper.fileAdviceReadSequential
	return helper
}

// thats a question to be answered
func (helper *FileHelper) ShouldFileAdviceReadSequential() bool {
	return helper.fileAdviceReadSequential
}

// please do file advice RANDOM_READ
func (helper *FileHelper) ToggleFileAdviceReadRandom() *FileHelper {
	helper.fileAdviceReadRandom = !helper.fileAdviceReadRandom
	return helper
}

// thats a question to be answered
func (helper *FileHelper) ShouldFileAdviceReadRandom() bool {
	return helper.fileAdviceReadRandom
}

// please do file advice WILLNEED
func (helper *FileHelper) ToggleFileAdviceWillNeed() *FileHelper {
	helper.fileAdviceWillNeed = !helper.fileAdviceWillNeed
	return helper
}

// thats a question to be answered
func (helper *FileHelper) ShouldFileAdviceWillNeed() bool {
	return helper.fileAdviceWillNeed
}

// please do file advice NO_REUSE
func (helper *FileHelper) ToggleFileAdviceNoReuse() *FileHelper {
	helper.fileAdviceNoReuse = !helper.fileAdviceNoReuse
	return helper
}

// thats a question to be answered
func (helper *FileHelper) ShouldFileAdviceNoReuse() bool {
	return helper.fileAdviceNoReuse
}

// @override toggle cache you should
func (helper *FileHelper) ToggleCache() {
	helper.BaseHelper.ToggleCache()
	if !helper.ShouldCache() {
		helper.ToggleFileAdviceDontNeed()
		helper.ToggleFileAdviceNoReuse()
	}
}
