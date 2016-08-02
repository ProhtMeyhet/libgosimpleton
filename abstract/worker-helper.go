package abstract

import(
	"runtime"
)

type WorkerHelper struct {
	// workers is the count of goroutines to start to work something
	workers			uint

	// how many worker buffers to be used. by default it's set to $workers * 4
	// may not be smaller then $workers
	workerBuffers		uint
}

// fresh and shiny
func NewWorkerHelper() (helper *WorkerHelper) {
	helper = &WorkerHelper{}
	helper.Initialise()
	return
}

// i n i t
func (helper *WorkerHelper) Initialise() {
	helper.SetWorkers(uint(runtime.GOMAXPROCS(0) * 2))
}

// workers
func (helper *WorkerHelper) Workers() uint {
	return helper.workers
}

// set workers
func (helper *WorkerHelper) SetWorkers(to uint) *WorkerHelper {
	// must be devideable by 2. minor error, fix it here
	if to % 2 != 0 { to++ }
	// ignore zero value
	if to == 0 { goto out }

	helper.workers = to
	helper.workerBuffers = to * 4

out:
	return helper
}

// get worker buffers
func (helper *WorkerHelper) WorkerBuffers() uint {
	return helper.workerBuffers
}

// set worker buffers
func (helper *WorkerHelper) SetWorkerBuffers(to uint) *WorkerHelper {
	// disallow setting worker buffers smaller then number of workers
	// however allow 0 as this makes the buffer channels unbuffered
	if to < helper.workers && to != 0 { goto out }

	helper.workerBuffers = to

out:
	return helper
}

func (helper *WorkerHelper) Copy(from interface{}) *WorkerHelper {
	if workerHelper, ok := from.(*WorkerHelper); !ok {
		panic("could not cast to *WorkerHelper !")
	} else {
		helper.workers = workerHelper.workers
		helper.workerBuffers = workerHelper.workerBuffers
	}

	return helper
}
