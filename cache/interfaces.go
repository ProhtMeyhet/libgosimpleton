package caching

import(

)

// TODO
// unfinished, check everything

type CacheInterface interface {
	Get() ([]byte, error)
	GetString() (string, error)

	// tell count of free bytes
	Free() uint
	GetSize() uint
	GetMaxSize() uint
	SetMaxSize(to uint)

	// get last error
	GetLastE() error
}


type StupidInterface interface {
	Reset()
	Remove(filename string)
}
