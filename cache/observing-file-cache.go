// +build linux

package caching

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

	// map[filename]chan stop
	threads map[string]chan bool
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
	cache.threads = make(map[string]chan bool)
	cache.StupidFileCache.init()
}

func (cache *ObservingFileCache) Get(filename string) (value []byte, e error) {
	cache.Lock() // do not defer cache.Unlock(); unlock as early as possible
	value, ok := cache.cache[filename]; if ok {
		cache.Unlock()
		return value, nil
	}

	// must start goroutine first and wait until watch is started, otherwise
	// events might not be catched due to a race condition
	// for example a write might happen before the inotify is started, but the
	// read already took place. this write would then go unnoticed.
	started := make(chan bool, 1)
	go cache.watch(filename, started)
	<-started
	cache.Unlock()
	return cache.StupidFileCache.Get(filename)
}

/* go */ func (cache *ObservingFileCache) watch(filename string, started chan bool) {
	if simpleton.DEBUG { logging.DebugFormat("inotify started: %v", filename) }

	watcher, e := inotify.NewWatcher()
	if e != nil { logging.ErrorFormat("inotify: %v", e.Error()); return }
	defer watcher.Close()

	//TODO determine if modify is enough
	e = watcher.AddWatch(filename, inotify.IN_MODIFY)

	// following line WILL produce an infinte loop.
	// you have been warned!
	// e = watcher.Watch(filename)

	started <-true
	if e != nil { logging.ErrorFormat("inotify: %v", e.Error());  return }

	cache.Lock()
	cache.threads[filename] = make(chan bool, 1)
	cache.Unlock()

infinite:
	for {
		select {
		case event := <-watcher.Event:
			if simpleton.DEBUG { logging.DebugFormat("inotify event: %v", event.String()) }
			cache.cacheIt(filename)
		case e := <-watcher.Error:
			logging.ErrorFormat("inotify: %v", e.Error())
			return
		case _, goOn := <-cache.threads[filename]:
			if !goOn {
				cache.Lock()
				delete(cache.threads, filename)
				cache.Unlock()
				break infinite
			}
		}
	}
}

// remove all files from cache
func (cache *ObservingFileCache) Reset() {
	cache.Lock()
	for key := range cache.cache {
		cache.remove(key)
	}
	cache.Unlock()
}

// remove from cache; stop watching goroutines.
func (cache *ObservingFileCache) Remove(filename string) {
	cache.Lock()
	cache.remove(filename)
	cache.Unlock()
}

// remember locks!
func (cache *ObservingFileCache) remove(filename string) {
	if _, ok := cache.threads[filename]; ok {
		close(cache.threads[filename])
	}
	cache.StupidFileCache.remove(filename)
}

// close
func (cache *ObservingFileCache) Close() {
	cache.Reset()
}
