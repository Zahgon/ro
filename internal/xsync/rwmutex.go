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

// RWMutex is a read-write mutex interface.
type RWMutex interface {
	TryLock() bool
	Lock()
	Unlock()
	TryRLock() bool
	RLock()
	RUnlock()
}

/************************
 *    Standard mutex    *
 ************************/

var _ RWMutex = (*RWMutexWithLock)(nil)

// NewRWMutexWithLock creates a new read-write mutex with a standard read-write mutex.
func NewRWMutexWithLock() *RWMutexWithLock { _ = "STUB: not implemented"; return nil }

// RWMutexWithLock is a read-write mutex with a standard read-write mutex.
type RWMutexWithLock struct {
	mu sync.RWMutex
}

// TryLock tries to lock the mutex.
func (m *RWMutexWithLock) TryLock() bool { _ = "STUB: not implemented"; return false }

// Lock locks the mutex.
func (m *RWMutexWithLock) Lock() {
	_ = "STUB: not implemented"

	// Unlock unlocks the mutex.
	return
}

func (m *RWMutexWithLock) Unlock() {
	_ = "STUB: not implemented"

	// TryRLock tries to lock the mutex for reading.
	return
}

func (m *RWMutexWithLock) TryRLock() bool { _ = "STUB: not implemented"; return false }

// RLock locks the mutex for reading.
func (m *RWMutexWithLock) RLock() {
	_ = "STUB: not implemented"

	// RUnlock unlocks the mutex for reading.
	return
}

func (m *RWMutexWithLock) RUnlock() {
	_ = "STUB: not implemented"

	/************************
	 *    Fast mutex        *
	 ************************/return
}

// @TODO

/************************
 *      Fake mutex      *
 ************************/

var _ RWMutex = (*RWMutexWithoutLock)(nil)

// NewRWMutexWithoutLock creates a new read-write mutex without a lock.
func NewRWMutexWithoutLock() *RWMutexWithoutLock { _ = "STUB: not implemented"; return nil }

// RWMutexWithoutLock is a read-write mutex without a lock.
type RWMutexWithoutLock struct{}

// TryLock always returns true.
func (m *RWMutexWithoutLock) TryLock() bool {
	_ = "STUB: not implemented"

	// Lock does nothing.
	return false
}

func (m *RWMutexWithoutLock) Lock() {
	_ = "STUB: not implemented"

	// Unlock does nothing.
	return
}

func (m *RWMutexWithoutLock) Unlock() {
	_ = "STUB: not implemented"

	// TryRLock always returns true.
	return
}

func (m *RWMutexWithoutLock) TryRLock() bool {
	_ = "STUB: not implemented"

	// RLock does nothing.
	return false
}

func (m *RWMutexWithoutLock) RLock() {
	_ = "STUB: not implemented"

	// RUnlock does nothing.
	return
}

func (m *RWMutexWithoutLock) RUnlock() { _ = "STUB: not implemented"; return }
