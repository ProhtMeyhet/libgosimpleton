package parallel

import(
	"runtime"

	"sync"
	"time"
)

/*
 * Work has locks and can utilize wait groups.
 * Locks are a feature for the working in the goroutines. Internally
 * locking is only used for Wait() to be thread safe.
 * Wait groups are used with Start() and it's wrapper Do() to ensure
 * all work is done. With Run() wait groups are not used and a call
 * to Wait() will return immidiatly. This however is not enforced to 
 * make it possible to use Start() and Run() at the same time for
 * different purposes.
 * 
 * See work_test.go for examples.
*/
type Work struct {
	sync.Mutex

	// worker of numbers
	workers uint

	// wait!
	waitGroup	sync.WaitGroup

	// final function. hint: use it to close channels
	finallyFunc	func()

	// ensure waitGroup.Wait() only gets called once.
	waitOnce	sync.Once

	// required to reset work.waitOnce
	reset		sync.Once

	// if 0, don't start anything. used by feeders
//	lenghtHint	uint
}

// New York; determine number of workers by number of CPU or amaxworkers, whichever smaller
func NewWork() (work *Work) {
	return NewWorkManual(SuggestNumberOfWorkers())
}

// manually set number of workers
func NewWorkManual(aworkers uint) (work *Work) {
	work = &Work{}
	work.Initialise(aworkers)
	return
}

// New York Finally
func NewWorkFinally(afinally func()) (work *Work) {
	return NewWorkFinallyManual(SuggestNumberOfWorkers(), afinally)
}

// New York Finally with manual setting of workers
func NewWorkFinallyManual(aworkers uint, afinally func()) (work *Work) {
	work = &Work{ finallyFunc: afinally }
	work.Initialise(aworkers)
	return
}

// i n i t i a l i z e
func (work *Work) Initialise(aworkers uint) {
	work.workers = aworkers
	if runtime.GOMAXPROCS(0) <= runtime.NumCPU() {
		runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	}
}

// start workers - 1 in separate goroutines and run one in this goroutine.
// intended for (web)servers
func (work *Work) Run(worker func()) {
	if work.workers == 0 { panic("work.workers is 0!") }
	if work.workers > 1 {
		for i := uint(0); i < work.workers -1; i++ {
			go worker()
		}
	}

	worker()
}

// do some work; this is equivalent to calling Start() and then Wait().
// calls finally func afterwards. guarantees all work is done. it blocks until
// all workers finish. should not be called with the go command.
func (work *Work) Do(worker func()) {
	work.start(worker)
	work.Wait()
}

// start workers; a worker should read from a channel and from that do work
// does not call finally as it doesn't record when or if a goroutine finishes.
// can guarantee all work is done, if work.Wait is called. does not block.
func (work *Work) Start(worker func()) WorkInterface {
	work.waitGroup.Add(1)
	go func() {
		work.start(worker)
		work.waitGroup.Done()
	}()
	return work
}

// go one worker with waitgroup
func (work *Work) Go(worker func()) WorkInterface {
	work.waitGroup.Add(1)
	go func() {
		worker()
		work.waitGroup.Done()
	}()
	return work
}

// start a bunch of workers with waitGroup
func (work *Work) start(worker func()) {
	if work.workers == 0 { panic("work.workers is 0!") }
	for i := uint(0); i < work.workers; i++ {
		work.Go(worker)
	}
}

// feed workers; yummy; loop over something and push each value on a channel.
// it is good practise to first start feeding and then to open the workers
// via Do(), Start() or Run().
func (work *Work) Feed(feeder func()) WorkInterface {
	go func() {
		defer RecoverClosedChannel()
		feeder()
	}()
	return work
}

// feed workers by ticks; the tick function must not loop indefinitly, but return as quickly as possible.
// returns a channel that must be closed or a boolean value send to, when not needed anymore
// otherwise the tick will run forever.
func (work *Work) Tick(duration time.Duration, tick func()) (cancel chan bool) {
	cancel = make(chan bool, 1)

	go func() {
		defer RecoverClosedChannel()
	infinite:
		for {
			select {
			case <-time.Tick(duration):
				tick()
			// if anything happens on the cancel channel, abort.
			case <-cancel:
				break infinite
			}
		}
	}()

	return
}

