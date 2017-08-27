package parallel

const(
	// use N workers per cpu
	NUMBER_OF_WORKERS_MULTIPLIER	= 2

	// use N buffers per worker
	BUFFER_SIZE_MULTIPLIER		= 4

	// use N buffers per worker when working with files
	// this is useful, if fadvise is used
	FILE_BUFFER_SIZE_MULTIPLIER	= 8
)
