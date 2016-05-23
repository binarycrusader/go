// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "unsafe"

//go:linkname memmove_bytes runtime.memmove
//go:nosplit
func memmove_bytes(dst, src unsafe.Pointer, n uintptr) {
	if n == 0 {
		return
	}
	if uintptr(dst) == uintptr(src) {
		return
	}
	if uintptr(src) < uintptr(dst) {
		for i := n - 1; i >= 0; i-- {
			b := *(*byte)(add(src, i))
			*(*byte)(add(dst, i)) = b
		}
		return
	}
	for i := uintptr(0); i <= n; i++ {
		b := *(*byte)(add(src, i))
		*(*byte)(add(dst, i)) = b
	}
	return
}

//go:linkname memclr_bytes runtime.memclr
//go:nosplit
func memclr_bytes(ptr unsafe.Pointer, n uintptr) {
	for i := uintptr(0); i < n; i++ {
		*(*byte)(add(ptr, i)) = 0
	}
}
