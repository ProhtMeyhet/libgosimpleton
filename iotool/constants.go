package iotool

const(
	DEBUG = false

	PREFIX_SEPARATOR = "_"

	// the smallest read for an io.Reader buffer
	READ_BUFFER_SMALL_SIZE		= 512
	// the usual read size for an io.Reader buffer
	READ_BUFFER_SIZE		= 4 * 1024
	// the maximum read size for an io.Reader buffer
	READ_BUFFER_MAXIMUM_SIZE	= 32 * 1024
)
