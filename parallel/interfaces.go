package parallel

import(
	"time"
)

type WorkInterface interface {
	Initialise(aworkers uint)

	Run(worker func())
	Do(worker func())
	Start(worker func())
	Wait()

	Feed(feeder func())

	Tick(duration time.Duration, tick func()) chan bool
	Timeout(duration time.Duration, cancel func())

	RecoverSendOnClosedChannel()

	SuggestBufferSize(max uint) uint
	SuggestFileBufferSize() uint
}
