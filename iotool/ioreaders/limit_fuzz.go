package ioreaders

// +build gofuzz

import(
	"bytes"
	"io"
)

func Fuzz(data []byte) int {
	buffer := bytes.NewBuffer(make([]byte, 0))
	buffer.Write(data); total := 0
	limited := NewLimitBytesReader(buffer, 14); into := make([]byte, 4)
	for {
		read, e := limited.Read(into)
		if e != nil {
			if e != io.EOF {
				panic("unexpected error! " + e.Error())
			}

			return 0
		}
		total += read
	}

	if total > 14 {
		return 0
	}

	return 1
}
