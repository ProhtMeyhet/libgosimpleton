package parallel

import(

)

type WorkUint struct {
	Work

	Talk	chan uint
}

func NewUintsFeeder(list ...uint) (work *WorkUint) {
	work = &WorkUint{}
	work.Initialise(SuggestNumberOfWorkers(uint(len(list))))
	work.Talk = make(chan uint, work.SuggestBufferSize(uint(len(list))))

	work.Feed(func() {
		for _, s := range list {
			work.Talk <-s
		}; close(work.Talk)
	})

	return
}
