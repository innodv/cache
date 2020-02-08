/**
 * Copyright 2019 Innodev LLC. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package cache

import (
	"time"

	"github.com/go-redis/cache/v7"
	"github.com/go-redis/redis/v7"
	"github.com/vmihailenco/msgpack/v4"
)

type cacheRedis struct {
	prefix string
	exp    time.Duration
	codec  *cache.Codec
}

type RedisConfig struct {
	Servers    map[string]string
	Password   string
	Database   int
	Prefix     string
	Expiration time.Duration
}

func NewRedis(conf RedisConfig) Cache {
	return &cacheRedis{
		codec: &cache.Codec{
			Redis: redis.NewRing(&redis.RingOptions{
				Addrs:    conf.Servers,
				Password: conf.Password,
				DB:       conf.Database,
			}),
			Marshal: func(v interface{}) ([]byte, error) {
				return msgpack.Marshal(v)
			},
			Unmarshal: func(b []byte, v interface{}) error {
				return msgpack.Unmarshal(b, v)
			},
		},
		exp:    conf.Expiration,
		prefix: conf.Prefix,
	}
}

func (c *cacheRedis) Add(key string, value interface{}) error {
	return c.codec.Set(&cache.Item{
		Key:        c.prefix + key,
		Object:     value,
		Expiration: c.exp,
	})
}

func (c *cacheRedis) Get(key string, out interface{}) (bool, error) {
	err := c.codec.Get(c.prefix+key, out)
	return err == nil, err
}

func (c *cacheRedis) Remove(key string) error {
	return c.codec.Delete(c.prefix + key)
}
