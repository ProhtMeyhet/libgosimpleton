package processes

const(
	PROC_PATH		= "/proc/"
	PROC_STAT_FILE		= PROC_PATH + "%v/stat"
	PROC_CMDLINE		= PROC_PATH + "%v/cmdline"
	PROC_GLOB		= PROC_PATH + "[0-9]*"
)
