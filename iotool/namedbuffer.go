package iotool

import(

)

type NamedBuffer struct {
	name string
	buffer []byte
	done bool
	read int
}

// shiny and fresh
func NewNamedBuffer(aname string) (namedBuffer NamedBuffer) {
	namedBuffer = NamedBuffer{ name: aname }
	return
}

// returns the name of the file
func (buffer *NamedBuffer) Name() string {
	return buffer.name
}

// returns the bytes read
func (buffer *NamedBuffer) Bytes() []byte {
	return buffer.buffer[:buffer.read]
}

// returns if this file is done
func (buffer *NamedBuffer) Done() bool {
	return buffer.done
}

// returns the number of bytes read
func (buffer *NamedBuffer)Read() int {
	return buffer.read
}
