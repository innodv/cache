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

func TestCache2(t *testing.T) {
	testVal := map[string]string{
		"foo": "bar",
	}
	c := New2(200, "/tmp/foobar")
	c.Add("hello", testVal)
	var val map[string]string
	ok, err := c.Get("hello", &val)
	assert.Equal(t, testVal, val)
	assert.True(t, ok)
	assert.NoError(t, err)
	c.Remove("hello")
	ok, _ = c.Get("hello", &val)
	assert.False(t, ok)
}
