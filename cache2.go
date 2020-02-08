/**
 * Copyright 2019 Innodev LLC. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package cache

import (
	"encoding/json"

	"github.com/gregjones/httpcache/diskcache"
	lru "github.com/hashicorp/golang-lru"
)

type cacher2 struct {
	cache  *lru.Cache
	cache2 *diskcache.Cache
}

func New2(size int, dir string) Cache2 {
	out := &cacher2{
		cache2: diskcache.New(dir),
	}
	c, err := lru.New(size)
	if err != nil {
		panic(err)
	}
	out.cache = c
	return out
}

func (c *cacher2) Add(key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	c.cache.Add(key, data)
	c.cache2.Set(key, data)
	return err
}

func (c *cacher2) Get(key string, out interface{}) (ok bool, err error) {
	val, ok := c.cache.Get(key)
	if ok {
		err = json.Unmarshal(val.([]byte), out)
		if err != nil {
			return false, err
		}
		return
	}
	data, ok := c.cache2.Get(key)
	if ok {
		err = json.Unmarshal(data, out)
		if err != nil {
			return false, err
		}
		return
	}
	return false, nil
}

func (c *cacher2) Remove(key string) {
	c.cache.Remove(key)
	c.cache2.Delete(key)
}
