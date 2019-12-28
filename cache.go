package cache

import (
	"github.com/gregjones/httpcache/diskcache"
	lru "github.com/hashicorp/golang-lru"
)

type Cache interface {
	Add(key, value string)
	Get(key string) (value string, ok bool)
	Remove(key string)
}

type cacher struct {
	cache  *lru.Cache
	cache2 *diskcache.Cache
}

func New(size int, dir string) Cache {
	out := &cacher{
		cache2: diskcache.New(dir),
	}
	c, err := lru.NewWithEvict(size, out.onEvicted)
	if err != nil {
		panic(err)
	}
	out.cache = c
	return out
}

func (c *cacher) onEvicted(key interface{}, value interface{}) {
	go c.cache2.Delete(key.(string))
}

func (c *cacher) Add(key, value string) {
	c.cache.Add(key, value)
	c.cache2.Set(key, []byte(value))
}

func (c *cacher) Get(key string) (value string, ok bool) {
	out, ok := c.cache.Get(key)
	if ok {
		return out.(string), ok
	}
	res, ok := c.cache2.Get(key)
	if ok {
		return string(res), ok
	}
	return "", false
}

func (c *cacher) Remove(key string) {
	c.cache.Remove(key)
	c.cache2.Delete(key)
}