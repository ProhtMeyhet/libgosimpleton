package system

import(
	"bufio"
	"path/filepath"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ProhtMeyhet/libgosimpleton/iotool"
)

// process info from /proc/[pid]/stat
type ProcessInfo struct {
	id				uint64
	name				string
	stat				[]byte

	globList			[]string
	globKey				uint

	// TODO add a set function
	scanAll				bool

	// TODO accessors
	state				string
	parentProcessId			int32
	processGroupId			int32
	sessionId			int32
	tty				int32
	foregroundProcessId		int32
	kernelFlags			uint32
	childrenMinorFaults		uint32
	memoryMajorFaults		uint32
	childrenMajorFaults		uint32
	memoryMinorFaults		uint32
	rawUserTime			uint32
	rawScheduledTime		uint32
	rawChildrenUserWaitTime		int32
	rawChildrenKernelWaitTime	int32
	priority			int32
	nice				int32
	numberOfThreads			int32
	// itrealvalue
	relativeStartTime		uint64
	virtualMemory			uint32
	residentSetSize			int32
	softLimitResidentSetSize	uint32
	startCode			uint32
	endCode				uint32
	startStack			uint32
	espStackPointer			uint32
	eipInstructionPointer		uint32
	// signal
	// blocked
	// sigignore
	// sigcatch

	// TODO the rest from `man 5 proc`
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
	process := &ProcessInfo{ name: aname }
	for process.findByName() {
		processes = append(processes, process.MakeCopy())
	}; return
}

// read from /proc/self/
func Self() (process *ProcessInfo) {
	process = &ProcessInfo{}

	handler, _ := iotool.Open(iotool.ReadOnly(), fmt.Sprintf(PROC_STAT_FILE, "self"))
	process.scanStat(handler)

	return
}

// copy values from a *ProcessInfo
func (info *ProcessInfo) Copy(from *ProcessInfo) {
	info.id = from.id; info.name = from.name; info.stat = from.stat
}

// make a copy of this *ProcessInfo
func (info *ProcessInfo) MakeCopy() (process *ProcessInfo) {
	return info.MakeCopyFrom(info)
}

// copy values from a given *ProcessInfo to a new *ProcessInfo
func (info *ProcessInfo) MakeCopyFrom(process *ProcessInfo) (processCopy *ProcessInfo) {
	processCopy = &ProcessInfo{}
	processCopy.id = process.id; processCopy.name = process.name; processCopy.stat = process.stat
	return
}

// find process by id
func (info *ProcessInfo) findById() (e error) {
	if info.id == 0 { return INVALID_PROCESS_ID_ERROR }

	handler, e := iotool.Open(iotool.ReadOnly(), fmt.Sprintf(PROC_STAT_FILE, info.id))
	if e != nil && os.IsNotExist(e) { e = NO_SUCH_PROCESS_ERROR; return }; defer handler.Close()
	return info.scanStat(handler)
}

// find a process by name. first call gives first result, second call second result ...
func (info *ProcessInfo) findByName() bool {
	if len(info.globList) == 0 { var e error
		info.globList, e = filepath.Glob(PROC_GLOB); if e != nil { return false }
	}

	for ; info.globKey < uint(len(info.globList)); info.globKey++ {
		id := info.globList[info.globKey][len(PROC_PATH):]
		process, e := FindProcessByStringId(id)
		// process ended in between
		if e != nil { continue }
		if strings.HasPrefix(process.name, info.name) {
			info.globKey++; info.Copy(process); return true
		}
	}

	return false
}

