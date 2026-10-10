// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package safeexpvar

import (
	"expvar"
	"sync"
)

var mu sync.Mutex

func SetInt(name string, value int64) {
	v := expvar.Get(name)
	if v == nil {
		mu.Lock()
		// expvar.Get followed by NewInt is not atomic; re-check under the lock.
		v = expvar.Get(name)
		if v == nil {
			v = expvar.NewInt(name)
		}
		mu.Unlock()
	}
	v.(*expvar.Int).Set(value)
}
