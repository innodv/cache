/**
 * Copyright 2019 Innodev LLC. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */


package cache

import(
	"github.com/innodv/errors"
)

// MultiCache makes multi-layer caching a breeze
type MultiCache []Cache


func (mc MultiCache) Add(key string, value interface{}) error {
	errChan := make(chan bool)
	for i := range mc {
		go func(i int) {
			errChan <- mc[i].Add(key, value)
		}(i)
	}
	return await.AwaitErrors(errChan, len(mc))
}

func (mc MultiCache) Get(key string, out interface{}) (bool,error) {
	for i := range mc {
		ok, err := mc[i].Get(key, out)
		if ok {
			return
		}
		if err != nil {
			return ok, err
		}
	}
}

func (mc MultiCache) Remove(key string) error {
	errChan := make(chan bool)
	for i := range mc {
		go func(i int) {
			errChan <- mc[i].Remove(key)
		}(i)
	}
	return await.AwaitErrors(errChan, len(mc))
}
