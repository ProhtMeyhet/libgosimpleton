package parallel

import(
	"time"
)

type WorkInterface interface {
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
	Start(worker func())

	// wrapper around sync.WaitGroup.Wait(). must be thread safe (for example with sync.Once).
	// only works if Start() or it's wrapper Do() are used.
	Wait()

	// start one feeding goroutine. can be used multiple times
	Feed(feeder func())

	// use time.Tick to execute a function periodically
	Tick(duration time.Duration, tick func()) chan bool

	// after duration, call cancel function
	Timeout(duration time.Duration, cancel func())

	// suggest a buffer size, best according to number of CPUs
	SuggestBufferSize(max uint) uint

	// suggest a buffer size for filesystem i/o
	SuggestFileBufferSize() uint
}
