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

var _ Subject[int] = (*behaviorSubjectImpl[int])(nil)

// NewBehaviorSubject emits the current value to new subscribers or initial value.
// After completion, new subscription won't receive the last value, but the error will eventually propagated.
func NewBehaviorSubject[T any](initial T) Subject[T] { _ = "STUB: not implemented"; return nil }

type behaviorSubjectImpl[T any] struct {
	mu     sync.Mutex // sync.RWMutex would be better, but it is too slow for high-volume subjects
	status Kind

	observers     sync.Map
	observerIndex uint32

	last lo.Tuple2[context.Context, T]
	err  lo.Tuple2[context.Context, error]
}

// Implements Observable.
func (s *behaviorSubjectImpl[T]) Subscribe(destination Observer[T]) Subscription {
	_ = "STUB: not implemented"
	return *new(Subscription)
}

// Implements Observable.
func (s *behaviorSubjectImpl[T]) SubscribeWithContext(subscriberCtx context.Context, destination Observer[T]) Subscription {
	_ = "STUB: not implemented"
	return *new(Subscription)
}

// fallthrough

// until we get a first value, should we send subscriberCtx or last.A (== context.TODO()) ?

func (s *behaviorSubjectImpl[T]) unsubscribeAll() { _ = "STUB: not implemented"; return }

// Implements Observer.
func (s *behaviorSubjectImpl[T]) Next(value T) { _ = "STUB: not implemented"; return }

// Implements Observer.
func (s *behaviorSubjectImpl[T]) NextWithContext(ctx context.Context, value T) {
	_ = "STUB: not implemented"
	return
}

// Implements Observer.
func (s *behaviorSubjectImpl[T]) Error(err error) { _ = "STUB: not implemented"; return }

// Implements Observer.
func (s *behaviorSubjectImpl[T]) ErrorWithContext(ctx context.Context, err error) {
	_ = "STUB: not implemented"
	return
}

// Implements Observer.
func (s *behaviorSubjectImpl[T]) Complete() { _ = "STUB: not implemented"; return }

// Implements Observer.
func (s *behaviorSubjectImpl[T]) CompleteWithContext(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (s *behaviorSubjectImpl[T]) HasObserver() (has bool) { _ = "STUB: not implemented"; return false }

func (s *behaviorSubjectImpl[T]) CountObservers() int { _ = "STUB: not implemented"; return 0 }

// Implements Observer.
func (s *behaviorSubjectImpl[T]) IsClosed() bool { _ = "STUB: not implemented"; return false }

// Implements Observer.
func (s *behaviorSubjectImpl[T]) HasThrown() bool { _ = "STUB: not implemented"; return false }

// Implements Observer.
func (s *behaviorSubjectImpl[T]) IsCompleted() bool { _ = "STUB: not implemented"; return false }

func (s *behaviorSubjectImpl[T]) AsObservable() Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

func (s *behaviorSubjectImpl[T]) AsObserver() Observer[T] { _ = "STUB: not implemented"; return nil }

func (s *behaviorSubjectImpl[T]) broadcastNext(ctx context.Context, value T) {
	_ = "STUB: not implemented"
	return
}

//nolint:errcheck,forcetypeassert

func (s *behaviorSubjectImpl[T]) broadcastError(ctx context.Context, err error) {
	_ = "STUB: not implemented"
	return
}

//nolint:errcheck,forcetypeassert

func (s *behaviorSubjectImpl[T]) broadcastComplete(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

//nolint:errcheck,forcetypeassert
