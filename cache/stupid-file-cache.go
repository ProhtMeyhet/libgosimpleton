package caching

import(
	"io/ioutil"
	"sync"

	simpleton "github.com/ProhtMeyhet/libgosimpleton"
	"github.com/ProhtMeyhet/libgosimpleton/logging"
)


/*
* a stupid cache that needs to be updated manually.
* good luck with that.
*
* it's meant to be abstract, but can be used on its own.
*
* for linux see the ObservingFileCache with inotify support.
*
*	TODO
*	check and correct interface definition
*	turn into StupidReaderCache for io.Reader
*	add a time based BoringFileCache
*	add a shared memory cache PromiscuousFileCache
*/
type StupidFileCache struct {
	sync.Mutex

	// map[filename]filecontent
	cache			map[string][]byte

	// count of total bytes read
	byteCount		uint

	// cache maximum size in bytes
	sizeMax			uint

	// at least free this amount
	// when sizeMax is reached
	freeThreshold		uint

	// maximum files that are cached. if hit, free freeMaxFilesThreshold
	filesMax		uint

	// at least free this amount fifo files from cache
	freeFilesMaxThreshold	uint

	itemChannel		chan *CacheItem
	numberOfManagerThreads	uint8

	// last error of inotify
	lastE			error
}

func NewStupidFileCache() (cache *StupidFileCache) {
	cache = &StupidFileCache{ }
	cache.init()
	return
}

func NewStupidFileCacheMaxSize(maxSize uint) (cache *StupidFileCache) {
	cache = NewStupidFileCache()
	cache.sizeMax = maxSize

	if cache.freeThreshold == 0 {
		cache.freeThreshold = maxSize / 4
	}

	return
}

func (cache *StupidFileCache) init() {
	cache.sizeMax = 1024 * 1024 // 1mb
	cache.freeThreshold = cache.sizeMax / 4
	cache.cache = make(map[string][]byte, 20)
	cache.itemChannel = make(chan *CacheItem, 5)
	cache.filesMax = 100
	cache.freeFilesMaxThreshold = cache.filesMax / 4

	if cache.numberOfManagerThreads == 0 {
		cache.numberOfManagerThreads = 2
	}

	for i := uint8(1); i <= cache.numberOfManagerThreads; i++ {
		go cache.manager()
	}
}

func (cache *StupidFileCache) GetString(filename string) (string, error) {
	bytes, e := cache.Get(filename)
	return string(bytes), e
}

func (cache *StupidFileCache) Get(filename string) (value []byte, e error) {
	// locking required:
	// golang fatal error: concurrent map read and map write
	cache.Lock()
	ok := false
	if value, ok = cache.cache[filename]; !ok {
		cache.Unlock()
		// cache miss
		value, e = cache.cacheIt(filename)
	} else {
		cache.Unlock()
	}

	return
}

func (cache *StupidFileCache) cacheIt(filename string) (contents []byte, e error) {
	item := &CacheItem{ name: filename }
	contents, e = ioutil.ReadFile(filename)
	if e != nil { logging.Log(logging.ERROR, "inotify: " + e.Error()); goto out }

	item.contents = contents
	cache.itemChannel <- item

out:
	return
}

// store and clean up thread
/* go */ func (cache *StupidFileCache) manager() {
	for {
		select {
		case item := <-cache.itemChannel:
			// lock protects cache.byteCount
			cache.Lock()
			cache.cache[item.name] = item.contents
			cache.byteCount += uint(len(cache.cache[item.name]))
			// unlock as early as possible
			// assume: cache size is not exhausted
			cache.Unlock()
			go cache.testCacheSize(item.name)
		}
	}
}

func (cache *StupidFileCache) testCacheSize(dontdelete string) {
	if uint(len(cache.cache)) > cache.filesMax {
		cache.Lock(); i := uint(0)
		if simpleton.DEBUG { logging.Log(logging.INFO, "cache: have to free! hit filesMax!") }
		for key, _ := range cache.cache {
			if i == cache.freeFilesMaxThreshold { break }
			cache.remove(key)
			i++
		}
		cache.Unlock()
	}

	if uint(len(cache.cache[dontdelete])) > cache.sizeMax {
		logging.Log(logging.ERROR, "file " + dontdelete + " is bigger then sizeMax! raising sizeMax!")
		if cache.sizeMax * 10 > uint(len(cache.cache[dontdelete])) &&
			cache.sizeMax * 10 < 104857600 {
			cache.sizeMax *= 10
		} else {
			logging.Log(logging.ERROR, "file " + dontdelete + ` is waaaay to big
					or cache.sizeMax *= 10 is bigger then 100mb!`)
		}
	}

	if cache.byteCount > cache.sizeMax {
		// lock protects cache.byteCount
		cache.Lock()
		if simpleton.DEBUG { logging.Log(logging.INFO, "cache: have to free!") }
		for key, _ := range cache.cache {
			if cache.Free() >= cache.freeThreshold { break }
			if key == dontdelete { continue }
			if simpleton.DEBUG { logging.Log(logging.INFO, "cache: removing " + key +
						" from cache! dontdelete: " + dontdelete) }
			cache.remove(key)
		}
		cache.Unlock()
	}
}

func (cache *StupidFileCache) Reset() {
	// lock protects cache.byteCount
	cache.Lock()
	cache.cache = make(map[string][]byte, 20)
	cache.byteCount = 0
	cache.Unlock()
}

func (cache *StupidFileCache) Remove(filename string) {
	// lock protects cache.byteCount
	cache.Lock()
	cache.remove(filename)
	cache.Unlock()
}

// remember locks!
func (cache *StupidFileCache) remove(filename string) {
	cache.byteCount -= uint(len(cache.cache[filename]))
	delete(cache.cache, filename)
}

func (cache *StupidFileCache) GetLastE() error {
	return cache.lastE
}

func (cache *StupidFileCache) Free() uint {
	if cache.byteCount <= cache.sizeMax {
		return cache.sizeMax - cache.byteCount
	} else {
		return 0
	}
}

func (cache *StupidFileCache) GetSize() uint {
	return cache.byteCount
}

func (cache *StupidFileCache) GetMaxSize() uint {
	return cache.sizeMax
}

func (cache *StupidFileCache) SetMaxSize(to uint) {
	cache.sizeMax = to
}