// recover the panic 'send on closed channel' and ignore it. otherwise panic some more.
func RecoverClosedChannel() {
	recovered := recover(); if recovered == nil { return }

	message := ""
	if e, ok := recovered.(error); ok {
		message = e.Error()
	} else if e, ok := recovered.(string); ok {
		message = e
	}

	if message != "send on closed channel" && message != "close of closed channel" {
		// panic some more
		panic(recovered)
	}
}

// a wrapper around the internally used *sync.WaitGroup.Wait(); thread safe
func (work *Work) Wait() {
	work.waitOnce.Do(func() {
		work.waitGroup.Wait()
		work.finally()
		// allow resetting of work.waitGroup.Wait()
		work.reset = sync.Once{}
	})

	// reset work.WaitOnce
	work.reset.Do(func() {
		// reset of work.waitGroup.Wait()
		work.waitOnce = sync.Once{}
	})
}

// TODO
func (work *Work) finally() {
	if work.finallyFunc != nil {
		work.finallyFunc()
	}
}

// cancel a work. this should be preferably done by closing a channel.
// alternativly use a second, bool typed, channel to send a cancel signal.
func (work *Work) Timeout(duration time.Duration, cancel func()) WorkInterface {
	go func() {
		select {
			case <-time.After(duration):
				cancel()
		}
	}()
	return work
}

/*
 * convinience functions
 */

// return number of configured workers
func (work *Work) Workers() uint {
	return work.workers
}

// usually it's a good idea to have 4 * more buffers then workers; please state a max
func (work *Work) SuggestBufferSize(_ uint) uint {
	// FIXME XXX TODO ignore uint argument for now
	// TODO remove uint argument from here and from interface
	return suggestBufferSize(work.workers * BUFFER_SIZE_MULTIPLIER, work.workers)
}

// suggest a minimum buffer size; please state a minimum
func (work *Work) SuggestMinimumBufferSize(min uint) uint {
	return suggestMinimumBufferSize(min, work.workers)
}

// suggest a size for filepaths to buffer
func (work *Work) SuggestFileBufferSize() uint {
	return SuggestFileBufferSize(work.workers)
}

// suggest a number of workers.
func SuggestNumberOfWorkers() uint {
	return uint(runtime.NumCPU()) * NUMBER_OF_WORKERS_MULTIPLIER
}

// suggest a number of workers.
func SuggestMaximumNumberOfWorkers(max uint) uint {
	if max > 0 && uint(runtime.NumCPU()) * NUMBER_OF_WORKERS_MULTIPLIER > max {
		return max
	}

	return uint(runtime.NumCPU()) * NUMBER_OF_WORKERS_MULTIPLIER
}

// usually it's a good idea to have 4 * more buffers then workers; please state a max
func SuggestBufferSize(max uint) uint {
	return suggestBufferSize(max, SuggestNumberOfWorkers())
}

func suggestBufferSize(max, workers uint) uint {
	if max > 0 && workers * BUFFER_SIZE_MULTIPLIER > max {
		return max
	}; return workers * BUFFER_SIZE_MULTIPLIER
}

// suggest a minimum buffer size; please state a minimum
func SuggestMinimumBufferSize(min uint) uint {
	return suggestMinimumBufferSize(min, SuggestNumberOfWorkers())
}

func suggestMinimumBufferSize(min, workers uint) uint {
	// allow zero for unbuffered channels
	if min < workers * BUFFER_SIZE_MULTIPLIER {
		return min
	}; return workers * BUFFER_SIZE_MULTIPLIER
}

// suggest a size for filepaths to buffer
func SuggestFileBufferSize(workers uint) uint {
	if workers == 0 {
		workers = SuggestNumberOfWorkers()
	}; return workers * FILE_BUFFER_SIZE_MULTIPLIER
}
