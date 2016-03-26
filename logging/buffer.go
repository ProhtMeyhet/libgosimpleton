package logging

import(

)

/*
* this logger buffers all messages in memory
* and can flush them to another logger
* used for early logging
*/

type bufferLogger struct {
	abstractLogger
	levels []uint8
	messages []string
}

// don't export a buffer logger as it is useless outside this package
func newBufferLogger(config LogConfigInterface) (buffer *bufferLogger) {
	buffer = &bufferLogger{}
	buffer.initialise(config)
	return
}

func (buffer *bufferLogger) Open() (e error) { return }

func (buffer *bufferLogger) Log(level uint8, message string) {
	buffer.levels = append(buffer.levels, level)
	buffer.messages = append(buffer.messages, message)
}

func (buffer *bufferLogger) Close() (e error) {
	buffer.levels = nil
	buffer.messages = nil
	return
}

func (buffer *bufferLogger) Flush(to logInterface) (e error) {
	for k, _ := range buffer.messages {
		to.Log(buffer.levels[k], buffer.messages[k])
	}

	return
}
