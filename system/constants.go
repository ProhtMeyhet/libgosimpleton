package system

const(
	PROC_PATH		= "/proc/"
	PROC_STAT_FILE		= PROC_PATH + "%v/stat"
	PROC_GLOB		= PROC_PATH + "[0-9]*"
	PROC_LOAD_AVERAGE	= "/proc/loadavg"
	PROC_ENTITIES_SEPARATOR = "/"
)
