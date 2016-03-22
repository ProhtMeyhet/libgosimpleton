package caching

// +build linux

import(
	"golang.org/x/exp/inotify"
	simpleton "github.com/ProhtMeyhet/libgosimpleton"
	"github.com/ProhtMeyhet/libgosimpleton/logging"
)


/*
* uses inotify to detect changes made to the cached files
* and if so update the cache
*
*	TODO
*	add support for optional inotify delete
*/
type ObservingFileCache struct {
	StupidFileCache
}

func NewObservingFileCache() (cache *ObservingFileCache) {
	cache = &ObservingFileCache{ }
	cache.init()

	return
}

func NewObservingFileCacheMaxSize(maxSize uint) (cache *ObservingFileCache) {
	cache = NewObservingFileCache()
	cache.sizeMax = maxSize
	cache.freeThreshold = cache.sizeMax / 4

	return
}

func (cache *ObservingFileCache) Get(filename string) (value []byte, e error) {
	value, e = cache.StupidFileCache.Get(filename)

	if e == nil {
		go cache.watch(filename)
	}

	return
}

/* go */ func (cache *ObservingFileCache) watch(filename string) {
	if simpleton.DEBUG { logging.Log(logging.INFO, "inotify started: " + filename) }

	watcher, e := inotify.NewWatcher()
	if e != nil { logging.Log(logging.ERROR, "inotify: " + e.Error()); cache.lastE = e; return }
	defer watcher.Close()

	//TODO determine if modify is enough
	e = watcher.AddWatch(filename, inotify.IN_MODIFY)

	// following line WILL produce an infinte loop.
	// you have been warned!
	// e = watcher.Watch(filename)

	if e != nil { logging.Log(logging.ERROR, "inotify: " + e.Error()); cache.lastE = e; return }

	for {
		select {
		case event := <-watcher.Event:
			if simpleton.DEBUG { logging.Log(logging.INFO, "inotify event: " + event.String()) }
			cache.cacheIt(filename)
		case e := <-watcher.Error:
			logging.Log(logging.ERROR, "inotify: " + e.Error())
			return
		}
	}
}
