package processes

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
	id				uint
	name				string

	owner				int
	group				int

	search				string

	globList			[]string
	globKey				uint

	commandLine			string

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
	// despite otherwise documented, softLimitResidentSetSize is uint64, not uint32 on x86_64
	residentSetSize			uint64
	softLimitResidentSetSize	uint64
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

// find process by id
func (info *ProcessInfo) findById() (e error) {
	if info.id == 0 { return INVALID_PROCESS_ID_ERROR }

	handler, e := iotool.Open(iotool.ReadOnly(), fmt.Sprintf(PROC_STAT_FILE, info.id))
	if e != nil && os.IsNotExist(e) { e = NO_SUCH_PROCESS_ERROR; return }; defer handler.Close()
	return info.scanStat(handler)
}

// find a process by filter. first call gives first result, second call second result ...
// panics if glob fails
func (info *ProcessInfo) findBy(filter func(*ProcessInfo) bool) bool {
	if filter == nil { return false }

	if len(info.globList) == 0 { var e error
		info.globList, e = filepath.Glob(PROC_GLOB); if e != nil { panic(e); return false }
	}

	for ; info.globKey < uint(len(info.globList)); info.globKey++ {
		id := info.globList[info.globKey][len(PROC_PATH):]
		process, e := FindByStringId(id)
		// process ended in between
		if e != nil { continue }
		process.search = info.search
		if filter(process) {
			info.globKey++; info.Copy(process); return true
		}
	}

	// save memory
	info.globList = nil; info.globKey = 0

	return false
}

