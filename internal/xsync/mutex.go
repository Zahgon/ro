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

package xsync

import (
	"sync"
)

// Mutex is a mutex interface.
type Mutex interface {
	TryLock() bool
	Lock()
	Unlock()
}

/************************
 *    Standard mutex    *
 ************************/

var _ Mutex = (*MutexWithLock)(nil)

// NewMutexWithLock creates a new mutex with a standard mutex.
func NewMutexWithLock() *MutexWithLock { _ = "STUB: not implemented"; return nil }

// MutexWithLock is a mutex with a standard mutex.
type MutexWithLock struct {
	mu sync.Mutex
}

// TryLock tries to lock the mutex.
func (m *MutexWithLock) TryLock() bool { _ = "STUB: not implemented"; return false }

// Lock locks the mutex.
func (m *MutexWithLock) Lock() {
	_ = "STUB: not implemented"

	// Unlock unlocks the mutex.
	return
}

func (m *MutexWithLock) Unlock() {
	_ = "STUB: not implemented"

	/************************
	 *    Fast mutex        *
	 ************************/return
}

var _ Mutex = (*MutexWithSpinlock)(nil)

// NewMutexWithSpinlock creates a new mutex with a spinlock.
// It is faster than the standard mutex, but it is CPU-intensive.
func NewMutexWithSpinlock() *MutexWithSpinlock { _ = "STUB: not implemented"; return nil }

// MutexWithSpinlock is a mutex with a spinlock.
type MutexWithSpinlock struct {
	lock int32 // 0 or 1
}

// TryLock tries to lock the mutex.
func (m *MutexWithSpinlock) TryLock() bool { _ = "STUB: not implemented"; return false }

// Lock locks the mutex.
func (m *MutexWithSpinlock) Lock() { _ = "STUB: not implemented"; return }

// Unlock unlocks the mutex.
func (m *MutexWithSpinlock) Unlock() { _ = "STUB: not implemented"; return }

/************************
 *      Fake mutex      *
 ************************/

var _ Mutex = (*MutexWithoutLock)(nil)

// NewMutexWithoutLock creates a new mutex without a lock.
func NewMutexWithoutLock() *MutexWithoutLock { _ = "STUB: not implemented"; return nil }

// MutexWithoutLock is a mutex without a lock.
type MutexWithoutLock struct{}

// TryLock always returns true.
func (m *MutexWithoutLock) TryLock() bool {
	_ = "STUB: not implemented"

	// Lock does nothing.
	return false
}

func (m *MutexWithoutLock) Lock() {
	_ = "STUB: not implemented"

	// Unlock does nothing.
	return
}

func (m *MutexWithoutLock) Unlock() { _ = "STUB: not implemented"; return }
