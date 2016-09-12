package iotool

import(
	"io"
)

type NamedBuffer struct {
	*io.PipeReader
	*io.PipeWriter

	name	string
	key	uint

	isCanceled bool
}

// shiny and fresh
func NewNamedBuffer(aname string) (namedBuffer *NamedBuffer) {
	namedBuffer = &NamedBuffer{ name: aname }
	namedBuffer.PipeReader, namedBuffer.PipeWriter = io.Pipe()
	return
}

// returns the name of the file
func (buffer *NamedBuffer) Name() string {
	return buffer.name
}

// if the file was in a sequence, eg a list, returns the key
func (buffer *NamedBuffer) Key() uint {
	return buffer.key
}

func (buffer *NamedBuffer) SetKey(to uint) {
	buffer.key = to
}

// cancel reading; clear buffers; blocks till cancel is done
func (buffer *NamedBuffer) Cancel() {
	buffer.isCanceled = true
	buffer.Close()
}

func (buffer *NamedBuffer) Close() {
	buffer.PipeWriter.Close()
}

func (buffer *NamedBuffer) IsCanceled() bool {
	return buffer.isCanceled
}
