/**
 * Copyright 2019 Innodev LLC. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package cache

type Cache interface {
	Add(key string, value interface{}) error
	Get(key string, out interface{}) (ok bool, err error)
	Remove(key string) error
}
