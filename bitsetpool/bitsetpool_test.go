// Copyright 2023 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bitsetpool

import (
	"math/rand"
	"testing"
	"time"

	"github.com/bits-and-blooms/bitset"
)

func TestBitset(t *testing.T) {

	start2 := time.Now()
	for j := 0; j < 100000; j++ {
		b2 := make(map[uint32]bool)
		for i := 0; i < 100; i++ {
			n := rand.Intn(1000000)
			b2[uint32(n)] = true
			m := rand.Intn(1000000)
			if b2[uint32(m)] == false {
			}
		}
	}
	stop2 := time.Since(start2)
	t.Logf("map done in %v", stop2.Seconds())

	start := time.Now()
	for j := 0; j < 100000; j++ {
		var b1 bitset.BitSet
		for i := 0; i < 100; i++ {
			n := rand.Intn(1000000)
			b1.Set(uint(n))
			m := rand.Intn(1000000)
			b1.Test(uint(m))
		}
	}
	stop := time.Since(start)
	t.Logf("bitset done in %v", stop.Seconds())

	start3 := time.Now()
	pool := New()
	for j := 0; j < 100000; j++ {
		id, b := pool.Get()
		for i := 0; i < 100; i++ {
			n := rand.Intn(1000000)
			b.Set(uint(n))
			m := rand.Intn(1000000)
			b.Test(uint(m))
		}
		pool.Free(id)
	}
	stop3 := time.Since(start3)
	t.Logf("bitset pool done in %v", stop3.Seconds())

	t.Logf("Performance boost %.2f%%", 100*(1-stop3.Seconds()/stop2.Seconds()))
}
