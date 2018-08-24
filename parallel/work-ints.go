package parallel

import(

)

type WorkInt struct {
	Work

	Talk	chan int
}

func NewIntsFeeder(list ...int) (work *WorkInt) {
	work = &WorkInt{}
	work.Initialise(SuggestNumberOfWorkers())
	work.Talk = make(chan int, work.SuggestBufferSize(uint(len(list))))

	work.Feed(func() {
		for _, s := range list {
			work.Talk <-s
		}; close(work.Talk)
	})

	return
}
