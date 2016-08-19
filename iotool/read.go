package iotool

import(
	"io"
)

// read a file into a channel buffer; set the read size in helper. error handling can also be done via helper
func ReadFileIntoBuffer(helper *FileHelper, path string, buffers chan NamedBuffer) (e error) {
	handler, e := Open(helper, path); if e != nil { return }; defer handler.Close()
	return readIntoBuffer(helper, path, handler, buffers)
}

// read reader into a channel buffer; set the read size in helper. error handling can also be done via helper
func ReadIntoBuffer(helper *FileHelper, reader io.Reader, buffers chan NamedBuffer) (e error) {
	return readIntoBuffer(helper, "", reader, buffers)
}

// read reader into a channel buffer; set the read size in helper. error handling can also be done via helper
func readIntoBuffer(helper *FileHelper, name string, reader io.Reader, buffers chan NamedBuffer) (e error) {
	namedBuffer := NewNamedBuffer(name); namedBuffer.buffer = make([]byte, helper.ReadSize())
infinite:
	for {
		namedBuffer.read, e = reader.Read(namedBuffer.buffer)
		if e != nil {
			if e == io.EOF { e = nil; break infinite }
			// TODO find out what errors can happen here and handle them
			break infinite
		}

		buffers <-namedBuffer
	}

	namedBuffer.done = true
	buffers <-namedBuffer

	return
}
