/**
 * Copyright 2019 Innodev LLC. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package cache

import (
	"encoding/json"

	"github.com/gregjones/httpcache/diskcache"
)

type disk struct {
	cache *diskcache.Cache
}

func NewDisk(dir string) Cache {
	return &disk{
		cache: diskcache.New(dir),
	}
}

func (c *disk) Add(key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	c.cache.Set(key, data)
	return nil
}

func (c *disk) Get(key string, out interface{}) (ok bool, err error) {
	data, ok := c.cache.Get(key)
	if ok {
		err = json.Unmarshal(data, out)
		if err != nil {
			return false, err
		}
		return
	}
	return false, nil
}

func (c *disk) Remove(key string) error {
	c.cache.Delete(key)
	return nil
}
