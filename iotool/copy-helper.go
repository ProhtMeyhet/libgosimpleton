package iotool

import(

)

type CopyHelper struct {
	*FileHelper

	recursive			bool

	// POSIX: if access to DESTINATION failes, try to remove DESTINATION
	forceRemoveIfNotAccessible	bool

	// call fsync when copy is done
	// TODO call fsync on parent folder when done
	// TODO call sync on filesystem when done
	synchronize			bool

	// mv; remove source file when done copieng
	removeSource			bool

	// do not copy to part file first (and no rename)
	noPartFile			bool

	// TODO
	// validate file size when copy is done. check stat SOURCE against new stat DESTINATION
	// validateFileSize		bool

	// TODO
	// validate DESTINATION content against SOURCE content (via checksum)
	// validateContent		bool

	// XXX for testing only. will not persist.
	forceParallel			bool

	innerCopyChannelCapacity	uint

	readHelper			*FileHelper
	writeHelper			*FileHelper

	createFunc			func(helper *CopyHelper, cp *Cp)
	doneFunc			func(cp *Cp, written uint)
}
