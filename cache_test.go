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


func TestCache(t *testing.T) {
	c := New(200, "/tmp/foobar")
	c.Add("hello","world")
	val, ok := c.Get("hello")
	assert.Equal(t, "world", val)
	assert.True(t, ok)
	c.Remove("hello")
	_, ok = c.Get("hello")
	assert.False(t, ok)
}
