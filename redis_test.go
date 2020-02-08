/**
 * Copyright 2019 Innodev LLC. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedis(t *testing.T) {
	testVal := map[string]string{
		"foo": "bar",
	}
	c := NewRedis(RedisConfig{
		Servers: map[string]string{
			"local": "127.0.0.1:6379",
		},
		Password: "password",
		Database: 1,
	})
	c.Add("hello", testVal)
	c.Add("hello2", testVal)
	c.Add("hello3", testVal)
	var val map[string]string
	ok, err := c.Get("hello", &val)
	assert.Equal(t, testVal, val)
	assert.True(t, ok)
	assert.NoError(t, err)
	c.Remove("hello")
	ok, _ = c.Get("hello", &val)
	assert.False(t, ok)
}