// scan /proc/[pid]/stat and read /proc/[pid]/cmdline
func (info *ProcessInfo) scanStat(handler iotool.FileInterface) (e error) {
	fileInfo, e := handler.Stat(); if e != nil { return }
	iofileinfo := iotool.NewFileInfo(handler.Name(), fileInfo)
	info.owner = iofileinfo.UserId(); info.group = iofileinfo.GroupId()

	scanner := bufio.NewScanner(handler); scanner.Split(bufio.ScanWords)

// 1
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		// pid
		if info.id == 0 {
			id, e := strconv.ParseUint(string(scanner.Bytes()), 10, 0); if e != nil { return e }
			info.id = uint(id)
		}

// 2
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		// name
		if info.name == "" {
			info.name = string(scanner.Bytes())[1:]
			info.name = info.name[:len(info.name)-1]
		}

	// read /proc/[pid]/cmdline
	info.commandLine, e = iotool.ReadFileAsString(fmt.Sprintf(PROC_CMDLINE, info.id))
	if e != nil { return UNEXPECTED_CMDLINE_ERROR }
	info.commandLine = strings.Replace(info.commandLine, "\x00", " ", -1)

// 3
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		info.state = string(scanner.Bytes())
// 4
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		toobig, _ := strconv.ParseInt(string(scanner.Bytes()), 10, 32)
		info.parentProcessId = int32(toobig)
// 5
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		toobig, _ = strconv.ParseInt(string(scanner.Bytes()), 10, 32)
		info.processGroupId = int32(toobig)
// 6
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		toobig, _ = strconv.ParseInt(string(scanner.Bytes()), 10, 32)
		info.sessionId = int32(toobig)
// 7
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		toobig, _ = strconv.ParseInt(string(scanner.Bytes()), 10, 32)
		info.tty = int32(toobig)
// 8
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		toobig, _ = strconv.ParseInt(string(scanner.Bytes()), 10, 32)
		info.foregroundProcessId = int32(toobig)
// 9
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		utoobig, _ := strconv.ParseUint(string(scanner.Bytes()), 10, 32)
		info.kernelFlags = uint32(utoobig)
// 10
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		utoobig, _ = strconv.ParseUint(string(scanner.Bytes()), 10, 32)
		info.memoryMinorFaults = uint32(utoobig)
// 11
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		utoobig, _ = strconv.ParseUint(string(scanner.Bytes()), 10, 32)
		info.childrenMinorFaults = uint32(utoobig)
// 12
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		utoobig, _ = strconv.ParseUint(string(scanner.Bytes()), 10, 32)
		info.memoryMajorFaults = uint32(utoobig)
// 13
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		utoobig, _ = strconv.ParseUint(string(scanner.Bytes()), 10, 32)
		info.childrenMajorFaults = uint32(utoobig)
// 14
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		utoobig, _ = strconv.ParseUint(string(scanner.Bytes()), 10, 32)
		info.rawUserTime = uint32(utoobig)
// 15
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		utoobig, _ = strconv.ParseUint(string(scanner.Bytes()), 10, 32)
		info.rawScheduledTime = uint32(utoobig)
// 16
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		toobig, _ = strconv.ParseInt(string(scanner.Bytes()), 10, 32)
		info.rawChildrenUserWaitTime = int32(toobig)
// 17
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		toobig, _ = strconv.ParseInt(string(scanner.Bytes()), 10, 32)
		info.rawChildrenKernelWaitTime = int32(toobig)
// 18
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		toobig, _ = strconv.ParseInt(string(scanner.Bytes()), 10, 32)
		info.priority = int32(toobig)
// 19
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		toobig, _ = strconv.ParseInt(string(scanner.Bytes()), 10, 32)
		info.nice = int32(toobig)
// 20
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		toobig, _ = strconv.ParseInt(string(scanner.Bytes()), 10, 32)
		info.numberOfThreads = int32(toobig)
// 21
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		// itrealvalue since kernel 2.6.17, this field is no longer maintained, and is hard coded as 0.
// 22
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		info.relativeStartTime, e = strconv.ParseUint(string(scanner.Bytes()), 10, 0)
// 23
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		utoobig, _ = strconv.ParseUint(string(scanner.Bytes()), 10, 32)
		info.virtualMemory = uint32(utoobig)
// 24
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		info.residentSetSize, _ = strconv.ParseUint(string(scanner.Bytes()), 10, 0)
// 25
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		info.softLimitResidentSetSize, _ = strconv.ParseUint(string(scanner.Bytes()), 10, 0)
// 26
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		utoobig, _ = strconv.ParseUint(string(scanner.Bytes()), 10, 32)
		info.startCode = uint32(utoobig)
// 27
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		utoobig, _ = strconv.ParseUint(string(scanner.Bytes()), 10, 32)
		info.endCode = uint32(utoobig)
// 28
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		utoobig, _ = strconv.ParseUint(string(scanner.Bytes()), 10, 32)
		info.startStack = uint32(utoobig)
// 29
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		utoobig, _ = strconv.ParseUint(string(scanner.Bytes()), 10, 32)
		info.espStackPointer = uint32(utoobig)
// 30
	if !scanner.Scan() { return UNEXPECTED_STAT_FORMAT_ERROR }
		utoobig, _ = strconv.ParseUint(string(scanner.Bytes()), 10, 32)
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

// make a copy of this *ProcessInfo
func (info *ProcessInfo) MakeCopy() (process *ProcessInfo) {
	return info.MakeCopyFrom(info)
}

// copy values from a given *ProcessInfo to a new *ProcessInfo
func (info *ProcessInfo) MakeCopyFrom(process *ProcessInfo) (processCopy *ProcessInfo) {
	processCopy = &ProcessInfo{}; processCopy.Copy(process); return
}

// copy values from a *ProcessInfo
func (info *ProcessInfo) Copy(from *ProcessInfo) {
	info.id = from.id; info.name = from.name; info.search = from.search
	info.owner = from.owner; info.group = from.group
	info.commandLine = from.commandLine

	info.state			= from.state
	info.parentProcessId		= from.parentProcessId
	info.processGroupId		= from.processGroupId
	info.sessionId			= from.sessionId
	info.tty			= from.tty
	info.foregroundProcessId	= from.foregroundProcessId
	info.kernelFlags		= from.kernelFlags
	info.childrenMinorFaults	= from.childrenMinorFaults
	info.memoryMajorFaults		= from.memoryMajorFaults
	info.childrenMajorFaults	= from.childrenMajorFaults
	info.memoryMinorFaults		= from.memoryMinorFaults
	info.rawUserTime		= from.rawUserTime
	info.rawScheduledTime		= from.rawScheduledTime
	info.rawChildrenUserWaitTime	= from.rawChildrenUserWaitTime
	info.rawChildrenKernelWaitTime	= from.rawChildrenKernelWaitTime
	info.priority			= from.priority
	info.nice			= from.nice
	info.numberOfThreads		= from.numberOfThreads
	// itrealvalue
	info.relativeStartTime		= from.relativeStartTime
	info.virtualMemory		= from.virtualMemory
	info.residentSetSize		= from.residentSetSize
	info.softLimitResidentSetSize	= from.softLimitResidentSetSize
	info.startCode			= from.startCode
	info.endCode			= from.endCode
	info.startStack			= from.startStack
	info.espStackPointer		= from.espStackPointer
	info.eipInstructionPointer	= from.eipInstructionPointer
	// signal
	// blocked
	// sigignore
	// sigcatch

	// TODO the rest from `man 5 proc`
}

func (info *ProcessInfo) String() string {
	s := " Name:\t%v\n" +
" Pid:\t%v\n" +
" State:\t%v\n" +
" ParentProcessId:\t%v\n" +
" ProcessGroupId:\t%v\n" +
" SessionId:\t%v\n" +
" Tty:\t%v\n" +
" ForegroundProcessId:\t%v\n" +
" KernelFlags:\t%v\n" +
" MemoryMinorFaults:\t%v\n" +
" ChildrenMinorFaults:\t%v\n" +
" MemoryMajorFaults:\t%v\n" +
" ChildrenMajorFaults:\t%v\n" +
" RawUserTime:\t%v\n" +
" RawSchduledTime:\t%v\n" +
" RawChildrenUserWaitTime:\t%v\n" +
" RawChildrenKernelWaitTime:\t%v\n" +
" Priority:\t%v\n" +
" Nice:\t%v\n" +
" NumberOfThreads:\t%v\n" +
" RelativeStartTime:\t%v\n" +
" VirtualMemory:\t%v\n" +
" ResidentSetSize:\t%v\n" +
" SoftLimitResidentSetSize:\t%v\n" +
" StartCode:\t%v\n" +
" EndCode:\t%v\n" +
" StartStack:\t%v\n" +
" EspStackPointer:\t%v\n" +
" EipInstructionPointer:\t%v\n"
    return fmt.Sprintf(s,
	info.name,
	info.id,
	info.state,
	info.parentProcessId,
	info.processGroupId,
	info.sessionId,
	info.tty,
	info.foregroundProcessId,
	info.kernelFlags,
	info.memoryMinorFaults,
	info.childrenMinorFaults,
	info.memoryMajorFaults,
	info.childrenMajorFaults,
	info.rawUserTime,
	info.rawScheduledTime,
	info.rawChildrenUserWaitTime,
	info.rawChildrenKernelWaitTime,
	info.priority,
	info.nice,
	info.numberOfThreads,
	// itrealvalue
	info.relativeStartTime,
	info.virtualMemory,
	info.residentSetSize,
	info.softLimitResidentSetSize,
	info.startCode,
	info.endCode,
	info.startStack,
	info.espStackPointer,
	info.eipInstructionPointer,
	// signal
	// blocked
	// sigignore
	// sigcatch
    )

	// TODO the rest from `man 5 proc`
}

// give process id
func (process *ProcessInfo) Id() uint {
	return process.id
}

// give process name
func (process *ProcessInfo) Name() string {
	return process.name
}

func (process *ProcessInfo) CommandLine() string {
	return process.commandLine
}

func (info *ProcessInfo) ParentProcessId() int32 {
	return info.parentProcessId
}

func (info *ProcessInfo) Priority() int32 {
	return info.priority
}

func (info *ProcessInfo) Nice() int32 {
	return info.nice
}

func (info *ProcessInfo) NumberOfThreads() int32 {
	return info.numberOfThreads
}

func (info *ProcessInfo) VirtualMemory() uint32 {
	return info.virtualMemory
}
