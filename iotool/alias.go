package iotool

import(
	"io"
)

// copy FROM -> TO
func Copy(source io.Reader, destination io.Writer) (written int64, e error) {
	return io.Copy(destination, source)
}

// copy FROM -> TO
func CopyBuffer(source io.Reader, destination io.Writer, buffer []byte) (written int64, e error) {
	return io.CopyBuffer(destination, source, buffer)
}

// copy FROM -> TO
func CopyN(source io.Reader, destination io.Writer, n int64) (written int64, e error) {
	return io.CopyN(destination, source, n)
}
