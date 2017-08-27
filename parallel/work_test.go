package parallel

import(
	"testing"

	"errors"
	"math/rand"
	"runtime"
	"sync/atomic"
	"time"
)

// Test Do(); counting the number of goroutines started; Do works with a waitGroup.
func TestDo(t *testing.T) {
	count, finallyCalled := uint(0), false

	// for test sake, give the number of workers manually. it is suggested to not do so.
	// the given function is called after all workers finished working.
	work := NewWorkFinally(func() {
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

	if count != work.Workers() {
		t.Errorf("count: got %v, expected %v", count, work.Workers())
	}

	if !finallyCalled {
		t.Error("finally wasn't called!")
	}
}

// Test feeding a group of workers with an integer list to sum it up.
func TestStartFeed(t *testing.T) {
	count, countList, sum := 0, []int{ 1, 5, 15, 0 }, 21
	work := NewWork()
	// the communication between goroutines is done via this channel.
	// make the channel buffered so the feeder does not have to block until
	// a value is read. the buffer is made double the number of workers so the feeder
	// can more probably write a value while the workers are processing earlier data
	from := make(chan int, work.SuggestBufferSize(1))

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
	work := NewWorkManual(numberOfWorkers)

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

// test nesting work.Start() [work.Do() cannot be nested!]
func TestNested(t *testing.T) {
	count, finallyCalled := uint(0), false
	work := NewWorkFinally(func() {
		finallyCalled = true
	})

	numberOfWorkers := work.Workers()

	// nest 3 Start()s
	// start 4 workers which each start 4 workers which each start 4 workers
	work.Start(func() {
		work.Start(func() {
			work.Start(func() {
				work.Lock()
				count++
				work.Unlock()
			})
		})
	}).Wait(); t.Logf("count %v", count)

	if count != numberOfWorkers << 4 {
		t.Errorf("count: got %v, expected %v", count, numberOfWorkers << 4)
	}

	if !finallyCalled {
		t.Error("finally wasn't called!")
	}

	// reset
	count = 0; finallyCalled = false

	// start 4 workers which each start 4 workers which each start 4 workers
	work.Start(func() {
		work.Start(func() {
			work.Start(func() {
				work.Lock()
				count++
				work.Unlock()
			})
		})
	})

	// work.Do() cannot be nested because it calls work.Wait()!
	work.Do(func() {
		work.Start(func() {
			work.Start(func() {
				work.Lock()
				count++
				work.Unlock()
			})
		})
	}); t.Logf("count2 %v", count)

	if count != numberOfWorkers << 5 {
		t.Errorf("count: got %v, expected %v", count, numberOfWorkers << 5)
	}

	if !finallyCalled {
		t.Error("finally wasn't called!")
	}
}

// test the run work. fast.
func TestRun(t *testing.T) {
	timeout := 1 * time.Second; numberOfWorkers := SuggestNumberOfWorkers()

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
	work := NewWorkManual(numberOfWorkers)

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

func TestRecoverClosedChannel(t *testing.T) {
	testRecoverClosedChannel(t, "send on closed channel", false)
	testRecoverClosedChannel(t, "close on closed channel", false)
	testRecoverClosedChannel(t, "See me, feel me, touch me, heal me", true)
	testRecoverClosedChannelWithError(t, errors.New("send on closed channel"), false)
	testRecoverClosedChannelWithError(t, errors.New("close on closed channel"), false)
	testRecoverClosedChannelWithError(t, errors.New("See me, feel me, touch me, heal me"), true)
}

func testRecoverClosedChannel(t *testing.T, message string, shouldInternalRecover bool) {
	defer func() {
		recovered := recover(); if recovered == nil { return }
		if e, ok := recovered.(string); ok {
			if e != message && !shouldInternalRecover {
				t.Errorf("unexpected panic! '%T' -> '%v'", e, e)
			}
		} else {
			t.Errorf("unexpected panic2! '%T' -> '%v'", recovered, recovered)
		}
	}()

	defer RecoverClosedChannel()

	// panic(errors.New("error"))
	// panic("error")
	panic(message)
}

func testRecoverClosedChannelWithError(t *testing.T, message error, shouldInternalRecover bool) {
	defer func() {
		recovered := recover(); if recovered == nil { return }
		if e, ok := recovered.(error); ok {
			if e != message && !shouldInternalRecover {
				t.Errorf("unexpected panic! '%T' -> '%v'", e, e)
			}
		} else {
			t.Errorf("unexpected panic2! '%T' -> '%v'", recovered, recovered)
		}
	}()

	defer RecoverClosedChannel()

	// panic(errors.New("error"))
	// panic("error")
	panic(message)
}

func TestMeta(t *testing.T) {
	workers := uint(runtime.NumCPU())
	work := NewWorkManual(workers)

	if workers != work.Workers() {
		t.Errorf("work.Workers(): expected '%v', got '%v'", workers, work.Workers())
	}

	max := workers * NUMBER_OF_WORKERS_MULTIPLIER; bufferSize := SuggestNumberOfWorkers()
	if max != bufferSize {
		t.Errorf("work.SuggestBufferSize() expected '%v', got '%v'", max, bufferSize)
	}

	max = uint(0); bufferSize = SuggestMaximumNumberOfWorkers(max); expected := workers * NUMBER_OF_WORKERS_MULTIPLIER
	if bufferSize == uint(0) {
		t.Error("work.SuggestBufferSize() is 0!")
	} else if bufferSize != expected {
		t.Errorf("work.SuggestNumberOfWorkers expected '%v', got '%v'", expected, bufferSize)
	}

	max = uint(20); bufferSize = SuggestMaximumNumberOfWorkers(max); expected = workers * NUMBER_OF_WORKERS_MULTIPLIER
	if max == bufferSize {
		t.Errorf("work.SuggestNumberOfWorkers expected '%v', got '%v'", expected, bufferSize)
	}

	max = uint(12); bufferSize = work.SuggestBufferSize(max)
	if bufferSize != workers * BUFFER_SIZE_MULTIPLIER {
		t.Errorf("work.SuggestBufferSize() expected '%v', got '%v'", max, bufferSize)
	}

	max = uint(0); bufferSize = work.SuggestBufferSize(max)
	if bufferSize == uint(0) {
		t.Error("work.SuggestBufferSize() is 0!")
	}

	max = uint(20); bufferSize = work.SuggestBufferSize(max); expected = workers * BUFFER_SIZE_MULTIPLIER
	if max == bufferSize {
		t.Errorf("work.SuggestBufferSize() expected '%v', got '%v'", expected , bufferSize)
	}
}

// JUST FOR TESTING, BAD EXAMPLE!
// same as TestDo, but with busy waiting!
func TestDo2(t *testing.T) {
	count, numberOfWorkers := uint(0), SuggestNumberOfWorkers()

	// the worker
	work := NewWorkManual(numberOfWorkers)

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
