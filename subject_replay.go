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

// ReplaySubjectUnlimitedBufferSize is the unlimited buffer size for a ReplaySubject.
const ReplaySubjectUnlimitedBufferSize = -1

var _ Subject[int] = (*replaySubjectImpl[int])(nil)

// NewReplaySubject emits old values to new subscribers.
// After error or completion, new subscriptions receive values from the buffer then the error or the completion.
func NewReplaySubject[T any](bufferSize int) Subject[T] { _ = "STUB: not implemented"; return nil }

type replaySubjectImpl[T any] struct {
	mu     sync.Mutex // sync.RWMutex would be better, but it is too slow for high-volume subjects
	status Kind

	observers     sync.Map
	observerIndex uint32

	err        lo.Tuple2[context.Context, error]
	values     []lo.Tuple2[context.Context, T]
	bufferSize int
}

// Implements Observable.
func (s *replaySubjectImpl[T]) Subscribe(destination Observer[T]) Subscription {
	_ = "STUB: not implemented"
	return *new(Subscription)
}

// Implements Observable.
func (s *replaySubjectImpl[T]) SubscribeWithContext(subscriberCtx context.Context, destination Observer[T]) Subscription {
	_ = "STUB: not implemented"
	return *new(Subscription)
}

// fallthrough

func (s *replaySubjectImpl[T]) unsubscribeAll() { _ = "STUB: not implemented"; return }

// Implements Observer.
func (s *replaySubjectImpl[T]) Next(value T) { _ = "STUB: not implemented"; return }

// Implements Observer.
func (s *replaySubjectImpl[T]) NextWithContext(ctx context.Context, value T) {
	_ = "STUB: not implemented"
	return
}

// Implements Observer.
func (s *replaySubjectImpl[T]) Error(err error) { _ = "STUB: not implemented"; return }

// Implements Observer.
func (s *replaySubjectImpl[T]) ErrorWithContext(ctx context.Context, err error) {
	_ = "STUB: not implemented"
	return
}

// Implements Observer.
func (s *replaySubjectImpl[T]) Complete() { _ = "STUB: not implemented"; return }

// Implements Observer.
func (s *replaySubjectImpl[T]) CompleteWithContext(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (s *replaySubjectImpl[T]) HasObserver() bool { _ = "STUB: not implemented"; return false }

func (s *replaySubjectImpl[T]) CountObservers() int { _ = "STUB: not implemented"; return 0 }

// Implements Observer.
func (s *replaySubjectImpl[T]) IsClosed() bool { _ = "STUB: not implemented"; return false }

// Implements Observer.
func (s *replaySubjectImpl[T]) HasThrown() bool { _ = "STUB: not implemented"; return false }

// Implements Observer.
func (s *replaySubjectImpl[T]) IsCompleted() bool { _ = "STUB: not implemented"; return false }

func (s *replaySubjectImpl[T]) AsObservable() Observable[T] { _ = "STUB: not implemented"; return nil }

func (s *replaySubjectImpl[T]) AsObserver() Observer[T] { _ = "STUB: not implemented"; return nil }

func (s *replaySubjectImpl[T]) broadcastNext(ctx context.Context, value T) {
	_ = "STUB: not implemented"
	return
}

//nolint:errcheck,forcetypeassert

func (s *replaySubjectImpl[T]) broadcastError(ctx context.Context, err error) {
	_ = "STUB: not implemented"
	return
}

//nolint:errcheck,forcetypeassert

func (s *replaySubjectImpl[T]) broadcastComplete(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

//nolint:errcheck,forcetypeassert
