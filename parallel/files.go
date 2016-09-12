package parallel

import(
	"io"

	"github.com/ProhtMeyhet/libgosimpleton/iotool"
	"github.com/ProhtMeyhet/libgosimpleton/simpleton"
)

// work := OpenFileDoWork(helper, path, func(buffered *NamedBuffer) {
//	into := make([]byte, 512)
//	for {
//		println("reading " + buffered.Name())
//		read, e := buffered.Read(into); if e != nil { /* handle errors and EOF */ }
// 		/* do work */
//	}
// })
//
// work.Wait()

// Open File and Do Work
func OpenFileDoWork(helper *iotool.FileHelper, path string, worker func(*iotool.NamedBuffer)) (work WorkInterface) {
	return openFileDoWork(helper, path, 0, worker)
}

func openFileDoWork(helper *iotool.FileHelper, path string, key uint, worker func(*iotool.NamedBuffer)) (work WorkInterface) {
	// use 1 inner worker, see below
	work = NewWorkManual(1); buffered := iotool.NewNamedBuffer(path); buffered.SetKey(key)

	// read in one thread
	work.Feed(func() {
		iotool.ReadFileIntoBuffer(helper, path, buffered)
	})

	// work in another
	work.Start(func() {
		worker(buffered)
	})

	return
}

// work := OpenFilesDoWork(helper, paths, func(buffered chan *iotool.NamedBuffer) {
//	/* create destination */
//	helper := iotool.WriteOnly().ToggleCreate()
//	myVeryOwnDestination, e := iotool.Open(helper, buffered.Name() + ".part")
//	io.Copy(myVeryOwnDestination, buffered)
// })
//
// work.Wait()

// Open N Files and Do Work
func OpenFilesDoWork(helper *iotool.FileHelper, paths <-chan string, worker func(*iotool.NamedBuffer)) (work WorkInterface) {
	work = NewWorkManual(2)

	// feed is chan

	work.Start(func() {
		for path := range paths {
			OpenFileDoWork(helper, path, worker).Wait()
		}
	})

	return
}

// hash files
//
// OpenFilesFromListDoWork(helper, func(buffered *iotool.NamedBuffer) {
//	hasher := sha256.New()
//	io.Copy(hasher, buffered)
//	println(hex.EncodeToString(hasher.Sum(nil)))	
// }, path1, path2, paths...).Wait()

// Open N Files and Do Work
func OpenFilesFromListDoWork(helper *iotool.FileHelper, worker func(*iotool.NamedBuffer), paths ...string) (work WorkInterface) {
	return OpenFilesDoWork(helper, simpleton.StringListToChannel(paths...), worker)
}

// hash files
//
// ReadFilesSequential(helper, path, func(buffers *iotool.NamedBuffer) {
//	hasher := sha256.New()
//	io.Copy(hasher, buffered)
//	println(hex.EncodeToString(hasher.Sum(nil)))	
// }).Wait()

// read files sequential to a buffer
func ReadFilesSequential(helper *iotool.FileHelper, paths []string, worker func(*iotool.NamedBuffer)) (work WorkInterface) {
	return ReadFilesFilteredSequential(helper, paths, nil, worker)
}

// print 10 lines to Stdout
//
// ReadFilesFilteredSequential(helper, paths, func(reader io.Reader) {
//	return iotool.LimitLinesReader(reader, 10)
// }, func(buffered *iotool.NamedBuffer) {
//	io.Copy(os.Stdout, buffered)
// })

// read files sequential to a buffer.
// uses file advice will need and thus relies on kernel page caching. this is good
// for being a team player, but may hurt performance especially in syntethic tests.
func ReadFilesFilteredSequential(helper *iotool.FileHelper, paths []string,
		filter func(reader io.Reader) io.Reader,
		worker func(*iotool.NamedBuffer)) (work WorkInterface) {

	work = NewWorkManual(1); handlers := make(chan iotool.FileInterface, work.SuggestBufferSize(0))
	ihandlers := make(chan iotool.FileInterface, 0)
	if !helper.ShouldFileAdviceWillNeed() { helper.ToggleFileAdviceWillNeed() }

	// in one thread open files. this allows parallel open syscalls (and file advice).
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
	work.Start(func() {
		for handler := range ihandlers {
			buffered := iotool.NewNamedBuffer(handler.Name())
			filtered, _ := handler.(io.Reader); if filter != nil { filtered = filter(handler) }
			go iotool.ReadIntoBuffer(helper, filtered, buffered)
			worker(buffered)
		}
	})

	return
}
