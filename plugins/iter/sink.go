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

package roiter

import (
	"iter"

	"github.com/samber/ro"
)

// ToSeq converts an observable to a Go sequence iterator.
func ToSeq[T any](source ro.Observable[T]) iter.Seq[T] { _ = "STUB: not implemented"; return nil }

// Create channels for synchronization

// Create a context for cancellation

// Subscribe to the observable

// Clean up subscription

// Yield values as they arrive

// ToSeq2 converts an observable to a Go sequence iterator with index-value pairs.
func ToSeq2[T any](source ro.Observable[T]) iter.Seq2[int, T] {
	_ = "STUB: not implemented"
	return nil
}

// Create channels for synchronization

// Create a context for cancellation

// Subscribe to the observable

// Clean up subscription

// Yield key-value pairs as they arrive
