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
*	add error handling
*	add support for optional inotify delete
*/
type ObservingFileCache struct {
	StupidFileCache

	stop chan bool
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

func (cache *ObservingFileCache) init() {
	cache.stop = make(chan bool, 1)
	cache.StupidFileCache.init()
}

func (cache *ObservingFileCache) Get(filename string) (value []byte, e error) {
	// must start goroutine first and wait until watch is started, otherwise
	// events might not be catched due to a race condition
	// for example a write might happen before the inotify is started, but the
	// read already took place. this write would then go unnoticed.
	started := make(chan bool, 1)
	go cache.watch(filename, started)
	<-started
	return cache.StupidFileCache.Get(filename)
}

// TODO howto stop watching when file is removed from cache?
/* go */ func (cache *ObservingFileCache) watch(filename string, started chan bool) {
	if simpleton.DEBUG { logging.DebugFormat("inotify started: %v", filename) }

	watcher, e := inotify.NewWatcher()
	if e != nil { logging.Error("inotify: %v", e.Error()); cache.lastE = e; return }
	defer watcher.Close()

	//TODO determine if modify is enough
	e = watcher.AddWatch(filename, inotify.IN_MODIFY)

	// following line WILL produce an infinte loop.
	// you have been warned!
	// e = watcher.Watch(filename)

	started <-true
	if e != nil { logging.ErrorFormat("inotify: %v", e.Error()); cache.lastE = e; return }

infinite:
	for {
		select {
		case event := <-watcher.Event:
			if simpleton.DEBUG { logging.DebugFormat("inotify event: %v", event.String()) }
			cache.cacheIt(filename)
		case e := <-watcher.Error:
			logging.ErrorFormat("inotify: %v", e.Error())
			return
		case _, stop := <-cache.stop:
			if stop { break infinite }
		}
	}
}

func (cache *ObservingFileCache) Close() {
	close(cache.stop)
}
