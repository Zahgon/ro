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

var _ Subject[int] = (*asyncSubjectImpl[int])(nil)

// NewAsyncSubject emits its last value on completion. If no value had been received, observer is completed without emitting value.
// The emitted value or error is stored for future subscriptions. 0 or 1 value is emitted.
func NewAsyncSubject[T any]() Subject[T] { _ = "STUB: not implemented"; return nil }

type asyncSubjectImpl[T any] struct {
	mu     sync.Mutex // sync.RWMutex would be better, but it is too slow for high-volume subjects
	status Kind

	observers     sync.Map
	observerIndex uint32

	hasValue bool
	value    lo.Tuple2[context.Context, T]
	err      lo.Tuple2[context.Context, error]
}

// Implements Observable.
func (s *asyncSubjectImpl[T]) Subscribe(destination Observer[T]) Subscription {
	_ = "STUB: not implemented"
	return *new(Subscription)
}

// Implements Observable.
func (s *asyncSubjectImpl[T]) SubscribeWithContext(subscriberCtx context.Context, destination Observer[T]) Subscription {
	_ = "STUB: not implemented"
	return *new(Subscription)
}

// fallthrough

func (s *asyncSubjectImpl[T]) unsubscribeAll() { _ = "STUB: not implemented"; return }

// Implements Observer.
func (s *asyncSubjectImpl[T]) Next(value T) { _ = "STUB: not implemented"; return }

// Implements Observer.
func (s *asyncSubjectImpl[T]) NextWithContext(ctx context.Context, value T) {
	_ = "STUB: not implemented"
	return
}

// A previous value might be erased. It won't be forwarded to `OnDroppedNotification`.

// Implements Observer.
func (s *asyncSubjectImpl[T]) Error(err error) { _ = "STUB: not implemented"; return }

// Implements Observer.
func (s *asyncSubjectImpl[T]) ErrorWithContext(ctx context.Context, err error) {
	_ = "STUB: not implemented"
	return
}

// Implements Observer.
func (s *asyncSubjectImpl[T]) Complete() { _ = "STUB: not implemented"; return }

// Implements Observer.
func (s *asyncSubjectImpl[T]) CompleteWithContext(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (s *asyncSubjectImpl[T]) HasObserver() bool { _ = "STUB: not implemented"; return false }

func (s *asyncSubjectImpl[T]) CountObservers() int { _ = "STUB: not implemented"; return 0 }

// Implements Observer.
func (s *asyncSubjectImpl[T]) IsClosed() bool { _ = "STUB: not implemented"; return false }

// Implements Observer.
func (s *asyncSubjectImpl[T]) HasThrown() bool { _ = "STUB: not implemented"; return false }

// Implements Observer.
func (s *asyncSubjectImpl[T]) IsCompleted() bool { _ = "STUB: not implemented"; return false }

func (s *asyncSubjectImpl[T]) AsObservable() Observable[T] { _ = "STUB: not implemented"; return nil }

func (s *asyncSubjectImpl[T]) AsObserver() Observer[T] { _ = "STUB: not implemented"; return nil }

func (s *asyncSubjectImpl[T]) broadcastNext(ctx context.Context, value T) {
	_ = "STUB: not implemented"
	return
}

//nolint:errcheck,forcetypeassert

func (s *asyncSubjectImpl[T]) broadcastError(ctx context.Context, err error) {
	_ = "STUB: not implemented"
	return
}

//nolint:errcheck,forcetypeassert

func (s *asyncSubjectImpl[T]) broadcastComplete(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

//nolint:errcheck,forcetypeassert
