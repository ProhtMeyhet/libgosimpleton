package parallel

import(

)

type WorkInt struct {
	Work

	Talk	chan int
}

func NewIntsFeeder(list ...int) (work *WorkInt) {
	work = &WorkInt{}
	work.Initialise(SuggestMaximumNumberOfWorkers(uint(len(list))))
	work.Talk = make(chan int, work.SuggestBufferSize(uint(len(list))))

	work.Feed(func() {
		for _, s := range list {
			work.Talk <-s
		}; close(work.Talk)
	})

	return
}
