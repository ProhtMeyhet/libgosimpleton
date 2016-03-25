package caching

import(

)


type CacheItem struct {
	name string
	contents []byte
	closing chan bool
}
