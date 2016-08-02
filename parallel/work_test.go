package parallel

import(
	"testing"

	"math/rand"
	"sync/atomic"
	"time"
)

// Test Do(); counting the number of goroutines started; Do works with a waitGroup.
func TestDo(t *testing.T) {
	count, numberOfWorkers, finallyCalled := uint(0), SuggestNumberOfWorkers(), false

	// the given function is called after all workers finished working.
	work := NewWorkFinally(numberOfWorkers, func() {
		finallyCalled = true
	})

	// the actual work is done in the function given to Do(); Do() blocks till all
	// workers finish.
	work.Do(func() {
		// locks are required, else additions can be lost. with simple math
		// sync.atomic.AddUint() and such are a lock free alternative.
		work.Lock()
		count++
		work.Unlock()
	})

	if count != numberOfWorkers {
		t.Errorf("count: got %v, expected %v", count, numberOfWorkers)
	}

	if !finallyCalled {
		t.Error("finally wasn't called!")
	}
}

// Test feeding a group of workers with an integer list to sum it up.
func TestStartFeed(t *testing.T) {
	count, numberOfWorkers := 0, SuggestNumberOfWorkers()
	countList, sum := []int{ 1, 5, 15, 0 }, 21
	// the communication between goroutines is done via this channel.
	// make the channel buffered so the feeder does not have to block until
	// a value is read. the buffer is made double the number of workers so the feeder
	// can more probably write a value while the workers are processing earlier data
	from := make(chan int, numberOfWorkers * 2)

	work := NewWork(numberOfWorkers)

	// start feeding the workers; Feed() does not block. 
	work.Feed(func() {
		for _, add := range countList {
			// this channel send will block this goroutine, when the channel is full
			from <-add
		}; close(from) // closing indicates the end of input
	})

	// the actual work is done in the function given to Start(); Start() does not block.
	work.Start(func() {
	// an infinity loop by design. its break condition is determined below
	infinite:
		for {
			// read from channel to get work to do
			select {
			// add is the work, ok is false when the channel is closed.
			// if the channel is closed, this means no more work to be done,
			// so break the loop to end the goroutine gracefully.
			case add, ok := <-from:
				if !ok { break infinite }
				// locks are required, else additions can be lost. with simple math
				// sync.atomic.AddUint() and such are a lock free alternative.
				work.Lock()
				count += add
				work.Unlock()
			}
		}
	})

	// wait until all goroutines are finished
	work.Wait()

	if count != sum {
		t.Errorf("count: got %v, expected %v", count, sum)
	}
}

// try the tick work. very slow.
// TODO test the tick
func TestTick(t *testing.T) {
	count, iterations, numberOfWorkers := uint32(0), uint64(0), SuggestNumberOfWorkers()
	timeout := 1 * time.Second
	// the communication between goroutines is done via this channel.
	// make the channel buffered so the feeder does not have to block until
	// a value is read. the buffer is made double the number of workers so the feeder
	// can more probably write a value while the workers are processing earlier data
	from := make(chan uint32, numberOfWorkers * 2)

	// *sigh* at least get some randomness... whoever figured this should be totally
	// deterministic should be shot.
	rand.Seed(time.Now().UnixNano())

	// the worker
	work := NewWork(numberOfWorkers)

	// start feeding the workers. the channel send will block, when the channel is full
	ticker := work.Tick(10 * time.Millisecond, func() {
		from <-rand.Uint32()
	})

	// do a timeout based cancel
	work.Timeout(timeout, func() {
		close(ticker)
		close(from)
	})

	// work.Run starts N workers and  blocks until the last started goroutine returns
	work.Run(func() {
	infinite:
		for {
			// read from channel to get work to do
			select {
			// add is the work, ok is false when the channel is closed.
			// if the channel is closed, this means no more work to be done,
			// so break the loop to end the goroutine gracefully.
			case add, ok := <-from:
				if !ok { break infinite }
				// since this is only adding and not doing something complex
				// two atomic calls can be used without locks.
				atomic.AddUint32(&count, add)
				atomic.AddUint64(&iterations, 1)
			}
		}
	})

	if count == 0 {
		t.Errorf("count: got %v, expected to be greater then 0 !", count)
	}

	t.Logf("iterations: %v, count: %v\n", iterations, count)
}

// test the run work. fast.
func TestRun(t *testing.T) {
	timeout := 1 * time.Second
	numberOfWorkers := SuggestNumberOfWorkers()

	testRun(t, numberOfWorkers, timeout)
}

