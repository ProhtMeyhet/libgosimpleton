package logging

import(

)

/*
* if you can't figure out what this logger does
* this comment won't help you
*/

type nullLogger struct {
	abstractLogger
}

func NewNullLogger(config LogConfigInterface) logInterface {
	return &nullLogger{}
}

func (null *nullLogger) Open() (e error) { return }
func (null *nullLogger) Log(level uint8, message string) {}
func (null *nullLogger) Close() (e error) { return }
func (null *nullLogger) run() {} // overwrite as it is uneeded to run a goroutine for this
