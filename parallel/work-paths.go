package parallel

/*
import(


)

type WorkPaths struct {
	Work

	Talk	chan string
}

func NewFileInfoFeeder(helper *iotool.FileHelper, list ...string) (work *WorkPaths) {
	strings := NewStringsFeeder(list)

	work = &WorkPaths{}
	work.Initialise(SuggestNumberOfWorkers())
	work.Talk = make(chan iotool.FileInfoInterface, work.SuggestBufferSize(uint(len(list))))

	strings.Start(func() {
		for _, s := range strings.Talk {
			
			work.Talk <-s
		}; close(work.Talk)
	})

	return
}*/
