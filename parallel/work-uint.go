package parallel

import(

)

type WorkUint64 struct {
	Work

	Talk	chan uint64
}

func NewUints64Feeder(list ...uint64) (work *WorkUint64) {
	work = &WorkUint64{}
	work.Initialise(SuggestNumberOfWorkers(uint(len(list))))
	work.Talk = make(chan uint64, work.SuggestBufferSize(uint(len(list))))

	work.Feed(func() {
		for _, s := range list {
			work.Talk <-s
		}; close(work.Talk)
	})

	return
}
