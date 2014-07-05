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

func NewBufferLogger(config LogConfigInterface) *bufferLogger {
	buffer := &bufferLogger{}
	inject(buffer, config)
	return buffer
}

func (buffer *bufferLogger) Open() (e error) {
	buffer.levels = make([]uint8, 20)
	buffer.messages = make([]string, 20)
	return
}

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
