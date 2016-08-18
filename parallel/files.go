package parallel

import(

	"github.com/ProhtMeyhet/libgosimpleton/iotool"
	"github.com/ProhtMeyhet/libgosimpleton/simpleton"
)

// Open File and Do Work
//
// OpenFileDoWork(helper, path, func(buffers chan NamedBuffer) {
// 	for buffered := range buffers {
//		if buffered.Done() {
//			fmt.Println("done!")
//			continue
//		}
// 		/* do work */
// 	}
// })
func OpenFileDoWork(helper *iotool.FileHelper, path string, worker func(chan iotool.NamedBuffer)) {
	// use 1 inner worker, see below
	work := NewWorkManual(1); buffers := make(chan iotool.NamedBuffer, work.SuggestFileBufferSize())

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

// Open N Files and Do Work
//
// OpenFilesDoWork(helper, paths, func(buffers chan NamedBuffer) {
// 	for buffered := range buffers {
//		if buffered.Done() {
//			fmt.Println("done!")
//			continue
//		}
// 		/* do work */
// 	}
// })
func OpenFilesDoWork(helper *iotool.FileHelper, paths <-chan string, worker func(chan iotool.NamedBuffer)) {
	work := NewWork(0)

	// feed is list

	work.Start(func() {
		for path := range paths {
			OpenFileDoWork(helper, path, worker)
		}
	})

	work.Wait()
}

// Open N Files and Do Work
//
// OpenFilesFromListDoWork(helper, func(buffers chan NamedBuffer) {
// 	for buffered := range buffers {
//		if buffered.Done() {
//			fmt.Println("done!")
//			continue
//		}
// 		/* do work */
// 	}
// }, path1, path2, paths...)
func OpenFilesFromListDoWork(helper *iotool.FileHelper, worker func(chan iotool.NamedBuffer), paths ...string) {
	OpenFilesDoWork(helper, simpleton.StringListToChannel(paths...), worker)
}