// scan /proc/[pid]/stat
func (info *ProcessInfo) scanStat(handler iotool.FileInterface) (e error) {
	scanner := bufio.NewScanner(handler); scanner.Split(bufio.ScanWords)

	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		// pid
		if info.id == 0 {
			info.id, e = strconv.ParseUint(string(scanner.Bytes()), 10, 0)
			if e != nil { return }
		}
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		// name
		if info.name == "" {
			info.name = string(scanner.Bytes())[1:]
			info.name = info.name[:len(info.name)-1]
		}


// 3
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		info.state = string(scanner.Bytes())
// 4
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		toobig, _ := strconv.ParseInt(string(scanner.Bytes()), 10, 0)
		info.parentProcessId = int32(toobig)
// 5
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		toobig, _ = strconv.ParseInt(string(scanner.Bytes()), 10, 0)
		info.processGroupId = int32(toobig)
// 6
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		toobig, _ = strconv.ParseInt(string(scanner.Bytes()), 10, 0)
		info.sessionId = int32(toobig)
// 7
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		toobig, _ = strconv.ParseInt(string(scanner.Bytes()), 10, 0)
		info.tty = int32(toobig)
// 8
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		toobig, _ = strconv.ParseInt(string(scanner.Bytes()), 10, 0)
		info.foregroundProcessId = int32(toobig)




	if !info.scanAll { return }




// 9
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		utoobig, _ := strconv.ParseUint(string(scanner.Bytes()), 10, 0)
		info.kernelFlags = uint32(toobig)
// 10
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		utoobig, _ = strconv.ParseUint(string(scanner.Bytes()), 10, 0)
		info.memoryMinorFaults = uint32(toobig)
// 11
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		utoobig, _ = strconv.ParseUint(string(scanner.Bytes()), 10, 0)
		info.childrenMinorFaults = uint32(toobig)
// 12
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		utoobig, _ = strconv.ParseUint(string(scanner.Bytes()), 10, 0)
		info.memoryMajorFaults = uint32(toobig)
// 13
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		utoobig, _ = strconv.ParseUint(string(scanner.Bytes()), 10, 0)
		info.childrenMajorFaults = uint32(toobig)
// 14
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		utoobig, _ = strconv.ParseUint(string(scanner.Bytes()), 10, 0)
		info.rawUserTime = uint32(utoobig)
// 15
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		utoobig, _ = strconv.ParseUint(string(scanner.Bytes()), 10, 0)
		info.rawScheduledTime = uint32(utoobig)
// 16
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		toobig, _ = strconv.ParseInt(string(scanner.Bytes()), 10, 0)
		info.rawChildrenUserWaitTime = int32(utoobig)
// 17
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		toobig, _ = strconv.ParseInt(string(scanner.Bytes()), 10, 0)
		info.rawChildrenKernelWaitTime = int32(toobig)
// 18
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		toobig, _ = strconv.ParseInt(string(scanner.Bytes()), 10, 0)
		info.priority = int32(toobig)
// 19
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		toobig, _ = strconv.ParseInt(string(scanner.Bytes()), 10, 0)
		info.nice = int32(toobig)
// 20
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		toobig, _ = strconv.ParseInt(string(scanner.Bytes()), 10, 0)
		info.numberOfThreads = int32(toobig)
// 21
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		// itrealvalue since kernel 2.6.17, this field is no longer maintained, and is hard coded as 0.
// 22
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		info.relativeStartTime, _ = strconv.ParseUint(string(scanner.Bytes()), 10, 0)
// 23
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		utoobig, _ = strconv.ParseUint(string(scanner.Bytes()), 10, 0)
		info.virtualMemory = uint32(toobig)
// 24
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		toobig, _ = strconv.ParseInt(string(scanner.Bytes()), 10, 0)
		info.residentSetSize = int32(toobig)
// 25
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		utoobig, _ = strconv.ParseUint(string(scanner.Bytes()), 10, 0)
		info.softLimitResidentSetSize = uint32(utoobig)
// 26
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		utoobig, _ = strconv.ParseUint(string(scanner.Bytes()), 10, 0)
		info.startCode = uint32(utoobig)
// 27
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		utoobig, _ = strconv.ParseUint(string(scanner.Bytes()), 10, 0)
		info.endCode = uint32(utoobig)
// 28
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		utoobig, _ = strconv.ParseUint(string(scanner.Bytes()), 10, 0)
		info.startStack = uint32(utoobig)
// 29
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		utoobig, _ = strconv.ParseUint(string(scanner.Bytes()), 10, 0)
		info.espStackPointer = uint32(utoobig)
// 30
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		utoobig, _ = strconv.ParseUint(string(scanner.Bytes()), 10, 0)
		info.eipInstructionPointer = uint32(utoobig)
// 31
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		// signal Obsolete, because it does not provide information on real-time signals;
		// use /proc/[pid]/status instead.
// 32
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		// blocked Obsolete, because it does not provide information on real-time signals;
		// use /proc/[pid]/status instead.
// 33
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		// sigignore  Obsolete, because it does not provide information on real-time signals;y2
		// use /proc/[pid]/status instead.
// 34
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		// sigcatch  Obsolete, because it does not provide information on real-time signals;y2
		// use /proc/[pid]/status instead.
// 35

	// TODO the rest from `man 5 proc`

	return
}

// give process id
func (process *ProcessInfo) Id() uint64 {
	return process.id
}

// give process name
func (process *ProcessInfo) Name() string {
	return process.name
}
