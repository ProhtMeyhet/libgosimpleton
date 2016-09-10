package iotool

import(

)

type NamedBuffer struct {
	name	string
	key	uint

	buffer	[]byte
	done	bool
	next	bool
	read	int
	buffers chan *NamedBuffer
	cancel	chan bool
	isCanceled bool
}

// shiny and fresh
func NewNamedBuffer(aname string, asize uint, abuffers chan *NamedBuffer, acancel chan bool) (namedBuffer *NamedBuffer) {
	namedBuffer = &NamedBuffer{ name: aname }
	namedBuffer.buffer = make([]byte, asize)
	namedBuffer.cancel = acancel
	namedBuffer.buffers = abuffers
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

// returns the bytes read
func (buffer *NamedBuffer) Bytes() []byte {
	return buffer.buffer[:buffer.read]
}

// returns the bytes read as string
func (buffer *NamedBuffer) String() string {
	return string(buffer.buffer[:buffer.read])
}

// is this buffer channel reused
func (buffer *NamedBuffer) Next() bool {
	return buffer.next
}

// returns if this file is done
func (buffer *NamedBuffer) Done() bool {
	return buffer.done
}

// returns the number of bytes read
func (buffer *NamedBuffer) Read() int {
	return buffer.read
}

// cancel reading; clear buffers; blocks till cancel is done
func (buffer *NamedBuffer) Cancel() {
	buffer.cancel <-true
	buffer.isCanceled = true

	// empty the buffers
	for _ = range buffer.buffers {}
}

func (buffer *NamedBuffer) IsCanceled() bool {
	return buffer.isCanceled
}
