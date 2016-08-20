package iotool

import(
	"errors"
	"fmt"
	"io"
	"os"
)


// quickndirty stdin stdout fake for file
type FakeStdFile struct {
	stdin io.Reader
	stdout io.Writer
}

// provide stdin and stdout
func NewFakeStdFile(astdin io.Reader, astdout io.Writer) *FakeStdFile {
	return &FakeStdFile{ stdin: astdin, stdout: astdout }
}

// os.Stdin and os.Stdout
func NewFakeStdFileCli() *FakeStdFile {
	return &FakeStdFile{ stdin: os.Stdin, stdout: os.Stdout }
}

var NotAvailableE = errors.New("That function is not available. This file is faking.")

func (fake *FakeStdFile) Chdir() error { return NotAvailableE }
func (fake *FakeStdFile) Chmod(mode os.FileMode) error { return nil }
func (fake *FakeStdFile) Chown(uid, gid int) error { return nil }
func (fake *FakeStdFile) Close() error { return nil }
func (fake *FakeStdFile) Fd() uintptr { return uintptr(0) }
func (fake *FakeStdFile) Name() string { return "stdin(-)" }
func (fake *FakeStdFile) Read(b []byte) (n int, err error) { return fake.stdin.Read(b) }
func (fake *FakeStdFile) ReadAt(b []byte, off int64) (n int, err error) { return fake.stdin.Read(b) }
func (fake *FakeStdFile) Readdir(n int) (fi []os.FileInfo, err error) { return nil, NotAvailableE }
func (fake *FakeStdFile) Readdirnames(n int) (names []string, err error) { return nil, NotAvailableE }
func (fake *FakeStdFile) Seek(offset int64, whence int) (ret int64, err error) { return 0, nil }
func (fake *FakeStdFile) Stat() (os.FileInfo, error) { return nil, NotAvailableE }
func (fake *FakeStdFile) Sync() error { return nil }
func (fake *FakeStdFile) Truncate(size int64) error { return nil }
func (fake *FakeStdFile) Write(b []byte) (n int, err error) { return fake.stdout.Write(b) }
func (fake *FakeStdFile) WriteAt(b []byte, off int64) (n int, err error)  { return fake.stdout.Write(b) }
func (fake *FakeStdFile) WriteString(s string) (n int, err error) { return fmt.Fprint(fake.stdout, s) }
