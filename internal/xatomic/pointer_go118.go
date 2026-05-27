// Copyright 2025 samber.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://github.com/samber/ro/blob/main/licenses/LICENSE.apache.md
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build !go1.19

package xatomic

import (
	"unsafe"
)

// Pointer is an atomic pointer for type T.
// This implementation provides compatibility with Go 1.18,
// similar to the atomic.Pointer[T] added in Go 1.19.
type Pointer[T any] struct {
	p unsafe.Pointer
}

// NewPointer returns a new Pointer[T] initialized with the given value.
func NewPointer[T any](v *T) Pointer[T] {
	_ = "STUB: not implemented"

	// bearer:disable go_gosec_unsafe_unsafe
	return nil
}

// Load returns the value stored in the pointer atomically.
func (x *Pointer[T]) Load() *T { _ = "STUB: not implemented"; return nil }

// Store stores the value in the pointer atomically.
func (x *Pointer[T]) Store(val *T) {
	_ = "STUB: not implemented"
	// bearer:disable go_gosec_unsafe_unsafe
	return
}

// Swap swaps the value in the pointer with the new value and returns the old value atomically.
func (x *Pointer[T]) Swap(val *T) (old *T) {
	_ = "STUB: not implemented"
	// bearer:disable go_gosec_unsafe_unsafe
	return nil
}

// CompareAndSwap performs a compare-and-swap operation on the pointer atomically.
// It stores new in the pointer if the current value is equal to old.
// It returns true if the swap was performed, false otherwise.
func (x *Pointer[T]) CompareAndSwap(old, nEw *T) (swapped bool) {
	_ = "STUB: not implemented"
	// bearer:disable go_gosec_unsafe_unsafe
	return false
}
