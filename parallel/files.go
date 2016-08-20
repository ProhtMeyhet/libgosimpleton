package parallel

import(
	"io"

	"github.com/ProhtMeyhet/libgosimpleton/iotool"
	"github.com/ProhtMeyhet/libgosimpleton/simpleton"
)

// OpenFileDoWork(helper, path, func(buffers chan NamedBuffer) {
// 	for buffered := range buffers {
//		if buffered.Done() {
//			fmt.Println("done!")
//			continue
//		}
// 		/* do work */
// 	}
// })

// Open File and Do Work
func OpenFileDoWork(helper *iotool.FileHelper, path string, worker func(chan *iotool.NamedBuffer)) {
	// use 1 inner worker, see below
	work := NewWorkManual(1); buffers := make(chan *iotool.NamedBuffer, work.SuggestFileBufferSize())

	// read in one thread
	work.Feed(func() {
		iotool.ReadFileIntoBuffer(helper, path, buffers)
		close(buffers)
	})

	// work in another
	work.Run(func() {
		worker(buffers)
	})
}

// OpenFilesDoWork(helper, paths, func(buffers chan NamedBuffer) {
// 	for buffered := range buffers {
//		if buffered.Done() {
//			fmt.Println("done!")
//			continue
//		}
// 		/* do work */
// 	}
// })

// Open N Files and Do Work
func OpenFilesDoWork(helper *iotool.FileHelper, paths <-chan string, worker func(chan *iotool.NamedBuffer)) {
	work := NewWork(0)

	// feed is list

	work.Start(func() {
		for path := range paths {
			OpenFileDoWork(helper, path, worker)
		}
	})

	work.Wait()
}

// OpenFilesFromListDoWork(helper, func(buffers chan NamedBuffer) {
// 	for buffered := range buffers {
//		if buffered.Done() {
//			fmt.Println("done!")
//			continue
//		}
// 		/* do work */
// 	}
// }, path1, path2, paths...)

// Open N Files and Do Work
func OpenFilesFromListDoWork(helper *iotool.FileHelper, worker func(chan *iotool.NamedBuffer), paths ...string) {
	OpenFilesDoWork(helper, simpleton.StringListToChannel(paths...), worker)
}

// ReadFilesSequential(helper, path, func(buffers chan NamedBuffer) {
// 	for buffered := range buffers {
//		if buffered.Done() {
//			fmt.Println("done!")
//			continue
//		}
// 		/* do work */
// 	}
// })

// read files sequential to a buffer
func ReadFilesSequential(helper *iotool.FileHelper, paths []string, worker func(chan *iotool.NamedBuffer)) {
	ReadFilesFilteredSequential(helper, paths, nil, worker)
}

// ReadFilesFilteredSequential(helper, paths, func(reader io.Reader) {
//	return iotool.LimitLinesReader(reader, 10)
// }, func(buffers chan NamedBuffer) {
// 	for buffered := range buffers {
//		if buffered.Done() {
//			fmt.Println("done!")
//			continue
//		}
// 		/* do work */
// 	}
// })

// read files sequential to a buffer
// uses file advice will need and thus relies on kernel page caching. this is good
// for being a team player, but may hurt performance especially in syntethic tests.
func ReadFilesFilteredSequential(helper *iotool.FileHelper, paths []string,
					filter func(reader io.Reader) io.Reader, worker func(chan *iotool.NamedBuffer)) {
	work := NewWorkManual(1); handlers := make(chan io.Reader, work.SuggestBufferSize(0))
	ihandlers := make(chan io.Reader, 0); if !helper.ShouldFileAdviceWillNeed() { helper.ToggleFileAdviceWillNeed() }

	// in one thread open files. this allows parallel open syscalls.
	work.Feed(func() {
		for _, path := range paths {
			handler, e := iotool.Open(helper, path)
			if e == nil { handlers <-handler }
		}; close(handlers)
	})

	// in another receive the handlers and send them one by one to the work function.
	// since channels are fifo, the sequence stays the same as in input paths.
	work.Feed(func() {
		for handler := range handlers {
			ihandlers <-handler
		}; close(ihandlers)
	})

	// and now read it
	work.Run(func() {
		for handler := range ihandlers {
			if filter != nil {
				handler = filter(handler)
			}

			readFileSequential(helper, handler, worker)
		}
	})
}

// read one file sequential
func readFileSequential(helper *iotool.FileHelper, reader io.Reader, innerWorker func(chan *iotool.NamedBuffer)) {
	work := NewWorkManual(0); buffers := make(chan *iotool.NamedBuffer, work.SuggestFileBufferSize())

	// read in one thread
	work.Feed(func() {
		iotool.ReadIntoBuffer(helper, reader, buffers)
		close(buffers)
	})

	// do work in another
	work.Run(func() {
		innerWorker(buffers)
	})
}
