package processes

import(
	"fmt"
	"strings"
	"strconv"

	"github.com/ProhtMeyhet/libgosimpleton/iotool"
	"github.com/ProhtMeyhet/libgosimpleton/parallel"
	"github.com/ProhtMeyhet/libgosimpleton/system/user"
)

// find processes by your custom order. iterates over every process, your function
// must give back true if found and false if not. if true, processes will contain
// this ProcessInfo. do not, i repeat, do not retain the *ProcessInfo - it is reused.
func FindBy(filter func(*ProcessInfo) bool) <-chan *ProcessInfo {
	process := &ProcessInfo{}
	processes := make(chan *ProcessInfo, parallel.SuggestBufferSize(0))
	go func() {
		for process.findBy(filter) {
			processes <-process.MakeCopy()
		}; close(processes)
	}(); return processes
}

func FindAll() chan *ProcessInfo {
	process := &ProcessInfo{}
	processes := make(chan *ProcessInfo, parallel.SuggestBufferSize(0))
	go func() {
		for process.findBy(func(*ProcessInfo) bool { return true } ) {
			processes <-process.MakeCopy()
		}; close(processes)
	}(); return processes
}

/* TODO
func FindFirstBy(filter func(*ProcessInfo) bool) (process *ProcessInfo) {
	process := &ProcessInfo{}
	process.findBy(filter)
	return
}

func FindLastBy(filter func(*ProcessInfo) bool) (process *ProcessInfo) {
	find := &ProcessInfo{}
	for find.findBy(filter) {
		process = find
	}; return
}*/

// find by a generating func
func FindByGenerator(generator func() func(*ProcessInfo) bool) <-chan *ProcessInfo {
	return FindBy(generator())
}

// walk over every process and do your thing
func Walk(stick func(*ProcessInfo)) {
	process := &ProcessInfo{}
	for process.findBy(func(process *ProcessInfo) bool {
		stick(process.MakeCopy()); return false
	}) {}
}

// walk by generating func
func WalkByGenerator(generator func() func(*ProcessInfo)) {
	Walk(generator())
}

// find a process by pid
func Find(aid uint) (process *ProcessInfo, e error) {
	process = &ProcessInfo{ id: aid }
	return process, process.findById()
}

// find a process by pid given as string
func FindByStringId(aid string) (process *ProcessInfo, e error) {
	pid, e := strconv.ParseUint(aid, 10, 0); if e != nil { return }
	return Find(uint(pid))
}

// find processes by name
func FindByName(aname string) <-chan *ProcessInfo {
	return FindBy(func(process *ProcessInfo) bool {
		return Contains(process, aname)
	})
}

// find processes by their exact name
func FindByExactName(aname string) <-chan *ProcessInfo {
	return FindBy(func(process *ProcessInfo) bool {
		return Exact(process, aname)
	})
}

// today is the oldest you've ever been ...
func FindOldestByName(aname string) (oldest *ProcessInfo) {
	min := uint64(0)
	Walk(func(process *ProcessInfo) {
		if !Contains(process, aname) { return }
		if min == 0 || min >= process.relativeStartTime {
			min = process.relativeStartTime
			oldest = process
		}
	}); return
}

// ... and the youngest you'll ever be again
func FindYoungestByName(aname string) (youngest *ProcessInfo) {
	max := uint64(0)
	Walk(func(process *ProcessInfo) {
		if !Contains(process, aname) { return }
		if process.relativeStartTime >= max {
			max = process.relativeStartTime
			youngest = process
		}
	}); return
}

// read from /proc/self/
func Self() (process *ProcessInfo) {
	process = &ProcessInfo{}
	handler, _ := iotool.Open(iotool.ReadOnly(), fmt.Sprintf(PROC_STAT_FILE, "self"))
	process.scanStat(handler); return
}

/***** current user *****/

// find all current users processes
// panics if current user can't be determined
func FindMyAll() <-chan *ProcessInfo {
	user, e := user.Current(); if e != nil { panic(e) }
	return FindBy(func(process *ProcessInfo) bool {
		return User(process, user)
	})
}

// find a process by pid
// panics if current user can't be determined
func FindMy(aid uint) (process *ProcessInfo) {
	user, e := user.Current(); if e != nil { panic(e) }
	process, e = Find(aid); if e != nil { return }
	if !User(process, user) { return }
	return
}

// find processes by name
// panics if current user can't be determined
func FindMyByName(aname string) <-chan *ProcessInfo {
	user, e := user.Current(); if e != nil { panic(e) }
	return FindBy(func(process *ProcessInfo) bool {
		return User(process, user) && Contains(process, aname)
	})
}

// find processes by their exact name
// panics if current user can't be determined
func FindMyByExactName(aname string) <-chan *ProcessInfo {
	user, e := user.Current(); if e != nil { panic(e) }
	return FindBy(func(process *ProcessInfo) bool {
		return User(process, user) && Exact(process, aname)
	})
}

/***** filters *****/

// contains
func Contains(process *ProcessInfo, aname string) bool {
	return strings.Contains(process.name, aname)
}

// exact name
func Exact(process *ProcessInfo, aname string) bool {
	return process.name == aname
}

// user id
func User(process *ProcessInfo, user user.UserInterface) bool {
	return uint32(process.owner) == user.Id()
}
