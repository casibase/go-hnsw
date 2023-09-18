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
	"sync"

	"github.com/bits-and-blooms/bitset"
)

type poolItem struct {
	b    bitset.BitSet
	busy bool
}

type BitsetPool struct {
	sync.RWMutex
	pool []poolItem
}

func New() *BitsetPool {
	var bp BitsetPool
	bp.pool = make([]poolItem, 0)
	return &bp
}

func (bp *BitsetPool) Free(i int) {
	bp.Lock()
	bp.pool[i].busy = false
	bp.Unlock()
}

func (bp *BitsetPool) Get() (int, *bitset.BitSet) {
	bp.Lock()
	for i := range bp.pool {
		if !bp.pool[i].busy {
			bp.pool[i].busy = true
			bp.pool[i].b.ClearAll()
			bp.Unlock()
			return i, &bp.pool[i].b
		}
	}
	id := len(bp.pool)
	bp.pool = append(bp.pool, poolItem{})
	bp.Unlock()
	return id, &bp.pool[id].b
}
