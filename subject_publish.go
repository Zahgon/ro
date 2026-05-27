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

var _ Subject[int] = (*publishSubjectImpl[int])(nil)

// NewPublishSubject broadcasts a value to observers (fanout).
// Values received before subscription are not transmitted.
func NewPublishSubject[T any]() Subject[T] { _ = "STUB: not implemented"; return nil }

type publishSubjectImpl[T any] struct {
	mu     sync.Mutex // sync.RWMutex would be better, but it is too slow for high-volume subjects
	status Kind

	observers     sync.Map
	observerIndex uint32

	err lo.Tuple2[context.Context, error]
}

// Implements Observable.
func (s *publishSubjectImpl[T]) Subscribe(destination Observer[T]) Subscription {
	_ = "STUB: not implemented"
	return *new(Subscription)
}

// Implements Observable.
func (s *publishSubjectImpl[T]) SubscribeWithContext(subscriberCtx context.Context, destination Observer[T]) Subscription {
	_ = "STUB: not implemented"
	return *new(Subscription)
}

// fallthrough

func (s *publishSubjectImpl[T]) unsubscribeAll() { _ = "STUB: not implemented"; return }

// Implements Observer.
func (s *publishSubjectImpl[T]) Next(value T) { _ = "STUB: not implemented"; return }

// Implements Observer.
func (s *publishSubjectImpl[T]) NextWithContext(ctx context.Context, value T) {
	_ = "STUB: not implemented"
	return
}

// Implements Observer.
func (s *publishSubjectImpl[T]) Error(err error) { _ = "STUB: not implemented"; return }

// Implements Observer.
func (s *publishSubjectImpl[T]) ErrorWithContext(ctx context.Context, err error) {
	_ = "STUB: not implemented"
	return
}

// Implements Observer.
func (s *publishSubjectImpl[T]) Complete() { _ = "STUB: not implemented"; return }

// Implements Observer.
func (s *publishSubjectImpl[T]) CompleteWithContext(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (s *publishSubjectImpl[T]) HasObserver() (has bool) { _ = "STUB: not implemented"; return false }

func (s *publishSubjectImpl[T]) CountObservers() int { _ = "STUB: not implemented"; return 0 }

// Implements Observer.
func (s *publishSubjectImpl[T]) IsClosed() bool { _ = "STUB: not implemented"; return false }

// Implements Observer.
func (s *publishSubjectImpl[T]) HasThrown() bool { _ = "STUB: not implemented"; return false }

// Implements Observer.
func (s *publishSubjectImpl[T]) IsCompleted() bool { _ = "STUB: not implemented"; return false }

func (s *publishSubjectImpl[T]) AsObservable() Observable[T] { _ = "STUB: not implemented"; return nil }

func (s *publishSubjectImpl[T]) AsObserver() Observer[T] { _ = "STUB: not implemented"; return nil }

func (s *publishSubjectImpl[T]) broadcastNext(ctx context.Context, value T) {
	_ = "STUB: not implemented"
	return
}

//nolint:errcheck,forcetypeassert

func (s *publishSubjectImpl[T]) broadcastError(ctx context.Context, err error) {
	_ = "STUB: not implemented"
	return
}

//nolint:errcheck,forcetypeassert

func (s *publishSubjectImpl[T]) broadcastComplete(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

//nolint:errcheck,forcetypeassert
