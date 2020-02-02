/**
 * Copyright 2019 Innodev LLC. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package cache

import (
	"encoding/json"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/gregjones/httpcache"
	"github.com/gregjones/httpcache/redis"
	lru "github.com/hashicorp/golang-lru"
)

type CacheRedis interface {
	Add(key string, value interface{}) error
	Get(key string, out interface{}) (ok bool, err error)
	Remove(key string)
}

type cacheRedis struct {
	cache  *lru.Cache
	cache2 httpcache.Cache
}

type RedisConfig struct {
	Addr     string
	Network  string
	Password string
	Database int
}

func NewRedis(size int, conf RedisConfig) CacheRedis {
	conn, err := redigo.Dial(conf.Network, conf.Addr,
		redigo.DialPassword(conf.Password), redigo.DialDatabase(conf.Database))
	if err != nil {
		panic(err)
	}
	out := &cacheRedis{
		cache2: redis.NewWithClient(conn),
	}
	c, err := lru.New(size)
	if err != nil {
		panic(err)
	}
	out.cache = c
	return out
}

func (c *cacheRedis) Add(key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	c.cache.Add(key, data)
	c.cache2.Set(key, data)
	return err
}

func (c *cacheRedis) Get(key string, out interface{}) (ok bool, err error) {
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

func (c *cacheRedis) Remove(key string) {
	c.cache.Remove(key)
	c.cache2.Delete(key)
}
