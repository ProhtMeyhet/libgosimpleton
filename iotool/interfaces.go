package iotool

import(
	"io"
	"os"
)

type FileInterface interface {
	io.Reader
	io.ReaderAt
	io.Writer
	io.WriterAt
	io.Seeker

	Chdir() error
	Chmod(mode os.FileMode) error
	Chown(uid, gid int) error
	Fd() uintptr
	Name() string
	Readdir(n int) (fi []os.FileInfo, err error)
	Readdirnames(n int) (names []string, err error)
	Stat() (os.FileInfo, error)
	Sync() error
	Truncate(size int64) error
	WriteString(s string) (n int, err error)
	Close() error
}
