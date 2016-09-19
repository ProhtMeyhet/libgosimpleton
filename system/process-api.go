package system

import(
	"fmt"
	"strconv"

	"github.com/ProhtMeyhet/libgosimpleton/iotool"
)

// find processes by your custom order. iterates over every process, your function
// must give back true if found and false if not. if true, processes will contain
// this ProcessInfo. do not, i repeat, do not retain the *ProcessInfo - it is reused.
func FindBy(filter func(*ProcessInfo) bool) (processes []*ProcessInfo) {
	process := &ProcessInfo{}
	for process.findBy(filter) {
		processes = append(processes, process.MakeCopy())
	}; return
}

// walk over every process and do your thing
func WalkProcesses(stick func(*ProcessInfo)) {
	process := &ProcessInfo{}
	for process.findBy(func(process *ProcessInfo) bool {
		stick(process.MakeCopy()); return false
	}) {}
}

// find a process by pid
func FindProcess(aid uint64) (process *ProcessInfo, e error) {
	process = &ProcessInfo{ id: aid }
	return process, process.findById()
}

// find a process by pid given as string
func FindProcessByStringId(aid string) (process *ProcessInfo, e error) {
	pid, e := strconv.ParseUint(aid, 10, 0)
	if e != nil { return }
	return FindProcess(pid)
}

// find processes by name
func FindProcessesByName(aname string) (processes []*ProcessInfo) {
	process := &ProcessInfo{ search: aname }
	for process.findByName() {
		processes = append(processes, process.MakeCopy())
	}; return
}

// today is the oldest you've ever been ...
func FindOldestProcessByName(aname string) (process *ProcessInfo) {
	list := FindProcessesByName(aname); min := uint64(0)
	for _, listProcess := range list {
		if min == 0 || min >= listProcess.relativeStartTime {
			min = listProcess.relativeStartTime
			process = listProcess
		}
	}; return
}

// ... and the youngest you'll ever be again
func FindYoungestProcessByName(aname string) (process *ProcessInfo) {
	list := FindProcessesByName(aname); max := uint64(0)
	for _, listProcess := range list {
		if listProcess.relativeStartTime >= max {
			max = listProcess.relativeStartTime
			process = listProcess
		}
	}; return
}

// read from /proc/self/
func Self() (process *ProcessInfo) {
	process = &ProcessInfo{}

	handler, _ := iotool.Open(iotool.ReadOnly(), fmt.Sprintf(PROC_STAT_FILE, "self"))
	process.scanStat(handler); return
}

/***** current user *****/

// find all current users processes
func FindMyProcesses() (processes []*ProcessInfo) {
	process := &ProcessInfo{}
	for process.findByCurrentUser() {
		processes = append(processes, process.MakeCopy())
	}; return
}

