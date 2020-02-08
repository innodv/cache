/**
 * Copyright 2019 Innodev LLC. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package cache

import (
	"encoding/json"
	golru "github.com/hashicorp/golang-lru"
)

type lru struct {
	cache *golru.Cache
}

func NewLRU(size int) Cache {
	c, err := golru.New(size)
	if err != nil {
		panic(err)
	}
	return &lru{
		cache: c,
	}
}

func (c *lru) Add(key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	c.cache.Add(key, data)
	return err
}

func (c *lru) Get(key string, out interface{}) (ok bool, err error) {
	val, ok := c.cache.Get(key)
	if ok {
		err = json.Unmarshal(val.([]byte), out)
		if err != nil {
			return false, err
		}
		return
	}
	return false, nil
}

func (c *lru) Remove(key string) error {
	c.cache.Remove(key)
	return nil
}
