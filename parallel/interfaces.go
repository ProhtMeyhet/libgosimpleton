package parallel

import(
	"time"
)

type WorkInterface interface {
	LockInterface

	// wrapper around sync.WaitGroup.Wait(). must be thread safe (for example with sync.Once).
	// only works if Start() or it's wrapper Do() are used.
	WaitInterface

	// used to enable other packages to implement a custom worker
	Initialise(aworkers uint)

	// start workers - 1 in separate goroutines and run one
	// worker in this goroutine. the worker should block till the work is done.
	// it is not guaranteed that all work will be processed nor that the last
	// worker is really the last finishing (which can mean, that workers still
	// processing work can be interrupted and shut down by program termination).
	// intended for webservers, where the workers run in endless loops.
	Run(worker func())

	// wrapper around Start() which Waits() until all workers finish and then returns
	Do(worker func())

	// Start() N workers
	Start(worker func()) WorkInterface

	// start one feeding goroutine. can be used multiple times
	Feed(feeder func()) WorkInterface

	// suggest a buffer size, best according to number of CPUs
	// FIXME remove max uint argument
	SuggestBufferSize(max uint) uint

	// suggest a buffer size for filesystem i/o
	SuggestFileBufferSize() uint

//	SetFeedRecover(to func())
}

type WorkHelpInterface interface {
	// suggest a buffer size, best according to number of CPUs
	SuggestBufferSize() uint

	// suggest a buffer size for filesystem i/o
	SuggestFileBufferSize() uint
}

type TickWorkInterface interface {
	// use time.Tick to execute a function periodically
	Tick(duration time.Duration, tick func()) chan bool
}

type TimeoutWorkInterface interface {
	// after duration, call cancel function
	Timeout(duration time.Duration, cancel func()) WorkInterface
}

// backup copy
type LockInterface interface {
	Lock()
	Unlock()
}

// wait for something
type WaitInterface interface {
	Wait()
}

type DoneInterface interface {
	Done()
}

type WaitGroupInterface interface {
	WaitInterface
	DoneInterface
	Add(delta int)
}
