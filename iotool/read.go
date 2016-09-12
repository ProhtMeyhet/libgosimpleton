package iotool

import(
	"io"
)

// go; read a file into a channel buffer
func ReadFileIntoBuffer(helper *FileHelper, path string, buffered *NamedBuffer) {
	handler, e := Open(helper, path); if e != nil { return }; defer handler.Close()
	ReadIntoBuffer(helper, handler, buffered)
}

// go; read a reader into a buffer
func ReadIntoBuffer(helper *FileHelper, reader io.Reader, buffered *NamedBuffer) {
	_, e := io.Copy(buffered, reader); if e != nil { helper.RaiseError(buffered.Name(), e) }
	buffered.Close()
}
