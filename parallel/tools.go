package parallel

import(
	"runtime"
)

func SETMAXPROCS() {
	// FIXME
	if runtime.GOMAXPROCS(0) < runtime.NumCPU() * NUMBER_OF_WORKERS_MULTIPLIER {
		runtime.GOMAXPROCS(runtime.NumCPU() * NUMBER_OF_WORKERS_MULTIPLIER)
	}
}
