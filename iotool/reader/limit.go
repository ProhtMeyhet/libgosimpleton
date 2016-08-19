package reader

import(
	"bufio"
	"bytes"
	"io"
)

// reads only n splitted; if reached return io.EOF even though its not EOF
type LimitReader struct {
	// interface
	scanner		*bufio.Scanner

	// have a guess
	max		uint

	// lines counted
	counted		uint

	// required if line is longer then cap(Read(readInto))
	buffered	*bytes.Buffer

	// reappend this byte sequence
	add		[]byte
}

// fresh and shiny
func NewLimitReader(areader io.Reader, amax uint, ascanner *bufio.Scanner, add []byte) (reader *LimitReader) {
	reader = &LimitReader{}
	reader.Initialise(amax, ascanner, add)
	return
}

// i n i t i a l i s e
func (reader *LimitReader) Initialise(amax uint, ascanner *bufio.Scanner, add []byte) {
	if ascanner == nil { panic("scanner cannot be nil!") } //TODO
	reader.max = amax
	reader.scanner = ascanner
	reader.add = add
}

// interface: io.Reader
func (reader *LimitReader) Read(readInto []byte) (read int, e error) {
	if reader.buffered != nil {
		read, e = reader.buffered.Read(readInto)
		if e != nil && e == io.EOF {
			reader.buffered = nil; e = nil
		}
		return
	}

	if !reader.scanner.Scan() || reader.counted == reader.max { return 0, io.EOF }

	readBytes := reader.scanner.Bytes(); readBytes = append(readBytes, reader.add...); reader.counted++
	if len(readBytes) > cap(readInto) {
		reader.buffered = bytes.NewBuffer(readBytes)
		return reader.buffered.Read(readInto)
	}

	copy(readInto, readBytes[:len(readBytes)])
	return len(readBytes), nil
}


// read maximum n lines
func NewLimitLineReader(areader io.Reader, amaxLines uint) (reader *LimitReader) {
	scanner := bufio.NewScanner(areader)
	scanner.Split(bufio.ScanLines)
	reader = NewLimitReader(areader, amaxLines, scanner, []byte("\n"))
	return
}

// read maximum n words
func NewLimitWordReader(areader io.Reader, amaxLines uint) (reader *LimitReader) {
	scanner := bufio.NewScanner(areader)
	scanner.Split(bufio.ScanWords)
	reader = NewLimitReader(areader, amaxLines, scanner, []byte(" "))
	return
}

// read maximum n bytes
func NewLimitBytesReader(areader io.Reader, amaxLines uint) (reader *LimitReader) {
	scanner := bufio.NewScanner(areader)
	scanner.Split(bufio.ScanBytes)
	reader = NewLimitReader(areader, amaxLines, scanner, []byte(""))
	return
}

// read maximum n runes
func NewLimitRuneReader(areader io.Reader, amaxLines uint) (reader *LimitReader) {
	scanner := bufio.NewScanner(areader)
	scanner.Split(bufio.ScanRunes)
	reader = NewLimitReader(areader, amaxLines, scanner, []byte(""))
	return
}
