package parallel

import(
	"io"

	"github.com/ProhtMeyhet/libgosimpleton/iotool"
	"github.com/ProhtMeyhet/libgosimpleton/simpleton"
)

// Open File and Do Work
//
// OpenFilesDoWork(helper, paths, func(buffers chan NamedBuffer) {
// 	for buffered := range buffers {
// 		/* do work */
// 	}
// })
func OpenFileDoWork(helper *iotool.FileHelper, path string, worker func(chan NamedBuffer)) {
	// use 1 inner worker, see below
	work := NewWorkManual(1); handlerBuffer := make([]byte, helper.ReadSize())
	buffers := make(chan NamedBuffer, work.SuggestFileBufferSize())

	// read in one thread
	work.Feed(func() {
		readFile(helper, buffers, handlerBuffer, path)
		close(buffers)
	})

	// start working
	work.Start(func() {
		worker(buffers)
	})

	work.Wait()
}

// Open N Files and Do Work
//
// OpenFilesDoWork(helper, paths, func(buffers chan NamedBuffer) {
// 	for buffered := range buffers {
// 		/* do work */
// 	}
// })
func OpenFilesDoWork(helper *iotool.FileHelper, paths <-chan string, worker func(chan NamedBuffer)) {
	work := NewWork(0)

	// feed is list

	work.Start(func() {
		for path := range paths {
			OpenFileDoWork(helper, path, worker)
		}
	})

	work.Wait()
}

func OpenFilesFromListDoWork(helper *iotool.FileHelper, worker func(chan NamedBuffer), paths ...string) {
	OpenFilesDoWork(helper, simpleton.StringListToChannel(paths...), worker)
}

func readFile(helper *iotool.FileHelper, buffers chan NamedBuffer, handlerBuffer []byte, path string) (e error) {
	handler, e := iotool.Open(helper, path); if e != nil { return }; defer handler.Close()
	namedBuffer := NewNamedBuffer(path); namedBuffer.buffer = make([]byte, len(handlerBuffer)); read := 0 // avoid e shadowing

infinite:
	for {
		read, e = handler.Read(handlerBuffer)
		if e != nil {
			if e == io.EOF { e = nil; break infinite }
			// TODO find out what errors can happen here and handle them
			break infinite
		}

	// FIXME i thought the following line would copy. doesn't seem to be or bug.
	//	buffer <-handlerBuffer[:read]
		namedBuffer.read = read
		copy(namedBuffer.buffer, handlerBuffer[:read])
		buffers <-namedBuffer
	}

	namedBuffer.done = true
	buffers <-namedBuffer

	return
}
