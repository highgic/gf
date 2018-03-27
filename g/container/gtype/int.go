// Copyright 2018 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

package gtype

import (
    "sync/atomic"
)

type Int struct {
    val int64
}

func NewInt(value...int) *Int {
    if len(value) > 0 {
        return &Int{val:int64(value[0])}
    }
    return &Int{}
}

func (t *Int)Set(value int) {
    atomic.StoreInt64(&t.val, int64(value))
}

func (t *Int)Get() int {
    return int(atomic.LoadInt64(&t.val))
}

func (t *Int)Add(delta int) int {
    return int(atomic.AddInt64(&t.val, int64(delta)))
}