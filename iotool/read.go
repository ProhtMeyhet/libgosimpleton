package iotool

import(
	"io"
//	"fmt"
)

// read a file into a channel buffer; set the read size in helper. error handling can also be done via helper
func ReadFileIntoBuffer(helper *FileHelper, path string, buffers chan *NamedBuffer) (e error) {
	handler, e := Open(helper, path); if e != nil { return }; defer handler.Close()
	return readIntoBuffer(helper, path, handler, buffers)
}

// read reader into a channel buffer; set the read size in helper. error handling can also be done via helper
func ReadIntoBuffer(helper *FileHelper, reader io.Reader, buffers chan *NamedBuffer) (e error) {
	return readIntoBuffer(helper, "", reader, buffers)
}

// read reader into a channel buffer; set the read size in helper. error handling can also be done via helper
func readIntoBuffer(helper *FileHelper, name string, reader io.Reader, buffers chan *NamedBuffer) (e error) {
	cancel := make(chan bool, 1)
	namedBuffer := NewNamedBuffer(name, uint(helper.ReadSize()), buffers, cancel)
	startBuffer := NewNamedBuffer(name, 1, buffers, cancel); startBuffer.next = true

	buffers <-startBuffer

infinite:
	for {
		namedBuffer.read, e = reader.Read(namedBuffer.buffer)
		if e != nil {
			if e == io.EOF { e = nil; break infinite }
			// TODO find out what errors can happen here and handle them
			break infinite
		}
		buffers <-namedBuffer
		namedBuffer = NewNamedBuffer(name, uint(helper.ReadSize()), buffers, cancel)

		select {
		case <-cancel:
			break
		default:
		}
	}

	doneBuffer := NewNamedBuffer(name, 1, buffers, cancel); doneBuffer.done = true
	buffers <-doneBuffer

	return
}