// FIXME: fixme
func BenchmarkRun(b *testing.B) {
	timeout := 1 * time.Second
	numberOfWorkers := SuggestNumberOfWorkers()

	testRun(b, numberOfWorkers, timeout)
}

// internal testRun runs Work.Run
func testRun(t testing.TB, numberOfWorkers uint, timeout time.Duration) {
	count, iterations := uint32(0), uint64(0)
	// the communication between goroutines is done via this channel.
	// make the channel buffered so the feeder does not have to block unless it is full
	// until a value is read. the buffer is made double the number of workers so the feeder
	// can more probably write a value while the workers are processing earlier data
	from := make(chan uint32, numberOfWorkers * 2)

	// *sigh* at least get some randomness... whoever figured this should be totally
	// deterministic should be shot.
	rand.Seed(time.Now().UnixNano())

	// the worker
	work := NewWork(numberOfWorkers)

	// start feeding the workers. the channel send will block, when the channel is full
	// work.Feed has a recover function that'll catch writings on a closed channel.
	work.Feed(func() {
		for {
			from <-rand.Uint32()
		}
	})

	// do a timeout based cancel
	work.Timeout(timeout, func() {
		close(from)
	})

	// work.Run blocks until the last started goroutine returns
	work.Run(func() {
	infinite:
		for {
			// read from channel from if there is work to do
			select {
			// add is the work, ok is false when the channel is closed.
			// if the channel is closed, this means no more work to be done,
			// so break the loop to end the goroutine gracefully.
			case add, ok := <-from:
				if !ok { break infinite }
				// since this is only adding and not doing something complex
				// two atomic calls can be used without locks.
				atomic.AddUint32(&count, add)
				atomic.AddUint64(&iterations, 1)
			}
		}
	})

	if count == 0 {
		t.Errorf("count: got %v, expected to be greater then 0 !", count)
	}

	t.Logf("iterations: %v, count: %v\n", iterations, count)
}

/*
// do panic!
// FIXME this doesn't work. see the next FIXME's
func TestPanic(t *testing.T) {
	// if anything gets here, BOOM!
	defer func() {
		if recovered := recover(); recovered != nil {
			t.Fatalf("unexpected: recovered panic: %v", recovered)
		}
	}()

	shoot := make(chan bool, 1)

	// a plain, simple and clear message is all that's required!
	panicMessage := "Listen to them, the children of the night. What music they make!"
	// the go message when sending on a closed channel
	closedChannel := "send on closed channel"

	// we wanna recover
	recoverFunc := func() {
		// FIXME: following line does nothing. it should fail the test.
		if recovered := recover(); recovered != nil {
			if e, ok := recovered.(string); ok {
				if e == closedChannel {
					// give the panic to the internal panic handler. it should ignore it.
					panic(e)
				} else if e != panicMessage {
					t.Errorf("expected '%v', got '%v'", panicMessage, e)
				}
			} else if e, ok := recovered.(error); ok {
				if e.Error() == closedChannel {
					// give the panic to the internal panic handler. it should ignore it.
					panic(e)
				} else if e.Error() != panicMessage {
					t.Errorf("expected '%v', got '%v'", panicMessage, e)
				}
			} else {
				t.Errorf("unexpected panic! '%v' \nthis calls for more panic!", recovered)
				panic(recovered)
			}
		}
	}

	// the worker
	work := NewWork(SuggestNumberOfWorkers())

	// feed panic
	work.Feed(func() {
		defer recoverFunc()

		panic(closedChannel)
	})

	// feed PaniC
	work.Feed(func() {
		defer recoverFunc()

		panic(panicMessage)
	})

	// feed PANIC
	work.Feed(func() {
		defer recoverFunc()

		close(shoot)

		panic(errors.New(panicMessage))
	})

	work.Do(func() {
		<-shoot
	})
}*/

// JUST FOR TESTING, BAD EXAMPLE!
// same as TestDo, but with busy waiting!
func TestDo2(t *testing.T) {
	count, numberOfWorkers := uint(0), SuggestNumberOfWorkers()

	// the worker
	work := NewWork(numberOfWorkers)

	// the actual work is done in the function given to Do()
	work.Do(func() {
		// locks are required, else additions can be lost
		work.Lock()
		count++
		work.Unlock()
	})

	// this is busy waiting and it's not good as it uses up lot's of cpu time
infinite:
	for {
		// a select with a default: checks all channels and if none
		// has any data, the default is called immidiatly without blocking. 
		select {
		case <-time.After(2 * time.Second):
			t.Errorf("timeout! count: %v, expected %v", count, numberOfWorkers)
		default:
			if count == numberOfWorkers {
				break infinite
			}
		}
	}

}
