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

package ro

import (
	"context"
	"sync"

	"github.com/samber/lo"
)

// UnicastSubjectUnlimitedBufferSize is the unlimited buffer size for a UnicastSubject.
const UnicastSubjectUnlimitedBufferSize = -1

var _ Subject[int] = (*unicastSubjectImpl[int])(nil)

// NewUnicastSubject queues up events until a single Observer subscribes to it,
// replays those events to it until the Observer catches up and then switches
// to relaying events live to this single Observer.
func NewUnicastSubject[T any](bufferSize int) Subject[T] { _ = "STUB: not implemented"; return nil }

type unicastSubjectImpl[T any] struct {
	mu     sync.Mutex // sync.RWMutex would be better, but it is too slow for high-volume subjects
	status Kind

	observer Observer[T]

	err        lo.Tuple2[context.Context, error]
	values     []lo.Tuple2[context.Context, T]
	bufferSize int
}

// Implements Observable.
func (s *unicastSubjectImpl[T]) Subscribe(destination Observer[T]) Subscription {
	_ = "STUB: not implemented"
	return *new(Subscription)
}

// Implements Observable.
func (s *unicastSubjectImpl[T]) SubscribeWithContext(subscriberCtx context.Context, destination Observer[T]) Subscription {
	_ = "STUB: not implemented"
	return *new(Subscription)
}

// fallthrough

// Implements Observer.
func (s *unicastSubjectImpl[T]) Next(value T) { _ = "STUB: not implemented"; return }

// Implements Observer.
func (s *unicastSubjectImpl[T]) NextWithContext(ctx context.Context, value T) {
	_ = "STUB: not implemented"
	return
}

//nolint:nestif

// out of lock

// Implements Observer.
func (s *unicastSubjectImpl[T]) Error(err error) { _ = "STUB: not implemented"; return }

// Implements Observer.
func (s *unicastSubjectImpl[T]) ErrorWithContext(ctx context.Context, err error) {
	_ = "STUB: not implemented"
	return
}

// Implements Observer.
func (s *unicastSubjectImpl[T]) Complete() { _ = "STUB: not implemented"; return }

// Implements Observer.
func (s *unicastSubjectImpl[T]) CompleteWithContext(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (s *unicastSubjectImpl[T]) HasObserver() bool { _ = "STUB: not implemented"; return false }

func (s *unicastSubjectImpl[T]) CountObservers() int { _ = "STUB: not implemented"; return 0 }

// Implements Observer.
func (s *unicastSubjectImpl[T]) IsClosed() bool { _ = "STUB: not implemented"; return false }

// Implements Observer.
func (s *unicastSubjectImpl[T]) HasThrown() bool { _ = "STUB: not implemented"; return false }

// Implements Observer.
func (s *unicastSubjectImpl[T]) IsCompleted() bool { _ = "STUB: not implemented"; return false }

func (s *unicastSubjectImpl[T]) AsObservable() Observable[T] { _ = "STUB: not implemented"; return nil }

func (s *unicastSubjectImpl[T]) AsObserver() Observer[T] { _ = "STUB: not implemented"; return nil }
